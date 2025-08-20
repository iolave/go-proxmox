package api

import "net/http"

func (s API) ClusterBackupInfoNotBackedUp() (res []struct {
	Type string  `json:"type"`
	VMID int     `json:"vmid"`
	Name *string `json:"name"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/backup-info/not-backed-up",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
