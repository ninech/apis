package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IngressNginx deploys an NGINX ingress controller to a cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=nginx
// +kubebuilder:object:root=true
type IngressNginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressNginxSpec   `json:"spec"`
	Status            IngressNginxStatus `json:"status,omitempty"`
}

// IngressNginxList contains a list of IngressNginx.
// +kubebuilder:object:root=true
type IngressNginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressNginx `json:"items"`
}

// IngressNginxSpec defines the desired state of an IngressNginx.
type IngressNginxSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            IngressNginxParameters `json:"forProvider"`
}

// IngressNginxParameters are the configurable fields of a IngressNginx.
type IngressNginxParameters struct {
	// Cluster specifies on which cluster this IngressNginx should be installed.
	Cluster meta.LocalReference `json:"cluster"`
	// Cache defines caching settings. If set, the cache will be enabled and
	// use the default settings. If not set, the cache will not be enabled.
	// Caching also needs to be enabled on the ingress resource by setting a
	// server-snippet annotation like this:
	//
	//   nginx.ingress.kubernetes.io/server-snippet: |
	//     proxy_cache mycache;
	//     proxy_cache_lock on;
	//     proxy_cache_valid any 60m;
	//     proxy_ignore_headers Cache-Control;
	//     add_header X-Cache-Status $upstream_cache_status;
	//
	// +optional
	Cache *IngressNginxCache `json:"cache,omitempty"`
	// Cloudflare defines the settings for having Cloudflare as a CDN in
	// front of this Ingress Controller.
	// Do NOT use this option in combination with the `AppendToXForwardedFor`
	// variable as this will break client IP detection.
	// Disabled if not set.
	// +optional
	Cloudflare *IngressNginxCloudflare `json:"cloudflare,omitempty"`
	// AppendToXForwardedFor defines if the ingress should take the incoming
	// X-Forwarded-For header and append IPs to it, instead of replacing the
	// whole header.
	// Do NOT use this option in combination with Cloudflare enabled
	// as this will break client IP detection.
	// +optional
	// +kubebuilder:default:=false
	AppendToXForwardedFor bool `json:"appendToXForwardedFor,omitempty"`
	// IngressClass sets the class of the ingress controller. Specifies which
	// ingress objects the controller should take care of. Please note that
	// the class `nginx` itself will handle all ingresses without a class
	// annotation set.
	// +optional
	// +kubebuilder:default:=nginx
	IngressClass string `json:"ingressClass,omitempty"`
	// IsDefaultIngressClass specifies if the IngressClass of this controller
	// should be the default IngressClass in the target cluster.
	// +optional
	// +kubebuilder:default:=false
	IsDefaultIngressClass bool `json:"isDefaultIngressClass"`
	// DefaultBackend sets the default backend that the ingress will proxy to
	// if the request does not match any configured ingress route. If
	// disabled, nginx will just return a generic 404 for such requests.
	// +optional
	DefaultBackend *IngressNginxDefaultBackend `json:"defaultBackend,omitempty"`
	// IPSharing enables sharing the IP address of the ingress LB with an
	// existing service type LoadBalancer. The reference should point to the
	// service with which the LB IP should be shared. The referenced service
	// has to have the annotation metallb.universe.tf/allow-shared-ip set and
	// the ports 80 and 443 cannot be already in use.
	// +optional
	IPSharing *meta.Reference `json:"ipSharing,omitempty"`
	// ScrapeConfigurations allows to overwrite which metrics of the ingress-nginx
	// instance are scraped by certain Prometheus instances.
	// +optional
	ScrapeConfigurations []meta.ScrapeConfig `json:"scrapeConfigurations,omitempty"`
	// SSLPassthrough flag enables the SSL Passthrough feature, which is
	// disabled by default. This is required to enable passthrough backends in
	// ingress objects. This feature is implemented by intercepting all
	// traffic on the configured HTTPS port (default: 443) and handing it over
	// to a local TCP proxy. This bypasses NGINX completely and introduces a
	// non-negligible performance penalty.
	// +optional
	// +kubebuilder:default:=false
	SSLPassthrough bool `json:"sslPassthrough,omitempty"`
	// DisableSnippetAnnotations disables all "*-snippet" annotations (like
	// "configuration-snippet" or "server-snippet") on ingress resources.
	// These annotations are potentially dangerous in an environment where
	// you do not trust all your users (multitenant environments). For
	// example, it might allow to see the generated nginx.conf file.
	// +optional
	// +kubebuilder:default:=false
	DisableSnippetAnnotations bool `json:"disableSnippetAnnotations,omitempty"`
	// AnnotationValueWordBlocklist is a list of comma separated words
	// which, if found in snippet annotations, will block the acceptance of
	// the ingress resource. This allows to block nginx configuration stanzas
	// which are potentially dangerous in a multitenant use case of the
	// ingress controller while still allowing snippet annotations to be
	// made.
	//
	// The suggested list of the ingress-nginx project is:
	// "load_module,lua_package,_by_lua,location,root,proxy_pass,serviceaccount,{,},',\""
	//
	// It is still recommended to disable snippet annotations in general
	// (see disableSnippetAnnotations) when using this ingress-nginx in a
	// multitenant scenario.
	// +optional
	AnnotationValueWordBlocklist string `json:"annotationValueWordBlocklist,omitempty"`
	// EnableModSecurity enables the owasp-modsecurity. Note this will enable
	// it for all paths, and each path must be disabled manually. ModSecurity
	// will run in "Detection-Only" mode using the recommended configuration.
	// You can enable the OWASP Core Rule Set by setting the following
	// annotation:
	//
	// nginx.ingress.kubernetes.io/enable-owasp-core-rules: "true"
	//
	// https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md#modsecurity
	//
	// +optional
	// +kubebuilder:default:=false
	EnableModSecurity bool `json:"enableModSecurity,omitempty"`
	// HSTS allows to configure settings for the "HTTP Strict Transport
	// Security" (HSTS) headers
	HSTS *HSTSConfiguration `json:"hsts,omitempty"`
	// DisableHTTP2 disables HTTP2 for all hosts.
	// +optional
	// +kubebuilder:default:=false
	DisableHTTP2 bool `json:"disableHTTP2"`
	// RateLimitStatusCode configures the status code nginx will return when
	// rate limits are hit.
	// https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#limit-req-status-code
	// +optional
	// +kubebuilder:default:=503
	RateLimitStatusCode int `json:"rateLimitStatusCode,omitempty"`
	// EnableUnderscoresInHeaders allows headers with underscores in the name.
	// +optional
	// +kubebuilder:default:=false
	EnableUnderscoresInHeaders bool `json:"enableUnderscoresInHeaders"`
	// EnableInvalidHeaders allows invalid header names.
	// +optional
	// +kubebuilder:default:=false
	EnableInvalidHeaders bool `json:"enableInvalidHeaders"`
	// TLSProtocols to use.
	// +optional
	TLSProtocols IngressNginxTLSProtocols `json:"tlsProtocols,omitempty"`
}

