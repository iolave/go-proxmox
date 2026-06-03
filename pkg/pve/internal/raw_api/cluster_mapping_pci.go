package rawapi

import (
	http "net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterMappingListPCIRequest struct {
	// If given, checks the configurations on the given node for correctness, and adds relevant diagnostics for the directory to the response.
	CheckNode string `in:"query=check-node;omitempty"`
}

type ClusterMappingGetPCIResponse struct {
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
}

// ClusterMappingListPCI List PCI mapping
//
// Required permissions:
//
//	Only lists entries where you have 'Mapping.Modify', 'Mapping.Use' or 'Mapping.Audit' permissions on '/mapping/pci/<id>'.
func ClusterMappingListPCI(c *pve.Client, req ClusterMappingListPCIRequest) (
	res []ClusterMappingGetPCIResponse,
	err error,
) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/pci",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingCreatePCIRequest struct {
	// The ID of the logical PCI mapping.
	ID string `in:"query=id;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>, node=<string>, path=<(?:[a-f0-9]{4,}:[a-f0-9]{2}:[a-f0-9]{2}(?:.[a-f0-9])?;)*[a-f0-9]{4,}:[a-f0-9]{2}:[a-f0-9]{2}(?:.[a-f0-9])?> [,description=<string>] [,iommugroup=<integer>] [,subsystem-id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>], ...]
	Map string `in:"query=map;omitempty"`

	// Description of the logical PCI mapping.
	Description string `in:"query=description;omitempty"`

	// Marks the device(s) as being able to be live-migrated (Experimental). This needs hardware and driver support to work.
	LiveMigrationCapable *int `in:"query=live-migration-capable;omitempty"`

	// Marks the device(s) as being capable of providing mediated devices.
	MDev *int `in:"query=mdev;omitempty"`
}

// ClusterMappingCreatePCI Create a new hardware mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/pci",["Mapping.Modify"]]
func ClusterMappingCreatePCI(c *pve.Client, req ClusterMappingCreatePCIRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/pci",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingGetPCI Get PCI Mapping.
//
// Parameters:
//
//   - id is the ID of the logical PCI mapping.
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/pci/{id}",["Mapping.Use"]],["perm","/mapping/pci/{id}",["Mapping.Modify"]],["perm","/mapping/pci/{id}",["Mapping.Audit"]]]
//
// TODO: Add response type. [Promox docs] doesn't have any response type.
//
// [Promox docs] https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/mapping/pci/{id}
func ClusterMappingGetPCI(c *pve.Client, id string) (res any, err error) {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/pci/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMappingUpdatePCIRequest struct {
	// The ID of the logical PCI mapping.
	ID string `in:"path=id;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Description of the logical PCI mapping.
	Description string `in:"query=description;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Marks the device(s) as being able to be live-migrated (Experimental). This needs hardware and driver support to work.
	LiveMigrationCapable *int `in:"query=live-migration-capable;omitempty"`

	// A list of maps for the cluster nodes.
	//
	// 	[id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>, node=<string>, path=<(?:[a-f0-9]{4,}:[a-f0-9]{2}:[a-f0-9]{2}(?:.[a-f0-9])?;)*[a-f0-9]{4,}:[a-f0-9]{2}:[a-f0-9]{2}(?:.[a-f0-9])?> [,description=<string>] [,iommugroup=<integer>] [,subsystem-id=<(?^:^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{4}$)>], ...]
	Map string `in:"query=map;omitempty"`

	// Marks the device(s) as being capable of providing mediated devices.
	MDev *int `in:"query=mdev;omitempty"`
}

// ClusterMappingUpdatePCI Update a hardware mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/pci/{id}",["Mapping.Modify"]]
func ClusterMappingUpdatePCI(c *pve.Client, req ClusterMappingUpdatePCIRequest) error {
	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/pci/{id}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterMappingDeletePCI Remove Hardware Mapping.
//
// Parameters:
//
//   - id is the ID of the logical PCI mapping.
//
// Required permissions:
//
//	Check: ["perm","/mapping/pci",["Mapping.Modify"]]
func ClusterMappingDeletePCI(c *pve.Client, id string) error {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/mapping/pci/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
