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
	IngressHAProxyKind             = reflect.TypeOf(IngressHAProxy{}).Name()
	IngressHAProxyGroupKind        = schema.GroupKind{Group: Group, Kind: IngressHAProxyKind}.String()
	IngressHAProxyKindAPIVersion   = IngressHAProxyKind + "." + SchemeGroupVersion.String()
	IngressHAProxyGroupVersionKind = SchemeGroupVersion.WithKind(IngressHAProxyKind)
	IngressNginxKind             = reflect.TypeOf(IngressNginx{}).Name()
	IngressNginxGroupKind        = schema.GroupKind{Group: Group, Kind: IngressNginxKind}.String()
	IngressNginxKindAPIVersion   = IngressNginxKind + "." + SchemeGroupVersion.String()
	IngressNginxGroupVersionKind = SchemeGroupVersion.WithKind(IngressNginxKind)
	ServiceConnectionKind             = reflect.TypeOf(ServiceConnection{}).Name()
	ServiceConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceConnectionKind}.String()
	ServiceConnectionKindAPIVersion   = ServiceConnectionKind + "." + SchemeGroupVersion.String()
	ServiceConnectionGroupVersionKind = SchemeGroupVersion.WithKind(ServiceConnectionKind)
	StaticEgressKind             = reflect.TypeOf(StaticEgress{}).Name()
	StaticEgressGroupKind        = schema.GroupKind{Group: Group, Kind: StaticEgressKind}.String()
	StaticEgressKindAPIVersion   = StaticEgressKind + "." + SchemeGroupVersion.String()
	StaticEgressGroupVersionKind = SchemeGroupVersion.WithKind(StaticEgressKind)
)

func init() {
	SchemeBuilder.Register(&IngressHAProxy{}, &IngressHAProxyList{})
	SchemeBuilder.Register(&IngressNginx{}, &IngressNginxList{})
	SchemeBuilder.Register(&ServiceConnection{}, &ServiceConnectionList{})
	SchemeBuilder.Register(&StaticEgress{}, &StaticEgressList{})
}
