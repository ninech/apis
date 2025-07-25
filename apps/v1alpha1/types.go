package v1alpha1

import (
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	meta "github.com/ninech/apis/meta/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

const (
	// LogLabelApplication is attached to all logs belonging to an app.
	LogLabelApplication = "app"
	// LogLabelRelease is attached to all logs belonging to a release.
	LogLabelRelease = "release"
	// LogLabelReplica is attached to all logs belonging to a replica.
	LogLabelReplica = "replica"
	// LogLabelProject attached to all logs belonging to a project.
	LogLabelProject = "project"
	// LogLabelWorkerJob is attached to all logs belonging to a worker job.
	LogLabelWorkerJob = "worker_job"
	// LogLabelScheduledJob is attached to all logs belonging to a scheduled job.
	LogLabelScheduledJob = "scheduled_job"
	// LogLabelDeployJob is attached to all logs belonging to a deploy job.
	LogLabelDeployJob = "deploy_job"
	// BuildProcessStatusError represents an unknown build status
	BuildProcessStatusUnknown BuildProcessStatus = "unknown"
	// BuildProcessStatusError represents the status of a failed build
	BuildProcessStatusError BuildProcessStatus = "error"
	// BuildProcessStatusImageUploadFailed represents the status of a failed build image upload
	BuildProcessStatusImageUploadFailed BuildProcessStatus = "imageUploadFailed"
	// BuildProcessStatusRunning represents the status of a running build
	BuildProcessStatusRunning BuildProcessStatus = "running"
	// BuildProcessStatusSuccess represents the status of a successful/finished build
	BuildProcessStatusSuccess BuildProcessStatus = "success"
	// LogLabelBuild is attached to all logs belonging to a build.
	LogLabelBuild = "build"
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
	// ReleaseProcessStatusPaused indicates the release is paused.
	ReleaseProcessStatusPaused ReleaseProcessStatus = "paused"
	// ReplicaStatusReady is the status for a ready running replica.
	ReplicaStatusReady ReplicaStatus = "ready"
	// ReplicaStatusSucceeded is the status for a succeeded running replica.
	ReplicaStatusSucceeded ReplicaStatus = "succeeded"
	// RepicaStatusProgressing is the status for a currently progressing
	// replica.
	ReplicaStatusProgressing ReplicaStatus = "progressing"
	// ReplicaStatusFailing describes a replica which is not in a ready state
	// and has restarted more than 2 times.
	ReplicaStatusFailing ReplicaStatus = "failing"
	// ReplicaStatusWaiting describes a state where a replica is waiting for
	// an event to happen. For example a currently running deploy job.
	ReplicaStatusWaiting ReplicaStatus = "waiting"
	// ReplicaStatusTerminating describes a state where a replica is terminating.
	ReplicaStatusTerminating ReplicaStatus = "terminating"
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
	DefaultConfig                                         = Config{Size: AppMicro, Replicas: ptr.To(int32(1)), Port: ptr.To(int32(8080)), EnableBasicAuth: ptr.To(false)}
	AppMicro      ApplicationSize                         = "micro"
	AppMini       ApplicationSize                         = "mini"
	AppStandard1  ApplicationSize                         = "standard-1"
	AppStandard2  ApplicationSize                         = "standard-2"
	AppResources  map[ApplicationSize]corev1.ResourceList = map[ApplicationSize]corev1.ResourceList{AppMicro: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("125m"), corev1.ResourceMemory: resource.MustParse("256Mi")}, AppMini: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("250m"), corev1.ResourceMemory: resource.MustParse("512Mi")}, AppStandard1: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("500m"), corev1.ResourceMemory: resource.MustParse("1Gi")}, AppStandard2: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("750m"), corev1.ResourceMemory: resource.MustParse("2Gi")}}
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
// +kubebuilder:validation:XValidation:rule="!(has(self.language) && self.dockerfileBuild.enabled)",message="language can't be set if dockerfile builds are enabled."
type ApplicationParameters struct {
	// Language specifies which kind of application/language should be
	// used for building the application. If left empty, an automatic
	// detection will be run.
	// +optional
	Language Language             `json:"language,omitempty"`
	Git      ApplicationGitConfig `json:"git"`
	Config   Config               `json:"config"`
	// Hosts is a list of host names where the application can be accessed. If
	// empty, the application will just be accessible on a generated host name
	// on the deploio.app domain.
	// +optional
	Hosts []string `json:"hosts,omitempty"`
	// Env variables which are passed to configure env variables required during
	// the build process.
	// +optional
	BuildEnv EnvVars `json:"buildEnv"`
	// DockerfileBuild contains settings for building with an existing
	// Dockerfile
	// +optional
	DockerfileBuild DockerfileBuild `json:"dockerfileBuild"`
	// BasicAuthPasswordChange can be used to create a new basic auth
	// password. If a password change should be done, the field needs to be
	// set to the current or a future timestamp. Once that time is reached,
	// the password will be changed to a new one. The timestamp need to be
	// in RFC3339 format (e.g. 2006-01-02T15:04:05Z).
	// +optional
	BasicAuthPasswordChange *metav1.Time `json:"basicAuthPasswordChange,omitempty"`
	// Paused pauses the release by stopping billing and disabling all runtime
	// workloads.
	// +optional
	Paused bool `json:"paused"`
}

