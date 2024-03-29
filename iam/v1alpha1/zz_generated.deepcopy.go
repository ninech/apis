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
	metav1alpha1 "github.com/ninech/apis/meta/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceAccount) DeepCopyInto(out *APIServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceAccount.
func (in *APIServiceAccount) DeepCopy() *APIServiceAccount {
	if in == nil {
		return nil
	}
	out := new(APIServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceAccountList) DeepCopyInto(out *APIServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceAccountList.
func (in *APIServiceAccountList) DeepCopy() *APIServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(APIServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceAccountParameters) DeepCopyInto(out *APIServiceAccountParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceAccountParameters.
func (in *APIServiceAccountParameters) DeepCopy() *APIServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(APIServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceAccountSpec) DeepCopyInto(out *APIServiceAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceAccountSpec.
func (in *APIServiceAccountSpec) DeepCopy() *APIServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceAccountStatus) DeepCopyInto(out *APIServiceAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceAccountStatus.
func (in *APIServiceAccountStatus) DeepCopy() *APIServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(APIServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBinding) DeepCopyInto(out *KubernetesClustersRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBinding.
func (in *KubernetesClustersRoleBinding) DeepCopy() *KubernetesClustersRoleBinding {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesClustersRoleBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBindingList) DeepCopyInto(out *KubernetesClustersRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubernetesClustersRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBindingList.
func (in *KubernetesClustersRoleBindingList) DeepCopy() *KubernetesClustersRoleBindingList {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesClustersRoleBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBindingObservation) DeepCopyInto(out *KubernetesClustersRoleBindingObservation) {
	*out = *in
	in.ReferenceStatus.DeepCopyInto(&out.ReferenceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBindingObservation.
func (in *KubernetesClustersRoleBindingObservation) DeepCopy() *KubernetesClustersRoleBindingObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBindingParameters) DeepCopyInto(out *KubernetesClustersRoleBindingParameters) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]RoleBindingSubject, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]metav1alpha1.LocalReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBindingParameters.
func (in *KubernetesClustersRoleBindingParameters) DeepCopy() *KubernetesClustersRoleBindingParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBindingSpec) DeepCopyInto(out *KubernetesClustersRoleBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBindingSpec.
func (in *KubernetesClustersRoleBindingSpec) DeepCopy() *KubernetesClustersRoleBindingSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClustersRoleBindingStatus) DeepCopyInto(out *KubernetesClustersRoleBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClustersRoleBindingStatus.
func (in *KubernetesClustersRoleBindingStatus) DeepCopy() *KubernetesClustersRoleBindingStatus {
	if in == nil {
		return nil
	}
	out := new(KubernetesClustersRoleBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccount) DeepCopyInto(out *KubernetesServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccount.
func (in *KubernetesServiceAccount) DeepCopy() *KubernetesServiceAccount {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountList) DeepCopyInto(out *KubernetesServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubernetesServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountList.
func (in *KubernetesServiceAccountList) DeepCopy() *KubernetesServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountObservation) DeepCopyInto(out *KubernetesServiceAccountObservation) {
	*out = *in
	in.ReferenceStatus.DeepCopyInto(&out.ReferenceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountObservation.
func (in *KubernetesServiceAccountObservation) DeepCopy() *KubernetesServiceAccountObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountParameters) DeepCopyInto(out *KubernetesServiceAccountParameters) {
	*out = *in
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountParameters.
func (in *KubernetesServiceAccountParameters) DeepCopy() *KubernetesServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountSpec) DeepCopyInto(out *KubernetesServiceAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountSpec.
func (in *KubernetesServiceAccountSpec) DeepCopy() *KubernetesServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountStatus) DeepCopyInto(out *KubernetesServiceAccountStatus) {
	*out = *in
	in.AtProvider.DeepCopyInto(&out.AtProvider)
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountStatus.
func (in *KubernetesServiceAccountStatus) DeepCopy() *KubernetesServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingSubject) DeepCopyInto(out *RoleBindingSubject) {
	*out = *in
	out.LocalReference = in.LocalReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingSubject.
func (in *RoleBindingSubject) DeepCopy() *RoleBindingSubject {
	if in == nil {
		return nil
	}
	out := new(RoleBindingSubject)
	in.DeepCopyInto(out)
	return out
}
