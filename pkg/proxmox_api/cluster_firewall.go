package proxmoxapi

type GetAliasesResponse struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasesResponse, error) {
	return sendGetRequest[[]GetAliasesResponse](api, "/cluster/firewall/aliases")
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