// Language specifies which kind of application/language should be used
// for building the application. It influences the buildpack used.
// +kubebuilder:validation:Enum:="";ruby;php;python;golang;nodejs;static
type Language string

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
	// code. If not given, the root directory of the git repo will be used. The
	// SubPath field needs to start with a letter.
	// +optional
	SubPath string `json:"subPath"`
	// Revision defines the revision of the source to deploy the application
	// to. This can be a commit, tag, or branch.
	// +kubebuilder:validation:MinLength=1
	Revision string `json:"revision"`
}
type GitAuth struct {
	// FromSecret is a reference to a Secret from which to read the credentials.
	// Should contain the following keys depending on the protocol used.
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
	Size ApplicationSize `json:"size"`
	// Env variables which are passed to the app at runtime.
	// +optional
	Env EnvVars `json:"env"`
	// Port the app is listening on.
	// +optional
	// +nullable
	Port *int32 `json:"port"`
	// Replicas sets the amount of replicas of the running app. If this is
	// increased, make sure your application can cope with running multiple
	// replicas and all state required is shared in some way.
	// +optional
	// +nullable
	Replicas *int32 `json:"replicas"`
	// EnableBasicAuth enables basic authentication for the application
	// +optional
	// +nullable
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty"`
	// +optional
	// +nullable
	DeployJob *DeployJob `json:"deployJob,omitempty"`
	// +optional
	// +nullable
	// +listType:="map"
	// +listMapKey:="name"
	// WorkerJob defines a list of WorkerJob. See WorkerJob definition for a
	// description.
	WorkerJobs []WorkerJob `json:"workerJobs,omitempty"`
	// +optional
	// +nullable
	// +listType:="map"
	// +listMapKey:="name"
	// ScheduledJobs defines a list of ScheduledJob. See ScheduledJob definition for a
	// description.
	ScheduledJobs []ScheduledJob `json:"scheduledJobs,omitempty"`
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

// WorkerJob is a separate process which uses the same image as the application,
// but a different entry point. One common use-case is to start a task scheduler
// which executes tasks on a regular base.
type WorkerJob struct {
	Job `json:",inline"`
	// Size defines the amount of CPU and memory which this job can make use of
	// +optional
	// +kubebuilder:default:="micro"
	Size *ApplicationSize `json:"size,omitempty" yaml:"size,omitempty"`
}

// ScheduledJob is a separate process that runs periodically which uses the same
// image as the application, but a different entry point.
type ScheduledJob struct {
	Job       `json:",inline"`
	FiniteJob `json:",inline"`
	// Schedule defines the schedule in crontab syntax
	Schedule string `json:"schedule"`
	// Size defines the amount of CPU and memory which this job can make use of
	// +optional
	// +kubebuilder:default:="micro"
	Size *ApplicationSize `json:"size,omitempty" yaml:"size,omitempty"`
}
type DockerfileBuild struct {
	// Enabled defines if the Dockerfile build should be enabled
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Dockerfile builds cannot be enabled/disabled after creation of the application (field is immutable)"
	// +optional
	Enabled bool `json:"enabled"`
	// DockerfilePath specifies the path to the Dockerfile. If left empty a
	// file named Dockerfile will be searched in the application code root
	// directory.
	// +optional
	DockerfilePath string `json:"dockerfilePath,omitempty"`
	// BuildContext defines the build context. If left empty, the
	// application code root directory will be used as build context.
	// +optional
	BuildContext string `json:"buildContext,omitempty"`
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
	Hosts meta.DNSVerificationStatusEntries `json:"hosts,omitempty"`
	// DefaultHostsCertificateStatus represents the latest Certificate status for the
	// default URLs where the app is available.
	// +optional
	DefaultHostsCertificateStatus CertificateStatus `json:"defaultHostsCertificateStatus,omitempty"`
	// CustomHostsCertificateStatus represents the latest Certificate status for the
	// defined custom hosts.
	// +optional
	CustomHostsCertificateStatus CertificateStatus `json:"customHostsCertificateStatus,omitempty"`
	// LatestRelease shows the latest release for this application
	// +optional
	LatestRelease string `json:"latestRelease,omitempty"`
	// LatestBuild shows the latest build for this application
	// +optional
	LatestBuild string `json:"latestBuild,omitempty"`
	// Size shows the effective application size which is currently in use
	// +optional
	Size ApplicationSize `json:"size,omitempty"`
	// Replicas shows the effective amount of replicas which are currently
	// deployed
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`
	// WorkerJobs shows the currently running workers.
	// +optional
	WorkerJobs []ApplicationWorkerJobStatus `json:"workerJobs,omitempty"`
	// ScheduledJobs shows the currently running ScheduledJobs.
	// +optional
	ScheduledJobs []ApplicationScheduledJobStatus `json:"scheduledJobs,omitempty"`
	// BasicAuthSecret references the secret which contains the basic auth
	// credentials
	// +optional
	BasicAuthSecret *meta.LocalReference `json:"basicAuthSecret,omitempty"`
	// LastBasicAuthPasswordUpdate shows when the basic auth credentials have been
	// updated the last time.
	// +optional
	LastBasicAuthPasswordUpdate *metav1.Time `json:"lastBasicAuthPasswordUpdate,omitempty"`
	// ReferenceStatus contains errors for wrongly referenced resources
	meta.ReferenceStatus `json:",inline"`
}

// CertificateStatus represents the Certificate status
type CertificateStatus string
type ApplicationWorkerJobStatus struct {
	Name string          `json:"name"`
	Size ApplicationSize `json:"size"`
}
type ApplicationScheduledJobStatus struct {
	Name     string          `json:"name"`
	Size     ApplicationSize `json:"size"`
	Schedule string          `json:"schedule"`
}

// BuildpackMetadata describes the binary that was used in the build phase.
// Copied from https://github.com/buildpacks-community/kpack/blob/v0.10.0/pkg/apis/core/v1alpha1/buildpack_metadata.go
type BuildpackMetadata struct {
	Id       string `json:"id"`
	Version  string `json:"version"`
	Homepage string `json:"homepage,omitempty"`
}

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
	// Env variables used at the build time
	// +optional
	Env EnvVars `json:"env,omitempty"`
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
	// BuildMetadata describes the list of buildpack binaries that were used in
	// the build phase
	// +optional
	BuildMetadata BuildpackMetadataList `json:"buildMetadata,omitempty"`
	// Stack identifies the kpack stack used to build the image
	// +optional
	Stack string `json:"stack,omitempty"`
}

// BuildProcessStatus describes the status of a build
type BuildProcessStatus string

// BuildpackMetadataList contains a list of BuildpackMetadata
type BuildpackMetadataList []BuildpackMetadata

// ConfigOrigin describes the origin of a config
type ConfigOrigin string
type OriginEnvVar struct {
	Value  EnvVar       `json:"value"`
	Origin ConfigOrigin `json:"origin"`
}
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GitExploreResponse describes the response to a GitExploreRequest.
// +kubebuilder:object:generate=false
type GitExploreResponse struct {
	// Warnings contains optional warning messages
	Warnings []string `json:"warnings,omitempty"`
	// Error contains an optional occurred error during the exploration of a
	// Git repository.
	Error string `json:"error,omitempty"`
	// RepositoryInfo contains the explored items of a Git repository.
	RepositoryInfo *RepositoryInfo `json:"repositoryInfo,omitempty"`
}

// RepositoryInfo contains information about an explored Git repository
// +kubebuilder:object:generate=false
type RepositoryInfo struct {
	// URL is the URL used to obtain the information
	URL string `json:"url"`
	// RevisionResponse returns information about the optional revision
	// which was requested in the GitExploreRequest
	// +optional
	RevisionResponse *RevisionResponse `json:"revisionResponse,omitempty"`
	// Branches are the available git branches
	Branches []string `json:"branches,omitempty"`
	// Tags are the available tags
	Tags []string `json:"tags,omitempty"`
}

// RevisionResponse contains information about a requested git revision.
// +kubebuilder:object:generate=false
type RevisionResponse struct {
	// RevisionRequested is the revision which was requested
	RevisionRequested string `json:"revisionRequested"`
	// Found is set to true if the requested revision was found, otherwise false
	Found bool `json:"found"`
}

// GitExploreRequest describes a request to our Git information service
// +kubebuilder:object:generate=false
type GitExploreRequest struct {
	// Repository describes the path to the GIT repository. It can be given
	// as a full URL ("http(s)://" or "ssh://" prefix) or just in
	// "<Hostname>/<Path>" format in which case HTTPS will be
	// used.
	Repository string `json:"repository"`
	// Auth defines the authentication which is needed. If not given, no
	// authentication will be used.
	Auth *Auth `json:"auth,omitempty"`
	// +optional
	// Revision is a git revision which will be checked for existence in
	// the specified repository. If not given, no revision check will be
	// done.
	Revision string `json:"revision,omitempty"`
}

// Auth defines the possible authentication methods.
// +kubebuilder:object:generate=false
type Auth struct {
	// PrivateKey is a PEM encoded SSH private key which must not be
	// encrypted. The key has to be base64 encoded.
	PrivateKey []byte `json:"privateKey,omitempty"`
	// BasicAuth describes basic auth credentials to be used with HTTPS
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`
}

// BasicAuth is a simple username/password pair used for authentication
// +kubebuilder:object:generate=false
type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	Configuration FieldOriginConfig `json:"configuration"`
	// Timeout of the release after it will be considered failed. This does
	// not include the time spent waiting for the deploy job and only concerns
	// the release rollout.
	// +optional
	// +kubebuilder:default:="10m"
	Timeout *metav1.Duration `json:"timeout,omitempty"`
	// DockerfileBuild indicates if the build has been built using a dockerfile.
	// +optional
	DockerfileBuild bool `json:"dockerfileBuild,omitempty"`
	// AdditionalResources specifies extra resource requests and limits
	// that are factored into the Application Deployment's resource requirements.
	// +optional
	AdditionalResources corev1.ResourceList `json:"additionalResources,omitempty"`
	// RunAsUser allows the UID of the entrypoint of the container process to be
	// overridden. If not specified, the default UID of the container image will be
	// used.
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty"`
	// Paused pauses the release by stopping billing and disabling all runtime
	// workloads.
	// +optional
	Paused bool `json:"paused"`
}

