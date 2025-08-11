package pve

import (
	"fmt"
	"net/http"
	"net/url"
)

type PVEClusterFirewallService struct {
	api *Client
}

func newPVEClusterFirewallService(api *Client) *PVEClusterFirewallService {
	service := new(PVEClusterFirewallService)
	service.api = api
	return service
}

type GetClusterFirewallAliasesResponse struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// GetAliases retrieves all cluster firewall aliases.
func (s *PVEClusterFirewallService) GetAliases() ([]GetClusterFirewallAliasesResponse, error) {
	method := http.MethodGet
	path := "/cluster/firewall/aliases"

	res := &[]GetClusterFirewallAliasesResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// GetAlias retrieves cluster firewall alias by it's name.
func (s *PVEClusterFirewallService) GetAlias(name string) (GetClusterFirewallAliasesResponse, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)

	res := &GetClusterFirewallAliasesResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// CreateAlias creates a cluster firewall IP or Network Alias.
func (s *PVEClusterFirewallService) CreateAlias(name, cidr string, comment *string) error {
	method := http.MethodPost
	path := "/cluster/firewall/aliases"

	payload := url.Values{}
	payload.Add("name", name)
	payload.Add("cidr", cidr)

	if comment != nil {
		payload.Add("comment", *comment)
	}

	err := s.api.client.sendReq(method, path, &payload, nil)
	return err
}

// UpdateAlias updates a cluster firewall IP or Network alias.
//
// Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
func (s *PVEClusterFirewallService) UpdateAlias(name, cidr string, comment *string, digest *string, rename *string) error {
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

	err := s.api.client.sendReq(method, path, &payload, nil)
	return err
}

// DeleteAlias removes a cluster firewall IP or Network alias.
//
// Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
func (s *PVEClusterFirewallService) DeleteAlias(name string, digest *string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/cluster/firewall/aliases/%s", name)

	payload := url.Values{}

	if digest != nil {
		payload.Add("digest", *digest)
	}

	err := s.api.client.sendReq(method, path, &payload, nil)
	return err
}

type GetClusterFirewallIPSetResponse struct {
	Digest  string `json:"digest"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// GetIPSet retrieves all cluster firewall IPSets.
func (s *PVEClusterFirewallService) GetIPSet() ([]GetClusterFirewallIPSetResponse, error) {
	method := http.MethodGet
	path := "/cluster/firewall/ipset"

	res := &[]GetClusterFirewallIPSetResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

type GetClusterFirewallRulesResponse struct {
	Pos int `json:"pos"`
}

// GetRules retrieves all cluster firewall rules.
func (s *PVEClusterFirewallService) GetRules() ([]GetClusterFirewallRulesResponse, error) {
	method := http.MethodGet
	path := "/cluster/firewall/rules"

	res := &[]GetClusterFirewallRulesResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}
