package pve

import (
	"fmt"
	"net/http"
	"net/url"
)

type GetClusterFirewallAliasesResponse struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// GetClusterFirewallAliases retrieves all cluster firewall aliases.
func (api *PVE) GetClusterFirewallAliases() ([]GetClusterFirewallAliasesResponse, error) {
	method := http.MethodPost
	path := "/cluster/firewall/aliases"

	res := &[]GetClusterFirewallAliasesResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// GetClusterFirewallAlias retrieves cluster firewall alias by it's name.
func (api *PVE) GetClusterFirewallAlias(name string) (GetClusterFirewallAliasesResponse, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)

	res := &GetClusterFirewallAliasesResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// CreateClusterFirewallAlias creates a cluster firewall IP or Network Alias.
func (api *PVE) CreateClusterFirewallAlias(name, cidr string, comment *string) error {
	method := http.MethodPost
	path := "/cluster/firewall/aliases"

	payload := url.Values{}
	payload.Add("name", name)
	payload.Add("cidr", cidr)

	if comment != nil {
		payload.Add("comment", *comment)
	}

	err := api.httpClient.sendReq(method, path, &payload, nil)
	return err
}

// UpdateClusterFirewallAlias updates a cluster firewall IP or Network alias.
//
// Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
func (api *PVE) UpdateClusterFirewallAlias(name, cidr string, comment *string, digest *string, rename *string) error {
	method := http.MethodPut
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)

	payload := url.Values{}
	payload.Add("cidr", cidr)

	if comment != nil {
		payload.Add("comment", *comment)
	}

	if digest != nil {
		payload.Add("digest", *digest)
	}

	if rename != nil {
		payload.Add("rename", *rename)
	}

	err := api.httpClient.sendReq(method, path, &payload, nil)
	return err
}

// DeleteClusterFirewallAlias removes a cluster firewall IP or Network alias.
//
// Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
func (api *PVE) DeleteClusterFirewallAlias(name string, digest *string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)

	payload := url.Values{}

	if digest != nil {
		payload.Add("digest", *digest)
	}

	err := api.httpClient.sendReq(method, path, &payload, nil)
	return err
}

type GetClusterFirewallIPSetResponse struct {
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// GetClusterFirewallIPSet retrieves all cluster firewall IPSets.
func (api *PVE) GetClusterFirewallIPSet() ([]GetClusterFirewallIPSetResponse, error) {
	method := http.MethodPost
	path := "/cluster/firewall/ipset"

	res := &[]GetClusterFirewallIPSetResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

type GetClusterFirewallRulesResponse struct {
	Pos int `json:"pos"`
}

// GetClusterFirewallRules retrieves all cluster firewall rules.
func (api *PVE) GetClusterFirewallRules() ([]GetClusterFirewallRulesResponse, error) {
	method := http.MethodGet
	path := "/cluster/firewall/aliases"

	res := &[]GetClusterFirewallRulesResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}
