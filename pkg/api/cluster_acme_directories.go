package api

import "net/http"

// ClusterACMEGetDirectories Get named known ACME directory endpoints.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (c API) ClusterACMEGetDirectories() (res []struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/acme/directories",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
