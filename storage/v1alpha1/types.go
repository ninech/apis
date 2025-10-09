package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	infra "github.com/ninech/apis/infrastructure/v1alpha1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const (
	// BucketRoleReader allows reads on all objects in a Bucket.
	BucketRoleReader BucketRole = "reader"
	// BucketRoleWriter allows reads and writes on all objects in a Bucket.
	BucketRoleWriter BucketRole = "writer"
	// BucketUserLocationDefault represents the default BucketUser location
	// if no explicit version was specified.
	BucketUserLocationDefault = meta.LocationNineCZ42
	// BucketUserCredentialAccessKey is the key name that the s3 access key will get in the
	// ConnectionCredential
	BucketUserCredentialAccessKey = "s3_access_key"
	// BucketUserCredentialSecretKey is the key name that the s3 secret key will get in the
	// ConnectionCredential
	BucketUserCredentialSecretKey                          = "s3_secret_key"
	FiveMins                      MigrationIntervalMinutes = "5"
	TenMins                       MigrationIntervalMinutes = "10"
	FifteenMins                   MigrationIntervalMinutes = "15"
	ThirtyMins                    MigrationIntervalMinutes = "30"
	SixtyMins                     MigrationIntervalMinutes = "60"
	// SyncStatusPending indicates the sync is pending and is scheduled to be started.
	SyncStatusPending SyncStatus = "pending"
	// SyncStatusSucceeded indicates that the sync job has
	// succeeded.
	SyncStatusSucceeded SyncStatus = "succeeded"
	// SyncStatusRunning indicates that the sync job is
	// currently running.
	SyncStatusRunning SyncStatus = "running"
	// SyncStatusFailed indicates that the sync job was unable
	// to successfully complete within the configured amount of retries.
	SyncStatusFailed SyncStatus = "failed"
	// SyncStatusUnknown indicates the status is unknown.
	SyncStatusUnknown SyncStatus = "unknown"
	// DBLocationDefault is the default location for database instances.
	DBLocationDefault = meta.LocationNineCZ41
	// DBDailyBackupsDefault is the default number of daily database backups to keep.
	DBKeepDailyBackupsDefault int = 10
	// DatabaseBackupStatePending indicates a scheduled but not yet started backup.
	DatabaseBackupStatePending DatabaseBackupState = "pending"
	// DatabaseBackupStateSucceeded indicates that the backup was completed successfully.
	DatabaseBackupStateSucceeded DatabaseBackupState = "succeeded"
	// DatabaseBackupStateRunning indicates that the backup is still running.
	DatabaseBackupStateRunning DatabaseBackupState = "running"
	// DatabaseBackupStateFailed indicates that the backup was unable to complete successfully.
	DatabaseBackupStateFailed DatabaseBackupState = "failed"
	// DatabaseBackupStateUnknown indicates the backup's state to be unknown.
	DatabaseBackupStateUnknown DatabaseBackupState = "unknown"
	// DatabaseBackupScheduleCalendarDisabled disables backup operations.
	DatabaseBackupScheduleCalendarDisabled DatabaseBackupScheduleCalendar = "disabled"
	// DatabaseBackupScheduleCalendarDaily sets up daily backups for a database.
	DatabaseBackupScheduleCalendarDaily DatabaseBackupScheduleCalendar = "daily"
	// DatabaseBackupScheduleTTLMin is the minimum duration after the backups creation until it can be deleted.
	DatabaseBackupScheduleTTLMin = 24 * time.Hour
	// DatabaseBackupScheduleTTLMax is the maximum duration after the backups creation until it can be deleted.
	DatabaseBackupScheduleTTLMax = 365 * 24 * time.Hour
	// DatabaseBackupMinimumInterval is the minimum interval between two backups.
	DatabaseBackupMinimumInterval = 1 * time.Hour
	// KeyValueStoreVersion7 KeyValueStore version 7
	KeyValueStoreVersion7 KeyValueStoreVersion = "7"
	// KeyValueStoreUser is the name of the KeyValueStore user account.
	KeyValueStoreUser string = "default"
	// KeyValueStorePort is the network port used by any KeyValueStore instance.
	KeyValueStorePort int32 = 6379
	// KeyValueStoreVersionDefault represents the default KeyValueStore version
	// if no explicit version was specified.
	KeyValueStoreVersionDefault = KeyValueStoreVersion7
	// KeyValueStoreLocationDefault represents the default KeyValueStore datacenter location
	// if no explicit version was specified.
	KeyValueStoreLocationDefault = meta.LocationNineES34
	// KeyValueStoreMaxMemoryPolicyDefault represents the default KeyValueStore max memory policy
	// if no explicit version was specified.
	KeyValueStoreMaxMemoryPolicyDefault KeyValueStoreMaxMemoryPolicy = "allkeys-lru"
	// KeyValueStoreMemorySizeDefault is the default memory size for a KeyValueStor.
	KeyValueStoreMemorySizeDefault = "1Gi"
	// KeyValueStoreMemorySizeMin is the minimum memory size for a KeyValueStore.
	KeyValueStoreMemorySizeMin = "256Mi"
	// KeyValueStoreMemorySizeMax is the maximum memory size for a KeyValueStore.
	KeyValueStoreMemorySizeMax = "4Gi"
	// MySQLVersion8 MySQL version 8
	MySQLVersion8 MySQLVersion = "8"
	// MySQLUser is the name of the MySQL user account.
	MySQLUser string = "dbadmin"
	// MySQLLocationDefault represents the default MySQL location
	// if no explicit version was specified.
	MySQLLocationDefault = meta.LocationNineCZ41
	// MySQLVersionDefault represents the default MySQL version used
	// if no explicit version was specified.
	MySQLVersionDefault MySQLVersion = MySQLVersion8
	// MySQLUserDefault is the default MySQL user name.
	MySQLUserDefault string = MySQLUser
	// MySQLLongQueryTimeDefault is the default value for long_query_time.
	MySQLLongQueryTimeDefault LongQueryTime = "10s"
	// MySQLCharsetDefault is the default character set.
	MySQLCharsetDefault string = "utf8mb4"
	// MySQLCollationDefault is the default collation.
	MySQLCollationDefault string = "utf8mb4_unicode_ci"
	// MySQLMinWordLengthDefault is the default value for `ft_min_word_len` and `innodb_ft_min_token_size`.
	MySQLMinWordLengthDefault int = 3
	// MySQLTransactionIsolationDefault is the default transaction isolation level.
	MySQLTransactionIsolationDefault MySQLTransactionCharacteristic = "REPEATABLE-READ"
	// MySQLBackupRetentionDays is the number of days to retain backups by default.
	MySQLBackupRetentionDaysDefault int = 10
	// MySQLDatabaseSizeMax is the maximum size of a [MySQLDatabase].
	// After reaching [MySQLDatabaseSizeMax], the permissions to writing additional
	// data to the database will be revoked.
	MySQLDatabaseSizeMax = "10Gi"
	// MySQLDatabaseLocationDefault represents the default MySQL database location,
	// if no explicit location was specified.
	MySQLDatabaseLocationDefault = meta.LocationNineES34
	// MySQLDatabaseVersionDefault represents the default MySQL version used
	// if no explicit version was specified.
	MySQLDatabaseVersionDefault MySQLVersion = MySQLVersionDefault
	// OpenSearchVersion2 OpenSearch version 2
	OpenSearchVersion2 OpenSearchVersion = "2"
	// OpenSearchVersion3 OpenSearch version 3
	OpenSearchVersion3 OpenSearchVersion = "3"
	// OpenSearchUser is the name of the OpenSearch user account.
	OpenSearchUser string = "admin"
	// OpenSearchVersionDefault is the default OpenSearch version.
	OpenSearchVersionDefault = OpenSearchVersion3
	// OpenSearchLocationDefault represents the default OpenSearch datacenter location
	// if no explicit version was specified.
	OpenSearchLocationDefault = meta.LocationNineES34
	// OpenSearchClusterTypeSingle represents a single node cluster.
	OpenSearchClusterTypeSingle OpenSearchClusterType = "single"
	// OpenSearchClusterTypeMulti represents a multi node cluster.
	OpenSearchClusterTypeMulti OpenSearchClusterType = "multi"
	// OpenSearchClusterTypeDefault represents the default cluster type.
	OpenSearchClusterTypeDefault = OpenSearchClusterTypeMulti
	// OpenSearchHealthStatusGreen means all primary shards and their replicas are allocated to nodes.
	OpenSearchHealthStatusGreen OpenSearchHealthStatus = "green"
	// OpenSearchHealthStatusYellow means all primary shards are allocated to nodes, but some replicas aren’t.
	OpenSearchHealthStatusYellow OpenSearchHealthStatus = "yellow"
	// OpenSearchHealthStatusRed means at least one primary shard is not allocated to any node.
	OpenSearchHealthStatusRed OpenSearchHealthStatus = "red"
	// PostgresVersion17 Postgres version 17
	PostgresVersion17 PostgresVersion = "17"
	// PostgresVersion16 Postgres version 16
	PostgresVersion16 PostgresVersion = "16"
	// PostgresVersion15 Postgres version 15
	PostgresVersion15 PostgresVersion = "15"
	// PostgresVersion14 Postgres version 14
	// Deprecated: Please use a newer version for new deployments.
	PostgresVersion14 PostgresVersion = "14"
	// PostgresVersion13 Postgres version 13
	// Deprecated: Please use a newer version for new deployments.
	PostgresVersion13 PostgresVersion = "13"
	// PostgresUser is the name of the Postgres user account.
	PostgresUser string = "dbadmin"
	// PostgresLocationDefault represents the default PostgreSQL datacenter location.
	// if no explicit version was specified.
	PostgresLocationDefault = meta.LocationNineCZ41
	// PostgresVersionDefault represents the default PostgreSQL version used
	// if no explicit version was specified.
	PostgresVersionDefault PostgresVersion = PostgresVersion17
	// PostgresBackupRetentionDaysDefault is the number of days to retain backups by default.
	PostgresBackupRetentionDaysDefault int = 10
	// PostgresDatabaseSizeMax is the maximum size of a PostgresDatabase.
	// After reaching [PostgresDatabaseSizeMax], the permissions to writing additional
	// data to the database will be revoked.
	PostgresDatabaseSizeMax = MySQLDatabaseSizeMax
	// PostgresDatabaseLocationDefault represents the default Postgres database location,
	// if no explicit location was specified.
	PostgresDatabaseLocationDefault = meta.LocationNineES34
	// PostgresDatabaseVersionDefault represents the default PostgreSQL version used
	// if no explicit version was specified.
	PostgresDatabaseVersionDefault PostgresVersion = PostgresVersionDefault
)

