package apiclient

import "net/http"

func (c HTTPClient) ClusterGetStatus() (res []struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	IP      *string `json:"ip"`
	Level   *string `json:"level"`
	Local   *int    `json:"local"`
	NodeID  *int    `json:"nodeid"`
	Nodes   *int    `json:"nodes"`
	Online  *int    `json:"online"`
	QuoRate *int    `json:"quorate"`
	Version *int    `json:"version"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/status",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
