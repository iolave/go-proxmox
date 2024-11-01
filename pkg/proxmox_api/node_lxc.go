package proxmoxapi

import (
	"net/http"
	"path"
)

type GetNodeLxcsResponse struct {
	Status  LxcStatus `json:"status"`
	VmID    int       `json:"vmid"`
	Cpus    *int      `json:"cpus"`
	Lock    *string   `json:"lock"`
	MaxDisk *int      `json:"maxdisk"`
	MaxMem  *int      `json:"maxmem"`
	MaxSwap *int      `json:"maxswap"`
	Name    *string   `json:"name"`
	Tags    *string   `json:"tags"`
	Uptime  *int      `json:"uptime"`
}

// GetNodeLxcs returns node's lxc index per node.
func (api *ProxmoxAPI) GetNodeLxcs(node string) ([]GetNodeLxcsResponse, error) {
	path := path.Join("/nodes", node, "/lxc")
	return sendRequest[[]GetNodeLxcsResponse](http.MethodGet, api, path, nil)
}