var (
	// BucketLocationOptions is a list of available locations for buckets.
	BucketLocationOptions = []string{string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// BucketUserLocationOptions is a list of available locations for bucket users.
	BucketUserLocationOptions = []string{string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// KeyValueStoreVersions represents the available versions for KeyValueStore instances.
	KeyValueStoreVersions = []KeyValueStoreVersion{KeyValueStoreVersion7}
	// KeyValueStoreLocationOptions represents the available locations for KeyValueStore instances.
	KeyValueStoreLocationOptions = []meta.LocationName{meta.LocationNineES34, meta.LocationNineCZ41}
	// MySQLMachineTypeDefault specifies the default machine type.
	MySQLMachineTypeDefault = infra.MachineTypeNineDBS
	// MySQLModeDefault is the list of enabled SQL modes.
	MySQLModeDefault = []string{"ONLY_FULL_GROUP_BY", "STRICT_TRANS_TABLES", "NO_ZERO_IN_DATE", "NO_ZERO_DATE", "ERROR_FOR_DIVISION_BY_ZERO", "NO_ENGINE_SUBSTITUTION"}
	// MySQLLocationOptions is a list of available datacenter locations.
	MySQLLocationOptions = []string{string(meta.LocationNineCZ41), string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// MySQLMachineTypes is a list of available machine types.
	MySQLMachineTypes []infra.MachineType = infra.MachineTypesDB
	// MySQLDatabaseLocationOptions is a list of available datacenter locations.
	MySQLDatabaseLocationOptions = []string{string(meta.LocationNineCZ41), string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// MySQLDatabaseVersions is a list of all available MySQLVersions.
	MySQLDatabaseVersions = []MySQLVersion{MySQLDatabaseVersionDefault}
	// OpenSearchMachineTypeDefault represents the default machine type for OpenSearch clusters.
	OpenSearchMachineTypeDefault = infra.MachineTypeNineSearchS
	// OpenSearchVersions represents the available versions of OpenSearch.
	OpenSearchVersions = []OpenSearchVersion{OpenSearchVersion2, OpenSearchVersion3}
	// OpenSearchLocationOptions represents the available locations for OpenSearch clusters.
	OpenSearchLocationOptions = []meta.LocationName{meta.LocationNineES34, meta.LocationNineCZ41}
	// OpenSearchClusterTypes represents the available cluster types.
	OpenSearchClusterTypes = []OpenSearchClusterType{OpenSearchClusterTypeSingle, OpenSearchClusterTypeMulti}
	// OpenSearchMachineTypes represents the available machine types.
	OpenSearchMachineTypes = []infra.MachineType{infra.MachineTypeNineSearchXS, infra.MachineTypeNineSearchS, infra.MachineTypeNineSearchM, infra.MachineTypeNineSearchL, infra.MachineTypeNineSearchXL}
	// OpenSearchClusterMachineTypes represents the available machine types for multi node clusters.
	OpenSearchMachineTypesMulti = []infra.MachineType{infra.MachineTypeNineSearchS, infra.MachineTypeNineSearchM, infra.MachineTypeNineSearchL, infra.MachineTypeNineSearchXL}
	// PostgresMachineTypeDefault specifies the default machine type.
	PostgresMachineTypeDefault = infra.MachineTypeNineDBS
	// PostgresLocationOptions is a list of available datacenter locations.
	PostgresLocationOptions = []string{string(meta.LocationNineCZ41), string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// PostgresMachineTypes is a list of available machine types.
	PostgresMachineTypes []infra.MachineType = infra.MachineTypesDB
	// PostgresVersions is a list of all available PostgresVersions.
	PostgresVersions = []PostgresVersion{PostgresVersion17, PostgresVersion16, PostgresVersion15}
	// PostgresVersionsDeprecated is a list of all deprecated PostgresVersions.
	PostgresVersionsDeprecated = []PostgresVersion{PostgresVersion14, PostgresVersion13}
	// PostgresDatabaseLocationOptions is a list of available datacenter locations.
	PostgresDatabaseLocationOptions = []string{string(meta.LocationNineCZ41), string(meta.LocationNineCZ42), string(meta.LocationNineES34)}
	// PostgresDatabaseVersions is a list of all available PostgresVersions.
	PostgresDatabaseVersions = []PostgresVersion{PostgresDatabaseVersionDefault}
)

// BucketRole defines what kind of access is possible.
// +kubebuilder:validation:Enum=reader;writer
type BucketRole string

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
	// +optional
	// +kubebuilder:default:="standard"
	StorageType BucketStorageType `json:"storageTier,omitempty"`
	// Permissions configures user access to the objects in this Bucket.
	// +optional
	Permissions []*BucketPermission `json:"permissions,omitempty"`
	// PublicRead sets this Buckets objects to be publicly readable.
	// +optional
	PublicRead bool `json:"publicRead,omitempty"`
	// PublicList sets this Buckets objects to be publicly listable.
	// +optional
	PublicList bool `json:"publicList,omitempty"`
	// Encryption enables encryption at rest for this Bucket. This is only
	// relevant for Buckets which use the v1 storage backend. Buckets with a
	// backend of v2 are always encrypted at rest.
	// Deprecated: Only affects v1 Buckets and will be removed in the future.
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
	// +optional
	// +kubebuilder:default:="v2"
	BackendVersion BucketBackendVersion `json:"backendVersion,omitempty"`
	// CustomHostnames are DNS entries under which the bucket should be
	// accessible. This can be used to serve public objects via an own
	// domain name.
	// +optional
	CustomHostnames []string `json:"customHostnames,omitempty"`
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

// BackendVersion specifies the bucket backend version to use. While the
// APIs work the same, buckets with v2 are only compatible with
// bucketusers also on v2.
// +kubebuilder:validation:Enum=v2
type BucketBackendVersion string

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
	// CustomHostnamesVerification shows the DNS verification status of all
	// custom hostnames.
	CustomHostnamesVerification meta.DNSVerificationStatus `json:"customHostnamesVerification,omitempty"`
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
	// +kubebuilder:default:="v2"
	BackendVersion BucketBackendVersion `json:"backendVersion,omitempty"`
}

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

// BucketMigration is an object to migrate a v1 Bucket's data to a v2 Bucket.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type BucketMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketMigrationSpec   `json:"spec"`
	Status            BucketMigrationStatus `json:"status,omitempty"`
}

// BucketMigrationList contains a list of BucketMigration
// +kubebuilder:object:root=true
type BucketMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketMigration `json:"items"`
}

// A BucketMigrationSpec defines the desired state of a BucketMigration.
type BucketMigrationSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            BucketMigrationParameters `json:"forProvider"`
}

// BucketMigrationParameters are the configurable fields of a BucketMigration.
type BucketMigrationParameters struct {
	// Source represents the source Bucket
	Source meta.TypedReference `json:"source"`
	// SourceUser is the BucketUser used for reading objects in the source
	// Bucket. It needs read permissions on the Source Bucket.
	SourceUser meta.TypedReference `json:"sourceUser"`
	// Destination represents the destination Bucket
	Destination meta.TypedReference `json:"destination"`
	// DestinationUser is the BucketUser used for reading objects in the source
	// Bucket. It needs read permissions on the Destination Bucket.
	DestinationUser meta.TypedReference `json:"destinationUser"`
	// If set to true, the BucketMigration will delete all objects in the
	// Destination that do not exist in the Source.
	// +kubebuilder:default:=false
	// +optional
	DeleteOutOfSyncObjects bool `json:"deleteOutOfSyncObjects"`
	// Interval defines how often the sync is run in minutes.
	// +kubebuilder:validation:Enum="5";"10";"15";"30";"60"
	// +kubebuilder:default:="15"
	// +optional
	Interval MigrationIntervalMinutes `json:"interval,omitempty"`
}
type MigrationIntervalMinutes string

// A BucketMigrationStatus represents the observed state of a BucketMigration.
type BucketMigrationStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               BucketMigrationObservation `json:"atProvider,omitempty"`
}

