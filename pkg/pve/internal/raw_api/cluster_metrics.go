package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterMetricsGetRequest struct {
	// Also return historic values. Returns full available metric history unless `start-time` is also set
	History *int `in:"query=history;omitempty"`

	// Only return metrics for the current node instead of the whole cluster
	LocalOnly *int `in:"query=local-only;omitempty"`

	// Only include metrics with a timestamp > start-time.
	StartTime *int `in:"query=start-time;omitempty"`
}

// ClusterMetricsGet Retrieve metrics of the cluster.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterMetricsGet(c *pve.Client, req *ClusterMetricsGetRequest) (res []struct {
	ID        string `json:"id"`
	Metric    string `json:"metric"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/metrics/export",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}
