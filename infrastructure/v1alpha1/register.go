package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "infrastructure.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	KedaKind             = reflect.TypeOf(Keda{}).Name()
	KedaGroupKind        = schema.GroupKind{Group: Group, Kind: KedaKind}.String()
	KedaKindAPIVersion   = KedaKind + "." + SchemeGroupVersion.String()
	KedaGroupVersionKind = SchemeGroupVersion.WithKind(KedaKind)
	KubernetesClusterKind             = reflect.TypeOf(KubernetesCluster{}).Name()
	KubernetesClusterGroupKind        = schema.GroupKind{Group: Group, Kind: KubernetesClusterKind}.String()
	KubernetesClusterKindAPIVersion   = KubernetesClusterKind + "." + SchemeGroupVersion.String()
	KubernetesClusterGroupVersionKind = SchemeGroupVersion.WithKind(KubernetesClusterKind)
)

func init() {
	SchemeBuilder.Register(&Keda{}, &KedaList{})
	SchemeBuilder.Register(&KubernetesCluster{}, &KubernetesClusterList{})
}
