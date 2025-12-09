package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	APIServiceAccountTokenKey      = "token"
	APIServiceAccountURLKey        = "url"
	APIServiceAccountKubeconfigKey = "kubeconfig"
	APIServiceAccountIDKey         = "client_id"
	APIServiceAccountSecretKey     = "client_secret"
	APIServiceAccountTokenURLKey   = "token_url"
	// APIServiceAccountV1 creates a service account based on Kubernetes service
	// account tokens. This is deprecated as all v1 service accounts will expire
	// on the 01.12.2025.
	APIServiceAccountV1 APIServiceAccountVersion = "v1"
	// APIServiceAccountV2 creates a service account based on oauth2 client
	// credentials.
	APIServiceAccountV2 APIServiceAccountVersion = "v2"
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
	// OrganizationAccess defines if the APIServiceAccount should have
	// access to all created projects with the defined role. This is only
	// possible for APIServiceAccounts which have been created in the
	// organizations default project.
	// +optional
	// +kubebuilder:default:=false
	OrganizationAccess bool `json:"organizationAccess"`
	// Version of the APIServiceAccount.
	// +optional
	// +kubebuilder:default:="v2"
	// +kubebuilder:validation:XValidation:message="downgrade from v2 to v1 is not allowed",rule="!(self == 'v1' && oldSelf == 'v2')"
	Version APIServiceAccountVersion `json:"version,omitempty"`
}

// +kubebuilder:validation:Enum:=admin;viewer;metrics-admin;metrics-viewer;internal-metrics;internal-ipaddress
type APIServiceAccountRole string

// +kubebuilder:validation:Enum:=v1;v2
type APIServiceAccountVersion string

// APIServiceAccountStatus represents the observed state of an APIServiceAccount.
type APIServiceAccountStatus struct {
	AtProvider               APIServiceAccountObservation `json:"atProvider"`
	runtimev1.ResourceStatus `json:",inline"`
}

// APIServiceAccountObservation are the observable fields of an APIServiceAccount.
type APIServiceAccountObservation struct {
	// Email assigned to this service account. Not a real email address but it's
	// used to identify this service account in RBAC.
	// +optional
	Email string `json:"email,omitempty"`
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