// BucketMigrationObservation are the observable fields of a BucketMigration.
type BucketMigrationObservation struct {
	// InitialSync indicates the status of the initial bucket data sync.
	InitialSync BucketMigrationSyncStatus `json:"initialSync"`
	// InitialSync indicates the status of the continuous bucket data sync.
	Resync               BucketMigrationSyncStatus `json:"resync"`
	meta.ReferenceStatus `json:",inline"`
}
type BucketMigrationSyncStatus struct {
	// SyncStatus indicates the status of the last sync run.
	// +optional
	Status SyncStatus `json:"status"`
	// SyncStartTime is the timestamp of the last sync run start.
	// +optional
	SyncStartTime *metav1.Time `json:"syncStartTime"`
	// SyncEndTime is the timestamp of the last sync run end.
	// +optional
	SyncEndTime *metav1.Time `json:"syncEndTime"`
	// Schedule in cron format.
	// +optional
	Schedule string `json:"schedule,omitempty"`
}

// SyncStatus represents the sync job status
type SyncStatus string

// DatabaseBackup creates a single backup in an object storage bucket.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type DatabaseBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseBackupSpec   `json:"spec"`
	Status            DatabaseBackupStatus `json:"status,omitempty"`
}

// DatabaseBackupList contains a list of DatabaseBackups.
// +kubebuilder:object:root=true
type DatabaseBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseBackup `json:"items"`
}

// A DatabaseBackupSpec defines the desired state of a DatabaseBackup.
type DatabaseBackupSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            DatabaseBackupParameters `json:"forProvider"`
}

// DatabaseBackupParameters are the configurable fields of a DatabaseBackup.
type DatabaseBackupParameters struct {
	// Source is a reference to a Postgres, MySQL or MySQLDatabase object, to
	// create the database backup from.
	Source meta.LocalTypedReference `json:"source"`
	// Name is the name of the database to be backed up. This is required for
	// MySQL and Postgres types as there can be multiple databases on these servers.
	// For shared databases like MySQLDatabase this field is ignored, as
	// the database is defined in the object.
	// +optional
	Name string `json:"name,omitempty"`
	// Bucket reference to an object storage bucket into which the backup is to be created.
	Bucket meta.LocalReference `json:"bucket"`
	// BucketUser is used to write database backups and read them again to restore.
	BucketUser meta.LocalReference `json:"bucketUser"`
	// Expiration is the time when the backup will be deleted.
	Expiration metav1.Time `json:"expiration"`
}

// A DatabaseBackupStatus represents the observed state of a DatabaseBackup.
type DatabaseBackupStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               DatabaseBackupObservation `json:"atProvider,omitempty"`
}

// DatabaseBackupObservation are the observable fields of a DatabaseBackup.
type DatabaseBackupObservation struct {
	// Type defines the database type that is contained in this backup.
	// +kubebuilder:default:=unknown
	// +optional
	Type DatabaseBackupType `json:"type,omitempty"`
	// State represents the backup state.
	// +kubebuilder:default:=unknown
	// +optional
	State DatabaseBackupState `json:"state,omitempty"`
	// Path is the object key name of the backup in the object store bucket.
	// +optional
	Path string `json:"path,omitempty"`
	// Start is the time when the backup operation started.
	// +optional
	Start metav1.Time `json:"start,omitempty"`
	// End is the time when the backup operation ended.
	// +optional
	End metav1.Time `json:"end,omitempty"`
	// Size is the size of the backup uncompressed.
	// +optional
	Size *resource.Quantity `json:"size,omitempty"`
	// Version contains the version information of the backup.
	// +optional
	Version DatabaseBackupVersion `json:"version,omitempty"`
}

// DatabaseBackupType represents the database management system a database backup is made from.
// This information is crucial, especially when it comes to restoring the backup.
// +kubebuilder:validation:Enum=mysql;postgres;unknown
type DatabaseBackupType string

// DatabaseBackupState represents the backup state.
// +kubebuilder:validation:Enum=pending;succeeded;running;failed;unknown
type DatabaseBackupState string

// DatabaseBackupVersion contains the version information of the backup.
type DatabaseBackupVersion struct {
	// Database is version of the database instance.
	// +optional
	Database string `json:"database,omitempty"`
	// Utility is version of the utility used to create the backup.
	// +optional
	Utility string `json:"utility,omitempty"`
}

// DatabaseBackupSchedule creates a single backup in an object storage bucket.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type DatabaseBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseBackupScheduleSpec   `json:"spec"`
	Status            DatabaseBackupScheduleStatus `json:"status,omitempty"`
}

// DatabaseBackupScheduleList contains a list of DatabaseBackupSchedules.
// +kubebuilder:object:root=true
type DatabaseBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseBackupSchedule `json:"items"`
}