// IngressNginxCache uses the nginx settings `proxy_cache_<x>` to cache
// content within the ingress controller.
type IngressNginxCache struct {
	// Key defines the key that is used to index the nginx cache.
	// +optional
	// +kubebuilder:default:="$scheme$proxy_host$request_uri"
	Key string `json:"key,omitempty"`
}

// IngressNginxCloudflare sets settings on the ingress-nginx controller to
// work with cloudflare. This means that it allows to read clientIPs out of
// HTTP headers that are added by Cloudflare. Note that all domains used in
// ingress objects should use cloudflare if this is enabled.
type IngressNginxCloudflare struct {
	// IPs is a list of CIDRs which should contain all Cloudflare Networks.
	// +optional
	// +kubebuilder:default:={"173.245.48.0/20","103.21.244.0/22","103.22.200.0/22","103.31.4.0/22","141.101.64.0/18","108.162.192.0/18","190.93.240.0/20","188.114.96.0/20","197.234.240.0/22","198.41.128.0/17","162.158.0.0/15","104.16.0.0/13","104.24.0.0/14","172.64.0.0/13","131.0.72.0/22"}
	IPs []string `json:"ips,omitempty"`
}

// IngressNginxDefaultBackend configuration
type IngressNginxDefaultBackend struct {
	// Image sets the image that is used for the backend pods. These are
	// responsible for returning HTTP error pages in case of a failure in the
	// service that an ingress resource points to.
	// +optional
	// +kubebuilder:default:={registry: "registry.k8s.io", repository: "defaultbackend-amd64", tag: "1.5"}
	Image *meta.Image `json:"defaultBackendImage,omitempty"`
	// Replicas sets the number of replicas for the default backend.
	// +optional
	// +kubebuilder:default:=2
	Replicas int `json:"defaultBackendReplicas,omitempty"`
}

// HSTSConfiguration allows to configure the "HTTP Strict Transport Security"
// headers. HSTS is enabled by default.
type HSTSConfiguration struct {
	// Enabled enables/disables all HSTS headers. If not given all HSTS
	// headers are enabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// IncludeSubdomains allows to toggle the "includeSubDomains" field in
	// the HSTS header (enabled by default).
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty"`
	// MaxAge sets the time that the browser should remember that this site
	// is only to be accessed using HTTPS.
	MaxAge *metav1.Duration `json:"maxAge,omitempty"`
	// Preload allows to toggle the HSTS preload feature.
	Preload *bool `json:"preload,omitempty"`
}

// +kubebuilder:default={TLSv1.2,TLSv1.3}
type IngressNginxTLSProtocols []IngressNginxTLSProtocol

// IngressNginxStatus represents the observed state of an IngressNginx.
type IngressNginxStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               IngressNginxObservation `json:"atProvider"`
}

