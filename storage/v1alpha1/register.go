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
	ObjectsBucketKind             = reflect.TypeOf(ObjectsBucket{}).Name()
	ObjectsBucketGroupKind        = schema.GroupKind{Group: Group, Kind: ObjectsBucketKind}.String()
	ObjectsBucketKindAPIVersion   = ObjectsBucketKind + "." + SchemeGroupVersion.String()
	ObjectsBucketGroupVersionKind = SchemeGroupVersion.WithKind(ObjectsBucketKind)
	PostgresKind             = reflect.TypeOf(Postgres{}).Name()
	PostgresGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresKind}.String()
	PostgresKindAPIVersion   = PostgresKind + "." + SchemeGroupVersion.String()
	PostgresGroupVersionKind = SchemeGroupVersion.WithKind(PostgresKind)
	RedisKind             = reflect.TypeOf(Redis{}).Name()
	RedisGroupKind        = schema.GroupKind{Group: Group, Kind: RedisKind}.String()
	RedisKindAPIVersion   = RedisKind + "." + SchemeGroupVersion.String()
	RedisGroupVersionKind = SchemeGroupVersion.WithKind(RedisKind)
	RegistryKind             = reflect.TypeOf(Registry{}).Name()
	RegistryGroupKind        = schema.GroupKind{Group: Group, Kind: RegistryKind}.String()
	RegistryKindAPIVersion   = RegistryKind + "." + SchemeGroupVersion.String()
	RegistryGroupVersionKind = SchemeGroupVersion.WithKind(RegistryKind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
	SchemeBuilder.Register(&BucketUser{}, &BucketUserList{})
	SchemeBuilder.Register(&ObjectsBucket{}, &ObjectsBucketList{})
	SchemeBuilder.Register(&Postgres{}, &PostgresList{})
	SchemeBuilder.Register(&Redis{}, &RedisList{})
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