// A DatabaseBackupScheduleSpec defines the desired state of a DatabaseBackupSchedule.
type DatabaseBackupScheduleSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            DatabaseBackupScheduleParameters `json:"forProvider"`
}

// DatabaseBackupScheduleParameters are the configurable fields of a DatabaseBackupSchedule.
type DatabaseBackupScheduleParameters struct {
	// Schedule is a cron description of how often a backup is to be created.
	// +kubebuilder:default:=daily
	// +optional
	Schedule DatabaseBackupScheduleCalendar `json:"schedule,omitempty"`
	// TTL is the duration after the backups creation until it can be deleted.
	// +kubebuilder:default:="240h"
	TTL *metav1.Duration `json:"ttl"`
	// BucketUsers is a list of bucket users that will gain read access to the
	// target bucket containing the database backups.
	// +optional
	BucketUsers []meta.LocalReference `json:"bucketUsers,omitempty"`
	// Name is the name of the database to be backed up. This is required for
	// MySQL and Posgres types as there can be multiple databases on these servers.
	// For shared databases like MySQLDatabase this field is ignored, as
	// the database is defined in the object.
	// +optional
	Name string `json:"name,omitempty"`
	// Source is a reference to a Postgres, MySQL or MySQLDatabase object, to
	// create the database backup from.
	Source meta.LocalTypedReference `json:"source"`
}

// DatabaseBackupScheduleCalendar is a limited cron or systemd calendar expression to define the
// frequency of database backup operations.
// +kubebuilder:validation:Enum=disabled;daily
type DatabaseBackupScheduleCalendar string

// A DatabaseBackupScheduleStatus represents the observed state of a DatabaseBackupSchedule.
type DatabaseBackupScheduleStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               DatabaseBackupScheduleObservation `json:"atProvider,omitempty"`
}

// DatabaseBackupScheduleObservation are the observable fields of a DatabaseBackupSchedule.
type DatabaseBackupScheduleObservation struct {
	// Type defines the database type that is contained in this backup.
	// +kubebuilder:default:=unknown
	// +optional
	Type DatabaseBackupType `json:"type,omitempty"`
	// TargetBucket is the object bucket into which backups are created.
	// +optional
	TargetBucket meta.LocalReference `json:"targetBucket,omitempty"`
	// Last is the time when the most recent backup operation was created.
	// +optional
	Last metav1.Time `json:"start,omitempty"`
	// Next is the time when the next backup operation is to be created.
	// +optional
	Next metav1.Time `json:"end,omitempty"`
}

// KeyValueStore deploys an on-demand KeyValueStore instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="MEMORYSIZE",type="string",JSONPath=".spec.forProvider.memorySize"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=kvs
// +kubebuilder:object:root=true
type KeyValueStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyValueStoreSpec   `json:"spec"`
	Status            KeyValueStoreStatus `json:"status,omitempty"`
}

// KeyValueStoreList contains a list of KeyValueStore instances.
// +kubebuilder:object:root=true
type KeyValueStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyValueStore `json:"items"`
}

// A KeyValueStoreSpec defines the desired state of a Key-Value document store.
type KeyValueStoreSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            KeyValueStoreParameters `json:"forProvider"`
}

// KeyValueStoreParameters are the configurable fields of a Key-Value document store.
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
type KeyValueStoreParameters struct {
	// Location specifies in which datacenter the in-memory data store will be spawned.
	//
	// +optional
	// +kubebuilder:default:="nine-es34"
	Location meta.LocationName `json:"location"`
	// Version specifies the KeyValueStore version.
	// Needs to match an available KeyValueStore Version.
	//
	// +optional
	// +kubebuilder:default:="7"
	Version KeyValueStoreVersion `json:"version,omitempty"`
	// MemorySize configures KeyValueStore to use a specified amount of memory for the data set.
	//
	// +optional
	// +kubebuilder:default:="1Gi"
	MemorySize *KeyValueStoreMemorySize `json:"memorySize,omitempty"`
	// MaxMemoryPolicy specifies the exact behavior KeyValueStore follows when the maxmemory limit is reached.
	//
	// +optional
	// +kubebuilder:default:="allkeys-lru"
	MaxMemoryPolicy KeyValueStoreMaxMemoryPolicy `json:"maxMemoryPolicy,omitempty"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to the instance.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	//
	// +listType:="set"
	// +optional
	AllowedCIDRs []meta.IPv4CIDR `json:"allowedCIDRs,omitempty"`
	// PublicNetworkingEnabled specifies if the service should be available without service connection.
	//
	// +optional
	// +kubebuilder:default:=true
	PublicNetworkingEnabled *bool `json:"publicNetworkingEnabled,omitempty"`
}

// KeyValueStoreVersion defines the KeyValueStore version to be used.
//
// +kubebuilder:validation:Enum="7"
type KeyValueStoreVersion string

// KeyValueStoreMemorySize configures KeyValueStore to use a specified amount of memory for the data set.
// When the specified amount of memory is reached, how eviction policies are configured determines the default behavior.
// KeyValueStore can return errors for commands that could result in more memory being used,
// or it can evict some old data to return back to the specified limit every time new data is added.
// https://redis.io/docs/reference/eviction/#maxmemory-configuration-directive
// +kubebuilder:validation:Type=string
// +kubebuilder:validation:XIntOrString
// +kubebuilder:validation:Pattern=`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`
type KeyValueStoreMemorySize struct {
	resource.Quantity `json:",inline"`
}

// KeyValueStoreMaxMemoryPolicy defines the exact behavior KeyValueStore follows when the maxmemory limit is reached:
// https://redis.io/docs/reference/eviction/#eviction-policies
// +kubebuilder:default:="allkeys-lru"
// +kubebuilder:validation:Enum="noeviction";"allkeys-lru";"allkeys-lfu";"volatile-lru";"volatile-lfu";"allkeys-random";"volatile-random";"volatile-ttl"
type KeyValueStoreMaxMemoryPolicy string

// A KeyValueStoreStatus represents the observed state of a Key-Value document store.
type KeyValueStoreStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               KeyValueStoreObservation `json:"atProvider"`
}

// KeyValueStoreObservation are the observable fields of a Key-Value document store.
type KeyValueStoreObservation struct {
	// FQDN is the fully qualified domain name, at which the instance is reachable at.
	//
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// PrivateNetworkingFQDN is the magic DNS name of a service connection destination.
	//
	// +optional
	PrivateNetworkingFQDN string `json:"privateNetworkingFQDN,omitempty"`
	// DiskSize specifies the total disk size used for persistence.
	// Note that the disk size cannot be decreased and is based
	// on the configured MemorySize.
	//
	// +optional
	DiskSize *resource.Quantity `json:"diskSize,omitempty"`
	// CACert is the certificate of the CA that is used by the service.
	// The value is a base64 encoded PEM.
	//
	// +optional
	CACert string `json:"caCert,omitempty"`
	// Status of all the child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// MySQL deploys a Self Service MySQL instance.
//
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type MySQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MySQLSpec   `json:"spec"`
	Status            MySQLStatus `json:"status,omitempty"`
}

// MySQLList contains a list of MySQL database.
// +kubebuilder:object:root=true
type MySQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQL `json:"items"`
}

// A MySQLSpec defines the desired state of a MySQL database.
type MySQLSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            MySQLParameters `json:"forProvider"`
}

// MySQLParameters are the configurable fields of a MySQL database.
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
type MySQLParameters struct {
	// MachineType defines the sizing for a particular machine and
	// implicitly the provider.
	//
	// +optional
	// +kubebuilder:default:="nine-db-s"
	MachineType infra.MachineType `json:"machineType,omitempty"`
	// Location specifies in which datacenter the database will be spawned.
	// Needs to match the available MachineTypes in that datacenter.
	//
	// +optional
	// +kubebuilder:default:="nine-cz41"
	Location meta.LocationName `json:"location,omitempty"`
	// Version specifies the MySQL version.
	// Needs to match an available MySQL Version.
	//
	// +optional
	// +kubebuilder:validation:Enum="8"
	// +kubebuilder:default:="8"
	Version MySQLVersion `json:"version,omitempty"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to the db.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	// Access from our Kubernetes products NKE and GKE as well as from deplo.io is already enabled.
	// See the documentation here: https://docs.nine.ch/docs/on-demand-databases/configuration-options#allowed-ip-addresses
	//
	// +listType:="set"
	// +optional
	AllowedCIDRs []meta.IPv4CIDR `json:"allowedCIDRs,omitempty"`
	// SSHKeys contains a list of SSH public keys, allowed to connect to the
	// db server, in order to up-/download and directly restore database backups.
	//
	// +optional
	SSHKeys []SSHKey `json:"sshKeys,omitempty"`
	// SQLMode configures the sql_mode setting.
	// Modes affect the SQL syntax MySQL supports and the data validation checks it performs.
	//
	// +listType:="set"
	// +optional
	// +kubebuilder:default={"ONLY_FULL_GROUP_BY","STRICT_TRANS_TABLES","NO_ZERO_IN_DATE","NO_ZERO_DATE","ERROR_FOR_DIVISION_BY_ZERO","NO_ENGINE_SUBSTITUTION"}
	SQLMode *[]MySQLMode `json:"sqlMode,omitempty"`
	// CharacterSet configures the `character_set_server` and collation_server` variables.
	//
	// +optional
	CharacterSet MySQLCharacterSet `json:"characterSet"`
	// LongQueryTime configures the `long_query_time` variable.
	// If a query takes longer than this many seconds, the the query is logged to the slow query log file.
	// This value is measured in real time, not CPU time, so a query that is under the threshold on a lightly
	// loaded system might be above the threshold on a heavily loaded one.
	// Smaller values of this variable result in more statements being considered long-running, with the result that more space is required for the slow query log.
	//
	// +optional
	// +kubebuilder:default="10s"
	LongQueryTime LongQueryTime `json:"longQueryTime,omitempty"`
	// MinWordLength configures the`ft_min_word_len` and `innodb_ft_min_token_size` variables.
	// Minimum length of words that are stored in an InnoDB or MyISAM FULLTEXT index.
	//
	// +optional
	// +kubebuilder:default=3
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16
	MinWordLength *int `json:"minWordLength,omitempty"`
	// TransactionIsolation configures the `transaction_isolation` variable.
	//
	// +optional
	// +kubebuilder:default="REPEATABLE-READ"
	TransactionIsolation MySQLTransactionCharacteristic `json:"transactionIsolation,omitempty"`
	// Number of daily database backups to keep.
	// Note, that setting this to 0,
	// the backup will be disabled and existing dumps will be deleted immediately.
	//
	// +optional
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=365
	KeepDailyBackups *int `json:"keepDailyBackups,omitempty"`
}

