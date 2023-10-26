package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	infrav1alpha1 "github.com/ninech/apis/infrastructure/v1alpha1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Bucket is an object storage bucket. It's used to group objects, defines
// who can access them and how they are stored.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// BucketList contains a list of Bucket
// +kubebuilder:object:root=true
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// A BucketSpec defines the desired state of a Bucket.
type BucketSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            BucketParameters `json:"forProvider"`
}

// BucketParameters are the configurable fields of a Bucket.
type BucketParameters struct {
	// Location specifies the physical location of the Bucket.
	// +immutable
	Location meta.LocationName `json:"location"`
	// StorageType defines the type of the backing storage for the Bucket.
	// +immutable
	// +kubebuilder:default:="standard"
	StorageType BucketStorageType `json:"storageTier"`
	// Permissions configures user access to the objects in this Bucket.
	// +optional
	Permissions []*BucketPermission `json:"permissions,omitempty"`
	// PublicRead sets this Buckets objects to be publicly readable.
	// +optional
	PublicRead bool `json:"publicRead,omitempty"`
	// PublicList sets this Buckets objects to be publicly listable.
	// +optional
	PublicList bool `json:"publicList,omitempty"`
	// Encryption enables encryption at rest for this Bucket.
	// +optional
	Encryption bool `json:"encryption,omitempty"`
	// Versioning enables object versioning for this Bucket.
	// +optional
	Versioning bool `json:"versioning,omitempty"`
	// LifecyclePolicies allows to define automatic expiry (deletion) of
	// objects using certain rules.
	// +optional
	LifecyclePolicies []*BucketLifecyclePolicy `json:"lifecyclePolicies,omitempty"`
	// CORS settings for this bucket. CORS is a mechanism to allow code
	// running in a browser to make requests to a domain other than the
	// one from where it originated.
	// +optional
	CORS *CORSConfig `json:"cors,omitempty"`
}

// BucketStorageType defines the backing storage for a Bucket.
// +kubebuilder:validation:Enum:=standard
type BucketStorageType string

// BucketPermission defines a role and users to assign permissions on a
// Bucket.
type BucketPermission struct {
	// Role that this permission will be assigned.
	Role BucketRole `json:"role"`
	// BucketUserRefs references users that will receive this permission.
	BucketUserRefs []*meta.LocalReference `json:"userRefs,omitempty"`
}

// BucketRole defines what kind of access is possible.
// +kubebuilder:validation:Enum=reader;writer
type BucketRole string

// BucketLifecyclePolicy defines a lifecycle policy of bucket objects.
type BucketLifecyclePolicy struct {
	// Prefix can be used to only expire objects with a certain prefix. Do not
	// specify if all objects should be expired.
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// ExpireAfter defines the time after an object will be expired (deleted).
	// Note that this is not minute-accurate and an object might exist for a
	// bit longer than specified until it is cleaned up. Usually it can take
	// around 30 minutes.
	// This field will only be used by Buckets with backend version 'v1'.
	// +optional
	ExpireAfter metav1.Duration `json:"expireAfter,omitempty"`
	// ExpireAfterDays defines the amount of days after an object will be expired (deleted).
	// This field will only be used by Buckets with backend version 'v2'.
	// +optional
	ExpireAfterDays int32 `json:"expireAfterDays,omitempty"`
	// IsLive specifies if this policy applies to live objects. If false, this
	// policy only applies to archived objects, so this is only useful when
	// the bucket has versioning enabled.
	// +kubebuilder:default:=true
	// +optional
	IsLive bool `json:"isLive"`
}

// CORSConfig defines the configuration for Cross-Origin Resource
// Sharing, see https://fetch.spec.whatwg.org/#http-cors-protocol
// for more information.
type CORSConfig struct {
	// Origins specifies the origins that should be allowed for
	// CORS.
	Origins []string `json:"origins"`
	// ResponseHeaders specifies which headers should be allowed
	// for CORS.
	//
	// +optional
	ResponseHeaders []string `json:"responseHeaders,omitempty"`
	// MaxAge is the maximum time for the origin to hold the preflight
	// results, in seconds. Also known as the cache-expiry time, it
	// defines the duration in which the browser is allowed to make
	// requests before it must repeat the preflight request.
	//
	// +optional
	// +kubebuilder:default:=3600
	MaxAge int `json:"maxAge"`
}

