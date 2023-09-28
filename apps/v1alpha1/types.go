package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

const (
	// BuildProcessStatusError represents an unknown build status
	BuildProcessStatusUnknown BuildProcessStatus = "unknown"
	// BuildProcessStatusError represents the status of a failed build
	BuildProcessStatusError BuildProcessStatus = "error"
	// BuildProcessStatusRunning represents the status of a running build
	BuildProcessStatusRunning BuildProcessStatus = "running"
	// BuildProcessStatusSuccess represents the status of a successful/finished build
	BuildProcessStatusSuccess BuildProcessStatus = "success"
	// The following config origins describe all possible sources where a
	// configuration can be given
	ConfigOriginDefault      = "default"
	ConfigOriginApplication  = "application"
	ConfigOriginProject      = "project"
	ConfigOriginOrganization = "organization"
	ConfigOriginGit          = "git"
	// ReleaseProcessStatusReplicaFailure is added in a Release when for the
	// collection of underlying resources (Deployment, Service, Ingress, Secret,
	// etc), one of its pods fails to be created or deleted.
	ReleaseProcessStatusReplicaFailure ReleaseProcessStatus = "replicaFailure"
	// ReleaseProcessStatusProgressing means the Release is progressing.
	// Progress for a Release is considered when for the collection of
	// underlying resources (Deployment, Service, Ingress, Secret, etc), a new
	// replica set is created or adopted, and when new pods scale up or old pods
	// scale down.
	ReleaseProcessStatusProgressing ReleaseProcessStatus = "progressing"
	// ReleaseProcessStatusAvailable means the Release is available, ie. for the
	// collection of underlying resources (Deployment, Service, Ingress, Secret,
	// etc), at least the minimum available replicas required are up and running
	// for at least minReadySeconds.
	ReleaseProcessStatusAvailable ReleaseProcessStatus = "available"
	// ReleaseProcessStatusSuperseded means that this release has been
	// superseded by a later one.
	ReleaseProcessStatusSuperseded ReleaseProcessStatus = "superseded"
	// ReleaseProcessStatusFailure indicates the release has failed.
	ReleaseProcessStatusFailure ReleaseProcessStatus = "failure"
	// ReplicaStatusReady is the status for a ready running replica.
	ReplicaStatusReady ReplicaStatus = "ready"
	// RepicaStatusProgressing is the status for a currently progressing
	// replica.
	ReplicaStatusProgressing ReplicaStatus = "progressing"
	// ReplicaStatusFailing describes a replica which is not in a ready state
	// and has restarted more than 2 times.
	ReplicaStatusFailing ReplicaStatus = "failing"
	// ReplicaStatusWaiting describes a state where a replica is waiting for
	// an event to happen. For example a currently running deploy job.
	ReplicaStatusWaiting ReplicaStatus = "waiting"
	// DeployJobProcessStatusSucceeded indicates that the deploy job has
	// succeeded.
	DeployJobProcessStatusSucceeded DeployJobProcessStatus = "succeeded"
	// DeployJobProcessStatusRunning indicates that the deploy job is
	// currently running.
	DeployJobProcessStatusRunning DeployJobProcessStatus = "running"
	// DeployJobProcessStatusFailed indicates that the deploy job was unable
	// to successfully complete within the configured amount of retries.
	DeployJobProcessStatusFailed DeployJobProcessStatus = "failed"
	// DeployJobProcessStatusUnknown indicates the status is unknown.
	DeployJobProcessStatusUnknown DeployJobProcessStatus = "unknown"
	DeployJobProcessReasonTimeout DeployJobProcessReason = "timeout"
	DeployJobProcessReasonBackoff DeployJobProcessReason = "backoffLimitExceeded"
)

var (
	// DefaultConfig defines the default values used for deplo.io
	// applications
	DefaultConfig                 = Config{Size: AppMicro, Replicas: pointer.Int32(1), Port: pointer.Int32(8080), EnableBasicAuth: pointer.Bool(false)}
	AppMicro      ApplicationSize = "micro"
	AppMini       ApplicationSize = "mini"
	AppStandard1  ApplicationSize = "standard-1"
	AppStandard2  ApplicationSize = "standard-2"
)

