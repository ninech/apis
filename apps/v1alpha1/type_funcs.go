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
