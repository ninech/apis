package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	"github.com/prometheus/common/model"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Alertmanager deploys a fully managed Alertmanager instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="CONFIG VALID",type="boolean",JSONPath=".status.atProvider.config.valid"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=am
// +kubebuilder:object:root=true
type Alertmanager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertmanagerSpec   `json:"spec"`
	Status            AlertmanagerStatus `json:"status,omitempty"`
}

// AlertmanagerList contains a list of Alertmanagers.
// +kubebuilder:object:root=true
type AlertmanagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alertmanager `json:"items"`
}

// AlertmanagerSpec defines the desired state of an Alertmanager.
type AlertmanagerSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            AlertmanagerParameters `json:"forProvider"`
}

// AlertmanagerParameters are the configurable fields of an Alertmanager.
type AlertmanagerParameters struct {
	// Config is a yaml string containing the  alertmanager configuration.
	// Reference:
	// https://www.prometheus.io/docs/alerting/latest/configuration/
	// +optional
	Config string `json:"config,omitempty"`
	// ConfigFromSecret specifies a secret name where the config is stored in.
	// This is mutually exclusive with the "Config" field.
	// Reference:
	// https://www.prometheus.io/docs/alerting/latest/configuration/
	// +optional
	ConfigFromSecret *meta.LocalReferenceSelector `json:"configFromSecret,omitempty"`
}

// AlertmanagerStatus represents the observed state of an Alertmanager.
type AlertmanagerStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               AlertmanagerObservation `json:"atProvider,omitempty"`
}

// AlertmanagerObservation are the observable fields of an Alertmanager.
type AlertmanagerObservation struct {
	// URL where alertmanager can be accessed.
	URL string `json:"url,omitempty"`
	// Targets are the target hosts that should be used for targeting
	// Alertmanager from Prometheus.
	// +optional
	Targets                  []string                 `json:"targets,omitempty"`
	Config                   AlertmanagerConfigStatus `json:"config,omitempty"`
	meta.ChildResourceStatus `json:",inline"`
}

// AlertmanagerConfigStatus contains information about the config status.
type AlertmanagerConfigStatus struct {
	// Valid indicates if the desired config is valid.
	Valid bool `json:"valid,omitempty"`
	// Message describes the status of the alertmanager config.
	Message string `json:"message,omitempty"`
}

// Grafana deploys a fully managed Grafana instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Grafana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrafanaSpec   `json:"spec"`
	Status            GrafanaStatus `json:"status,omitempty"`
}

// GrafanaList contains a list of Grafana
// +kubebuilder:object:root=true
type GrafanaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grafana `json:"items"`
}

// An GrafanaSpec defines the desired state of a Grafana.
type GrafanaSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            GrafanaParameters `json:"forProvider"`
}

// GrafanaParameters are the configurable fields of a Grafana.
type GrafanaParameters struct {
	// DataSource defines search parameters for DataSources.
	DataSource DataSourceSelection `json:"dataSource"`
	// EnableAdminAccess gives you admin permissions in the Grafana instance.
	// Please note: Nine reserves the right to revoke admin access in future
	// versions if the admin access is deemed a security or an operational risk.
	// +optional
	EnableAdminAccess bool `json:"enableAdminAccess,omitempty"`
}

// DataSourceSelection specifies parameters on how to select / search for DataSources.
type DataSourceSelection struct {
	// SearchNamespaces is a list of namespaces that will be searched
	// for data sources that will be added to the Grafana deployment.
	// If empty, the default is to search in the same namespace that
	// Grafana resides in.
	SearchNamespaces []string `json:"searchNamespaces,omitempty"`
	// FilterLabels is a list of labels that will be filtered for
	// when searching for data sources.
	// +optional
	FilterLabels map[string]string `json:"filterLabels,omitempty"`
}

// An GrafanaStatus represents the observed state of a Grafana.
type GrafanaStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               GrafanaObservation `json:"atProvider"`
}

// GrafanaObservation are the observable fields of a Grafana.
type GrafanaObservation struct {
	// URL where Grafana can be accessed.
	URL                      string `json:"url,omitempty"`
	meta.ChildResourceStatus `json:",inline"`
}

// Loki deploys a fully managed Loki instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Loki struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LokiSpec   `json:"spec"`
	Status            LokiStatus `json:"status,omitempty"`
}

// LokiList contains a list of Loki
// +kubebuilder:object:root=true
type LokiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Loki `json:"items"`
}

// An LokiSpec defines the desired state of a Loki.
type LokiSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            LokiParameters `json:"forProvider"`
}

