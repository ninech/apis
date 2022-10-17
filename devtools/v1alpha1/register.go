package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "devtools.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	ArgoCDKind             = reflect.TypeOf(ArgoCD{}).Name()
	ArgoCDGroupKind        = schema.GroupKind{Group: Group, Kind: ArgoCDKind}.String()
	ArgoCDKindAPIVersion   = ArgoCDKind + "." + SchemeGroupVersion.String()
	ArgoCDGroupVersionKind = SchemeGroupVersion.WithKind(ArgoCDKind)
)

func init() {
	SchemeBuilder.Register(&ArgoCD{}, &ArgoCDList{})
}
