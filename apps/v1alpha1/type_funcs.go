package v1alpha1

// WithoutOrigin removes the origin information and returns just the pure config.
func (m FieldOriginConfig) WithoutOrigin() Config {
	result := Config{
		Size: m.Size.Value,
		Env:  m.Env.WithoutOrigin(),
	}
	if m.Port != nil {
		result.Port = &m.Port.Value
	}
	if m.Replicas != nil {
		result.Replicas = &m.Replicas.Value
	}
	if m.EnableBasicAuth != nil {
		result.EnableBasicAuth = &m.EnableBasicAuth.Value
	}
	if m.DeployJob != nil {
		result.DeployJob = &m.DeployJob.Value
	}
	for _, job := range m.WorkerJobs {
		result.WorkerJobs = append(result.WorkerJobs, job.Value)
	}
	for _, job := range m.ScheduledJobs {
		result.ScheduledJobs = append(result.ScheduledJobs, job.Value)
	}

	return result
}

// WithoutOrigin removes the origin information and returns just the pure EnvVars.
func (o OriginEnvVarList) WithoutOrigin() EnvVars {
	result := make(EnvVars, len(o))
	for i, content := range o {
		result[i] = EnvVar{Name: content.Value.Name, Value: content.Value.Value}
	}
	return result
}

// WithOrigin turns the config into a FieldOriginConfig by annotating the
// fields with the given origin
func (c Config) WithOrigin(origin ConfigOrigin) FieldOriginConfig {
	result := FieldOriginConfig{
		Size: OriginApplicationSize{
			Value:  c.Size,
			Origin: origin,
		},
		Env: c.Env.WithOrigin(origin),
	}
	if c.Port != nil {
		result.Port = &OriginInt32{
			Value:  *c.Port,
			Origin: origin,
		}
	}
	if c.Replicas != nil {
		result.Replicas = &OriginInt32{
			Value:  *c.Replicas,
			Origin: origin,
		}
	}
	if c.EnableBasicAuth != nil {
		result.EnableBasicAuth = &OriginBool{
			Value:  *c.EnableBasicAuth,
			Origin: origin,
		}
	}
	if c.DeployJob != nil {
		result.DeployJob = &OriginDeployJob{
			Value:  *c.DeployJob,
			Origin: origin,
		}
	}
	for _, job := range c.WorkerJobs {
		result.WorkerJobs = append(result.WorkerJobs, OriginWorkerJob{Value: job, Origin: origin})
	}
	for _, job := range c.ScheduledJobs {
		result.ScheduledJobs = append(result.ScheduledJobs, OriginScheduledJob{Value: job, Origin: origin})
	}
	return result
}

// WithOrigin creates a list of env variables annotated with the given origin
func (envs EnvVars) WithOrigin(origin ConfigOrigin) OriginEnvVarList {
	if envs == nil {
		return nil
	}

	result := make(OriginEnvVarList, len(envs))
	for i, envVar := range envs {
		result[i] = OriginEnvVar{
			Value:  envVar,
			Origin: origin,
		}
	}
	return result
}
