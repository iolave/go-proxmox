package pve

import (
	"errors"
	"net/http"
)

type PVENodeService struct {
	api      *Client
	APT      *PVENodeAPTService
	Firewall *PVENodeFirewallService
	Storage  *PVENodeStorageService
}

func newPVENodeService(api *Client) *PVENodeService {
	service := new(PVENodeService)
	service.api = api
	service.APT = newPVENodeAPTService(api)
	service.Firewall = newPVENodeFirewallService(api)
	service.Storage = newPVENodeStorageService(api)
	return service
}

// Proxmox availabe node statuses
type NodeStatus string

const (
	NODE_STATUS_ONLINE  NodeStatus = "online"
	NODE_STATUS_OFFLINE NodeStatus = "offline"
	NODE_STATUS_UNKNOWN NodeStatus = "unknown"
)

type GetNodeResponse struct {
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
func (service *PVENodeService) GetAll() ([]GetNodeResponse, error) {
	method := http.MethodGet
	path := "/nodes"

	res := &[]GetNodeResponse{}
	err := service.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// Get retrieves a single nodes.
func (service *PVENodeService) Get(node string) (GetNodeResponse, error) {
	nodes, err := service.GetAll()

	if err != nil {
		return GetNodeResponse{}, err
	}

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Node == "node" {
			return nodes[i], nil
		}
	}

	return GetNodeResponse{}, errors.New("Node not found")
}
