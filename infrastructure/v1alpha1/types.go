package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	backup "github.com/ninech/apis/backup/v1alpha1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineTypeNineStandard1 is a VM running on Nine Infrastructure with:
	// 1 CPU Cores
	// 4GB RAM
	MachineTypeNineStandard1 MachineType = "nine-standard-1"
	// MachineTypeNineStandard2 is a VM running on Nine Infrastructure with:
	// 2 CPU Cores
	// 8GB RAM
	MachineTypeNineStandard2 MachineType = "nine-standard-2"
	// MachineTypeNineStandard4 is a VM running on Nine Infrastructure with:
	// 4 CPU Cores
	// 16GB RAM
	MachineTypeNineStandard4 MachineType = "nine-standard-4"
	// MachineTypeHighMem2 is a VM running on Nine Infrastructure with:
	// 2 CPU Cores
	// 16GB RAM
	MachineTypeNineHighMem2 MachineType = "nine-highmem-2"
	// MachineTypeHighMem4 is a VM running on Nine Infrastructure with:
	// 4 CPU Cores
	// 32GB RAM
	MachineTypeNineHighMem4 MachineType = "nine-highmem-4"
	// MachineTypeHighCPU2 is a VM running on Nine Infrastructure with:
	// 2 CPU Cores
	// 4GB RAM
	MachineTypeNineHighCPU2 MachineType = "nine-highcpu-2"
	// MachineTypeHighCPU4 is a VM running on Nine Infrastructure with:
	// 4 CPU Cores
	// 8GB RAM
	MachineTypeNineHighCPU4 MachineType = "nine-highcpu-4"
	// MachineTypeHighCPU8 is a VM running on Nine Infrastructure with:
	// 8 CPU Cores
	// 16GB RAM
	MachineTypeNineHighCPU8 MachineType = "nine-highcpu-8"
	// MachineTypeNineDBXS is a VM running on Nine Infrastructure with:
	// 2 CPU Cores
	// 4GB RAM
	MachineTypeNineDBXS MachineType = "nine-db-xs"
	// MachineTypeNineDBS is a VM running on Nine Infrastructure with:
	// 4 CPU Cores
	// 8GB RAM
	MachineTypeNineDBS MachineType = "nine-db-s"
	// MachineTypeNineDBM is a VM running on Nine Infrastructure with:
	// 4 CPU Cores
	// 12GB RAM
	MachineTypeNineDBM MachineType = "nine-db-m"
	// MachineTypeNineDBL is a VM running on Nine Infrastructure with:
	// 6 CPU Cores
	// 16GB RAM
	MachineTypeNineDBL MachineType = "nine-db-l"
	// MachineTypeNineDBXL is a VM running on Nine Infrastructure with:
	// 8 CPU Cores
	// 24GB RAM
	MachineTypeNineDBXL MachineType = "nine-db-xl"
	// MachineTypeNineDBXXL is a VM running on Nine Infrastructure with:
	// 10 CPU Cores
	// 32GB RAM
	MachineTypeNineDBXXL MachineType     = "nine-db-xxl"
	Rocky9               OperatingSystem = "rocky9"
	// Ubuntu LTS
	// http://releases.ubuntu.com/
	Ubuntu24_04 OperatingSystem = "ubuntu24.04"
	Ubuntu22_04 OperatingSystem = "ubuntu22.04"
	Ubuntu20_04 OperatingSystem = "ubuntu20.04"
	Ubuntu18_04 OperatingSystem = "ubuntu18.04"
)

var (
	// CloudVirtualMachineOperatingSystems lists all cloud VM operating systems.
	CloudVirtualMachineOperatingSystems = []CloudVirtualMachineOS{CloudVirtualMachineOS(Ubuntu20_04), CloudVirtualMachineOS(Ubuntu22_04), CloudVirtualMachineOS(Ubuntu24_04), CloudVirtualMachineOS(Rocky9)}
	// MachineTypes is a list of all machine types.
	MachineTypes = []MachineType{MachineTypeNineStandard1, MachineTypeNineStandard2, MachineTypeNineStandard4, MachineTypeNineHighMem2, MachineTypeNineHighMem4, MachineTypeNineHighCPU2, MachineTypeNineHighCPU4, MachineTypeNineHighCPU8, MachineTypeNineDBXS, MachineTypeNineDBS, MachineTypeNineDBM, MachineTypeNineDBL, MachineTypeNineDBXL, MachineTypeNineDBXXL}
	// MachineTypesDB is a list of all database machine types.
	MachineTypesDB = []MachineType{MachineTypeNineDBXS, MachineTypeNineDBS, MachineTypeNineDBM, MachineTypeNineDBL, MachineTypeNineDBXL, MachineTypeNineDBXXL}
)

