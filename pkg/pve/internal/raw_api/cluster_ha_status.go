package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterHAStatusCurrent Get HA manger status.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHAStatusCurrent(c *pve.Client) (res []struct {
	ID           string  `json:"id"`
	Node         string  `json:"node"`
	Status       string  `json:"status"`
	Type         string  `json:"type"`
	CRMState     *string `json:"crm_state"`
	Failback     *string `json:"failback"`
	MaxRelocate  *string `json:"max_relocate"`
	MaxRestart   *string `json:"max_restart"`
	QuoRate      *string `json:"quorate"`
	RequestState *string `json:"request_state"`
	SID          *string `json:"sid"`
	State        *string `json:"state"`
	Timestamp    *string `json:"timestamp"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/ha/status/current",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterHAStatusManager Get full HA manger status, including LRM status.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// TODO: Add response type (not provided in [proxmox docs])
//
// [proxmox docs]: https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ha/status/manager_status
func ClusterHAStatusManager(c *pve.Client) (res any, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/ha/status/manager_status",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
