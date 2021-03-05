/*
Copyright 2021 The Kubernetes Authors.

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

package rollout

import (
	"testing"

	"github.com/spf13/cobra"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"
)

func TestCompletionRolloutRestartNoArg(t *testing.T) {
	tf, streams := cmdtesting.PreparePodsForCompletion(t)
	cmd := NewCmdRolloutRestart(tf, streams)
	comps, directive := cmd.ValidArgsFunction(cmd, []string{}, "")
	cmdtesting.CheckCompletion(t, comps, []string{"deployment", "daemonset", "statefulset"}, directive, cobra.ShellCompDirectiveNoFileComp)
}

func TestCompletionRolloutRestartOneArg(t *testing.T) {
	tf, streams := cmdtesting.PreparePodsForCompletion(t)
	cmd := NewCmdRolloutRestart(tf, streams)
	comps, directive := cmd.ValidArgsFunction(cmd, []string{"pods"}, "b")
	cmdtesting.CheckCompletion(t, comps, []string{"bar"}, directive, cobra.ShellCompDirectiveNoFileComp)
}