// MySQLVersion Version of MySQL
// +kubebuilder:validation:Enum="8"
type MySQLVersion string

// SSHKey Public SSH key without options
// +kubebuilder:validation:Pattern=`\A(?:sk-(?:ecdsa-sha2-nistp256|ssh-ed25519)@openssh\.com|ecdsa-sha2-nistp(?:256|384|521))|ssh-(?:ed25519|dss|rsa) [a-z0-9A-Z+\/]+={0,2}(?: [^\n]+)?\z`
type SSHKey string

// MySQLMode represents a single possible SQL mode.
// Modes affect the SQL syntax MySQL supports and the data validation checks it performs.
// This makes it easier to use MySQL in different environments and to use
// MySQL together with other database servers.
// https://dev.mysql.com/doc/refman/8.0/en/sql-mode.html
//
// +kubebuilder:validation:Enum="ALLOW_INVALID_DATES";"ANSI_QUOTES";"ERROR_FOR_DIVISION_BY_ZERO";"HIGH_NOT_PRECEDENCE";"IGNORE_SPACE";"NO_AUTO_VALUE_ON_ZERO";"NO_BACKSLASH_ESCAPES";"NO_DIR_IN_CREATE";"NO_ENGINE_SUBSTITUTION";"NO_UNSIGNED_SUBTRACTION";"NO_ZERO_DATE";"NO_ZERO_IN_DATE";"ONLY_FULL_GROUP_BY";"PAD_CHAR_TO_FULL_LENGTH";"PIPES_AS_CONCAT";"REAL_AS_FLOAT";"STRICT_ALL_TABLES";"STRICT_TRANS_TABLES";"TIME_TRUNCATE_FRACTIONAL"
type MySQLMode string

