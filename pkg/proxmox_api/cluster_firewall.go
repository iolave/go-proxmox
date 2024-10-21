package proxmoxapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type GetAliasResponse struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasResponse, error) {
	return sendRequest[[]GetAliasResponse](http.MethodGet, api, "/cluster/firewall/aliases", nil)
}

func (api *ProxmoxAPI) GetClusterFirewallAlias(name string) (GetAliasResponse, error) {
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)
	return sendRequest[GetAliasResponse](http.MethodGet, api, path, nil)
}

func (api *ProxmoxAPI) CreateClusterFirewallAlias(name, cidr string, comment *string) error {
	payload := url.Values{}
	payload.Add("name", name)
	payload.Add("cidr", cidr)

	if comment != nil {
		payload.Add("comment", *comment)
	}

	_, err := sendRequest[any](http.MethodPost, api, "/cluster/firewall/aliases", &payload)
	return err
}

func (api *ProxmoxAPI) DeleteClusterFirewallAlias(name string, digest *string) error {
	payload := url.Values{}
	if digest != nil {
		payload.Add("digest", *digest)
	}

	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)
	_, err := sendRequest[any](http.MethodDelete, api, path, &payload)
	return err
}

type GetIPSetResponse struct {
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (api *ProxmoxAPI) GetClusterFirewallIPSet() ([]GetIPSetResponse, error) {
	return sendRequest[[]GetIPSetResponse](http.MethodGet, api, "/cluster/firewall/ipset", nil)
}

type GetRulesResponse struct {
	Pos int `json:"pos"`
}

func (api *ProxmoxAPI) GetClusterFirewallRules() ([]GetRulesResponse, error) {
	return sendRequest[[]GetRulesResponse](http.MethodGet, api, "/cluster/firewall/rules", nil)
}
