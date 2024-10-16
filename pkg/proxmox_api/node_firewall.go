package proxmoxapi

import (
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

func (api *ProxmoxAPI) GetNodeFirewallRules(node string) ([]GetNodeFirewallRulesResponse[int], error) {
	return sendGetRequest[[]GetNodeFirewallRulesResponse[int]](api, path.Join("/nodes", node, "/firewall/rules"))
}

func (api *ProxmoxAPI) GetNodeFirewallRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error) {
	return sendGetRequest[GetNodeFirewallRulesResponse[string]](api, path.Join("/nodes", node, "/firewall/rules", strconv.Itoa(pos)))
}

func (api *ProxmoxAPI) ReadNodeFirewallLog(node string) ([]FirewallLogEntry, error) {
	return sendGetRequest[[]FirewallLogEntry](api, path.Join("/nodes", node, "/firewall/log"))
}
