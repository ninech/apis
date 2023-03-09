package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "iam.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	APIServiceAccountKind             = reflect.TypeOf(APIServiceAccount{}).Name()
	APIServiceAccountGroupKind        = schema.GroupKind{Group: Group, Kind: APIServiceAccountKind}.String()
	APIServiceAccountKindAPIVersion   = APIServiceAccountKind + "." + SchemeGroupVersion.String()
	APIServiceAccountGroupVersionKind = SchemeGroupVersion.WithKind(APIServiceAccountKind)
	KubernetesClustersRoleBindingKind             = reflect.TypeOf(KubernetesClustersRoleBinding{}).Name()
	KubernetesClustersRoleBindingGroupKind        = schema.GroupKind{Group: Group, Kind: KubernetesClustersRoleBindingKind}.String()
	KubernetesClustersRoleBindingKindAPIVersion   = KubernetesClustersRoleBindingKind + "." + SchemeGroupVersion.String()
	KubernetesClustersRoleBindingGroupVersionKind = SchemeGroupVersion.WithKind(KubernetesClustersRoleBindingKind)
	KubernetesServiceAccountKind             = reflect.TypeOf(KubernetesServiceAccount{}).Name()
	KubernetesServiceAccountGroupKind        = schema.GroupKind{Group: Group, Kind: KubernetesServiceAccountKind}.String()
	KubernetesServiceAccountKindAPIVersion   = KubernetesServiceAccountKind + "." + SchemeGroupVersion.String()
	KubernetesServiceAccountGroupVersionKind = SchemeGroupVersion.WithKind(KubernetesServiceAccountKind)
)

func init() {
	SchemeBuilder.Register(&APIServiceAccount{}, &APIServiceAccountList{})
	SchemeBuilder.Register(&KubernetesClustersRoleBinding{}, &KubernetesClustersRoleBindingList{})
	SchemeBuilder.Register(&KubernetesServiceAccount{}, &KubernetesServiceAccountList{})
}
