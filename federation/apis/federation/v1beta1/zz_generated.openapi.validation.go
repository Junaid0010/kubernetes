// +build !ignore_autogenerated

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

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	field "k8s.io/kubernetes/pkg/util/validation/field"
	validation "k8s.io/kubernetes/pkg/validation"
)

func (s Cluster) Validate(meta *validation.FieldMeta, op validation.OperationType) field.ErrorList {
	allErrs := field.ErrorList{}
	if errs := s.ObjectMeta.Validate(&validation.FieldMeta{Path: meta.Child("ObjectMeta")}, op); len(errs) != 0 {
		allErrs = append(allErrs, errs...)
	}
	return allErrs
}
