package api

import "net/http"

type ClusterBackupGetJobResponse struct {
	// ID is the unique identifier of the backup job.
	ID *string `json:"id,omitempty"`

	// Backup all known guest systems on this host
	All *int `json:"all,omitempty"`

	// Limit I/O bandwidth (in KiB/s).
	BWLimit *int `json:"bwlimit,omitempty"`

	// Description for the Job.
	Comment *string `json:"comment,omitempty"`

	// Compress dump file.
	//	0 | 1 | gzip | lzo | zstd
	Compress *string `json:"compress,omitempty"`

	// Store resulting files to specified directory.
	DumpDir *string `json:"dumpdir,omitempty"`

	// Enable or disable the job.
	Enabled *int `json:"enabled,omitempty"`

	// Exclude specified guest systems (assumes --all)
	Exclude *string `json:"exclude,omitempty"`

	// Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root, other paths match relative to each subdirectory.
	ExcludePath []string `json:"exclude-path,omitempty"`

	// Options for backup fleecing (VM only).
	Fleecing struct {
		Enabled *string `json:"enabled,omitempty"`
		Storage *string `json:"storage,omitempty"`
	} `json:"fleecing"`

	// Set IO priority when using the BFQ scheduler. For snapshot and suspend mode backups of VMs, this only affects the compressor. A value of 8 means the idle priority is used, otherwise the best-effort priority is used with the specified value.
	//
	// 	0-8
	Ionice *int `json:"ionice,omitempty"`

	// Maximal time to wait for the global lock (minutes).
	LockWait *int `json:"lockwait,omitempty"`

	// Deprecated: use notification targets/matchers instead. Specify when to send a notification mail
	MailNotification *string `json:"mailnotification,omitempty"`

	// Deprecated: Use notification targets/matchers instead. Comma-separated list of email addresses or users that should receive email notifications.
	Mailto *string `json:"mailto,omitempty"`

	// Backup mode.
	//
	// 	snapshot | suspend | stop
	Mode *string `json:"mode,omitempty"`

	// Time of next run (unix timestamp).
	NextRun *int `json:"next-run,omitempty"`

	// Only run if executed on this node.
	Node *string `json:"node,omitempty"`

	// Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	NotesTemplate *string `json:"notes-template,omitempty"`

	// Determine which notification system to use. If set to 'legacy-sendmail', vzdump will consider the mailto/mailnotification parameters and send emails to the specified address(es) via the 'sendmail' command. If set to 'notification-system', a notification will be sent via PVE's notification system, and the mailto and mailnotification will be ignored. If set to 'auto' (default setting), an email will be sent if mailto is set, and the notification system will be used if not.
	//
	// 	auto | legacy-sendmail | notification-system
	NotificationMode *string `json:"notification-mode,omitempty"`

	// Other performance-related settings
	Performance struct {
		PbsEntriesMax *string `json:"pbs-entries-max,omitempty"`
		MaxWorkers    *string `json:"max-workers,omitempty"`
	} `json:"performance"`

	// Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pigz *int `json:"pigz,omitempty"`

	// Backup all known guest systems included in the specified pool.
	Pool *string `json:"pool,omitempty"`

	// If true, mark backup(s) as protected.
	Protected *int `json:"protected,omitempty"`

	// Use these retention options instead of those from the storage configuration.
	PruneBackups struct {
		KeepLast    *string `json:"keep-last,omitempty"`
		KeepAll     *string `json:"keep-all,omitempty"`
		KeepDaily   *string `json:"keep-daily,omitempty"`
		KeepHourly  *string `json:"keep-hourly,omitempty"`
		KeepMonthly *string `json:"keep-monthly,omitempty"`
		KeepWeekly  *string `json:"keep-weekly,omitempty"`
		KeepYearly  *string `json:"keep-yearly,omitempty"`
	} `json:"prune-backups"`

	// Be quiet.
	Quiet *int `json:"quiet,omitempty"`

	// Prune older backups according to 'prune-backups'.
	Remove *int `json:"remove,omitempty"`

	// If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	RepeatMissed *int `json:"repeat-missed,omitempty"`

	// Backup schedule. The format is a subset of `systemd` calendar even
	Schedule *string `json:"schedule,omitempty"`

	// Use specified hook script.
	Script *string `json:"script,omitempty"`

	// Exclude temporary files and logs.
	STDExcludes *int `json:"stdexcludes,omitempty"`

	// Stop running backup jobs on this host.
	Stop *int `json:"stop,omitempty"`

	// Maximal time to wait until a guest system is stopped (minutes).
	StopWait *int `json:"stopwait,omitempty"`

	// Store resulting file to this storage.
	Storage *string `json:"storage,omitempty"`

	// Backup type.
	Type *string `json:"type,omitempty"`

	// Store temporary files to specified directory.
	TMPDir *string `json:"tmpdir,omitempty"`

	// The ID of the guest system you want to backup.
	VMID *string `json:"vmid,omitempty"`

	// Zstd threads. N=0 uses half of the available cores, if N is set to a value bigger than 0, N is used as thread count.
	Zstd *int `json:"zstd,omitempty"`
}

