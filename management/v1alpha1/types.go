package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A Project represents a project/environment of a organization.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// ProjectList contains a list of Projects
// +kubebuilder:object:root=true
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// A ProjectSpec defines the desired state of a Project.
type ProjectSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	// DisplayName is the name that, if specified, is displayed in the UI.
	// DisplayName is a user-friendly Project name. DisplayName is not unique.
	// DisplayName may differ from the Project name.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
}

// A ProjectStatus represents the observed state of a Project.
type ProjectStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ProjectObservation `json:"atProvider,omitempty"`
}

// ProjectObservation are the observable fields of an Project.
type ProjectObservation struct {
	// Status of all our child resources.
	meta.ChildResourceStatus `json:",inline"`
}