// MySQLCharacterSet defines the `character_set_server` and `collation_server` variables.
type MySQLCharacterSet struct {
	// Name configures the `character_set_server` variable.
	// The servers default character set.
	// See section 10.15 (https://dev.mysql.com/doc/refman/8.0/en/charset-configuration.html), "Character Set Configuration".
	// If you set this variable, you should also set collation_server to specify the collation for the character set.
	//
	// +optional
	// +kubebuilder:default="utf8mb4"
	// +kubebuilder:validation:Enum="armscii8";"ascii";"big5";"binary";"cp1250";"cp1251";"cp1256";"cp1257";"cp850";"cp852";"cp866";"cp932";"dec8";"eucjpms";"euckr";"gb2312";"gbk";"geostd8";"greek";"hebrew";"hp8";"keybcs2";"koi8r";"koi8u";"latin1";"latin2";"latin5";"latin7";"macce";"macroman";"sjis";"swe7";"tis620";"ucs2";"ujis";"utf16";"utf16le";"utf32";"utf8mb3";"utf8mb4"
	Name string `json:"name,omitempty"`
	// Collation configures the `collation_server` variable.
	// The server's default collation.
	// See section 10.15 (https://dev.mysql.com/doc/refman/8.0/en/charset-configuration.html), "Character Set Configuration".
	// This should be aligned with the configured character set.
	//
	// +optional
	// +kubebuilder:default="utf8mb4_unicode_ci"
	// +kubebuilder:validation:Enum="armscii8_bin";"armscii8_general_ci";"armscii8_general_nopad_ci";"armscii8_nopad_bin";"ascii_bin";"ascii_general_ci";"ascii_general_nopad_ci";"ascii_nopad_bin";"big5_bin";"big5_chinese_ci";"big5_chinese_nopad_ci";"big5_nopad_bin";"binary";"cp1250_bin";"cp1250_croatian_ci";"cp1250_czech_cs";"cp1250_general_ci";"cp1250_general_nopad_ci";"cp1250_nopad_bin";"cp1250_polish_ci";"cp1251_bin";"cp1251_bulgarian_ci";"cp1251_general_ci";"cp1251_general_cs";"cp1251_general_nopad_ci";"cp1251_nopad_bin";"cp1251_ukrainian_ci";"cp1256_bin";"cp1256_general_ci";"cp1256_general_nopad_ci";"cp1256_nopad_bin";"cp1257_bin";"cp1257_general_ci";"cp1257_general_nopad_ci";"cp1257_lithuanian_ci";"cp1257_nopad_bin";"cp850_bin";"cp850_general_ci";"cp850_general_nopad_ci";"cp850_nopad_bin";"cp852_bin";"cp852_general_ci";"cp852_general_nopad_ci";"cp852_nopad_bin";"cp866_bin";"cp866_general_ci";"cp866_general_nopad_ci";"cp866_nopad_bin";"cp932_bin";"cp932_japanese_ci";"cp932_japanese_nopad_ci";"cp932_nopad_bin";"dec8_bin";"dec8_nopad_bin";"dec8_swedish_ci";"dec8_swedish_nopad_ci";"eucjpms_bin";"eucjpms_japanese_ci";"eucjpms_japanese_nopad_ci";"eucjpms_nopad_bin";"euckr_bin";"euckr_korean_ci";"euckr_korean_nopad_ci";"euckr_nopad_bin";"gb2312_bin";"gb2312_chinese_ci";"gb2312_chinese_nopad_ci";"gb2312_nopad_bin";"gbk_bin";"gbk_chinese_ci";"gbk_chinese_nopad_ci";"gbk_nopad_bin";"geostd8_bin";"geostd8_general_ci";"geostd8_general_nopad_ci";"geostd8_nopad_bin";"greek_bin";"greek_general_ci";"greek_general_nopad_ci";"greek_nopad_bin";"hebrew_bin";"hebrew_general_ci";"hebrew_general_nopad_ci";"hebrew_nopad_bin";"hp8_bin";"hp8_english_ci";"hp8_english_nopad_ci";"hp8_nopad_bin";"keybcs2_bin";"keybcs2_general_ci";"keybcs2_general_nopad_ci";"keybcs2_nopad_bin";"koi8r_bin";"koi8r_general_ci";"koi8r_general_nopad_ci";"koi8r_nopad_bin";"koi8u_bin";"koi8u_general_ci";"koi8u_general_nopad_ci";"koi8u_nopad_bin";"latin1_bin";"latin1_danish_ci";"latin1_general_ci";"latin1_general_cs";"latin1_german1_ci";"latin1_german2_ci";"latin1_nopad_bin";"latin1_spanish_ci";"latin1_swedish_ci";"latin1_swedish_nopad_ci";"latin2_bin";"latin2_croatian_ci";"latin2_czech_cs";"latin2_general_ci";"latin2_general_nopad_ci";"latin2_hungarian_ci";"latin2_nopad_bin";"latin5_bin";"latin5_nopad_bin";"latin5_turkish_ci";"latin5_turkish_nopad_ci";"latin7_bin";"latin7_estonian_cs";"latin7_general_ci";"latin7_general_cs";"latin7_general_nopad_ci";"latin7_nopad_bin";"macce_bin";"macce_general_ci";"macce_general_nopad_ci";"macce_nopad_bin";"macroman_bin";"macroman_general_ci";"macroman_general_nopad_ci";"macroman_nopad_bin";"sjis_bin";"sjis_japanese_ci";"sjis_japanese_nopad_ci";"sjis_nopad_bin";"swe7_bin";"swe7_nopad_bin";"swe7_swedish_ci";"swe7_swedish_nopad_ci";"tis620_bin";"tis620_nopad_bin";"tis620_thai_ci";"tis620_thai_nopad_ci";"ucs2_bin";"ucs2_croatian_ci";"ucs2_croatian_mysql561_ci";"ucs2_czech_ci";"ucs2_danish_ci";"ucs2_esperanto_ci";"ucs2_estonian_ci";"ucs2_general_ci";"ucs2_general_mysql500_ci";"ucs2_general_nopad_ci";"ucs2_german2_ci";"ucs2_hungarian_ci";"ucs2_icelandic_ci";"ucs2_latvian_ci";"ucs2_lithuanian_ci";"ucs2_myanmar_ci";"ucs2_nopad_bin";"ucs2_persian_ci";"ucs2_polish_ci";"ucs2_roman_ci";"ucs2_romanian_ci";"ucs2_sinhala_ci";"ucs2_slovak_ci";"ucs2_slovenian_ci";"ucs2_spanish2_ci";"ucs2_spanish_ci";"ucs2_swedish_ci";"ucs2_thai_520_w2";"ucs2_turkish_ci";"ucs2_unicode_520_ci";"ucs2_unicode_520_nopad_ci";"ucs2_unicode_ci";"ucs2_unicode_nopad_ci";"ucs2_vietnamese_ci";"ujis_bin";"ujis_japanese_ci";"ujis_japanese_nopad_ci";"ujis_nopad_bin";"utf16_bin";"utf16_croatian_ci";"utf16_croatian_mysql561_ci";"utf16_czech_ci";"utf16_danish_ci";"utf16_esperanto_ci";"utf16_estonian_ci";"utf16_general_ci";"utf16_general_nopad_ci";"utf16_german2_ci";"utf16_hungarian_ci";"utf16_icelandic_ci";"utf16_latvian_ci";"utf16le_bin";"utf16le_general_ci";"utf16le_general_nopad_ci";"utf16le_nopad_bin";"utf16_lithuanian_ci";"utf16_myanmar_ci";"utf16_nopad_bin";"utf16_persian_ci";"utf16_polish_ci";"utf16_roman_ci";"utf16_romanian_ci";"utf16_sinhala_ci";"utf16_slovak_ci";"utf16_slovenian_ci";"utf16_spanish2_ci";"utf16_spanish_ci";"utf16_swedish_ci";"utf16_thai_520_w2";"utf16_turkish_ci";"utf16_unicode_520_ci";"utf16_unicode_520_nopad_ci";"utf16_unicode_ci";"utf16_unicode_nopad_ci";"utf16_vietnamese_ci";"utf32_bin";"utf32_croatian_ci";"utf32_croatian_mysql561_ci";"utf32_czech_ci";"utf32_danish_ci";"utf32_esperanto_ci";"utf32_estonian_ci";"utf32_general_ci";"utf32_general_nopad_ci";"utf32_german2_ci";"utf32_hungarian_ci";"utf32_icelandic_ci";"utf32_latvian_ci";"utf32_lithuanian_ci";"utf32_myanmar_ci";"utf32_nopad_bin";"utf32_persian_ci";"utf32_polish_ci";"utf32_roman_ci";"utf32_romanian_ci";"utf32_sinhala_ci";"utf32_slovak_ci";"utf32_slovenian_ci";"utf32_spanish2_ci";"utf32_spanish_ci";"utf32_swedish_ci";"utf32_thai_520_w2";"utf32_turkish_ci";"utf32_unicode_520_ci";"utf32_unicode_520_nopad_ci";"utf32_unicode_ci";"utf32_unicode_nopad_ci";"utf32_vietnamese_ci";"utf8mb3_bin";"utf8mb3_croatian_ci";"utf8mb3_croatian_mysql561_ci";"utf8mb3_czech_ci";"utf8mb3_danish_ci";"utf8mb3_esperanto_ci";"utf8mb3_estonian_ci";"utf8mb3_general_ci";"utf8mb3_general_mysql500_ci";"utf8mb3_general_nopad_ci";"utf8mb3_german2_ci";"utf8mb3_hungarian_ci";"utf8mb3_icelandic_ci";"utf8mb3_latvian_ci";"utf8mb3_lithuanian_ci";"utf8mb3_myanmar_ci";"utf8mb3_nopad_bin";"utf8mb3_persian_ci";"utf8mb3_polish_ci";"utf8mb3_roman_ci";"utf8mb3_romanian_ci";"utf8mb3_sinhala_ci";"utf8mb3_slovak_ci";"utf8mb3_slovenian_ci";"utf8mb3_spanish2_ci";"utf8mb3_spanish_ci";"utf8mb3_swedish_ci";"utf8mb3_thai_520_w2";"utf8mb3_turkish_ci";"utf8mb3_unicode_520_ci";"utf8mb3_unicode_520_nopad_ci";"utf8mb3_unicode_ci";"utf8mb3_unicode_nopad_ci";"utf8mb3_vietnamese_ci";"utf8mb4_bin";"utf8mb4_croatian_ci";"utf8mb4_croatian_mysql561_ci";"utf8mb4_czech_ci";"utf8mb4_danish_ci";"utf8mb4_esperanto_ci";"utf8mb4_estonian_ci";"utf8mb4_general_ci";"utf8mb4_general_nopad_ci";"utf8mb4_german2_ci";"utf8mb4_hungarian_ci";"utf8mb4_icelandic_ci";"utf8mb4_latvian_ci";"utf8mb4_lithuanian_ci";"utf8mb4_myanmar_ci";"utf8mb4_nopad_bin";"utf8mb4_persian_ci";"utf8mb4_polish_ci";"utf8mb4_roman_ci";"utf8mb4_romanian_ci";"utf8mb4_sinhala_ci";"utf8mb4_slovak_ci";"utf8mb4_slovenian_ci";"utf8mb4_spanish2_ci";"utf8mb4_spanish_ci";"utf8mb4_swedish_ci";"utf8mb4_thai_520_w2";"utf8mb4_turkish_ci";"utf8mb4_unicode_520_ci";"utf8mb4_unicode_520_nopad_ci";"utf8mb4_unicode_ci";"utf8mb4_unicode_nopad_ci";"utf8mb4_vietnamese_ci"
	Collation string `json:"collation,omitempty"`
}

