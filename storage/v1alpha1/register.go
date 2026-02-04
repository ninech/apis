package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "storage.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	BucketKind             = reflect.TypeOf(Bucket{}).Name()
	BucketGroupKind        = schema.GroupKind{Group: Group, Kind: BucketKind}.String()
	BucketKindAPIVersion   = BucketKind + "." + SchemeGroupVersion.String()
	BucketGroupVersionKind = SchemeGroupVersion.WithKind(BucketKind)
	BucketUserKind             = reflect.TypeOf(BucketUser{}).Name()
	BucketUserGroupKind        = schema.GroupKind{Group: Group, Kind: BucketUserKind}.String()
	BucketUserKindAPIVersion   = BucketUserKind + "." + SchemeGroupVersion.String()
	BucketUserGroupVersionKind = SchemeGroupVersion.WithKind(BucketUserKind)
	BucketMigrationKind             = reflect.TypeOf(BucketMigration{}).Name()
	BucketMigrationGroupKind        = schema.GroupKind{Group: Group, Kind: BucketMigrationKind}.String()
	BucketMigrationKindAPIVersion   = BucketMigrationKind + "." + SchemeGroupVersion.String()
	BucketMigrationGroupVersionKind = SchemeGroupVersion.WithKind(BucketMigrationKind)
	DatabaseBackupKind             = reflect.TypeOf(DatabaseBackup{}).Name()
	DatabaseBackupGroupKind        = schema.GroupKind{Group: Group, Kind: DatabaseBackupKind}.String()
	DatabaseBackupKindAPIVersion   = DatabaseBackupKind + "." + SchemeGroupVersion.String()
	DatabaseBackupGroupVersionKind = SchemeGroupVersion.WithKind(DatabaseBackupKind)
	DatabaseBackupScheduleKind             = reflect.TypeOf(DatabaseBackupSchedule{}).Name()
	DatabaseBackupScheduleGroupKind        = schema.GroupKind{Group: Group, Kind: DatabaseBackupScheduleKind}.String()
	DatabaseBackupScheduleKindAPIVersion   = DatabaseBackupScheduleKind + "." + SchemeGroupVersion.String()
	DatabaseBackupScheduleGroupVersionKind = SchemeGroupVersion.WithKind(DatabaseBackupScheduleKind)
	KeyValueStoreKind             = reflect.TypeOf(KeyValueStore{}).Name()
	KeyValueStoreGroupKind        = schema.GroupKind{Group: Group, Kind: KeyValueStoreKind}.String()
	KeyValueStoreKindAPIVersion   = KeyValueStoreKind + "." + SchemeGroupVersion.String()
	KeyValueStoreGroupVersionKind = SchemeGroupVersion.WithKind(KeyValueStoreKind)
	MySQLKind             = reflect.TypeOf(MySQL{}).Name()
	MySQLGroupKind        = schema.GroupKind{Group: Group, Kind: MySQLKind}.String()
	MySQLKindAPIVersion   = MySQLKind + "." + SchemeGroupVersion.String()
	MySQLGroupVersionKind = SchemeGroupVersion.WithKind(MySQLKind)
	MySQLDatabaseKind             = reflect.TypeOf(MySQLDatabase{}).Name()
	MySQLDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: MySQLDatabaseKind}.String()
	MySQLDatabaseKindAPIVersion   = MySQLDatabaseKind + "." + SchemeGroupVersion.String()
	MySQLDatabaseGroupVersionKind = SchemeGroupVersion.WithKind(MySQLDatabaseKind)
	OpenSearchKind             = reflect.TypeOf(OpenSearch{}).Name()
	OpenSearchGroupKind        = schema.GroupKind{Group: Group, Kind: OpenSearchKind}.String()
	OpenSearchKindAPIVersion   = OpenSearchKind + "." + SchemeGroupVersion.String()
	OpenSearchGroupVersionKind = SchemeGroupVersion.WithKind(OpenSearchKind)
	PostgresKind             = reflect.TypeOf(Postgres{}).Name()
	PostgresGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresKind}.String()
	PostgresKindAPIVersion   = PostgresKind + "." + SchemeGroupVersion.String()
	PostgresGroupVersionKind = SchemeGroupVersion.WithKind(PostgresKind)
	PostgresDatabaseKind             = reflect.TypeOf(PostgresDatabase{}).Name()
	PostgresDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresDatabaseKind}.String()
	PostgresDatabaseKindAPIVersion   = PostgresDatabaseKind + "." + SchemeGroupVersion.String()
	PostgresDatabaseGroupVersionKind = SchemeGroupVersion.WithKind(PostgresDatabaseKind)
	RegistryKind             = reflect.TypeOf(Registry{}).Name()
	RegistryGroupKind        = schema.GroupKind{Group: Group, Kind: RegistryKind}.String()
	RegistryKindAPIVersion   = RegistryKind + "." + SchemeGroupVersion.String()
	RegistryGroupVersionKind = SchemeGroupVersion.WithKind(RegistryKind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
	SchemeBuilder.Register(&BucketUser{}, &BucketUserList{})
	SchemeBuilder.Register(&BucketMigration{}, &BucketMigrationList{})
	SchemeBuilder.Register(&DatabaseBackup{}, &DatabaseBackupList{})
	SchemeBuilder.Register(&DatabaseBackupSchedule{}, &DatabaseBackupScheduleList{})
	SchemeBuilder.Register(&KeyValueStore{}, &KeyValueStoreList{})
	SchemeBuilder.Register(&MySQL{}, &MySQLList{})
	SchemeBuilder.Register(&MySQLDatabase{}, &MySQLDatabaseList{})
	SchemeBuilder.Register(&OpenSearch{}, &OpenSearchList{})
	SchemeBuilder.Register(&Postgres{}, &PostgresList{})
	SchemeBuilder.Register(&PostgresDatabase{}, &PostgresDatabaseList{})
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
