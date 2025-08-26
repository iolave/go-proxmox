package api

import "net/http"

type ClusterJobsScheduleListRequest struct {
	Schedule   string `in:"query=schedule;omitempty"`
	Iterations *int   `in:"query=iterations;omitempty"`
	StartTime  *int   `in:"query=starttime;omitempty"`
}

// ClusterJobsScheduleList Returns a list of future schedule runtimes.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (c API) ClusterJobsScheduleList(req ClusterJobsScheduleListRequest) (res []struct {
	Timestamp int    `json:"timestamp"`
	UTC       string `json:"utc"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/jobs/schedule-analyze",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterJobsGetRealmSyncResponse struct {
	// If the job is enabled or not.
	Enabled int `json:"enabled"`

	// The ID of the entry.
	ID string `json:"id"`

	// Authentication domain ID
	Realm string `json:"realm"`

	// The configured sync schedule.
	Schedule string `json:"schedule"`

	// A comment for the job.
	Comment string `json:"comment"`

	// Last execution time of the job in seconds since the beginning of the UNIX epoch
	LastRun *int `json:"last-run"`

	// Next planned execution time of the job in seconds since the beginning of the UNIX epoch.
	NextRun *int `json:"next-run"`

	// A semicolon-separated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	//
	// 	([acl];[properties];[entry])|none
	RemoveVanished *string `json:"remove-vanished"`

	// Select what to sync.
	//
	// 	users | groups | both
	Scope *string `json:"scope"`
}

// ClusterJobsListRealmSync List configured realm-sync-jobs.
//
// Required permissions:
//
//	'Realm.AllocateUser' on '/access/realm/<realm>' and 'User.Modify' permissions to '/access/groups/'.
//	Check: ["and",["perm","/access/realm/{realm}",["Realm.AllocateUser"]],["perm","/access/groups",["User.Modify"]]
func (c API) ClusterJobsListRealmSync() (res []ClusterJobsGetRealmSyncResponse, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/jobs/realm-sync",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterJobsGetRealmSync Read realm-sync job definition.
//
// Parameters:
//
//   - id is the unique identifier of the realm-sync job.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// TODO: Add response type (not provided in [proxmox docs])
//
// [proxmox docs]: https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/jobs/realm-sync/{id}
func (c API) ClusterJobsGetRealmSync(id string) (res ClusterJobsGetRealmSyncResponse, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/jobs/realm-sync/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterJobsPostRealmSyncRequest struct {
	// The ID of the job.
	ID string `in:"path=id;omitempty"`

	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule string `in:"query=schedule;omitempty"`

	// Description for the Job.
	Comment string `in:"query=comment;omitempty"`

	// Enable newly synced users immediately.
	EnableNew *int `in:"query=enable-new;omitempty"`

	// Determines if the job is enabled.
	Enabled *int `in:"query=enabled;omitempty"`

	// Authentication domain ID
	Realm string `in:"query=realm;omitempty"`

	// A semicolon-separated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	//
	// 	([acl];[properties];[entry])|none
	RemoveVanished string `in:"query=remove-vanished;omitempty"`

	// Select what to sync.
	//
	// 	users | groups | both
	Scope string `in:"query=scope;omitempty"`
}

// ClusterJobsPostRealmSync Create new realm-sync job.
//
// Required permissions:
//
//	     'Realm.AllocateUser' on '/access/realm/<realm>' and 'User.Modify' permissions to '/access/groups/'.
//		Check: ["and",["perm","/access/realm/{realm}",["Realm.AllocateUser"]],["perm","/access/groups",["User.Modify"]]]
func (c API) ClusterJobsPostRealmSync(req ClusterJobsPostRealmSyncRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/jobs/realm-sync",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterJobsPutRealmSyncRequest struct {
	// The ID of the job.
	ID string `in:"path=id;omitempty"`

	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule string `in:"query=schedule;omitempty"`

	// Description for the Job.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Enable newly synced users immediately.
	EnableNew *int `in:"query=enable-new;omitempty"`

	// Determines if the job is enabled.
	Enabled *int `in:"query=enabled;omitempty"`

	// A semicolon-separated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	//
	// 	([acl];[properties];[entry])|none
	RemoveVanished string `in:"query=remove-vanished;omitempty"`

	// Select what to sync.
	//
	// 	users | groups | both
	Scope string `in:"query=scope;omitempty"`
}

// ClusterJobsPutRealmSync Update realm-sync job.
//
// Required permissions:
//
//	'Realm.AllocateUser' on '/access/realm/<realm>' and 'User.Modify' permissions to '/access/groups/'.
//	Check: ["and",["perm","/access/realm/{realm}",["Realm.AllocateUser"]],["perm","/access/groups",["User.Modify"]]]
func (c API) ClusterJobsPutRealmSync(req ClusterJobsPutRealmSyncRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/jobs/realm-sync/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterJobsDeleteRealmSync Delete realm-sync job.
//
// Parameters:
//
//   - id is the unique identifier of the realm-sync job.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterJobsDeleteRealmSync(id string) (err error) {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/jobs/realm-sync/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
