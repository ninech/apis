package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExternalSecrets deploys the ExternalSecrets controller to a cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type ExternalSecrets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalSecretsSpec   `json:"spec"`
	Status            ExternalSecretsStatus `json:"status,omitempty"`
}

// ExternalSecretsList contains a list of ExternalSecrets instances.
// +kubebuilder:object:root=true
type ExternalSecretsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalSecrets `json:"items"`
}

// ExternalSecretsSpec defines the desired state of a ExternalSecrets instance.
type ExternalSecretsSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ExternalSecretsParameters `json:"forProvider"`
}

// ExternalSecretsParameters are the configurable fields of a ExternalSecrets instance.
type ExternalSecretsParameters struct {
	// Cluster is the cluster where the external-secrets controller should be deployed.
	Cluster meta.LocalReference `json:"cluster"`
}

// ExternalSecretsStatus represents the observed state of a ExternalSecrets instance.
type ExternalSecretsStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ExternalSecretsObservation `json:"atProvider"`
}

// ExternalSecretsObservation are the observable fields of a ExternalSecrets instance.
type ExternalSecretsObservation struct {
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}

// SealedSecrets deploys the SealedSecrets controller to a cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type SealedSecrets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SealedSecretsSpec   `json:"spec"`
	Status            SealedSecretsStatus `json:"status,omitempty"`
}

// SealedSecretsList contains a list of SealedSecrets instances.
// +kubebuilder:object:root=true
type SealedSecretsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SealedSecrets `json:"items"`
}

// SealedSecretsSpec defines the desired state of a SealedSecrets instance.
type SealedSecretsSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            SealedSecretsParameters `json:"forProvider"`
}

// SealedSecretsParameters are the configurable fields of a SealedSecrets instance.
type SealedSecretsParameters struct {
	// Cluster is the cluster where the sealed-secrets controller should be deployed.
	Cluster meta.LocalReference `json:"cluster"`
}

// SealedSecretsStatus represents the observed state of a SealedSecrets instance.
type SealedSecretsStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               SealedSecretsObservation `json:"atProvider"`
}

// SealedSecretsObservation are the observable fields of a SealedSecrets instance.
type SealedSecretsObservation struct {
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}