// LongQueryTime configures the `long_query_time` variable.
type LongQueryTime string

// MySQLTransactionCharacteristic value sets the transaction isolation level or access mode.
// The isolation level is used for operations on InnoDB tables.
// The access mode specifies whether transactions operate in read/write or read-only mode.
//
// +kubebuilder:validation:Enum="READ-UNCOMMITTED";"READ-COMMITTED";"REPEATABLE-READ";"SERIALIZABLE"
type MySQLTransactionCharacteristic string

// A MySQLStatus represents the observed state of a MySQL database.
type MySQLStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               MySQLObservation `json:"atProvider"`
}

// MySQLObservation are the observable fields of a MySQL database.
type MySQLObservation struct {
	// FQDN is the fully qualified domain name at which the database is accessible.
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// Size indicates the total size of the disk.
	// +optional
	Size *resource.Quantity `json:"size,omitempty"`
	// Databases contains the databases that exist on the instance.
	// +optional
	Databases map[string]DatabaseObservation `json:"databases,omitempty"`
	// CACert is the certificate of the CA that is used by the service.
	// The value is a base64 encoded PEM.
	CACert string `json:"caCert,omitempty"`
	// Status of all child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// DatabaseObservation are the observable fields of a database.
type DatabaseObservation struct {
	// Size indicates the total size of the database.
	// +kubebuilder:default=0
	Size *resource.Quantity `json:"size"`
	// Connections indicates the number of connections on a database.
	// +kubebuilder:default=0
	Connections uint16 `json:"connections"`
}

// MySQLDatabase deploys a MySQL database.
//
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type MySQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MySQLDatabaseSpec   `json:"spec"`
	Status            MySQLDatabaseStatus `json:"status,omitempty"`
}

// MySQLList contains a list of MySQL database
// +kubebuilder:object:root=true
type MySQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabase `json:"items"`
}

// A MySQLDatabaseSpec defines the desired state of a MySQLDatabase.
type MySQLDatabaseSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            MySQLDatabaseParameters `json:"forProvider"`
}

// MySQLDatabaseParameters are the configurable fields of a MySQLDatabase.
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
type MySQLDatabaseParameters struct {
	// Location specifies in which data center the database will be spawned.
	// +optional
	// +kubebuilder:default:="nine-es34"
	Location meta.LocationName `json:"location,omitempty"`
	// Version specifies the MySQL version.
	// Needs to match an available MySQL Version.
	// +immutable
	// +optional
	// +kubebuilder:default:="8"
	Version MySQLVersion `json:"version,omitempty"`
	// CharacterSet configures the `character_set_server` and collation_server` variables.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CharacterSet is immutable"
	// +optional
	CharacterSet MySQLCharacterSet `json:"characterSet"`
	// InstanceRef is the instance that contains the database.
	// It is not possible to configure this value manually.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="InstanceReference is immutable"
	// +optional
	InstanceReference *meta.Reference `json:"instanceRef,omitempty"`
}

// A MySQLDatabaseStatus represents the observed state of a MySQLDatabase.
type MySQLDatabaseStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               MySQLDatabaseObservation `json:"atProvider"`
}

// MySQLDatabaseObservation are the observable fields of a MySQLDatabase.
type MySQLDatabaseObservation struct {
	// FQDN is the fully qualified domain name at which the database is accessible.
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// Name of the database and user.
	// +optional
	Name string `json:"name,omitempty"`
	// Locked reports why a database has been locked.
	// +optional
	Locked string `json:"locked,omitempty"`
	// Database details
	DatabaseObservation `json:",inline"`
	// CACert is the certificate of the CA that is used by the service.
	// The value is a base64 encoded PEM.
	CACert string `json:"caCert,omitempty"`
	// Status of all child resources.
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

// OpenSearch deploys an on-demand OpenSearch instance.
//
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="CLUSTERTYPE",type="string",JSONPath=".spec.forProvider.clusterType"
// +kubebuilder:printcolumn:name="MEMORYSIZE",type="string",JSONPath=".spec.forProvider.memorySize"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=osearch
// +kubebuilder:object:root=true
type OpenSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenSearchSpec   `json:"spec"`
	Status            OpenSearchStatus `json:"status,omitempty"`
}

// OpenSearchList contains a list of OpenSearch instances.
// +kubebuilder:object:root=true
type OpenSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenSearch `json:"items"`
}

// A OpenSearchSpec defines the desired state of an OpenSearch instance.
type OpenSearchSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            OpenSearchParameters `json:"forProvider"`
}

