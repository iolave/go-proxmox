package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

func ClusterBackupInfoNotBackedUp(c *pve.Client) (res []struct {
	Type string  `json:"type"`
	VMID int     `json:"vmid"`
	Name *string `json:"name"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/backup-info/not-backed-up",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