// LokiParameters are the configurable fields of a Loki.
type LokiParameters struct {
	// Retention is the amount of time the logs will be kept on the backend
	// storage of Loki. Defaults to 30 days.
	// +kubebuilder:default:="720h"
	Retention *metav1.Duration `json:"retention,omitempty"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to Loki.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	// In addition to your IPs, we add all operational necessary IPs as well.
	//
	// +listType:="set"
	// +kubebuilder:default:={"0.0.0.0/0"}
	// +optional
	AllowedCIDRs []meta.IPv4CIDR `json:"allowedCIDRs"`
}

// An LokiStatus represents the observed state of a Loki.
type LokiStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               LokiObservation `json:"atProvider"`
}

// LokiObservation are the observable fields of a Loki.
type LokiObservation struct {
	meta.ChildResourceStatus `json:",inline"`
}

// MetricsAgent deploys a Metrics collection agent.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=ma
// +kubebuilder:object:root=true
type MetricsAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetricsAgentSpec   `json:"spec"`
	Status            MetricsAgentStatus `json:"status,omitempty"`
}

// MetricsAgentList contains a list of MetricsAgent.
// +kubebuilder:object:root=true
type MetricsAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricsAgent `json:"items"`
}

// MetricsAgentSpec defines the desired state of a MetricsAgent.
type MetricsAgentSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            MetricsAgentParameters `json:"forProvider"`
}

// MetricsAgentParameters are the configurable fields of a MetricsAgent.
type MetricsAgentParameters struct {
	// Cluster specifies on which cluster agent should be started on.
	Cluster meta.LocalReference `json:"cluster"`
	// Replicas sets the amount of agent replicas to start.
	// +kubebuilder:default:=2
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// EnableDefaultMetrics specifies if this MetricsAgent will scrape
	// default metrics.
	// +optional
	// +kubebuilder:default:=true
	EnableDefaultMetrics *bool `json:"enableDefaultMetrics"`
	// ExternalLabels are labels which are attached to every scraped
	// metric.
	// +optional
	ExternalLabels map[string]string `json:"externalLabels,omitempty"`
	// ScrapeInterval sets the default scrape interval.
	// +optional
	// +kubebuilder:default:="30s"
	// TODO: set reasonable minimum or even move it to cfg?
	ScrapeInterval *metav1.Duration `json:"scrapeInterval,omitempty"`
	// Alertmanagers MetricsAgent should send alerts to.
	// +optional
	Alertmanagers []meta.Reference `json:"alertmanagers,omitempty"`
}

// MetricsAgentStatus represents the observed state of a MetricsAgent.
type MetricsAgentStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               MetricsAgentObservation `json:"atProvider"`
}

// MetricsAgentObservation are the observable fields of a MetricsAgent.
type MetricsAgentObservation struct {
	meta.ReferenceStatus `json:",inline"`
	// ExternalURL provides the URL to access this MetricsAgent instance from
	// external of the cluster. The corresponding basic auth credentials
	// are exposed in the referenced connection secret
	// (`spec.resourceSpec.writeConnectionSecretToRef`). Access can be
	// disabled by using `spec.forProvider.access.noExternalAccess`.
	ExternalURL string `json:"externalURL,omitempty"`
	// InternalURL provides the URL to access the MetricsAgent instance from
	// internal of the referenced cluster. Needs to be enabled via
	// `spec.forProvider.access.internal.enabled`.
	InternalURL string `json:"internalURL,omitempty"`
}

// Prometheus deploys a fully managed Prometheus server.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=prom
// +kubebuilder:object:root=true
type Prometheus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrometheusSpec   `json:"spec"`
	Status            PrometheusStatus `json:"status,omitempty"`
}

// PrometheusList contains a list of Prometheus.
// +kubebuilder:object:root=true
type PrometheusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Prometheus `json:"items"`
}

// PrometheusSpec defines the desired state of a Prometheus.
type PrometheusSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            PrometheusParameters `json:"forProvider"`
}

// PrometheusParameters are the configurable fields of a Prometheus.
type PrometheusParameters struct {
	// Cluster specifies on which cluster prometheus should be started on.
	Cluster meta.LocalReference `json:"cluster"`
	// DiskSpace to request for storing metrics.
	// +kubebuilder:default:="50Gi"
	// +optional
	DiskSpace *resource.Quantity `json:"diskSpace,omitempty"`
	// Replicas sets the amount of prometheus replicas to start.
	// Using more than 1 replica results in the usage of promxy
	// in front of the prometheus instances.
	// +kubebuilder:default:=1
	// +optional
	Replicas int `json:"replicas"`
	// EnableDefaultMetrics specifies if this Prometheus will scrape
	// default metrics.
	// +optional
	// +kubebuilder:default:=true
	EnableDefaultMetrics *bool `json:"enableDefaultMetrics"`
	// ExternalLabels are labels which are attached to every scraped
	// metric.
	// +optional
	ExternalLabels map[string]string `json:"externalLabels,omitempty"`
	// RetentionTime is the amount of time we store metrics.
	// +optional
	// +kubebuilder:default:="720h"
	RetentionTime *metav1.Duration `json:"retentionTime,omitempty"`
	// ScrapeInterval sets the default scrape interval.
	// +optional
	// +kubebuilder:default:="30s"
	ScrapeInterval *metav1.Duration `json:"scrapeInterval,omitempty"`
	// Alertmanagers Prometheus should send alerts to.
	// +optional
	Alertmanagers []meta.Reference `json:"alertmanagers,omitempty"`
	// Access configures the access to the Prometheus instance.
	// +optional
	Access PrometheusAccess `json:"access"`
}
type PrometheusAccess struct {
	// Internal selects from which Pods in the cluster this Promtheus
	// instance can be accessed.
	// +optional
	Internal PrometheusInternalAccess `json:"internal"`
	// NoExternalAccess removes the possibility to access this
	// Prometheus instance from outside of the cluster. The Prometheus
	// instance can not be used in Nine Grafana instances anymore. The
	// Prometheus instance can still be accessed from Pods running in the
	// cluster, if enabled.
	// +optional
	NoExternalAccess bool `json:"noExternalAccess"`
}
type PrometheusInternalAccess struct {
	// Enabled can be used to allow selected pods (see below) to connect to
	// the Prometheus instance by using the internal Service URL found in
	// the connection secret and status of the Prometheus resource.
	// If `Enabled` is set to false, no internal access will be possible,
	// no matter what was set in the `podSelector` or `namespaceSelector`
	// fields.
	// +optional
	Enabled bool `json:"enabled"`
	// PodSelector is a label selector which selects pods permitted to
	// connect to Prometheus. If empty, all pods are allowed to connect to
	// Prometheus.
	//
	// If NamespaceSelector is also set with some value, then only pods
	// from within the namespaces selected by it are permitted.
	// +optional
	PodSelector metav1.LabelSelector `json:"podSelector"`
	// NamespaceSelector selects the namespaces from which the
	// `PodSelector` field will select. If left empty, all namespaces will
	// be considered.
	// +optional
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector"`
}

