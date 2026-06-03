package rawapi

import (
	http "net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterMappingListResourceTypes list resource types
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterMappingListResourceTypes(c *pve.Client) (res []struct {
	Name string `json:"name"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/mapping",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