// ClusterBackupGetJobs List vzdump backup schedule.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterBackupGetJobs() (res []ClusterBackupGetJobResponse, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/backup",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterBackupDeleteJob Delete vzdump backup job definition.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterBackupDeleteJob(id string) (err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/backup/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

// ClusterBackupGetJob List vzdump backup schedule.
//
// Parameters:
//
//   - id is the unique identifier of the backup job.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterBackupGetJob(id string) (res ClusterBackupGetJobResponse, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/backup/{id}",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

type ClusterBackupPostJobRequest struct {
	// Backup all known guest systems on this host.
	All *int `in:"query=all;omitempty"`

	// Limit I/O bandwidth (in KiB/s).
	BWLimit *int `in:"query=bwlimit;omitempty"`

	// Description for the Job.
	Comment string `in:"query=comment;omitempty"`

	// Compress dump file.
	//	0 | 1 | gzip | lzo | zstd
	Compress string `in:"query=compress;omitempty"`

	// Day of week selection.
	//
	// 	mon,tue,wed,thu,fri,sat,sun
	DOW string `in:"query=dow;omitempty"`

	// Store resulting files to specified directory.
	DumpDir string `in:"query=dumpdir;omitempty"`

	// Enable or disable the job.
	Enabled *int `in:"query=enabled;omitempty"`

	// Exclude specified guest systems (assumes --all).
	Exclude string `in:"query=exclude;omitempty"`

	// Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root, other paths match relative to each subdirectory.
	//
	// 	<path>[:<type>][<string>, ...]
	ExcludePath string `in:"query=exclude-path;omitempty"`

	// Options for backup fleecing (VM only).
	//
	// 	[[enabled=]<1|0>] [,storage=<storage ID>]
	Fleecing string `in:"query=fleecing;omitempty"`

	// Job ID (will be autogenerated).
	ID string `in:"query=id;omitempty"`

	// Set I/O priority when using the BFQ scheduler. For snapshot and suspend mode backups of VMs, this only affects the compressor. A value of 8 means the idle priority is used, otherwise the best-effort priority is used with the specified value.
	//
	// 	0-8
	Ionice *int `in:"query=ionice;omitempty"`

	// Maximal time to wait for the global lock (minutes).
	LockWait *int `in:"query=lockwait;omitempty"`

	// Deprecated: use notification targets/matchers instead. Specify when to send a notification mail.
	MailNotification string `in:"query=mailnotification;omitempty"`

	// Deprecated: Use notification targets/matchers instead. Comma-separated list of email addresses or users that should receive email notifications.
	Mailto string `in:"query=mailto;omitempty"`

	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	MaxFiles *int `in:"query=maxfiles;omitempty"`

	// Backup mode.
	//
	// 	snapshot | suspend | stop
	Mode string `in:"query=mode;omitempty"`

	// Only run if executed on this node.
	Node string `in:"query=node;omitempty"`

	// Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	NotesTemplate string `in:"query=notes-template;omitempty"`

	// Determine which notification system to use. If set to 'legacy-sendmail', vzdump will consider the mailto/mailnotification parameters and send emails to the specified address(es) via the 'sendmail' command. If set to 'notification-system', a notification will be sent via PVE's notification system, and the mailto and mailnotification will be ignored. If set to 'auto' (default setting), an email will be sent if mailto is set, and the notification system will be used if not.
	//
	// 	auto | legacy-sendmail | notification-system
	NotificationMode string `in:"query=notification-mode;omitempty"`

	// PBS mode used to detect file changes and switch encoding format for container backups.
	//
	// 	legacy | data | metadata
	PBSChangeDetectionMode string `in:"query=pbs-change-detection-mode;omitempty"`

	// Other performance-related settings.
	//
	// 	[pbs-entries-max=<INTEGER>] [,max-workers=<INTEGER>]
	Performance string `in:"query=performance;omitempty"`

	// Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pigz *int `in:"query=pigz;omitempty"`

	// Backup all known guest systems included in the specified pool.
	Pool string `in:"query=pool;omitempty"`

	// If true, mark backup(s) as protected.
	Protected *int `in:"query=protected;omitempty"`

	// Use these retention options instead of those from the storage configuration.
	//
	// 	[keep-all=<1|0>] [,keep-daily=<N>] [,keep-hourly=<N>] [,keep-last=<N>] [,keep-monthly=<N>] [,keep-weekly=<N>] [,keep-yearly=<N>]
	PruneBackups string `in:"query=prune-backups;omitempty"`

	// Be quiet.
	Quiet *int `in:"query=quiet;omitempty"`

	// Prune older backups according to 'prune-backups'.
	Remove *int `in:"query=remove;omitempty"`

	// If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	RepeatMissed *int `in:"query=repeat-missed;omitempty"`

	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule string `in:"query=schedule;omitempty"`

	// Use specified hook script.
	Script string `in:"query=script;omitempty"`

	// Job Start time.
	//
	//	HH:MM
	StartTime string `in:"query=starttime;omitempty"`

	// Exclude temporary files and logs.
	STDExcludes string `in:"query=stdexcludes;omitempty"`

	// Stop running backup jobs on this host.
	Stop *int `in:"query=stop;omitempty"`

	// Maximal time to wait until a guest system is stopped (minutes).
	Stopwait *int `in:"query=stopwait;omitempty"`

	// Store resulting file to this storage.
	Storage string `in:"query=storage;omitempty"`

	// Store temporary files to specified directory.
	TMPDir string `in:"query=tmpdir;omitempty"`

	// The ID of the guest system you want to backup.
	VMID *int `in:"query=vmid;omitempty"`

	// Zstd threads. N=0 uses half of the available cores, if N is set to a value bigger than 0, N is used as thread count.
	Zstd *int `in:"query=zstd;omitempty"`
}

// ClusterBackupPostJob Create new vzdump backup job.
//
// Required permissions:
//
//	The 'tmpdir', 'dumpdir' and 'script' parameters are additionally restricted to the 'root@pam' user.
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterBackupPostJob(req ClusterBackupPostJobRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/backup",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterBackupUpdateJobRequest struct {
	ClusterBackupPostJobRequest

	// Job ID
	ID string `in:"path=id;nonzero"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`
}

// ClusterBackupUpdateJobUpdate vzdump backup job definition.
//
// Required permissions:
//
//	The 'tmpdir', 'dumpdir' and 'script' parameters are additionally restricted to the 'root@pam' user.
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterBackupUpdateJob(req ClusterBackupUpdateJobRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/backup/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterBackupGetJobIncludedVolumes Returns included guests and the backup status of their disks. Optimized to be used in ExtJS tree views.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterBackupGetJobIncludedVolumes(id string) (res struct {
	Children []struct {
		VMID     int    `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Children []struct {
			ID       string `json:"id"`
			Included int    `json:"included"`
			Name     string `json:"name"`
			Reason   string `json:"reason"`
		} `json:"children"`
	} `json:"children"`
}, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/backup/{id}/included_volumes",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}
