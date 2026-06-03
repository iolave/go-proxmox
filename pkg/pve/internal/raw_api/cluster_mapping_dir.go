package rawapi

import (
	http "net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterMappingListDirRequest struct {
	// If given, checks the configurations on the given node for correctness, and adds relevant diagnostics for the directory to the response.
	CheckNode string `in:"query=check-node;omitempty"`
}

type ClusterMappingGetDirResponse struct {
}

// ClusterMappingListDir List directory mapping
//
// Required permissions:
//
//	Only lists entries where you have 'Mapping.Modify', 'Mapping.Use' or 'Mapping.Audit' permissions on '/mapping/dir/<id>'.
func ClusterMappingListDir(c *pve.Client, req ClusterMappingListDirRequest) (
	res []struct {
		// A description of the logical mapping.
		Description string `json:"description"`

		// The logical ID of the mapping.
		ID int `json:"id"`

		// The entries of the mapping.
		Map []string `json:"map"`

		// A list of checks, only present if 'check-node' is set.
		Checks []struct {
			// The message of the error
			Message string `json:"message"`

			// The severity of the error
			//
			// 	error | warning
			Severity string `json:"severity"`
		} `json:"checks"`
	},
	err error,
) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/dir",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingCreateDirRequest struct {
	// The ID of the directory mapping.
	ID string `in:"query=id;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[node=<string>, path=<string> , ...]
	Map string `in:"query=map;omitempty"`

	// Description of the directory mapping.
	Description string `in:"query=description;omitempty"`
}

// ClusterMappingCreateDir Create a new directory mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/dir",["Mapping.Modify"]]
func ClusterMappingCreateDir(c *pve.Client, req ClusterMappingCreateDirRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/dir",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingGetDir Get directory mapping.
//
// Parameters:
//
//   - id is the ID of the directory mapping.
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/dir/{id}",["Mapping.Use"]],["perm","/mapping/dir/{id}",["Mapping.Modify"]],["perm","/mapping/dir/{id}",["Mapping.Audit"]]]
//
// TODO: Add response type. [Promox docs] doesn't have any response type.
//
// [Promox docs] https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/mapping/dir/{id}
func ClusterMappingGetDir(c *pve.Client, id string) (res any, err error) {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/dir/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingUpdateDirRequest struct {
	// The ID of the directory mapping.
	ID string `in:"path=id;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Description of the directory mapping.
	Description string `in:"query=description;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[node=<string>, path=<string> , ...]
	Map string `in:"query=map;omitempty"`
}

// ClusterMappingCreateDir Update a directory mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/dir/{id}",["Mapping.Modify"]]
func ClusterMappingUpdateDir(c *pve.Client, req ClusterMappingUpdateDirRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/dir/{id}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingDeleteDir Remove directory mapping.
//
// Parameters:
//
//   - id is the ID of the directory mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/dir",["Mapping.Modify"]]
func ClusterMappingDeleteDir(c *pve.Client, id string) error {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/dir/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
