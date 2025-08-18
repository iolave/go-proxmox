package api

import "net/http"

// ClusterGetTasks List recent tasks (cluster wide).
func (c API) ClusterGetTasks() (res []struct {
	UpID      string  `json:"upid"`
	Node      *string `json:"node"`
	Status    *string `json:"status"`
	ID        *string `json:"id"`
	StartTime *int    `json:"starttime"`
	Saved     *string `json:"saved"`
	User      *string `json:"user"`
	EndTime   *int    `json:"endtime"`
	Type      *string `json:"type"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/tasks",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
