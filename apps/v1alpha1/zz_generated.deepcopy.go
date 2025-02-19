//go:build !ignore_autogenerated

/*
Copyright 2022 Nine Internet Solutions AG.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	metav1alpha1 "github.com/ninech/apis/meta/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationGitConfig) DeepCopyInto(out *ApplicationGitConfig) {
	*out = *in
	out.GitTarget = in.GitTarget
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(GitAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationGitConfig.
func (in *ApplicationGitConfig) DeepCopy() *ApplicationGitConfig {
	if in == nil {
		return nil
	}
	out := new(ApplicationGitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationObservation) DeepCopyInto(out *ApplicationObservation) {
	*out = *in
	if in.DefaultURLs != nil {
		in, out := &in.DefaultURLs, &out.DefaultURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]VerificationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.WorkerJobs != nil {
		in, out := &in.WorkerJobs, &out.WorkerJobs
		*out = make([]ApplicationWorkerJobStatus, len(*in))
		copy(*out, *in)
	}
	if in.ScheduledJobs != nil {
		in, out := &in.ScheduledJobs, &out.ScheduledJobs
		*out = make([]ApplicationScheduledJobStatus, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuthSecret != nil {
		in, out := &in.BasicAuthSecret, &out.BasicAuthSecret
		*out = new(metav1alpha1.LocalReference)
		**out = **in
	}
	if in.LastBasicAuthPasswordUpdate != nil {
		in, out := &in.LastBasicAuthPasswordUpdate, &out.LastBasicAuthPasswordUpdate
		*out = (*in).DeepCopy()
	}
	in.ReferenceStatus.DeepCopyInto(&out.ReferenceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationObservation.
func (in *ApplicationObservation) DeepCopy() *ApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParameters) DeepCopyInto(out *ApplicationParameters) {
	*out = *in
	in.Git.DeepCopyInto(&out.Git)
	in.Config.DeepCopyInto(&out.Config)
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BuildEnv != nil {
		in, out := &in.BuildEnv, &out.BuildEnv
		*out = make(EnvVars, len(*in))
		copy(*out, *in)
	}
	out.DockerfileBuild = in.DockerfileBuild
	if in.BasicAuthPasswordChange != nil {
		in, out := &in.BasicAuthPasswordChange, &out.BasicAuthPasswordChange
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParameters.
func (in *ApplicationParameters) DeepCopy() *ApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationScheduledJobStatus) DeepCopyInto(out *ApplicationScheduledJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationScheduledJobStatus.
func (in *ApplicationScheduledJobStatus) DeepCopy() *ApplicationScheduledJobStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationScheduledJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationWorkerJobStatus) DeepCopyInto(out *ApplicationWorkerJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationWorkerJobStatus.
func (in *ApplicationWorkerJobStatus) DeepCopy() *ApplicationWorkerJobStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationWorkerJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Build) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildList) DeepCopyInto(out *BuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Build, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildList.
func (in *BuildList) DeepCopy() *BuildList {
	if in == nil {
		return nil
	}
	out := new(BuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildObservation) DeepCopyInto(out *BuildObservation) {
	*out = *in
	if in.StepsCompleted != nil {
		in, out := &in.StepsCompleted, &out.StepsCompleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BuildMetadata != nil {
		in, out := &in.BuildMetadata, &out.BuildMetadata
		*out = make(BuildpackMetadataList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildObservation.
func (in *BuildObservation) DeepCopy() *BuildObservation {
	if in == nil {
		return nil
	}
	out := new(BuildObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildParameters) DeepCopyInto(out *BuildParameters) {
	*out = *in
	out.BuildReference = in.BuildReference
	out.SourceConfig = in.SourceConfig
	out.Image = in.Image
	in.Config.DeepCopyInto(&out.Config)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(EnvVars, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildParameters.
func (in *BuildParameters) DeepCopy() *BuildParameters {
	if in == nil {
		return nil
	}
	out := new(BuildParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildReference) DeepCopyInto(out *BuildReference) {
	*out = *in
	out.Cluster = in.Cluster
	out.Build = in.Build
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildReference.
func (in *BuildReference) DeepCopy() *BuildReference {
	if in == nil {
		return nil
	}
	out := new(BuildReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildSpec) DeepCopyInto(out *BuildSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildSpec.
func (in *BuildSpec) DeepCopy() *BuildSpec {
	if in == nil {
		return nil
	}
	out := new(BuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStatus) DeepCopyInto(out *BuildStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStatus.
func (in *BuildStatus) DeepCopy() *BuildStatus {
	if in == nil {
		return nil
	}
	out := new(BuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackMetadata) DeepCopyInto(out *BuildpackMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackMetadata.
func (in *BuildpackMetadata) DeepCopy() *BuildpackMetadata {
	if in == nil {
		return nil
	}
	out := new(BuildpackMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BuildpackMetadataList) DeepCopyInto(out *BuildpackMetadataList) {
	{
		in := &in
		*out = make(BuildpackMetadataList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackMetadataList.
func (in BuildpackMetadataList) DeepCopy() BuildpackMetadataList {
	if in == nil {
		return nil
	}
	out := new(BuildpackMetadataList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(EnvVars, len(*in))
		copy(*out, *in)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.DeployJob != nil {
		in, out := &in.DeployJob, &out.DeployJob
		*out = new(DeployJob)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerJobs != nil {
		in, out := &in.WorkerJobs, &out.WorkerJobs
		*out = make([]WorkerJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledJobs != nil {
		in, out := &in.ScheduledJobs, &out.ScheduledJobs
		*out = make([]ScheduledJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployJob) DeepCopyInto(out *DeployJob) {
	*out = *in
	out.Job = in.Job
	in.FiniteJob.DeepCopyInto(&out.FiniteJob)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployJob.
func (in *DeployJob) DeepCopy() *DeployJob {
	if in == nil {
		return nil
	}
	out := new(DeployJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployJobStatus) DeepCopyInto(out *DeployJobStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.ExitTime != nil {
		in, out := &in.ExitTime, &out.ExitTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployJobStatus.
func (in *DeployJobStatus) DeepCopy() *DeployJobStatus {
	if in == nil {
		return nil
	}
	out := new(DeployJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerfileBuild) DeepCopyInto(out *DockerfileBuild) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerfileBuild.
func (in *DockerfileBuild) DeepCopy() *DockerfileBuild {
	if in == nil {
		return nil
	}
	out := new(DockerfileBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EnvVars) DeepCopyInto(out *EnvVars) {
	{
		in := &in
		*out = make(EnvVars, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVars.
func (in EnvVars) DeepCopy() EnvVars {
	if in == nil {
		return nil
	}
	out := new(EnvVars)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldOriginConfig) DeepCopyInto(out *FieldOriginConfig) {
	*out = *in
	out.Size = in.Size
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(OriginEnvVarList, len(*in))
		copy(*out, *in)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(OriginInt32)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(OriginInt32)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(OriginBool)
		**out = **in
	}
	if in.DeployJob != nil {
		in, out := &in.DeployJob, &out.DeployJob
		*out = new(OriginDeployJob)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerJobs != nil {
		in, out := &in.WorkerJobs, &out.WorkerJobs
		*out = make([]OriginWorkerJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledJobs != nil {
		in, out := &in.ScheduledJobs, &out.ScheduledJobs
		*out = make([]OriginScheduledJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldOriginConfig.
func (in *FieldOriginConfig) DeepCopy() *FieldOriginConfig {
	if in == nil {
		return nil
	}
	out := new(FieldOriginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiniteJob) DeepCopyInto(out *FiniteJob) {
	*out = *in
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int32)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiniteJob.
func (in *FiniteJob) DeepCopy() *FiniteJob {
	if in == nil {
		return nil
	}
	out := new(FiniteJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitAuth) DeepCopyInto(out *GitAuth) {
	*out = *in
	if in.FromSecret != nil {
		in, out := &in.FromSecret, &out.FromSecret
		*out = new(metav1alpha1.LocalReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAuth.
func (in *GitAuth) DeepCopy() *GitAuth {
	if in == nil {
		return nil
	}
	out := new(GitAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTarget) DeepCopyInto(out *GitTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTarget.
func (in *GitTarget) DeepCopy() *GitTarget {
	if in == nil {
		return nil
	}
	out := new(GitTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginApplicationSize) DeepCopyInto(out *OriginApplicationSize) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginApplicationSize.
func (in *OriginApplicationSize) DeepCopy() *OriginApplicationSize {
	if in == nil {
		return nil
	}
	out := new(OriginApplicationSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginBool) DeepCopyInto(out *OriginBool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginBool.
func (in *OriginBool) DeepCopy() *OriginBool {
	if in == nil {
		return nil
	}
	out := new(OriginBool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginDeployJob) DeepCopyInto(out *OriginDeployJob) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginDeployJob.
func (in *OriginDeployJob) DeepCopy() *OriginDeployJob {
	if in == nil {
		return nil
	}
	out := new(OriginDeployJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginEnvVar) DeepCopyInto(out *OriginEnvVar) {
	*out = *in
	out.Value = in.Value
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginEnvVar.
func (in *OriginEnvVar) DeepCopy() *OriginEnvVar {
	if in == nil {
		return nil
	}
	out := new(OriginEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OriginEnvVarList) DeepCopyInto(out *OriginEnvVarList) {
	{
		in := &in
		*out = make(OriginEnvVarList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginEnvVarList.
func (in OriginEnvVarList) DeepCopy() OriginEnvVarList {
	if in == nil {
		return nil
	}
	out := new(OriginEnvVarList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginInt32) DeepCopyInto(out *OriginInt32) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginInt32.
func (in *OriginInt32) DeepCopy() *OriginInt32 {
	if in == nil {
		return nil
	}
	out := new(OriginInt32)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginScheduledJob) DeepCopyInto(out *OriginScheduledJob) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginScheduledJob.
func (in *OriginScheduledJob) DeepCopy() *OriginScheduledJob {
	if in == nil {
		return nil
	}
	out := new(OriginScheduledJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginWorkerJob) DeepCopyInto(out *OriginWorkerJob) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginWorkerJob.
func (in *OriginWorkerJob) DeepCopy() *OriginWorkerJob {
	if in == nil {
		return nil
	}
	out := new(OriginWorkerJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfigList) DeepCopyInto(out *ProjectConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfigList.
func (in *ProjectConfigList) DeepCopy() *ProjectConfigList {
	if in == nil {
		return nil
	}
	out := new(ProjectConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfigParameters) DeepCopyInto(out *ProjectConfigParameters) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfigParameters.
func (in *ProjectConfigParameters) DeepCopy() *ProjectConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfigSpec) DeepCopyInto(out *ProjectConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfigSpec.
func (in *ProjectConfigSpec) DeepCopy() *ProjectConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfigStatus) DeepCopyInto(out *ProjectConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfigStatus.
func (in *ProjectConfigStatus) DeepCopy() *ProjectConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseObservation) DeepCopyInto(out *ReleaseObservation) {
	*out = *in
	if in.ReplicaObservation != nil {
		in, out := &in.ReplicaObservation, &out.ReplicaObservation
		*out = make([]ReplicaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployJobStatus != nil {
		in, out := &in.DeployJobStatus, &out.DeployJobStatus
		*out = new(DeployJobStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerJobStatus != nil {
		in, out := &in.WorkerJobStatus, &out.WorkerJobStatus
		*out = make([]WorkerJobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledJobStatus != nil {
		in, out := &in.ScheduledJobStatus, &out.ScheduledJobStatus
		*out = make([]ScheduledJobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseObservation.
func (in *ReleaseObservation) DeepCopy() *ReleaseObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseParameters) DeepCopyInto(out *ReleaseParameters) {
	*out = *in
	out.Build = in.Build
	out.Image = in.Image
	in.Config.DeepCopyInto(&out.Config)
	if in.DefaultHosts != nil {
		in, out := &in.DefaultHosts, &out.DefaultHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VerifiedHosts != nil {
		in, out := &in.VerifiedHosts, &out.VerifiedHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuthSecret != nil {
		in, out := &in.BasicAuthSecret, &out.BasicAuthSecret
		*out = new(metav1alpha1.LocalReference)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(FieldOriginConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RunAsUser != nil {
		in, out := &in.RunAsUser, &out.RunAsUser
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseParameters.
func (in *ReleaseParameters) DeepCopy() *ReleaseParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaObservation) DeepCopyInto(out *ReplicaObservation) {
	*out = *in
	if in.LastExitCode != nil {
		in, out := &in.LastExitCode, &out.LastExitCode
		*out = new(int32)
		**out = **in
	}
	if in.RestartCount != nil {
		in, out := &in.RestartCount, &out.RestartCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaObservation.
func (in *ReplicaObservation) DeepCopy() *ReplicaObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledJob) DeepCopyInto(out *ScheduledJob) {
	*out = *in
	out.Job = in.Job
	in.FiniteJob.DeepCopyInto(&out.FiniteJob)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(ApplicationSize)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledJob.
func (in *ScheduledJob) DeepCopy() *ScheduledJob {
	if in == nil {
		return nil
	}
	out := new(ScheduledJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledJobStatus) DeepCopyInto(out *ScheduledJobStatus) {
	*out = *in
	if in.ReplicaObservation != nil {
		in, out := &in.ReplicaObservation, &out.ReplicaObservation
		*out = make([]ReplicaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledJobStatus.
func (in *ScheduledJobStatus) DeepCopy() *ScheduledJobStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfig) DeepCopyInto(out *SourceConfig) {
	*out = *in
	out.Git = in.Git
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfig.
func (in *SourceConfig) DeepCopy() *SourceConfig {
	if in == nil {
		return nil
	}
	out := new(SourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerificationError) DeepCopyInto(out *VerificationError) {
	*out = *in
	in.Timestamp.DeepCopyInto(&out.Timestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerificationError.
func (in *VerificationError) DeepCopy() *VerificationError {
	if in == nil {
		return nil
	}
	out := new(VerificationError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerificationStatus) DeepCopyInto(out *VerificationStatus) {
	*out = *in
	if in.LatestSuccess != nil {
		in, out := &in.LatestSuccess, &out.LatestSuccess
		*out = (*in).DeepCopy()
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(VerificationError)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerificationStatus.
func (in *VerificationStatus) DeepCopy() *VerificationStatus {
	if in == nil {
		return nil
	}
	out := new(VerificationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerJob) DeepCopyInto(out *WorkerJob) {
	*out = *in
	out.Job = in.Job
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(ApplicationSize)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerJob.
func (in *WorkerJob) DeepCopy() *WorkerJob {
	if in == nil {
		return nil
	}
	out := new(WorkerJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerJobStatus) DeepCopyInto(out *WorkerJobStatus) {
	*out = *in
	if in.ReplicaObservation != nil {
		in, out := &in.ReplicaObservation, &out.ReplicaObservation
		*out = make([]ReplicaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerJobStatus.
func (in *WorkerJobStatus) DeepCopy() *WorkerJobStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerJobStatus)
	in.DeepCopyInto(out)
	return out
}
