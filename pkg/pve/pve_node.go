package pve

import (
	"errors"
	"net/http"
)

// Proxmox availabe node statuses
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

// GetAll retrieves all nodes.
func (api *PVE) GetAll() ([]GetNodesResponse, error) {
	method := http.MethodGet
	path := "/nodes"

	res := &[]GetNodesResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// Get retrieves a single nodes.
func (api *PVE) Get(node string) (GetNodesResponse, error) {
	nodes, err := api.GetAll()

	if err != nil {
		return GetNodesResponse{}, err
	}

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Node == "node" {
			return nodes[i], nil
		}
	}

	return GetNodesResponse{}, errors.New("Node not found")
}
