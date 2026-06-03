package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClisterReplicationListReplicationJobs List replication jobs.
//
// Required permissions:
//
//	Will only return replication jobs for which the calling
//	user has VM.Audit permission on /vms/<vmid>.
func ClusterReplicationListReplicationJobs(c *pve.Client) (res []any, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/replication",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterReplicationNewReplicationJobRequest struct {
	// Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	//	format: pve-replication-job-id
	ID string `in:"query=id;omitempty"`

	// Target node.
	Target string `in:"query=target;omitempty"`

	// Section type.
	//	local
	Type string `in:"query=type;omitempty"`

	// Description.
	Comment string `in:"query=comment;omitempty"`

	// Flag to disable/deactivate the entry.
	Disable *int `in:"query=disable;omitempty"`

	// Rate limit in mbps (megabytes per second) as floating point number.
	Rate *int `in:"query=rate;omitempty"`

	// Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
	//	local | full
	RemoveJob string `in:"query=remove_job;omitempty"`

	// Storage replication schedule. The format is a subset of `systemd` calendar events.
	//	format: */15
	Schedule string `in:"query=schedule;omitempty"`

	// For internal use, to detect if the guest was stolen.
	Source string `in:"query=source;omitempty"`
}

// ClusterReplicationNewReplicationJob Create a new replication job.
//
// Required permissions:
//
//	Requires the VM.Replicate permission on /vms/<vmid>.
func ClusterReplicationNewReplicationJob(c *pve.Client, req ClusterReplicationNewReplicationJobRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/replication",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterReplicationGetReplicationJob Read replication job configuration.
//
// Parameters:
//
//   - id is the replication job ID.
//
// Required permissions:
//
//	Requires the VM.Audit permission on /vms/<vmid>.
func ClusterReplicationGetReplicationJob(c *pve.Client, id string) (res struct {
	Comment   string `json:"comment"`
	Disable   *int   `json:"disable"`
	ID        string `json:"id"`
	Rate      *int   `json:"rate"`
	RemoveJob string `json:"remove_job"`
	Schedule  string `json:"schedule"`
	Source    string `json:"source"`
	Target    string `json:"target"`
	Type      string `json:"type"`
}, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/replication/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterReplicationUpdateReplicationJobRequest struct {
	// Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	//	format: pve-replication-job-id
	ID string `in:"query=id;omitempty"`

	// Description.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Flag to disable/deactivate the entry.
	Disable *int `in:"query=disable;omitempty"`

	// Rate limit in mbps (megabytes per second) as floating point number.
	Rate *int `in:"query=rate;omitempty"`

	// Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
	//	local | full
	RemoveJob string `in:"query=remove_job;omitempty"`

	// Storage replication schedule. The format is a subset of `systemd` calendar events.
	//	format: */15
	Schedule string `in:"query=schedule;omitempty"`

	// For internal use, to detect if the guest was stolen.
	Source string `in:"query=source;omitempty"`
}

// ClusterReplicationUpdateReplicationJob Update replication job configuration.
//
// Required permissions:
//
//	Requires the VM.Replicate permission on /vms/<vmid>.
func ClusterReplicationUpdateReplicationJob(c *pve.Client, req ClusterReplicationUpdateReplicationJobRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/replication/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterReplicationDeleteReplicationJob Mark replication job for removal.
//
// Parameters:
//
//   - id: Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
//   - force (0|1): Will remove the jobconfig entry, but will not cleanup.
//   - keep (0|1): Keep replicated data at target (do not remove).
//
// Required permissions:
//
//	Requires the VM.Replicate permission on /vms/<vmid>.
func ClusterReplicationDeleteReplicationJob(c *pve.Client, id string, force, keep int) (err error) {
	req := struct {
		ID    string `in:"path=id;nonzero"`
		Force int    `in:"query=force"`
		Keep  int    `in:"query=keep"`
	}{
		ID:    id,
		Force: force,
		Keep:  keep,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/replication/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
