package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "observability.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	AlertmanagerKind             = reflect.TypeOf(Alertmanager{}).Name()
	AlertmanagerGroupKind        = schema.GroupKind{Group: Group, Kind: AlertmanagerKind}.String()
	AlertmanagerKindAPIVersion   = AlertmanagerKind + "." + SchemeGroupVersion.String()
	AlertmanagerGroupVersionKind = SchemeGroupVersion.WithKind(AlertmanagerKind)
	GrafanaKind             = reflect.TypeOf(Grafana{}).Name()
	GrafanaGroupKind        = schema.GroupKind{Group: Group, Kind: GrafanaKind}.String()
	GrafanaKindAPIVersion   = GrafanaKind + "." + SchemeGroupVersion.String()
	GrafanaGroupVersionKind = SchemeGroupVersion.WithKind(GrafanaKind)
	LokiKind             = reflect.TypeOf(Loki{}).Name()
	LokiGroupKind        = schema.GroupKind{Group: Group, Kind: LokiKind}.String()
	LokiKindAPIVersion   = LokiKind + "." + SchemeGroupVersion.String()
	LokiGroupVersionKind = SchemeGroupVersion.WithKind(LokiKind)
	PrometheusKind             = reflect.TypeOf(Prometheus{}).Name()
	PrometheusGroupKind        = schema.GroupKind{Group: Group, Kind: PrometheusKind}.String()
	PrometheusKindAPIVersion   = PrometheusKind + "." + SchemeGroupVersion.String()
	PrometheusGroupVersionKind = SchemeGroupVersion.WithKind(PrometheusKind)
	PromtailKind             = reflect.TypeOf(Promtail{}).Name()
	PromtailGroupKind        = schema.GroupKind{Group: Group, Kind: PromtailKind}.String()
	PromtailKindAPIVersion   = PromtailKind + "." + SchemeGroupVersion.String()
	PromtailGroupVersionKind = SchemeGroupVersion.WithKind(PromtailKind)
)

func init() {
	SchemeBuilder.Register(&Alertmanager{}, &AlertmanagerList{})
	SchemeBuilder.Register(&Grafana{}, &GrafanaList{})
	SchemeBuilder.Register(&Loki{}, &LokiList{})
	SchemeBuilder.Register(&Prometheus{}, &PrometheusList{})
	SchemeBuilder.Register(&Promtail{}, &PromtailList{})
}
