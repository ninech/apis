package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	backup "github.com/ninech/apis/backup/v1alpha1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Keda deploys Keda to a KubernetesCluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Keda struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KedaSpec   `json:"spec"`
	Status            KedaStatus `json:"status,omitempty"`
}

// KedaList contains a list of Keda instances
// +kubebuilder:object:root=true
type KedaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keda `json:"items"`
}

// A KedaSpec defines the desired state of a Keda instance.
type KedaSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            KedaParameters `json:"forProvider"`
}

// KedaParameters are the configurable fields of a Keda instance.
type KedaParameters struct {
	// Cluster is the cluster where the keda should be deployed
	// to.
	Cluster meta.LocalReference `json:"cluster"`
}

// A KedaStatus represents the observed state of a Keda instance.
type KedaStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               KedaObservation `json:"atProvider"`
}

// KedaObservation are the observable fields of a Keda instance.
type KedaObservation struct {
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}

// A KubernetesCluster is a fully managed Kubernetes cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=kc
// +kubebuilder:object:root=true
type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

// KubernetesClusterList contains a list of KubernetesClusters.
// +kubebuilder:object:root=true
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesCluster `json:"items"`
}

// A KubernetesClusterSpec defines the desired state of a KubernetesCluster.
type KubernetesClusterSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            KubernetesClusterParameters `json:"forProvider"`
}

// KubernetesClusterParameters are the configurable fields of a KubernetesCluster.
type KubernetesClusterParameters struct {
	// Location of the KubernetesCluster. Note that Clusters are currently
	// only available in the location nine-es34.
	Location meta.LocationName `json:"location"`
	// NKE represents a KubernetesCluster in Nine's datacentres.
	// +optional
	NKE *NKEClusterSettings `json:"nke,omitempty"`
	// VCluster is a virtual KubernetesCluster running on top of NKE.
	// Experimental and should only be used for development and testing.
	// +optional
	VCluster *VClusterSettings `json:"vcluster,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	NodePools []NodePool `json:"nodePools"`
	// AdditionalBackupSchedules allows custom backup schedules to be setup.
	// The daily full cluster backup won't be affected by this.
	// +kubebuilder:validation:MaxItems:=3
	// +optional
	AdditionalBackupSchedules []backup.VeleroSchedule `json:"additionalBackupSchedules,omitempty"`
	// ScrapeConfigurations allows to overwrite which metrics of this cluster are scraped
	// by certain Prometheus instances
	// +optional
	ScrapeConfigurations *KubernetesClusterScrapeConfiguration `json:"scrapeConfiguration,omitempty"`
}

// NKEClusterSettings defines additional fields that a nine KubernetesCluster
// can have.
type NKEClusterSettings struct{}

// VClusterSettings defines additional fields that a nine KubernetesCluster
// based on VCluster can have.
type VClusterSettings struct {
	// Version specifies the Kubernetes version that will be used for this
	// cluster, e.g. "1.24". The patch version cannot be specified and the
	// latest supported one will be used.
	// +optional
	// +kubebuilder:default:="1.24"
	// +kubebuilder:validation:Enum:="1.21";"1.22";"1.23";"1.24"
	Version string `json:"version,omitempty"`
	// CertManager enables cert-manager on the vcluster. Currently just
	// Certificate resources are supported.
	// +optional
	CertManager bool `json:"certManager,omitempty"`
}

// NodePool configures a pool of nodes which are added to the cluster.
type NodePool struct {
	// Name of the node pool. Changing this results in a new rollout of all
	// nodes in the pool.
	Name string `json:"name"`
	// MinNodes describes the lower bound of nodes in this pool. If MinNodes
	// == MaxNodes, autoscaling is disabled.
	// +kubebuilder:validation:Minimum=0
	MinNodes int `json:"minNodes"`
	// MinNodes describes the upper bound of nodes in this pool. If MinNodes
	// == MaxNodes, autoscaling is disabled.
	// +kubebuilder:validation:Minimum=1
	MaxNodes int `json:"maxNodes"`
	// Labels specifies the node labels. Changing this results in a new
	// rollout of all nodes in the pool.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations specifies the node annotations. Changing this results in a
	// new rollout of all nodes in the pool.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Taints specifies the node taints. Changing this results in a new
	// rollout of all nodes in the pool.
	// +optional
	Taints []v1.Taint `json:"taints,omitempty"`
	// MachineType identifies the machine sizing. Changing this results in a
	// new rollout of all nodes in the pool.
	// +kubebuilder:validation:Enum=nine-standard-1;nine-standard-2;nine-standard-4;nine-highmem-2;nine-highmem-4
	MachineType MachineType `json:"machineType"`
	// DiskSize specifies the size of the disk for the nodes in this pool.
	// Changing this results in a new rollout of all nodes in the pool.
	// Allowed range is 20Gi - 100Gi.
	// +kubebuilder:default:="20Gi"
	// +optional
	DiskSize *resource.Quantity `json:"diskSize,omitempty"`
}

// MachineType is a name for a particular machine sizing.
type MachineType string
type KubernetesClusterScrapeConfiguration struct {
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	NodeExporter []meta.ScrapeConfig `json:"nodeExporter,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	KubeStateMetrics []meta.ScrapeConfig `json:"kubeStateMetrics,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	KubeletCAdvisor []meta.ScrapeConfig `json:"kubeletCAdvisor,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	Kubelet []meta.ScrapeConfig `json:"kubelet,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	Velero []meta.ScrapeConfig `json:"velero,omitempty"`
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	CertManager []meta.ScrapeConfig `json:"certManager,omitempty"`
}

// A KubernetesClusterStatus represents the observed state of a KubernetesCluster.
type KubernetesClusterStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               KubernetesClusterObservation `json:"atProvider,omitempty"`
}

// KubernetesClusterObservation are the observable fields of a KubernetesCluster.
type KubernetesClusterObservation struct {
	ClusterObservation `json:",inline"`
	// KubernetesVersion represents the version of Kubernetes that this cluster
	// is running at.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	// VCluster exposes vcluster specific status fields.
	// +optional
	VCluster *VClusterSpecificStatus `json:"vcluster,omitempty"`
}

// ClusterObservation are the observable fields of a Cluster.
type ClusterObservation struct {
	// APIEndpoint is the URL under which the Kubernetes API is reachable at.
	APIEndpoint string `json:"apiEndpoint,omitempty"`
	// APICACert is the base64 encoded ca certificate of the kube-apiserver
	APICACert string `json:"apiCACert,omitempty"`
	// OIDCClientID is the client ID for the OIDC login flow to this cluster.
	OIDCClientID string `json:"oidcClientID,omitempty"`
	// OIDCIssuerURL is the issuer URL for the OIDC login flow to this cluster.
	OIDCIssuerURL string `json:"oidcIssuerURL,omitempty"`
	// NodePools lists the name of the node pools plus their associated status.
	NodePools map[string]NodePoolStatus `json:"nodePools,omitempty"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}
type NodePoolStatus struct {
	// NumNodes describes the current number of nodes in the node pool.
	NumNodes int `json:"numNodes"`
}
type VClusterSpecificStatus struct {
	// DefaultIngress that Ingress objects within the vcluster can use.
	DefaultIngress VClusterIngress `json:"defaultIngress"`
}
type VClusterIngress struct {
	// Host is the fully qualified hostname that points to the Ingress
	// Loadbalancer.
	Host string `json:"host"`
	// Class is the name of the IngressClass that can be referenced within an
	// Ingress resource.
	Class string `json:"class"`
}
