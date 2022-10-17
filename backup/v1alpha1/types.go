package v1alpha1

import velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"

// VeleroSchedule defines a Velero backup schedule.
// +kubebuilder:object:generate=true
type VeleroSchedule struct {
	// Name of the backup schedule
	Name string `json:"name"`
	// CronExpression is a POSIX cron expression defining when to run the
	// backup. For example every night at 3 am UTC: "0 3 * * *". Note that it
	// does not allow to configure a backup to run every minute, meaning the
	// first term of the expression (minute) has to be a number.
	// +kubebuilder:validation:Pattern:=`^(\d+)(\d+)?(\s+(\d+|\*)(\d+)?){4}$`
	CronExpression string `json:"cronExpression"`
	// Spec defines what to back up and how long to retain it
	Spec velerov1.BackupSpec `json:"spec"`
}
