package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// LocationNineCZ41 is the name for our CZ41 location
	LocationNineCZ41 LocationName = NinePrefix + "cz41"
	// LocationNineCZ42 is the name for our CZ42 location
	LocationNineCZ42 LocationName = NinePrefix + "cz42"
	// LocationNineES34 is the name for our ES34 location
	LocationNineES34 LocationName = NinePrefix + "es34"
	// ClusterDataDeploioName is the name of the ClusterData resource which
	// exposes connection details of the deplo.io cluster to customers
	ClusterDataDeploioName = "deploio"
	// NinePrefix contains a prefix which all nine owned
	// resources begin their name with
	NinePrefix = "nine-"
	// NineOwnedLabelKey specifies the key for the label which
	// identifies nine resources.
	NineOwnedLabelKey = "nine.ch/owned-by"
	// NineOwnedLabelValue specifies the value for the label which
	// identifies nine resources.
	NineOwnedLabelValue = "nine"
)

// DNSCheckType describes the type of a DNS check.
// +kubebuilder:object:generate=true
type DNSCheckType string

// DNSVerificationStatusEntries is a list of DNSVerificationStatusEntry with added functions.
// +kubebuilder:object:generate=true
type DNSVerificationStatusEntries []DNSVerificationStatusEntry

// DNSVerificationStatus reflects the DNS verification status.
// +kubebuilder:object:generate=true
type DNSVerificationStatus struct {
	// CNAMETarget shows to which target the CNAME entry should point to
	// for the DNS CNAME verification method to succeed.
	// +optional
	CNAMETarget string `json:"cnameTarget,omitempty"`
	// TXTRecordValue shows which TXT DNS record value needs to be created
	// for the DNS TXT verification method to succeed.
	// +optional
	TXTRecordValue string `json:"txtRecordValue,omitempty"`
	// StatusEntries show the status of the DNS verification methods.
	// +optional
	StatusEntries DNSVerificationStatusEntries `json:"statusEntries,omitempty"`
}

// DNSVerificationStatusEntry is a single entry used in the status of the DNSVerification.
// +kubebuilder:object:generate=true
type DNSVerificationStatusEntry struct {
	// the hostname of the verification entry
	Name string `json:"name"`
	// CheckType describes which kind of DNS check this entry is about.
	// +optional
	CheckType DNSCheckType `json:"checkType,omitempty"`
	// LatestSuccess specifies when this host was last verified
	// successfully.
	// +optional
	LatestSuccess *metav1.Time `json:"latestSuccess,omitempty"`
	// Error displays a potential error which happened during the
	// verification.
	// +optional
	Error *VerificationError `json:"error,omitempty"`
}

// VerificationError describes a verification error.
// +kubebuilder:object:generate=true
type VerificationError struct {
	// Message refers to the error message.
	Message string `json:"message"`
	// Timestamp refers to the time when this error happened.
	Timestamp metav1.Time `json:"timestamp"`
}

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

// IPv4CIDR represents a IPv4 address in CIDR notation
// +kubebuilder:validation:Pattern=`\A([0-9]{1,3}\.){3}[0-9]{1,3}\/([0-9]|[1-2][0-9]|3[0-2])\z`
type IPv4CIDR string

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
type TypedReference struct {
	// Reference to the resource
	Reference `json:",inline"`
	// Type info about the resource.
	metav1.GroupKind `json:",inline"`
}
type LocalTypedReference struct {
	// Reference to the resource
	LocalReference `json:",inline"`
	// Type info about the resource.
	metav1.GroupKind `json:",inline"`
}

// URL represents a URI reference.
// The general form represented is:
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//
// +kubebuilder:validation:Pattern=`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
type URL string
