package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArgoCD deploys a fully managed Argo CD instance.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CLUSTER CONNECTION ERROR",type="string",JSONPath=".status.atProvider.clusterConnectionError"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type ArgoCD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArgoCDSpec   `json:"spec"`
	Status            ArgoCDStatus `json:"status,omitempty"`
}

// ArgoCDList contains a list of ArgoCDs
// +kubebuilder:object:root=true
type ArgoCDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoCD `json:"items"`
}

// An ArgoCDSpec defines the desired state of a ArgoCD.
type ArgoCDSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ArgoCDParameters `json:"forProvider"`
}

// ArgoCDParameters are the configurable fields of a ArgoCD.
type ArgoCDParameters struct {
	// Clusters that Argo CD has access to.
	// +kubebuilder:validation:MinItems:=1
	Clusters []meta.Reference `json:"clusters"`
	// EnableApplicationSets specifies if Argo CD should support
	// ApplicationSet CRDs
	// +optional
	// +kubebuilder:default:=false
	EnableApplicationSets bool `json:"enableApplicationSets"`
	// CustomSecrets allows to pass secrets to ArgoCD which can then be
	// used (e.g. in ApplicationSets)
	// +optional
	CustomSecret *meta.LocalReference `json:"customSecret,omitempty"`
}

// An ArgoCDStatus represents the observed state of a ArgoCD.
type ArgoCDStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ArgoCDObservation `json:"atProvider"`
}

// ArgoCDObservation are the observable fields of a ArgoCD.
type ArgoCDObservation struct {
	// URL points to the web UI of Argo CD.
	// +optional
	URL string `json:"url,omitempty"`
	// CLIURL points to the URL used for logging into the CLI.
	// +optional
	CLIURL string `json:"cliURL,omitempty"`
	// Webhooks outputs the URLs for configured webhooks
	// +optional
	Webhooks *ArgoCDWebhooks `json:"webhooks,omitempty"`
	// CustomSecret is the name of the secret which can be used in ArgoCD
	// (e.g. in ApplicationSets). It contains the same keys as the
	// referenced custom secret.
	// +optional
	CustomSecret *meta.Reference `json:"customSecret,omitempty"`
	// ClusterConnectionError is an error that occurs if any of the references
	// point to a cluster that does not exist.
	ClusterConnectionError   string `json:"clusterConnectionError,omitempty"`
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}
type ArgoCDWebhooks struct {
	// ArgoCD specifies the default webhook used to notify ArgoCD about git
	// repository changes
	//+optional
	ArgoCD string `json:"argoCD,omitempty"`
	// ApplicationSet specifies the special application set webhook used to
	// notify the ArgoCD ApplicationSet generators about changes in git
	//+optional
	ApplicationSet string `json:"applicationSet,omitempty"`
}
