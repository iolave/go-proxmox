package pve

import (
	"net/http"
	"path"
	"strconv"
)

type PVENodeFirewallService struct {
	api *PVE
}

func newPVENodeFirewallService(api *PVE) *PVENodeFirewallService {
	service := new(PVENodeFirewallService)
	service.api = api
	return service
}

type GetNodeFirewallRulesResponse[Position interface{ int | string }] struct {
	Action          string           `json:"action"`
	Comment         string           `json:"comment"`
	Destination     string           `json:"dest"`
	DestinationPort string           `json:"dport"`
	Enable          int              `json:"enable"`
	ICMPType        string           `json:"icmp-type"`
	Interface       string           `json:"iface"`
	IPVersion       int              `json:"ipversion"`
	LogLevel        FirewallLogLevel `json:"log"`
	Macro           string           `json:"macro"`
	Pos             Position         `json:"pos"`
	Proto           string           `json:"proto"`
	Source          string           `json:"source"`
	Sport           string           `json:"sport"`
	Type            string           `json:"type"`
}

// GetRules retrieves node's firewall rules.
func (s *PVENodeFirewallService) GetRules(node string) ([]GetNodeFirewallRulesResponse[int], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules")

	res := &[]GetNodeFirewallRulesResponse[int]{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// GetRulesByPos Retrieves a single node's firewall rule using rule's position (pos) as an index.
func (s *PVENodeFirewallService) GetRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules", strconv.Itoa(pos))

	res := &GetNodeFirewallRulesResponse[string]{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// ReadLog Retrieves node's firewall log entries.
//
// TODO: Add missing limit, since, start, until parameters shown in [docs].
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log
func (s *PVENodeFirewallService) ReadLog(node string) ([]FirewallLogEntry, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/log")

	res := &[]FirewallLogEntry{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}
