package api

import "net/http"

// ClusterACMEGetChallSchema Get schema of ACME challenge types.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (c API) ClusterACMEGetChallSchema() (res []struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Schema struct {
		Name   string `json:"name,omitempty"`
		Fields map[string]struct {
			Type string `json:"type"`
			Desc string `json:"description"`
		} `json:"fields,omitempty"`
	} `json:"schema"`
	Type string `json:"type"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/acme/challenge-schema",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