// A FieldOriginConfig contains the fields of a normal config, but with an
// origin indicator for that field.
type FieldOriginConfig struct {
	// Size describes the CPU and memory requirements of the application
	Size OriginApplicationSize `json:"size,omitempty"`
	// Env variables which are passed to the app at runtime.
	// +optional
	Env OriginEnvVarList `json:"env,omitempty"`
	// Port the app is listening on.
	// +optional
	Port *OriginInt32 `json:"port,omitempty"`
	// Replicas sets the amount of replicas of the running app.
	// +optional
	Replicas *OriginInt32 `json:"replicas,omitempty"`
	// EnableBasicAuth enables basic authentication for the application
	// +optional
	EnableBasicAuth *OriginBool `json:"enableBasicAuth,omitempty"`
	// +optional
	DeployJob *OriginDeployJob `json:"deployJob,omitempty"`
	// +optional
	WorkerJobs []OriginWorkerJob `json:"workerJobs,omitempty"`
	// +optional
	ScheduledJobs []OriginScheduledJob `json:"scheduledJobs,omitempty"`
}
type OriginApplicationSize struct {
	// +optional
	// +kubebuilder:default:=""
	Value  ApplicationSize `json:"value"`
	Origin ConfigOrigin    `json:"origin"`
}
type OriginEnvVarList []OriginEnvVar
type OriginInt32 struct {
	Value  int32        `json:"value"`
	Origin ConfigOrigin `json:"origin"`
}
type OriginBool struct {
	Value  bool         `json:"value"`
	Origin ConfigOrigin `json:"origin"`
}
type OriginDeployJob struct {
	Value  DeployJob    `json:"value"`
	Origin ConfigOrigin `json:"origin"`
}
type OriginWorkerJob struct {
	Value  WorkerJob    `json:"value"`
	Origin ConfigOrigin `json:"origin"`
}
type OriginScheduledJob struct {
	Value  ScheduledJob `json:"value"`
	Origin ConfigOrigin `json:"origin"`
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
	// WorkerJobStatus describes the status of all workerjobs
	// +optional
	// +listType:="map"
	// +listMapKey:="name"
	WorkerJobStatus []WorkerJobStatus `json:"workerJobStatus,omitempty"`
	// ScheduledJobStatus describes the status of all workerjobs
	// +optional
	// +listType:="map"
	// +listMapKey:="name"
	ScheduledJobStatus []ScheduledJobStatus `json:"scheduledJobStatus,omitempty"`
	// CustomHostsCertificateStatus represents the latest Certificate status for
	// the custom hosts where the app is available.
	// +optional
	CustomHostsCertificateStatus CertificateStatus `json:"customHostsCertificateStatus,omitempty"`
	// Replicas describes the amount of rolled out replicas, ie. for the
	// underlying Deployment, it shows number of non-terminated pods targeted by
	// this Release.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Owning indicates if this release is currently owning all resources.
	// +optional
	Owning bool `json:"owning"`
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
	// ReplicaName is the name of the replica.
	// +optional
	ReplicaName string `json:"replicaName,omitempty"`
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
type WorkerJobStatus struct {
	Name               string               `json:"name"`
	ReplicaObservation []ReplicaObservation `json:"replicaObservation"`
}
type ScheduledJobStatus struct {
	Name               string               `json:"name"`
	ReplicaObservation []ReplicaObservation `json:"replicaObservation"`
}