// Application takes source code as the input and fully builds and deploys
// the application.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=app
// +kubebuilder:object:root=true
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// ApplicationList contains a list of Applications
// +kubebuilder:object:root=true
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// An ApplicationSpec defines the desired state of an Application.
type ApplicationSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ApplicationParameters `json:"forProvider"`
}

// ApplicationParameters are the configurable fields of a Application.
type ApplicationParameters struct {
	Git    ApplicationGitConfig `json:"git"`
	Config Config               `json:"config"`
	// Hosts is a list of host names where the application can be accessed. If
	// empty, the application will just be accessible on a generated host name
	// on the deploio.app domain.
	// +optional
	Hosts []string `json:"hosts,omitempty"`
	// Env variables which are passed to configure env variables required during
	// the build process.
	// +optional
	BuildEnv EnvVars `json:"buildEnv"`
}

// ApplicationGitConfig configures the git repo to connect to.
type ApplicationGitConfig struct {
	// GitTarget specifies the git repository target
	GitTarget `json:",inline"`
	// Auth configures the authentication to a secured git repository.
	// Can be omitted for publicly accessible git repositories.
	// +optional
	Auth *GitAuth `json:"auth,omitempty"`
}
type GitTarget struct {
	// URL is the URL to the Git repository containing the application source.
	// Both HTTPS and SSH formats are supported.
	// +kubebuilder:validation:MinLength=1
	URL string `json:"url"`
	// SubPath is a path in the git repo which contains the application
	// code. If not given, the root directory of the git repo will be used.
	// +optional
	SubPath string `json:"subPath"`
	// Revision defines the revision of the source to deploy the application
	// to. This can be a commit, tag, or branch.
	// +kubebuilder:validation:MinLength=1
	Revision string `json:"revision"`
}
type GitAuth struct {
	// Username is the username to use when connecting to the git repository
	// over HTTPS.
	// +optional
	Username string `json:"username,omitempty"`
	// Password is the password to use when connecting to the git repository
	// over HTTPS. In case of GitHub or GitLab, this can also be an access
	// token.
	// +optional
	Password string `json:"password,omitempty"`
	// SSHPrivateKey is a private key in PEM format to connect to the git
	// repository via SSH.
	// +optional
	SSHPrivateKey string `json:"sshPrivateKey,omitempty"`
	// FromSecret is a reference to a Secret to read the credentials from
	// instead of using the inline fields. Should contain the following keys
	// depending on the protocol used.
	//
	// HTTPS:
	// data:
	//   username: <username>
	//   password: <password>
	// SSH:
	// data:
	//   privatekey: <pem-private-key>
	//
	// +optional
	FromSecret *meta.LocalReference `json:"fromSecret,omitempty"`
}

// Config contains all parameters configuring the deployment of an Application.
type Config struct {
	// +optional
	// +kubebuilder:default:=""
	Size ApplicationSize `json:"size" yaml:"size"`
	// Env variables which are passed to the app at runtime.
	// +optional
	Env EnvVars `json:"env" yaml:"env"`
	// Port the app is listening on.
	// +optional
	// +nullable
	Port *int32 `json:"port" yaml:"port"`
	// Replicas sets the amount of replicas of the running app. If this is
	// increased, make sure your application can cope with running multiple
	// replicas and all state required is shared in some way.
	// +optional
	// +nullable
	Replicas *int32 `json:"replicas" yaml:"replicas"`
	// EnableBasicAuth enables basic authentication for the application
	// +optional
	// +nullable
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" yaml:"enableBasicAuth,omitempty"`
	// +optional
	// +nullable
	DeployJob *DeployJob `json:"deployJob,omitempty"`
}

// ApplicationSize defines the size of an application and the resources that
// will be allocated for it.
// +kubebuilder:validation:Enum:="";micro;mini;standard-1;standard-2
type ApplicationSize string
type EnvVars []EnvVar

// DeployJob defines a command which is executed before a new release gets
// deployed. The deployment will only continue if the job finished
// successfully.
type DeployJob struct {
	Job       `json:",inline"`
	FiniteJob `json:",inline"`
}

