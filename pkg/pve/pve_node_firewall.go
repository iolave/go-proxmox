package pve

import (
	"net/http"
	"path"
	"strconv"
)

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

// GetNodeRules retrieves node's firewall rules.
func (api *PVE) GetNodeRules(node string) ([]GetNodeFirewallRulesResponse[int], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules")

	res := &[]GetNodeFirewallRulesResponse[int]{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// GetNodeRulesByPos Retrieves a single node's firewall rule using rule's position (pos) as an index.
func (api *PVE) GetNodeRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules", strconv.Itoa(pos))

	res := &GetNodeFirewallRulesResponse[string]{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// ReadNodeLog Retrieves node's firewall log entries.
//
// TODO: Add missing limit, since, start, until parameters shown in [docs].
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log
func (api *PVE) ReadNodeLog(node string) ([]FirewallLogEntry, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/log")

	res := &[]FirewallLogEntry{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}
