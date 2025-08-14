package apiclient

import "net/http"

type ClusterGetResourcesRequest struct {
	Type string `in:"query=type"`
}

func (c HTTPClient) ClusterGetResources(req ClusterGetResourcesRequest) (res []struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	CGroupMode *int    `json:"cgroup-mode"`
	Content    *string `json:"content"`
	CPU        *int    `json:"cpu"`
	Disk       *int    `json:"disk"`
	DiskRead   *int    `json:"diskread"`
	DiskWrite  *int    `json:"diskwrite"`
	HAState    *string `json:"hastate"`
	Level      *string `json:"level"`
	Lock       *string `json:"lock"`
	MaxCPU     *int    `json:"maxcpu"`
	MaxDisk    *int    `json:"maxdisk"`
	MaxMem     *int    `json:"maxmem"`
	Mem        *int    `json:"mem"`
	MemHost    *int    `json:"memhost"`
	Name       *string `json:"name"`
	NetIn      *int    `json:"netin"`
	NetOut     *int    `json:"netout"`
	Node       *string `json:"node"`
	PluginType *string `json:"plugintype"`
	Pool       *string `json:"pool"`
	Status     *string `json:"status"`
	Storage    *string `json:"storage"`
	Tags       *string `json:"tags"`
	Template   *int    `json:"template"`
	Uptime     *int    `json:"uptime"`
	VMID       *int    `json:"vmid"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/resources",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}
