package proxmoxapi

import "fmt"

type GetAliasResponse struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasResponse, error) {
	return sendGetRequest[[]GetAliasResponse](api, "/cluster/firewall/aliases")
}

func (api *ProxmoxAPI) GetClusterFirewallAlias(name string) (GetAliasResponse, error) {
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)
	return sendGetRequest[GetAliasResponse](api, path)
}

type GetIPSetResponse struct {
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (api *ProxmoxAPI) GetClusterFirewallIPSet() ([]GetIPSetResponse, error) {
	return sendGetRequest[[]GetIPSetResponse](api, "/cluster/firewall/ipset")
}

type GetRulesResponse struct {
	Pos int `json:"pos"`
}

func (api *ProxmoxAPI) GetClusterFirewallRules() ([]GetRulesResponse, error) {
	return sendGetRequest[[]GetRulesResponse](api, "/cluster/firewall/rules")
}
