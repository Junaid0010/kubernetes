// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package audit

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_Event, InType: reflect.TypeOf(&Event{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_EventList, InType: reflect.TypeOf(&EventList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_GroupKinds, InType: reflect.TypeOf(&GroupKinds{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_ObjectReference, InType: reflect.TypeOf(&ObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_audit_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
	)
}

func DeepCopy_audit_Event(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Event)
		out := out.(*Event)
		*out = *in
		out.Timestamp = in.Timestamp.DeepCopy()
		if newVal, err := c.DeepCopy(&in.User); err != nil {
			return err
		} else {
			out.User = *newVal.(*UserInfo)
		}
		if in.Impersonate != nil {
			in, out := &in.Impersonate, &out.Impersonate
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*UserInfo)
			}
		}
		if in.ObjectRef != nil {
			in, out := &in.ObjectRef, &out.ObjectRef
			*out = new(ObjectReference)
			**out = **in
		}
		if in.ResponseStatus != nil {
			in, out := &in.ResponseStatus, &out.ResponseStatus
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.Status)
			}
		}
		// in.RequestObject is kind 'Interface'
		if in.RequestObject != nil {
			if newVal, err := c.DeepCopy(&in.RequestObject); err != nil {
				return err
			} else {
				out.RequestObject = *newVal.(*runtime.Object)
			}
		}
		// in.ResponseObject is kind 'Interface'
		if in.ResponseObject != nil {
			if newVal, err := c.DeepCopy(&in.ResponseObject); err != nil {
				return err
			} else {
				out.ResponseObject = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}

func DeepCopy_audit_EventList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EventList)
		out := out.(*EventList)
		*out = *in
		if in.Events != nil {
			in, out := &in.Events, &out.Events
			*out = make([]Event, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Event)
				}
			}
		}
		return nil
	}
}

func DeepCopy_audit_GroupKinds(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupKinds)
		out := out.(*GroupKinds)
		*out = *in
		if in.Kinds != nil {
			in, out := &in.Kinds, &out.Kinds
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_audit_ObjectReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectReference)
		out := out.(*ObjectReference)
		*out = *in
		return nil
	}
}

func DeepCopy_audit_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*PolicyRule)
				}
			}
		}
		return nil
	}
}

func DeepCopy_audit_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.UserGroups != nil {
			in, out := &in.UserGroups, &out.UserGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ResourceKinds != nil {
			in, out := &in.ResourceKinds, &out.ResourceKinds
			*out = make([]GroupKinds, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*GroupKinds)
				}
			}
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_audit_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserInfo)
		out := out.(*UserInfo)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		}
		return nil
	}
}