// A BucketStatus represents the observed state of a Bucket.
type BucketStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               BucketObservation `json:"atProvider,omitempty"`
}

// BucketObservation are the observable fields of a Bucket.
type BucketObservation struct {
	// API endpoint to use with S3 compatible clients.
	Endpoint string `json:"endpoint"`
	// PublicURL where the bucket content is accessible if set to PublicRead.
	PublicURL string `json:"publicURL"`
	// BytesUsed shows the amount of bytes a bucket is currently using.
	BytesUsed int64 `json:"bytesUsed"`
	// ObjectCount shows the amount of objects a bucket has.
	ObjectCount int64 `json:"objectCount"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// BucketUser defines a user who can access Buckets.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type BucketUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketUserSpec   `json:"spec"`
	Status            BucketUserStatus `json:"status,omitempty"`
}

// BucketUserList contains a list of BucketUser
// +kubebuilder:object:root=true
type BucketUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketUser `json:"items"`
}

// A BucketUserSpec defines the desired state of a BucketUser.
type BucketUserSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            BucketUserParameters `json:"forProvider"`
}

// BucketUserParameters are the configurable fields of a BucketUser.
type BucketUserParameters struct {
	// ResetCredentials if set to true, the credentials of this user will be reset
	// in the next reconciliation loop.
	// +optional
	ResetCredentials *bool `json:"resetCredentials,omitempty"`
	// Location specifies the physical location of the BucketUser.
	Location meta.LocationName `json:"location"`
	// +optional
	// +kubebuilder:default:="v1"
	BackendVersion BucketBackendVersion `json:"backendVersion,omitempty"`
}

// BackendVersion specifies the bucket backend version to use. While the
// APIs work the same, buckets with v1 are only compatible with
// bucketusers also on v1 and the same applies to v2.
// +kubebuilder:validation:Enum=v1;v2
type BucketBackendVersion string

// A BucketUserStatus represents the observed state of a BucketUser.
type BucketUserStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               BucketUserObservation `json:"atProvider,omitempty"`
}

// BucketUserObservation are the observable fields of a Bucket.
type BucketUserObservation struct {
	// ResettingCredentials indicates if a reset is in progress. If it is true,
	// the current credentials can be considered out of date.
	ResettingCredentials bool `json:"resettingCredentials"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// ObjectsBucket defines a Nutanix Objects Bucket.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type ObjectsBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectsBucketSpec   `json:"spec"`
	Status            ObjectsBucketStatus `json:"status,omitempty"`
}

// ObjectsBucketList contains a list of ObjectsBucket
// +kubebuilder:object:root=true
type ObjectsBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectsBucket `json:"items"`
}

// A ObjectsBucketSpec defines the desired state of an ObjectsBucket.
type ObjectsBucketSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ObjectsBucketParameters `json:"forProvider"`
}

// ObjectsBucketParameters are the configurable fields of an ObjectsBucket.
type ObjectsBucketParameters struct {
	// Location specifies the physical location of the ObjectsBucket.
	Location meta.LocationName `json:"location"`
	// Permissions configures user access to the objects in this Bucket.
	// +optional
	Permissions []*BucketPermission `json:"permissions,omitempty"`
	// PublicRead sets this Buckets objects to be publicly readable.
	// +optional
	PublicRead bool `json:"publicRead,omitempty"`
	// PublicList sets this Buckets objects to be publicly listable.
	// +optional
	PublicList bool `json:"publicList,omitempty"`
	// Versioning enables object versioning for this Bucket.
	// +optional
	Versioning bool `json:"versioning,omitempty"`
	// LifecyclePolicies allows to define automatic expiry (deletion) of
	// objects using certain rules.
	// +optional
	LifecyclePolicies []*BucketLifecyclePolicy `json:"lifecyclePolicies,omitempty"`
	// CORS settings for this bucket. CORS is a mechanism to allow code
	// running in a browser to make requests to a domain other than the
	// one from where it originated.
	// +optional
	CORS *CORSConfig `json:"cors,omitempty"`
}

// A ObjectsBucketStatus represents the observed state of an ObjectsBucket.
type ObjectsBucketStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ObjectsBucketObservation `json:"atProvider,omitempty"`
}

