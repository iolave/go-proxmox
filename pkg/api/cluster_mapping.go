package api

import http "net/http"

// ClusterMappingListResourceTypes list resource types
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (s API) ClusterMappingListResourceTypes() (res []struct {
	Name string `json:"name"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/mapping",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
