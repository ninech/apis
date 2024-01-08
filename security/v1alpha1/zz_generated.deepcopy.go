//go:build !ignore_autogenerated

/*
Copyright 2022 Nine Internet Solutions AG.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecrets) DeepCopyInto(out *ExternalSecrets) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecrets.
func (in *ExternalSecrets) DeepCopy() *ExternalSecrets {
	if in == nil {
		return nil
	}
	out := new(ExternalSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecrets) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretsList) DeepCopyInto(out *ExternalSecretsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecrets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretsList.
func (in *ExternalSecretsList) DeepCopy() *ExternalSecretsList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretsObservation) DeepCopyInto(out *ExternalSecretsObservation) {
	*out = *in
	in.ChildResourceStatus.DeepCopyInto(&out.ChildResourceStatus)
	in.ReferenceStatus.DeepCopyInto(&out.ReferenceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretsObservation.
func (in *ExternalSecretsObservation) DeepCopy() *ExternalSecretsObservation {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretsParameters) DeepCopyInto(out *ExternalSecretsParameters) {
	*out = *in
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretsParameters.
func (in *ExternalSecretsParameters) DeepCopy() *ExternalSecretsParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretsSpec) DeepCopyInto(out *ExternalSecretsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretsSpec.
func (in *ExternalSecretsSpec) DeepCopy() *ExternalSecretsSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretsStatus) DeepCopyInto(out *ExternalSecretsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretsStatus.
func (in *ExternalSecretsStatus) DeepCopy() *ExternalSecretsStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKey) DeepCopyInto(out *SSHKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKey.
func (in *SSHKey) DeepCopy() *SSHKey {
	if in == nil {
		return nil
	}
	out := new(SSHKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyList) DeepCopyInto(out *SSHKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyList.
func (in *SSHKeyList) DeepCopy() *SSHKeyList {
	if in == nil {
		return nil
	}
	out := new(SSHKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyObservation) DeepCopyInto(out *SSHKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyObservation.
func (in *SSHKeyObservation) DeepCopy() *SSHKeyObservation {
	if in == nil {
		return nil
	}
	out := new(SSHKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyParameters) DeepCopyInto(out *SSHKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyParameters.
func (in *SSHKeyParameters) DeepCopy() *SSHKeyParameters {
	if in == nil {
		return nil
	}
	out := new(SSHKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeySpec) DeepCopyInto(out *SSHKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeySpec.
func (in *SSHKeySpec) DeepCopy() *SSHKeySpec {
	if in == nil {
		return nil
	}
	out := new(SSHKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyStatus) DeepCopyInto(out *SSHKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyStatus.
func (in *SSHKeyStatus) DeepCopy() *SSHKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SSHKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecrets) DeepCopyInto(out *SealedSecrets) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecrets.
func (in *SealedSecrets) DeepCopy() *SealedSecrets {
	if in == nil {
		return nil
	}
	out := new(SealedSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SealedSecrets) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsList) DeepCopyInto(out *SealedSecretsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SealedSecrets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsList.
func (in *SealedSecretsList) DeepCopy() *SealedSecretsList {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SealedSecretsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsObservation) DeepCopyInto(out *SealedSecretsObservation) {
	*out = *in
	in.ChildResourceStatus.DeepCopyInto(&out.ChildResourceStatus)
	in.ReferenceStatus.DeepCopyInto(&out.ReferenceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsObservation.
func (in *SealedSecretsObservation) DeepCopy() *SealedSecretsObservation {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsParameters) DeepCopyInto(out *SealedSecretsParameters) {
	*out = *in
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsParameters.
func (in *SealedSecretsParameters) DeepCopy() *SealedSecretsParameters {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsSpec) DeepCopyInto(out *SealedSecretsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsSpec.
func (in *SealedSecretsSpec) DeepCopy() *SealedSecretsSpec {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsStatus) DeepCopyInto(out *SealedSecretsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsStatus.
func (in *SealedSecretsStatus) DeepCopy() *SealedSecretsStatus {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsStatus)
	in.DeepCopyInto(out)
	return out
}