// Job defines fields which all jobs have in common
type Job struct {
	// Name of the Job.
	Name string `json:"name"`
	// Command to execute. This command will be executed by a bash shell which
	// gets started by the cloud native buildpacks launcher binary.
	Command string `json:"command"`
}

// FiniteJob defines fields for all jobs which have a finite runtime
type FiniteJob struct {
	// Retries defines how many times the job will be restarted on failure.
	// +optional
	// +kubebuilder:default:=0
	// +kubebuilder:validation:max:=5
	Retries *int32 `json:"retries,omitempty"`
	// Timeout of the job. Minimum is 1 minute and maximum is 30 minutes.
	// +optional
	// +kubebuilder:default:="5m"
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=duration
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// An ApplicationStatus represents the observed state of an Application.
type ApplicationStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ApplicationObservation `json:"atProvider"`
}

// ApplicationObservation are the observable fields of an Application.
type ApplicationObservation struct {
	// DefaultURLs are the URLs at which the application is avalilable.
	// +optional
	DefaultURLs []string `json:"defaultURLs"`
	// CNAMETarget specifies to which DNS entry all custom hosts should point to
	// (via a CNAME entry)
	// +optional
	CNAMETarget string `json:"cnameTarget,omitempty"`
	// TXTRecordContent specifies the content of the DNS TXT record which
	// will be used for host validation
	// +optional
	TXTRecordContent string `json:"txtRecordContent,omitempty"`
	// Hosts represents the latest status of the verification of each
	// custom host.
	// +optional
	Hosts []VerificationStatus `json:"hosts,omitempty"`
	// DefaultHostsCertificateStatus represents the latest Certificate status for the
	// default URLs where the app is available.
	// +optional
	DefaultHostsCertificateStatus CertificateStatus `json:"defaultHostsCertificateStatus,omitempty"`
	// LatestRelease shows the latest release for this application
	// +optional
	LatestRelease string `json:"latestRelease,omitempty"`
	// Size shows the effective application size which is currently in use
	// +optional
	Size ApplicationSize `json:"size,omitempty"`
	// Replicas shows the effective amount of replicas which are currently
	// deployed
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`
	// BasicAuthSecret references the secret which contains the basic auth
	// credentials
	// +optional
	BasicAuthSecret *meta.LocalReference `json:"basicAuthSecret,omitempty"`
	// ReferenceStatus contains errors for wrongly referenced resources
	meta.ReferenceStatus `json:",inline"`
}
type VerificationStatus struct {
	// the hostname of the verification entry
	Name string `json:"name"`
	// CheckType describes which kind of DNS check this entry is about
	// (CNAME or TXT)
	// +optional
	CheckType dnsCheckType `json:"checkType,omitempty"`
	// LatestSuccess specifies when this host was last verified successfully
	// +optional
	LatestSuccess *metav1.Time `json:"latestSuccess,omitempty"`
	// Error displays a potential error which happened during the
	// verification
	// +optional
	Error *VerificationError `json:"error,omitempty"`
}
type dnsCheckType string
type VerificationError struct {
	// Message refers to the error message
	Message string `json:"message"`
	// Timestamp refers to the time when this error happened
	Timestamp metav1.Time `json:"timestamp"`
}

// CertificateStatus represents the Certificate status
type CertificateStatus string

// A Build represents an OCI image build of some referenced source code
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="BUILD-STATUS",type="string",JSONPath=".status.atProvider.buildStatus"
// +kubebuilder:object:root=true
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BuildSpec   `json:"spec"`
	Status            BuildStatus `json:"status,omitempty"`
}

// BuildList contains a list of Builds
// +kubebuilder:object:root=true
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}

// A BuildSpec defines the desired state of an Build.
type BuildSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            BuildParameters `json:"forProvider"`
}

// BuildParameters define the desired state of a build
type BuildParameters struct {
	// BuildReference references the original build resource
	BuildReference BuildReference `json:"buildReference"`
	// SourceConfig refers to the source of the build
	SourceConfig SourceConfig `json:"sourceConfig"`
	// Images defines the image spec of the final built image
	Image meta.Image `json:"image"`
	// Config contains deployment related config
	// +optional
	Config Config `json:"config"`
}

// BuildReference describes a reference to a kpack build on a cluster
type BuildReference struct {
	// Cluster references the cluster of the build
	Cluster meta.Reference `json:"cluster"`
	// Build references the original build resource on the cluster of the
	// build
	Build meta.Reference `json:"build"`
}

// SourceConfig describes the source of a build
type SourceConfig struct {
	// Git specifies a target git repo with a revision to use
	Git GitTarget `json:"git"`
}

// An BuildStatus represents the observed state of an Build.
type BuildStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               BuildObservation `json:"atProvider"`
}

// BuildObservation are the observable fields of a Build.
type BuildObservation struct {
	// Status describes the status of the build
	// +optional
	BuildStatus BuildProcessStatus `json:"buildStatus,omitempty"`
	// StepsCompleted describes all the completed build steps
	// +optional
	StepsCompleted []string `json:"stepsCompleted,omitempty"`
}

// BuildProcessStatus describes the status of a build
type BuildProcessStatus string

// ConfigOrigin describes the origin of a config
type ConfigOrigin string
type OriginEnvVar struct {
	Value  EnvVar       `json:"value" yaml:"value"`
	Origin ConfigOrigin `json:"origin" yaml:"origin"`
}
type EnvVar struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

// A ProjectConfig defines the default config of an application in a project.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
type ProjectConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectConfigSpec   `json:"spec"`
	Status            ProjectConfigStatus `json:"status,omitempty"`
}

// ProjectConfigList contains a list of ProjectConfigs
// +kubebuilder:object:root=true
type ProjectConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectConfig `json:"items"`
}

// An ProjectConfigSpec defines the desired state of a ProjectConfig.
type ProjectConfigSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ProjectConfigParameters `json:"forProvider"`
}

// ProjectConfigParameters are the configurable fields of a ProjectConfig.
type ProjectConfigParameters struct {
	Config Config `json:"config"`
}

// An ProjectConfigStatus represents the observed state of a ProjectConfig.
type ProjectConfigStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
}

// A Release is created from a "successful" Build.
// It contains all the information needed to deploy the application.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="APPLICATION",type="string",JSONPath=".metadata.labels.application\\.apps\\.nine\\.ch/name"
// +kubebuilder:printcolumn:name="SIZE",type="string",JSONPath=".spec.forProvider.config.size"
// +kubebuilder:printcolumn:name="REPLICAS",type="integer",JSONPath=".spec.forProvider.config.replicas"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.atProvider.releaseStatus"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// CreationTimestampNano is the unix nano timestamp of the release when
	// it was created. It can be used to order releases and find the most
	// current one. It provides a higher accuracy than the normal
	// kubernetes resource CreationTimestamp which only has a resolution
	// down to a second.
	CreationTimestampNano int64         `json:"creationTimestampNano"`
	Spec                  ReleaseSpec   `json:"spec"`
	Status                ReleaseStatus `json:"status,omitempty"`
}

// ReleaseList contains a list of Releases
// +kubebuilder:object:root=true
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

// A ReleaseSpec defines the desired Release state
type ReleaseSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            ReleaseParameters `json:"forProvider"`
}

// ReleaseParameters define the desired Release state
type ReleaseParameters struct {
	// Build references the Build resource, that is the Release source.
	Build meta.LocalReference `json:"build"`
	// Image defines the image spec of the built image
	Image meta.Image `json:"image"`
	// Config contains all configurations from the various configuration
	// sources (project level, application level, etc) merged into one.
	Config Config `json:"config"`
	// DefaultHosts are the URLs at which the application is available.
	// +optional
	DefaultHosts []string `json:"defaultHosts"`
	// VerifiedHosts are the custom hosts which have been verified and can be
	// used in the release
	// +optional
	VerifiedHosts []string `json:"verifiedHosts,omitempty"`
	// BasicAuthSecret references a local secret which contains the basic
	// auth credentials
	// +optional
	BasicAuthSecret *meta.LocalReference `json:"basicAuthSecret,omitempty"`
	// Configuration contains all configurations from the various configuration
	// sources (project level, application level, etc) merged into one.
	// +optional
	Configuration *FieldOriginConfig `json:"configuration,omitempty"`
	// Timeout of the release after it will be considered failed. This does
	// not include the time spent waiting for the deploy job and only concerns
	// the release rollout.
	// +optional
	// +kubebuilder:default:="10m"
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// A FieldOriginConfig contains the fields of a normal config, but with an
// origin indicator for that field.
type FieldOriginConfig struct {
	// Size describes the CPU and memory requirements of the application
	Size OriginApplicationSize `json:"size,omitempty" yaml:"size,omitempty"`
	// Env variables which are passed to the app at runtime.
	// +optional
	Env OriginEnvVarList `json:"env,omitempty" yaml:"env,omitempty"`
	// Port the app is listening on.
	// +optional
	Port *OriginInt32 `json:"port,omitempty" yaml:"port,omitempty"`
	// Replicas sets the amount of replicas of the running app.
	// +optional
	Replicas *OriginInt32 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	// EnableBasicAuth enables basic authentication for the application
	// +optional
	EnableBasicAuth *OriginBool `json:"enableBasicAuth,omitempty" yaml:"enableBasicAuth,omitempty"`
	// +optional
	DeployJob *OriginDeployJob `json:"deployJob,omitempty" yaml:"deployJob,omitempty"`
}
type OriginApplicationSize struct {
	// +optional
	// +kubebuilder:default:=""
	Value  ApplicationSize `json:"value" yaml:"size"`
	Origin ConfigOrigin    `json:"origin" yaml:"origin"`
}
type OriginEnvVarList []OriginEnvVar
type OriginInt32 struct {
	Value  int32        `json:"value" yaml:"value"`
	Origin ConfigOrigin `json:"origin" yaml:"origin"`
}
type OriginBool struct {
	Value  bool         `json:"value" yaml:"value"`
	Origin ConfigOrigin `json:"origin" yaml:"origin"`
}
type OriginDeployJob struct {
	Value  DeployJob    `json:"value" yaml:"value"`
	Origin ConfigOrigin `json:"origin" yaml:"origin"`
}

// An ReleaseStatus represents the observed Release state
type ReleaseStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               ReleaseObservation `json:"atProvider"`
}

// ReleaseObservation are the observable fields of a Release
type ReleaseObservation struct {
	// ReleaseStatus describes the status of the Release
	// +optional
	ReleaseStatus ReleaseProcessStatus `json:"releaseStatus,omitempty"`
	// ReplicaObservation shows details about all replicas of the release
	// +optional
	ReplicaObservation []ReplicaObservation `json:"replicaObservation,omitempty"`
	// DeployJobStatus describes the status of the deploy job of a release
	// +optional
	DeployJobStatus *DeployJobStatus `json:"deployJobStatus,omitempty"`
	// CustomHostsCertificateStatus represents the latest Certificate status for
	// the custom hosts where the app is available.
	// +optional
	CustomHostsCertificateStatus CertificateStatus `json:"customHostsCertificateStatus,omitempty"`
	// Replicas describes the amount of rolled out replicas, ie. for the
	// underlying Deployment, it shows number of non-terminated pods targeted by
	// this Release.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
}

// ReleaseProcessStatus represents the Release status
type ReleaseProcessStatus string

// ReplicaObservation describes a replica
type ReplicaObservation struct {
	// Status describes the status of the replica.
	// +optional
	Status ReplicaStatus `json:"status"`
	// LastExitCode shows the last exit code of the replica.
	// +optional
	LastExitCode *int32 `json:"lastExitCode,omitempty"`
	// RestartCount indicates how often the replica was already restarted.
	// +optional
	RestartCount *int32 `json:"restartCount,omitempty"`
}

// ReplicaStatus is a status of a replica
type ReplicaStatus string
type DeployJobStatus struct {
	// Name of the deploy job.
	Name string `json:"name"`
	// Status indicates the status of the deploy job.
	Status DeployJobProcessStatus `json:"status"`
	// Reason indicates the failure reason in case of a failure.
	// +optional
	Reason DeployJobProcessReason `json:"reason,omitempty"`
	// StartTime is the timestamp the job has started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// ExitTime is the timestamp the job has exited.
	// +optional
	ExitTime *metav1.Time `json:"exitTime,omitempty"`
}

// DeployJobProcessStatus represents the deploy job status
type DeployJobProcessStatus string
type DeployJobProcessReason string
