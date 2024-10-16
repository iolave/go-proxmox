package proxmoxapi

type NodeStatus string

const (
	NODE_STATUS_ONLINE  NodeStatus = "online"
	NODE_STATUS_OFFLINE NodeStatus = "offline"
	NODE_STATUS_UNKNOWN NodeStatus = "unknown"
)

type GetNodesResponse struct {
	Node           string     `json:"node"`
	Status         NodeStatus `json:"status"`
	CPU            float64    `json:"cpu"`
	Level          string     `json:"level"`
	MaxCpu         int        `json:"maxcpu"`
	MaxMem         int        `json:"maxmem"`
	Mem            int        `json:"mem"`
	SSLFingerprint string     `json:"ssl_fingerprint"`
	Uptime         int        `json:"uptime"`
}

func (api *ProxmoxAPI) GetNodes() ([]GetNodesResponse, error) {
	return sendGetRequest[[]GetNodesResponse](api, "/nodes")
}