// CloudVirtualMachine is a virtual machine instance providing flexible scaling and a
// variety of Linux distributions.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="POWERSTATE",type="string",JSONPath=".status.atProvider.powerState"
// +kubebuilder:printcolumn:name="IP",type="string",JSONPath=".status.atProvider.ipAddress"
// +kubebuilder:printcolumn:name="FQDN",type="string",JSONPath=".status.atProvider.fqdn"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName=cloudvm
// +kubebuilder:object:root=true
type CloudVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudVirtualMachineSpec   `json:"spec"`
	Status            CloudVirtualMachineStatus `json:"status,omitempty"`
}

// CloudVirtualMachineList contains a list of CloudVirtualMachines
// +kubebuilder:object:root=true
type CloudVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudVirtualMachine `json:"items"`
}

// CloudVirtualMachineSpec defines the desired state of a cloud VM.
type CloudVirtualMachineSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            CloudVirtualMachineParameters `json:"forProvider"`
}

// CloudVirtualMachineParameters are the configurable fields of a CloudVirtualMachine.
type CloudVirtualMachineParameters struct {
	// MachineType defines the sizing for a particular cloud vm.
	// +optional
	// +kubebuilder:default:=nine-standard-1
	MachineType MachineType `json:"machineType,omitempty"`
	// Location specifies in which datacenter the VM will be spawned.
	// Needs to match the available MachineTypes in that datacenter.
	Location meta.LocationName `json:"location"`
	// Hostname allows to set the hostname explicitly. If unset, the name
	// of the resource will be used as the hostname. This does not affect
	// the DNS name.
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Hostname is immutable after creation"
	Hostname string `json:"hostname,omitempty"`
	// OS which should be used to boot the VM.
	// +optional
	// +kubebuilder:default:=ubuntu24.04
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="OS is immutable after creation"
	OS CloudVirtualMachineOS `json:"os,omitempty"`
	// BootDisk that will be used to boot the VM from.
	// +optional
	// +kubebuilder:default:={name:"root",size:"20Gi"}
	BootDisk *Disk `json:"bootDisk,omitempty"`
	// Disks specifies which additional disks to mount to the machine.
	// +listType:="map"
	// +listMapKey:="name"
	// +optional
	Disks []Disk `json:"disks,omitempty"`
	// PowerState specifies the power state of the cloud VM. A value of
	// "On" turns the VM on, shutdown sends an ACPI signal to the VM to
	// perform a clean shutdown and off forces the power off immediately.
	// +optional
	// +kubebuilder:default:="on"
	// +kubebuilder:validation:Enum=on;shutdown;off
	PowerState VirtualMachinePowerState `json:"powerState,omitempty"`
	// PublicKeys specifies the SSH Public Keys that can be used to connect to
	// the VM as root. The keys are expected to be in SSH format as defined in
	// RFC4253.
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Public Key is immutable after creation"
	PublicKeys []string `json:"publicKeys,omitempty"`
	// CloudConfig allows to pass custom cloud config data (https://cloudinit.readthedocs.io/en/latest/topics/format.html#cloud-config-data)
	// to the cloud VM. If a CloudConfig is passed, the PublicKey parameter is ignored.
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Cloud Config is immutable after creation"
	CloudConfig string `json:"cloudConfig,omitempty"`
	// Rescue configures booting into a rescue live-OS for fixing a VM that is
	// in an unbootable state.
	Rescue *CloudVirtualMachineRescue `json:"rescue,omitempty"`
}

// MachineType is a name for a particular machine sizing.
// +nine:public:definition
// +kubebuilder:validation:Enum=nine-standard-1;nine-standard-2;nine-standard-4;nine-highmem-2;nine-highmem-4;nine-highcpu-2;nine-highcpu-4;nine-highcpu-8;nine-db-xs;nine-db-s;nine-db-m;nine-db-l;nine-db-xl;nine-db-xxl
type MachineType string

// CloudVirtualMachineOS is an operating system for a cloud VM.
// +kubebuilder:validation:Enum=ubuntu20.04;ubuntu22.04;ubuntu24.04;rocky9
// +nine:public:definition
type CloudVirtualMachineOS OperatingSystem

// Disk describes a Disk that can be attached to a VM.
type Disk struct {
	// Name specifies the name of the disk. Used to identify a disk, changing
	// the name of a disk means the old disk will be deleted and a new one
	// will be created.
	Name string `json:"name"`
	// Size specifies the disk size.
	Size resource.Quantity `json:"size"`
}
type VirtualMachinePowerState string
type CloudVirtualMachineRescue struct {
	// Enable enables booting into rescue. This will trigger an immediate
	// reboot of the VM into a rescue live-OS. Set this to false to configure
	// boot from the root disk again.
	Enabled bool `json:"enabled"`
	// PublicKeys specifies additional SSH Public Keys that can be used to
	// connect to the rescue OS as root. The keys are expected to be in SSH
	// format as defined in RFC4253. If not specified, just the PublicKeys
	// from the parameters will be used.
	PublicKeys []string `json:"publicKeys,omitempty"`
}

