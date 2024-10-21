package proxmoxapi

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

func (api *ProxmoxAPI) GetNodeFirewallRules(node string) ([]GetNodeFirewallRulesResponse[int], error) {
	return sendRequest[[]GetNodeFirewallRulesResponse[int]](http.MethodGet, api, path.Join("/nodes", node, "/firewall/rules"), nil)
}

func (api *ProxmoxAPI) GetNodeFirewallRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error) {
	return sendRequest[GetNodeFirewallRulesResponse[string]](http.MethodGet, api, path.Join("/nodes", node, "/firewall/rules", strconv.Itoa(pos)), nil)
}

func (api *ProxmoxAPI) ReadNodeFirewallLog(node string) ([]FirewallLogEntry, error) {
	return sendRequest[[]FirewallLogEntry](http.MethodGet, api, path.Join("/nodes", node, "/firewall/log"), nil)
}
