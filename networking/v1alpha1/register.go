package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "networking.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	IngressNginxKind             = reflect.TypeOf(IngressNginx{}).Name()
	IngressNginxGroupKind        = schema.GroupKind{Group: Group, Kind: IngressNginxKind}.String()
	IngressNginxKindAPIVersion   = IngressNginxKind + "." + SchemeGroupVersion.String()
	IngressNginxGroupVersionKind = SchemeGroupVersion.WithKind(IngressNginxKind)
)

func init() {
	SchemeBuilder.Register(&IngressNginx{}, &IngressNginxList{})
}
