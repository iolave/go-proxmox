package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterNotificationsGetMacherFieldValues Returns known notification metadata fields and their known values
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsGetMacherFieldValues(c *pve.Client) (res []struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/notifications/matcher-field-values",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterNotificationsGetMacherFields Returns known notification metadata fields
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsGetMacherFields(c *pve.Client) (res []struct {
	Name string `json:"name"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/notifications/matcher-fields",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