// ObjectsBucketObservation are the observable fields of an ObjectsBucket.
type ObjectsBucketObservation struct {
	meta.ReferenceStatus `json:",inline"`
	// URL where the bucket can be accessed for path-style access.
	URL string `json:"url"`
	// API endpoint to use with S3 compatible clients.
	Endpoint string `json:"endpoint"`
	// BytesUsed shows the amount of bytes a bucket is currently using.
	BytesUsed int64 `json:"bytesUsed"`
	// ObjectCount shows the amount of objects a bucket has.
	ObjectCount int64 `json:"objectCount"`
}

// Postgres deploys a Self Service Postgres instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Postgres struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresSpec   `json:"spec"`
	Status            PostgresStatus `json:"status,omitempty"`
}

// PostgresList contains a list of Postgres database
// +kubebuilder:object:root=true
type PostgresList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Postgres `json:"items"`
}

// A PostgresSpec defines the desired state of a Postgres database.
type PostgresSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            PostgresParameters `json:"forProvider"`
}

// PostgresParameters are the configurable fields of a Postgres database.
type PostgresParameters struct {
	// MachineType defines the sizing for a particular machine and
	// implicitly the provider.
	// +kubebuilder:validation:Enum=nine-standard-1;nine-standard-2
	MachineType infrav1alpha1.MachineType `json:"machineType"`
	// Location specifies in which Datacenter the database will be spawned.
	// Needs to match the available MachineTypes in that datacenter.
	// +immutable
	Location meta.LocationName `json:"location"`
	// Version specifies the Postgres version.
	// Needs to match an available Postgres Version.
	// +immutable
	// +optional
	// +kubebuilder:default:="15"
	Version PostgresVersion `json:"version"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to the db.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	// +listType:="set"
	// +optional
	AllowedCIDRs []IPv4CIDR `json:"allowedCIDRs"`
	// SSHKeys contains a list of SSH public keys, allowed to connect to the
	// db server, in order to up-/download and directly restore database backups.
	// +optional
	SSHKeys []SSHKey `json:"sshKeys"`
	// Number of daily database backups to keep. Note, that setting this
	// to 0, backup will be disabled and existing dumps deleted immediately.
	// +optional
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=365
	KeepDailyBackups *int `json:"keepDailyBackups,omitempty"`
}

// PostgresVersion Version of Postgres
// +kubebuilder:validation:Enum="13";"14";"15"
type PostgresVersion string

// IPv4CIDR represents a IPv4 address in CIDR notation
// +kubebuilder:validation:Pattern=`\A([0-9]{1,3}\.){3}[0-9]{1,3}\/([0-9]|[1-2][0-9]|3[0-2])\z`
type IPv4CIDR string

// SSHKey Public SSH key without options
// +kubebuilder:validation:Pattern=`\A(?:sk-(?:ecdsa-sha2-nistp256|ssh-ed25519)@openssh\.com|ecdsa-sha2-nistp(?:256|384|521))|ssh-(?:ed25519|dss|rsa) [a-z0-9A-Z+\/]+={0,2}(?: [^\n]+)?\z`
type SSHKey string

// A PostgresStatus represents the observed state of a Postgres database.
type PostgresStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               PostgresObservation `json:"atProvider"`
}

// PostgresObservation are the observable fields of a Postgres database.
type PostgresObservation struct {
	// FQDN is the fully qualified domain name, at which the database is reachable at.
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// Size specifies the total disk size
	// +optional
	Size *resource.Quantity `json:"size,omitempty"`
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// Registry describes and deploys Docker Container Registry to a cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistrySpec   `json:"spec"`
	Status            RegistryStatus `json:"status,omitempty"`
}

// RegistryList contains a list of Registry
// +kubebuilder:object:root=true
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// An RegistrySpec defines the desired state of a Registry.
type RegistrySpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	// +optional
	ForProvider RegistryParameters `json:"forProvider,omitempty"`
}

// RegistryParameters are the configurable fields of a Registry.
type RegistryParameters struct {
	// NineAuthServer adds an auth server in front of the registry to allow
	// authentication and authorization using Nine API credentials. If
	// disabled auth will be done using basic auth.
	// +kubebuilder:default:=false
	// +optional
	NineAuthServer bool `json:"nineAuthServer,omitempty"`
}

// An RegistryStatus represents the observed state of a Registry.
type RegistryStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               RegistryObservation `json:"atProvider"`
}

// RegistryObservation are the observable fields of a Registry.
type RegistryObservation struct {
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
	// URL shows the URL of the deployed registry
	URL string `json:"url,omitempty"`
}
