package apiclient

import "net/http"

func (c HTTPClient) CoreGetVersion() (res struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}, err error) {
	err = c.sendPVERequest(PVERequest{
		Path:   "/api2/json/version",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
