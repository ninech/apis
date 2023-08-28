package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "security.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	ExternalSecretsKind             = reflect.TypeOf(ExternalSecrets{}).Name()
	ExternalSecretsGroupKind        = schema.GroupKind{Group: Group, Kind: ExternalSecretsKind}.String()
	ExternalSecretsKindAPIVersion   = ExternalSecretsKind + "." + SchemeGroupVersion.String()
	ExternalSecretsGroupVersionKind = SchemeGroupVersion.WithKind(ExternalSecretsKind)
	SealedSecretsKind             = reflect.TypeOf(SealedSecrets{}).Name()
	SealedSecretsGroupKind        = schema.GroupKind{Group: Group, Kind: SealedSecretsKind}.String()
	SealedSecretsKindAPIVersion   = SealedSecretsKind + "." + SchemeGroupVersion.String()
	SealedSecretsGroupVersionKind = SchemeGroupVersion.WithKind(SealedSecretsKind)
	SSHKeyKind             = reflect.TypeOf(SSHKey{}).Name()
	SSHKeyGroupKind        = schema.GroupKind{Group: Group, Kind: SSHKeyKind}.String()
	SSHKeyKindAPIVersion   = SSHKeyKind + "." + SchemeGroupVersion.String()
	SSHKeyGroupVersionKind = SchemeGroupVersion.WithKind(SSHKeyKind)
)

func init() {
	SchemeBuilder.Register(&ExternalSecrets{}, &ExternalSecretsList{})
	SchemeBuilder.Register(&SealedSecrets{}, &SealedSecretsList{})
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