// PrometheusStatus represents the observed state of a Prometheus.
type PrometheusStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               PrometheusObservation `json:"atProvider"`
}

// PrometheusObservation are the observable fields of a Prometheus.
type PrometheusObservation struct {
	meta.ReferenceStatus `json:",inline"`
	// ExternalURL provides the URL to access this Prometheus instance from
	// external of the cluster. The corresponding basic auth credentials
	// are exposed in the referenced connection secret
	// (`spec.resourceSpec.writeConnectionSecretToRef`). Access can be
	// disabled by using `spec.forProvider.access.noExternalAccess`.
	ExternalURL string `json:"externalURL,omitempty"`
	// InternalURL provides the URL to access the Prometheus instance from
	// internal of the referenced cluster. Needs to be enabled via
	// `spec.forProvider.access.internal.enabled`.
	InternalURL string `json:"internalURL,omitempty"`
}

// Promtail deploys Promtail to a cluster and pushes logs to the configured Loki instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Promtail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PromtailSpec   `json:"spec"`
	Status            PromtailStatus `json:"status,omitempty"`
}

// PromtailList contains a list of Promtails.
// +kubebuilder:object:root=true
type PromtailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Promtail `json:"items"`
}

// An PromtailSpec defines the desired state of a Promtail.
type PromtailSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            PromtailParameters `json:"forProvider"`
}

// PromtailParameters are the configurable fields of a Promtail.
type PromtailParameters struct {
	// Cluster specifies on which cluster Promtail should be installed to.
	Cluster meta.LocalReference `json:"cluster"`
	// The Loki instance where promtail should push logs to.
	Loki meta.Reference `json:"loki"`
	// ExternalLabels are static labels to add to all logs being sent to Loki.
	// +optional
	ExternalLabels model.LabelSet `json:"externalLabels,omitempty"`
}

// An PromtailStatus represents the observed state of a Promtail.
type PromtailStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               PromtailObservation `json:"atProvider"`
}

// PromtailObservation are the observable fields of a Promtail.
type PromtailObservation struct {
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}