// CloudVirtualMachineStatus represents the observed state of a cloud VM.
type CloudVirtualMachineStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               CloudVirtualMachineObservation `json:"atProvider"`
}

// CloudVirtualMachineObservation are the observable fields of a cloud VM.
type CloudVirtualMachineObservation struct {
	// IPAddress is the public IPAddress for the VM.
	IPAddress string `json:"ipAddress,omitempty"`
	// PowerState indicates the observed power state of the VM.
	PowerState VirtualMachinePowerState `json:"powerState,omitempty"`
	// FQDN is the fully qualified domain name at which the VM is reachable at.
	FQDN string `json:"fqdn,omitempty"`
	// UUID of the underlying virtual machine.
	UUID string `json:"uuid,omitempty"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// ClusterData provides cluster information of the referenced KubernetesCluster resource.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type ClusterData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterDataSpec   `json:"spec"`
	Status            ClusterDataStatus `json:"status,omitempty"`
}

// ClusterDataList contains a list of ClusterData resources.
// +kubebuilder:object:root=true
type ClusterDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterData `json:"items"`
}

// ClusterDataSpec defines the desired state of ClusterData resource.
type ClusterDataSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ClusterDataParameters `json:"forProvider"`
}

// ClusterDataParameters are the configurable fields of a ClusterData resource.
type ClusterDataParameters struct {
	// ClusterReference selects the KubernetesCluster of which the cluster
	// data should be exposed
	ClusterReference meta.Reference `json:"clusterReference"`
}

// ClusterDataStatus represents the observed state of a ClusterData resource.
type ClusterDataStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ClusterDataObservation `json:"atProvider"`
}

// ClusterDataObservation are the observable fields of a ClusterData resource.
type ClusterDataObservation struct {
	// APIEndpoint is the URL under which the Kubernetes API is reachable at.
	// +optional
	APIEndpoint string `json:"apiEndpoint,omitempty"`
	// APICACert is the base64 encoded ca certificate of the kube-apiserver
	// +optional
	APICACert string `json:"apiCACert,omitempty"`
}

// Keda deploys Keda to a KubernetesCluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,path=kedas
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

// KubernetesCluster is a fully managed Kubernetes cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="API_READY",type="string",JSONPath=".status.atProvider.apiReady"
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
type NKEClusterSettings struct {
	// StaticEgress defines settings for the static egress feature
	// +optional
	StaticEgress StaticEgress `json:"staticEgress"`
	// AuditLog configures audit logging.
	// +optional
	AuditLog AuditLogConfiguration `json:"auditLog"`
}
type StaticEgress struct {
	// Enabled defines if the static egress feature should be enabled or
	// disabled
	// +optional
	// +kubebuilder:default:=false
	Enabled bool `json:"enabled"`
}
type AuditLogConfiguration struct {
	// Targets to send the audit log to. The only supported target is Loki.
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=0
	Targets []AuditLogTarget `json:"targets,omitempty"`
}
type AuditLogTarget struct {
	meta.LocalTypedReference `json:",inline"`
}

// VClusterSettings defines additional fields that a nine KubernetesCluster
// based on VCluster can have.
type VClusterSettings struct {
	// Version specifies the Kubernetes version that will be used for this
	// cluster, e.g. "1.26". The patch version cannot be specified and the
	// latest supported one will be used.
	// +optional
	// +kubebuilder:default:="1.28"
	// +kubebuilder:validation:Enum:="1.25";"1.26";"1.27";"1.28";"1.29"
	// +kubebuilder:validation:XValidation:message="downgrade is not allowed",rule="double(self) >= double(oldSelf)"
	// +kubebuilder:validation:XValidation:message="only one minor upgrade is allowed",rule="double(self) - double(oldSelf) < 0.02"
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
	MachineType MachineType `json:"machineType"`
	// DiskSize specifies the size of the disk for the nodes in this pool.
	// Changing this results in a new rollout of all nodes in the pool.
	// Allowed range is 20Gi - 100Gi.
	// +kubebuilder:default:="20Gi"
	// +optional
	DiskSize *resource.Quantity `json:"diskSize,omitempty"`
}
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
	// KubernetesVersion is the version of Kubernetes that this cluster is running.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	// APIReady indicates if the API is ready for consumption.
	// +optional
	APIReady bool `json:"apiReady"`
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
	// KubernetesVersion is the version of Kubernetes that this cluster is running.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	// NodePools lists the name of the node pools plus their associated status.
	NodePools map[string]NodePoolStatus `json:"nodePools,omitempty"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}
type NodePoolStatus struct {
	// NumNodes describes the current number of nodes in the node pool.
	NumNodes int `json:"numNodes"`
	// MachineType shows the current machine type of the node pool.
	MachineType *MachineType `json:"machineType,omitempty"`
	// DiskSize shows the current disk size of the node pool.
	DiskSize *resource.Quantity `json:"diskSize,omitempty"`
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

// OperatingSystem is an Operating System for a VM.
type OperatingSystem string
