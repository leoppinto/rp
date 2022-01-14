package model

type JobResouceConfiguration struct {
	ID   string
	Name string
}

func NewJobResourceConfiguration() *JobResouceConfiguration {
	jobResouceConfiguration := JobResouceConfiguration{}
	jobResouceConfiguration.ID = "AHHSA"
	return &JobResouceConfiguration{}
}

type JobResouceConfigurations struct {
	JobResouceConfiguration []JobResouceConfiguration
}

func (j *JobResouceConfigurations) Add(jobResouceConfiguration *JobResouceConfiguration) {
	j.JobResouceConfiguration = append(j.JobResouceConfiguration, *jobResouceConfiguration)
}

func NewJobResouceConfigurations() *JobResouceConfigurations {
	return &JobResouceConfigurations{}
}
