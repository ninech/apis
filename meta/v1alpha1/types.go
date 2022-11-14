package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ChildResourceStatus reflects observed errors from child resources
// of a managed resource.
// +kubebuilder:object:generate=true
type ChildResourceStatus struct {
	// ChildResourceErrors of a managed resource.
	// +optional
	ChildResourceErrors []ChildResourceError `json:"childResourceErrors,omitempty"`
}

// ChildResourceError is an error that occurred on a child resource of
// a managed resource.
type ChildResourceError struct {
	// Resource references the child resource that errored.
	Resource Reference `json:"resource"`
	// Type specifies the type of the resource.
	Type metav1.GroupVersionKind `json:"type"`
	// Message that describes why the resource failed.
	Message string `json:"message"`
}

// Reference references another object in a specific namespace.
type Reference struct {
	// Name of the target.
	Name string `json:"name"`
	// Namespace of the target.
	Namespace string `json:"namespace"`
}

// Image specifies general container image options which describe
// from where and in which version a container image should be used
type Image struct {
	// Registry specifies the registry from where the image should
	// be pulled
	Registry string `json:"registry,omitempty"`
	// Repository specifies the repository from where the image should
	// be pulled
	Repository string `json:"repository,omitempty"`
	// Tag specifies the image tag to use
	// +optional
	Tag string `json:"tag,omitempty"`
	// Digest specifies the image digest to use
	// +optional
	Digest string `json:"digest,omitempty"`
	// PullPolicy specifies the image pull policy to use
	// +optional
	// +kubebuilder:default:="Always"
	PullPolicy string `json:"pullPolicy,omitempty"`
	// PullSecret specifies a image pull secret name
	// +optional
	PullSecret string `json:"pullSecret,omitempty"`
}

// LocationName specifies the physical location of resources.
// Not all locations implement the same functionality.
// +kubebuilder:validation:Enum:=nine-cz41;nine-cz42;nine-es34
type LocationName string

// ScrapeConfig configures a metrics scrape config.
// +kubebuilder:object:generate:=true
type ScrapeConfig struct {
	// Name uniquely identifies an ExporterScrapeConfig.
	Name string `json:"name"`
	// Enabled specifies if metrics of a corresponding
	// exporter should be scraped by certain Prometheus
	// instances
	// +optional
	// +kubebuilder:default:=true
	Enabled *bool `json:"enabled"`
	// AdditionalMetrics specifies which additional
	// metrics should be scraped from this exporter
	// +optional
	AdditionalMetrics []string `json:"additionalMetrics"`
	// AllMetrics specifies that all metrics of this
	// specific component should be scraped
	// +optional
	// +kubebuilder:default:=false
	AllMetrics bool `json:"allMetrics,omitempty"`
	// ScrapedBy defines which prometheus instances will
	// target this scrape config.
	ScrapedBy []LocalReference `json:"scrapedBy,omitempty"`
}

// LocalReference references another object in the same namespace.
type LocalReference struct {
	// Name of the target.
	Name string `json:"name"`
}

// +kubebuilder:object:generate=true
type ReferenceStatus struct {
	ReferenceErrors ReferenceErrors `json:"referenceErrors,omitempty"`
}
type ReferenceErrors []ReferenceError
type ReferenceError struct {
	Reference `json:",inline"`
	Kind      string `json:"kind"`
	Message   string `json:"message"`
}
type LocalReferenceSelector struct {
	LocalReference `json:",inline"`
	// Key describes the key within the secret.
	Key string `json:"key"`
}