// OpenSearchParameters are the configurable fields of an OpenSearch cluster.
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.clusterType == oldSelf.clusterType",message="Cluster type is immutable and cannot be unset"
type OpenSearchParameters struct {
	// Location specifies in which datacenter the document store will be spawned.
	//
	// +optional
	// +kubebuilder:default:="nine-es34"
	Location meta.LocationName `json:"location,omitempty"`
	// Version specifies the OpenSearch version.
	// Needs to match an available OpenSearch Version.
	//
	// +optional
	// +kubebuilder:default:="3"
	Version OpenSearchVersion `json:"version,omitempty"`
	// ClusterType specifies the type of OpenSearch cluster.
	// A cluster can be upgraded from single-node to multi-node, but not vice versa.
	//
	// +optional
	// +kubebuilder:validation:Enum="single";"multi"
	// +kubebuilder:default:="multi"
	ClusterType OpenSearchClusterType `json:"clusterType,omitempty"`
	// MachineType defines the sizing for the OpenSearch instance.
	// +optional
	// +kubebuilder:default:="nine-search-s"
	MachineType infra.MachineType `json:"machineType,omitempty"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to the instance.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	//
	// +listType:="set"
	// +optional
	AllowedCIDRs []meta.IPv4CIDR `json:"allowedCIDRs,omitempty"`
	// PublicNetworkingEnabled specifies if the service should be available without service connection.
	//
	// +optional
	// +kubebuilder:default:=true
	PublicNetworkingEnabled *bool `json:"publicNetworkingEnabled,omitempty"`
	// BucketUsers is a list of bucket users that will gain read access to the
	// SnapshotsBucket containing the OpenSearch snapshots.
	// +optional
	BucketUsers []meta.LocalReference `json:"bucketUsers,omitempty"`
}

// OpenSearchVersion defines the Major OpenSearch version to be used.
//
// +kubebuilder:validation:Enum="2";"3"
type OpenSearchVersion string

// OpenSearchClusterType defines the cluster type of an OpenSearch installation.
type OpenSearchClusterType string

// A OpenSearchStatus represents the observed state of an OpenSearch instance.
type OpenSearchStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               OpenSearchObservation `json:"atProvider"`
}

// OpenSearchObservation are the observable fields of an OpenSearch cluster.
type OpenSearchObservation struct {
	// URL is the public address at which the instance can be reached.
	//
	// +optional
	URL meta.URL `json:"url,omitempty"`
	// FQDN is the fully qualified domain name at which the instance can be reached.
	//
	// Deprecated: FQDN exists for historical compatibility.
	// URL should be used instead as it includes the protocol and port.
	//
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// PrivateNetworkingURL is the private address at which the instance can be reached.
	// It requires a service connection to be configured.
	//
	// +optional
	PrivateNetworkingURL meta.URL `json:"privateNetworkingURL,omitempty"`
	// PrivateNetworkingFQDN is the magic DNS name of a service connection destination.
	//
	// Deprecated: PrivateNetworkingFQDN exists for historical compatibility.
	// PrivateNetworkingURL should be used instead as it includes the protocol and port.
	//
	// +optional
	PrivateNetworkingFQDN string `json:"privateNetworkingFQDN,omitempty"`
	// DiskSize specifies the total storage used for persistence.
	// It includes storage used for all nodes if a multi node cluster is deployed.
	// Disk size cannot be decreased.
	//
	// +optional
	DiskSize *resource.Quantity `json:"diskSize,omitempty"`
	// CACert is the base64 certificate of the CA that signed the certificates of OpenSearch.
	// The value is base64 a encoded PEM.
	//
	// +optional
	CACert string `json:"caCert,omitempty"`
	// OpenSearchClusterHealth holds a basic overview over the health of the cluster.
	//
	// +optional
	ClusterHealth OpenSearchClusterHealth `json:"clusterHealth,omitempty"`
	// SnapshotsBucket is the object bucket into which index snapshots are created.
	// +optional
	SnapshotsBucket meta.LocalReference `json:"snapshotsBucket,omitempty"`
	// Status of all the child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// OpenSearchClusterHealth holds a basic overview over the health of the cluster.
// https://opensearch.org/docs/latest/api-reference/cluster-api/cluster-health/
type OpenSearchClusterHealth struct {
	// ClusterName, the name of the cluster.
	//
	// +optional
	ClusterName string `json:"clusterName,omitempty"`
	// NumberOfNodes is the count of cluster nodes.
	//
	// +optional
	NumberOfNodes int `json:"numberOfNodes,omitempty"`
	// Indices is a map of all indices on the cluster.
	//
	// +optional
	Indices map[string]OpenSearchClusterIndex `json:"indices,omitempty"`
}

// OpenSearchClusterIndex holds information about an OpenSearch index health status.
type OpenSearchClusterIndex struct {
	// Status is the index's health status,
	// which represents the state of shard allocation in the cluster.
	//
	// +optional
	Status OpenSearchHealthStatus `json:"status,omitempty"`
	// Size indicates the total size of the index.
	//
	// +kubebuilder:default=0
	Size *resource.Quantity `json:"size"`
}

// OpenSearchHealthStatus expresses health in three colors: green, yellow, and red.
// A green status means all primary shards and their replicas are allocated to nodes.
// A yellow status means all primary shards are allocated to nodes, but some replicas aren’t.
// A red status means at least one primary shard is not allocated to any node.
//
// +kubebuilder:validation:Enum="green";"yellow";"red"
type OpenSearchHealthStatus string

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
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannoth be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
type PostgresParameters struct {
	// MachineType defines the sizing for a particular machine and
	// implicitly the provider.
	// +optional
	// +kubebuilder:default:="nine-db-s"
	MachineType infra.MachineType `json:"machineType,omitempty"`
	// Location specifies in which datacenter the database will be spawned.
	// Needs to match the available MachineTypes in that datacenter.
	// +optional
	// +kubebuilder:default:="nine-cz41"
	Location meta.LocationName `json:"location,omitempty"`
	// Version specifies the Postgres version.
	// Needs to match an available Postgres Version.
	// +optional
	// +kubebuilder:default:="17"
	Version PostgresVersion `json:"version,omitempty"`
	// AllowedCIDRs specify the allowed IP addresses, connecting to the db.
	// IPs are in CIDR format, e.g. 192.168.1.1/24
	// Access from our Kubernetes products NKE and GKE as well as from deplo.io is already enabled.
	// See the documentation here: https://docs.nine.ch/docs/on-demand-databases/configuration-options#allowed-ip-addresses
	//
	// +listType:="set"
	// +optional
	AllowedCIDRs []meta.IPv4CIDR `json:"allowedCIDRs,omitempty"`
	// SSHKeys contains a list of SSH public keys, allowed to connect to the
	// db server, in order to up-/download and directly restore database backups.
	// +optional
	SSHKeys []SSHKey `json:"sshKeys,omitempty"`
	// Number of daily database backups to keep.
	// Note, that setting this to 0,
	// the backup will be disabled and existing dumps will be deleted immediately.
	// +optional
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=365
	KeepDailyBackups *int `json:"keepDailyBackups,omitempty"`
}

// PostgresVersion Version of Postgres
// Please use Version >=15 for new deployments.
// https://www.postgresql.org/support/versioning/
type PostgresVersion string

// A PostgresStatus represents the observed state of a Postgres database.
type PostgresStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               PostgresObservation `json:"atProvider"`
}

// PostgresObservation are the observable fields of a Postgres database.
type PostgresObservation struct {
	// FQDN is the fully qualified domain name at which the database is accessible.
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// Size indicates the total size of the disk.
	// +optional
	Size *resource.Quantity `json:"size,omitempty"`
	// Databases contains the databases that exist on the instance.
	// +optional
	Databases map[string]DatabaseObservation `json:"databases,omitempty"`
	// CACert is the certificate of the CA that is used by the service.
	// The value is a base64 encoded PEM.
	CACert string `json:"caCert,omitempty"`
	// Status of all child resources.
	meta.ChildResourceStatus `json:",inline"`
}

// PostgresDatabase deploys a Postgres database.
//
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type PostgresDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresDatabaseSpec   `json:"spec"`
	Status            PostgresDatabaseStatus `json:"status,omitempty"`
}

// PostgresDatabaseList contains a list of Postgres database
// +kubebuilder:object:root=true
type PostgresDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresDatabase `json:"items"`
}

// A PostgresDatabaseSpec defines the desired state of a PostgresDatabase.
type PostgresDatabaseSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            PostgresDatabaseParameters `json:"forProvider"`
}

// PostgresDatabaseParameters are the configurable fields of a PostgresDatabase.
// +kubebuilder:validation:XValidation:rule="self.location == oldSelf.location",message="Location is immutable and cannot be unset"
// +kubebuilder:validation:XValidation:rule="self.version == oldSelf.version",message="Version is immutable and cannot be unset"
type PostgresDatabaseParameters struct {
	// Location specifies in which data center the database will be spawned.
	// +optional
	// +kubebuilder:default:="nine-es34"
	Location meta.LocationName `json:"location"`
	// Version specifies the Postgres version.
	// Needs to match an available Postgres Version.
	// +optional
	// +kubebuilder:default:="17"
	Version PostgresVersion `json:"version,omitempty"`
	// InstanceRef is the instance that contains the database.
	// It is not possible to configure this value manually.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="InstanceReference is immutable"
	// +optional
	InstanceReference *meta.Reference `json:"instanceRef,omitempty"`
}

// A PostgresDatabaseStatus represents the observed state of a PostgresDatabase.
type PostgresDatabaseStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               PostgresDatabaseObservation `json:"atProvider"`
}

// PostgresDatabaseObservation are the observable fields of a PostgresDatabase.
type PostgresDatabaseObservation struct {
	// FQDN is the fully qualified domain name at which the database is accessible.
	// +optional
	FQDN string `json:"fqdn,omitempty"`
	// Name of the database and user.
	// +optional
	Name string `json:"name,omitempty"`
	// Locked reports why a database has been locked.
	// +optional
	Locked string `json:"locked,omitempty"`
	// Database details
	DatabaseObservation `json:",inline"`
	// CACert is the certificate of the CA that is used by the service.
	// The value is a base64 encoded PEM.
	CACert string `json:"caCert,omitempty"`
	// Status of all child resources.
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
