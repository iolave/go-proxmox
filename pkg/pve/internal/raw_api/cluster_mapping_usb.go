package rawapi

import (
	http "net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterMappingListUSBRequest struct {
	// If given, checks the configurations on the given node for correctness, and adds relevant diagnostics for the directory to the response.
	CheckNode string `in:"query=check-node;omitempty"`
}

// ClusterMappingListUSB List USB Hardware Mappings
//
// Required permissions:
//
//	Only lists entries where you have 'Mapping.Modify', 'Mapping.Use' or 'Mapping.Audit' permissions on '/mapping/usb/<id>'.
func ClusterMappingListUSB(c *pve.Client, req ClusterMappingListUSBRequest) (
	res []struct {
		Description string   `json:"description"`
		Error       string   `json:"error"`
		ID          int      `json:"id"`
		Map         []string `json:"map"`
	},
	err error,
) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/usb",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingCreateUSBRequest struct {
	// The ID of the logical USB mapping.
	ID string `in:"query=id;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>, node=<string> [,description=<string>] [,path=<(?^:^(\d+)\-(\d+(\.\d+)*)$)>], ...]
	Map string `in:"query=map;omitempty"`

	// Description of the logical USB mapping.
	Description string `in:"query=description;omitempty"`
}

// ClusterMappingCreateUSB Create a new hardware mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/usb",["Mapping.Modify"]]
func ClusterMappingCreateUSB(c *pve.Client, req ClusterMappingCreateUSBRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/usb",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingGetUSB Get USB Mapping.
//
// Parameters:
//
//   - id is the ID of the logical USB mapping.
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/usb/{id}",["Mapping.Audit"]],["perm","/mapping/usb/{id}",["Mapping.Use"]],["perm","/mapping/usb/{id}",["Mapping.Modify"]]]
//
// TODO: Add response type. [Promox docs] doesn't have any response type.
//
// [Promox docs] https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/mapping/usb/{id}
func ClusterMappingGetUSB(c *pve.Client, id string) (res any, err error) {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/usb/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingUpdateUSBRequest struct {
	// The ID of the logical USB mapping.
	ID string `in:"path=id;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>, node=<string> [,description=<string>] [,path=<(?^:^(\d+)\-(\d+(\.\d+)*)$)>], ...]
	Map string `in:"query=map;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Description of the logical USB mapping.
	Description string `in:"query=description;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`
}

// ClusterMappingUpdateUSB Update a hardware mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/usb/{id}",["Mapping.Modify"]]
func ClusterMappingUpdateUSB(c *pve.Client, req ClusterMappingUpdateUSBRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/usb/{id}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingDeleteUSB Remove Hardware Mapping.
//
// Parameters:
//
//   - id is the ID of the logical USB mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/usb",["Mapping.Modify"]]
func ClusterMappingDeleteUSB(c *pve.Client, id string) error {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/usb/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
