package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIServiceAccount is a service account to access the API.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=asa
// +kubebuilder:object:root=true
type APIServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIServiceAccountSpec   `json:"spec"`
	Status            APIServiceAccountStatus `json:"status,omitempty"`
}

// APIServiceAccountList contains a list of APIServiceAccounts
// +kubebuilder:object:root=true
type APIServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIServiceAccount `json:"items"`
}

// APIServiceAccountSpec defines the desired state of a APIServiceAccount.
type APIServiceAccountSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            APIServiceAccountParameters `json:"forProvider"`
}

// APIServiceAccountParameters are the configurable fields of an APIServiceAccount.
type APIServiceAccountParameters struct {
	// A predefined Role the service account will get.
	// +optional
	// +kubebuilder:default:="admin"
	Role APIServiceAccountRole `json:"role,omitempty"`
}

// +kubebuilder:validation:Enum:=admin;viewer;metrics-admin;internal-metrics
type APIServiceAccountRole string

// APIServiceAccountStatus represents the observed state of a APIServiceAccount.
type APIServiceAccountStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
}

// KubernetesClustersRoleBinding binds a role to subjects and KubernetesClusters.
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

// KubernetesClusterRole are the permissions for KubernetesClusters. See
// https://docs.nine.ch/a/b25bM5HIh5M/#rbac for more information.
// +kubebuilder:validation:Enum:=admin;viewer;user;argocd;nine-admin;nine-viewer
type KubernetesClusterRole string

// KubernetesClustersRoleBindingStatus represents the observed state of a KubernetesClustersRoleBinding.
type KubernetesClustersRoleBindingStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               KubernetesClustersRoleBindingObservation `json:"atProvider,omitempty"`
}
type KubernetesClustersRoleBindingObservation struct {
	meta.ReferenceStatus `json:",inline"`
}

// KubernetesServiceAccount is a service account to access a KubernetesCluster.
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
	Name string `json:"name,omitempty"`
	// Namespace of the service account on the cluster.
	Namespace string `json:"namespace,omitempty"`
	// FullName is the full username of the service account on the cluster.
	FullName             string `json:"fullName,omitempty"`
	meta.ReferenceStatus `json:",inline"`
}
