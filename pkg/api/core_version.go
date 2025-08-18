package api

import "net/http"

// CoreGetVersion API version details,
// including some parts of the global datacenter config.
func (c API) CoreGetVersion() (res struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/version",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
