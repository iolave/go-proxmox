package api

import "net/http"

type ClusterGetLogRequest struct {
	// Maximum number of entries (1 - N)
	MaxEntries int `in:"query=max;omitempty"`
}

// ClusterGetLog Read cluster log.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (c API) ClusterGetLog(req ClusterGetLogRequest) (res []struct {
	ID      string `json:"id"`
	Message string `json:"msg"`
	UID     string `json:"uid"`
	Time    int    `json:"time"`
	Node    string `json:"node"`
	PID     int    `json:"pid"`
	Pri     int    `json:"pri"`
	User    string `json:"user"`
	Tag     string `json:"tag"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/log",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}