// IngressNginxObservation are the observable fields of an IngressNginx.
type IngressNginxObservation struct {
	// DNSName is the name of the ingress A-Record, pointing to IPAddress.
	// +optional
	DNSName string `json:"dnsName,omitempty"`
	// IPAddress is the address where the ingress controller is reachable.
	// +optional
	IPAddress string `json:"ipAddress,omitempty"`
	// IPSharingError indicates errors with the IPSharing configuration.
	// +optional
	IPSharingError           string `json:"ipSharingError,omitempty"`
	meta.ChildResourceStatus `json:",inline"`
	meta.ReferenceStatus     `json:",inline"`
}

// IngressNginxTLSProtocol is a SSL/TLS protocol that nginx can use.
// +kubebuilder:validation:Enum=TLSv1.2;TLSv1.3
type IngressNginxTLSProtocol string

// ServiceConnection connects a source to destination securely.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=sc
// +kubebuilder:object:root=true
type ServiceConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceConnectionSpec   `json:"spec"`
	Status            ServiceConnectionStatus `json:"status,omitempty"`
}

// ServiceConnectionList contains a list of ServiceConnection.
// +kubebuilder:object:root=true
type ServiceConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceConnection `json:"items"`
}

// ServiceConnectionSpec defines the desired state of an ServiceConnection.
type ServiceConnectionSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ServiceConnectionParameters `json:"forProvider"`
}

// ServiceConnectionParameters are the configurable fields of a ServiceConnection.
type ServiceConnectionParameters struct {
	// Source is the source of the connection
	Source Source `json:"source"`
	// Destination is a reference to the destination of the connection
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="destination is immutable after creation"
	Destination meta.TypedReference `json:"destination"`
}
type Source struct {
	// Reference is a reference to the source of the connection
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="source reference is immutable after creation"
	Reference meta.TypedReference `json:"reference"`
	// KubernetesClusterOptions can only be set when the source is a
	// KubernetesCluster. They allow for more configuration.
	// +optional
	KubernetesClusterOptions *KubernetesClusterOptions `json:"kubernetesClusterOptions,omitempty"`
}
type KubernetesClusterOptions struct {
	// PodSelector can be used to restrict which pods of the
	// KubernetesCluster can connect to the service connection destination.
	// If left empty, all pods are allowed. If namespaceSelector is also
	// set, then the podSelector as a whole selects the pods matching
	// podSelector in the namespaces selected by namespaceSelector.
	// +optional
	PodSelector metav1.LabelSelector `json:"podSelector"`
	// NamespaceSelector selects namespaces using labels set on namespaces.
	// If left empty, it selects all namespaces. It can be used to further
	// restrict the pods selected by the PodSelector.
	// +optional
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector"`
}

// ServiceConnectionStatus represents the observed state of an ServiceConnection.
type ServiceConnectionStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ServiceConnectionObservation `json:"atProvider"`
}

// ServiceConnectionObservation are the observable fields of an ServiceConnection.
type ServiceConnectionObservation struct {
	// +optional
	// DestinationDNS is the DNS to which this ServiceConnection resource
	// allows to connect.
	DestinationDNS string `json:"destinationDNS,omitempty"`
	// ReferenceStatus contains errors for wrongly referenced resources
	meta.ReferenceStatus `json:",inline"`
}

// StaticEgress describes a static egress configuration
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EGRESS-ADDRESS",type="string",JSONPath=".status.atProvider.address"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type StaticEgress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StaticEgressSpec   `json:"spec"`
	Status            StaticEgressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// StaticEgressList contains a list of StaticEgress resources
type StaticEgressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticEgress `json:"items"`
}

// A StaticEgressSpec defines the desired state of a StaticEgress
type StaticEgressSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            StaticEgressParameters `json:"forProvider"`
}

// StaticEgressParameters are the configurable fields of a StaticEgress
type StaticEgressParameters struct {
	// Disabled can be set to true to quickly disable the static egress
	// without loosing the registered egress IP.
	// egress IP.
	// +optional
	// +kubebuilder:default:=false
	Disabled bool `json:"disabled"`
	// Target specifies for which resource the static egress should be
	// configured
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	Target meta.LocalTypedReference `json:"target"`
}

// A StaticEgressStatus represents the observed state of a StaticEgress
type StaticEgressStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               StaticEgressObservation `json:"atProvider"`
}

// StaticEgressObservation are the observable fields of a static egress
type StaticEgressObservation struct {
	// Address is the egress IP which will be used for all egressing
	// traffic
	// +optional
	Address string `json:"address,omitempty"`
	// SelectorLabel needs to be set on pods in target clusters, so that
	// all egressing traffic of the pod will be emitted with the egress IP
	// address. Only supported by certain targets.
	// +optional
	SelectorLabel        *StaticEgressSelectorLabel `json:"selectionLabel,omitempty"`
	meta.ReferenceStatus `json:",inline"`
}

// StaticEgressSelectorLabel needs to be used to select pods for matching
// static egress traffic
type StaticEgressSelectorLabel struct {
	// Name of the label.
	Name string `json:"name"`
	// Value of the label.
	Value string `json:"value"`
}
