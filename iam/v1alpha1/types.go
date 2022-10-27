package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A KubernetesClustersRoleBinding binds a role to subjects and KubernetesClusters.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=kcrb
// +kubebuilder:object:root=true
type KubernetesClustersRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClustersRoleBindingSpec   `json:"spec"`
	Status            KubernetesClustersRoleBindingStatus `json:"status,omitempty"`
}

// KubernetesClustersRoleBindingList contains a list of KubernetesClustersRoleBindings.
// +kubebuilder:object:root=true
type KubernetesClustersRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesClustersRoleBinding `json:"items"`
}

// KubernetesClustersRoleBindingSpec defines the desired state of a KubernetesClustersRoleBinding.
type KubernetesClustersRoleBindingSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            KubernetesClustersRoleBindingParameters `json:"forProvider"`
}

// KubernetesClustersRoleBindingParameters are the parameters of a KubernetesClustersRoleBinding.
type KubernetesClustersRoleBindingParameters struct {
	// Subjects define who gets access to the specified clusters.
	Subjects []RoleBindingSubject `json:"subjects"`
	// Clusters references the KubernetesClusters to which the subjects will
	// be given access to.
	Clusters []meta.LocalReference `json:"clusters"`
	// Role defines the role that the subjects will be given on the specified clusters.
	// +optional
	// +kubebuilder:default:=user
	Role KubernetesClusterRole `json:"role,omitempty"`
}
type RoleBindingSubject struct {
	// Kind of the object being referenced.
	Kind                RoleBindingKind `json:"kind"`
	meta.LocalReference `json:",inline"`
}

// +kubebuilder:validation:Enum:=User;Group;ServiceAccount
type RoleBindingKind string

// KubernetesClusterRole are the permissions for KubernetesClusters.
// See https://support.nine.ch/articles/users-permissions for more information.
// +kubebuilder:validation:Enum:=admin;viewer;user;argocd;nine-admin;nine-viewer
type KubernetesClusterRole string

// KubernetesClustersRoleBindingStatus represents the observed state of a KubernetesClustersRoleBinding.
type KubernetesClustersRoleBindingStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               KubernetesClustersRoleBindingObservation `json:"atProvider,omitempty"`
}
type KubernetesClustersRoleBindingObservation struct {
	// +optional
	ReferenceErrors meta.ReferenceErrors `json:"referenceErrors,omitempty"`
}

// A KubernetesServiceAccount is a service account to access KubernetesClusters.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=ksa
// +kubebuilder:object:root=true
type KubernetesServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesServiceAccountSpec   `json:"spec"`
	Status            KubernetesServiceAccountStatus `json:"status,omitempty"`
}

// KubernetesServiceAccountList contains a list of KubernetesServiceAccounts.
// +kubebuilder:object:root=true
type KubernetesServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesServiceAccount `json:"items"`
}

// KubernetesServiceAccountSpec defines the desired state of a KubernetesServiceAccount.
type KubernetesServiceAccountSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            KubernetesServiceAccountParameters `json:"forProvider"`
}

// KubernetesServiceAccountParameters are the parameters of a KubernetesServiceAccount.
type KubernetesServiceAccountParameters struct {
	// Cluster references the KubernetesCluster to which the
	// service account will be scoped to.
	Cluster meta.LocalReference `json:"cluster"`
}

// KubernetesServiceAccountStatus represents the observed state of a KubernetesServiceAccount.
type KubernetesServiceAccountStatus struct {
	AtProvider               KubernetesServiceAccountObservation `json:"atProvider,omitempty"`
	runtimev1.ResourceStatus `json:",inline"`
}

// KubernetesServiceAccountObservation are the observable fields of a KubernetesServiceAccount.
type KubernetesServiceAccountObservation struct {
	// Name of the service account on the cluster.
	Name string `json:"name"`
	// Namespace of the service account on the cluster.
	Namespace string `json:"namespace"`
	// FullName is the full username of the service account on the cluster.
	FullName string `json:"fullName"`
}
