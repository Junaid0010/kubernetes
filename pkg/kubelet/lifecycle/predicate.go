/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lifecycle

import (
	"fmt"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubelet/util/format"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	"k8s.io/kubernetes/pkg/scheduler/schedulercache"
)

type getNodeAnyWayFuncType func() (*v1.Node, error)

type pluginResourceUpdateFuncType func(*schedulercache.NodeInfo, *PodAdmitAttributes) error

// AdmissionFailureHandler is an interface which defines how to deal with a failure to admit a pod.
// This allows for the graceful handling of pod admission failure.
type AdmissionFailureHandler interface {
	HandleAdmissionFailure(pod *v1.Pod, failureReasons []algorithm.PredicateFailureReason) (bool, []algorithm.PredicateFailureReason, error)
}

type predicateAdmitHandler struct {
	getNodeAnyWayFunc        getNodeAnyWayFuncType
	pluginResourceUpdateFunc pluginResourceUpdateFuncType
	admissionFailureHandler  AdmissionFailureHandler
}

var _ PodAdmitHandler = &predicateAdmitHandler{}

func NewPredicateAdmitHandler(getNodeAnyWayFunc getNodeAnyWayFuncType, admissionFailureHandler AdmissionFailureHandler, pluginResourceUpdateFunc pluginResourceUpdateFuncType) *predicateAdmitHandler {
	return &predicateAdmitHandler{
		getNodeAnyWayFunc,
		pluginResourceUpdateFunc,
		admissionFailureHandler,
	}
}

func (w *predicateAdmitHandler) Admit(attrs *PodAdmitAttributes) PodAdmitResult {
	node, err := w.getNodeAnyWayFunc()
	if err != nil {
		glog.Errorf("Cannot get Node info: %v", err)
		return PodAdmitResult{
			Admit:   false,
			Reason:  "InvalidNodeInfo",
			Message: "Kubelet cannot get node info.",
		}
	}
	pod := attrs.Pod
	pods := attrs.OtherPods
	nodeInfo := schedulercache.NewNodeInfo(pods...)
	nodeInfo.SetNode(node)
	// Ignore missing extended resources. This is required to support cluster level resources which are not tied to Nodes.
	// If there happens to be a node level resource that is exposed via device plugin as an extended resource,
	// it will also be ignored if it is not found in the Node Capacity and Allocatable.
	// There is an implicit assumption that the lifecycle of device plugins will affect node
	// health and nodes marked unhealthy should not be accepting new pods and
	// possibly evicting existing pods too.
	predicates.RegisterPodFitsResourcesPredicateMetadataProducer(true /* ignoreMissingExtendedResources */)
	// ensure the node has enough plugin resources for that required in pods
	if err = w.pluginResourceUpdateFunc(nodeInfo, attrs); err != nil {
		message := fmt.Sprintf("Update plugin resources failed due to %v, which is unexpected.", err)
		glog.Warningf("Failed to admit pod %v - %s", format.Pod(pod), message)
		return PodAdmitResult{
			Admit:   false,
			Reason:  "UnexpectedAdmissionError",
			Message: message,
		}
	}

	fit, reasons, err := predicates.GeneralPredicates(pod, nil, nodeInfo)
	if err != nil {
		message := fmt.Sprintf("GeneralPredicates failed due to %v, which is unexpected.", err)
		glog.Warningf("Failed to admit pod %v - %s", format.Pod(pod), message)
		return PodAdmitResult{
			Admit:   fit,
			Reason:  "UnexpectedAdmissionError",
			Message: message,
		}
	}
	if !fit {
		fit, reasons, err = w.admissionFailureHandler.HandleAdmissionFailure(pod, reasons)
		if err != nil {
			message := fmt.Sprintf("Unexpected error while attempting to recover from admission failure: %v", err)
			glog.Warningf("Failed to admit pod %v - %s", format.Pod(pod), message)
			return PodAdmitResult{
				Admit:   fit,
				Reason:  "UnexpectedAdmissionError",
				Message: message,
			}
		}
	}
	if !fit {
		var reason string
		var message string
		if len(reasons) == 0 {
			message = fmt.Sprint("GeneralPredicates failed due to unknown reason, which is unexpected.")
			glog.Warningf("Failed to admit pod %v - %s", format.Pod(pod), message)
			return PodAdmitResult{
				Admit:   fit,
				Reason:  "UnknownReason",
				Message: message,
			}
		}
		// If there are failed predicates, we only return the first one as a reason.
		r := reasons[0]
		switch re := r.(type) {
		case *predicates.PredicateFailureError:
			reason = re.PredicateName
			message = re.Error()
			glog.V(2).Infof("Predicate failed on Pod: %v, for reason: %v", format.Pod(pod), message)
		case *predicates.InsufficientResourceError:
			reason = fmt.Sprintf("OutOf%s", re.ResourceName)
			message = re.Error()
			glog.V(2).Infof("Predicate failed on Pod: %v, for reason: %v", format.Pod(pod), message)
		case *predicates.FailureReason:
			reason = re.GetReason()
			message = fmt.Sprintf("Failure: %s", re.GetReason())
			glog.V(2).Infof("Predicate failed on Pod: %v, for reason: %v", format.Pod(pod), message)
		default:
			reason = "UnexpectedPredicateFailureType"
			message = fmt.Sprintf("GeneralPredicates failed due to %v, which is unexpected.", r)
			glog.Warningf("Failed to admit pod %v - %s", format.Pod(pod), message)
		}
		return PodAdmitResult{
			Admit:   fit,
			Reason:  reason,
			Message: message,
		}
	}
	return PodAdmitResult{
		Admit: true,
	}
}
