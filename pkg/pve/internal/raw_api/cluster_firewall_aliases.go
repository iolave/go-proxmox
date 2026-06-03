package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type GETClusterFirewallAliasesResponse struct {
	// Network/IP specification in CIDR format.
	CIDR string `json:"cidr"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `json:"digest"`

	// Alias name.
	Name string `json:"name"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// GETClusterFirewallAliases List aliases.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func GETClusterFirewallAliases(c *pve.Client) (res []GETClusterFirewallAliasesResponse, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/firewall/aliases",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// GETClusterFirewallAliases_name Get IP or Network Alias.
//
// Parameters:
//
//   - name is the alias name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func GETClusterFirewallAliases_name(c *pve.Client, name string) (res GETClusterFirewallAliasesResponse, err error) {
	req := struct {
		Name string `in:"path=name"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type POSTClusterFirewallAliasesRequest struct {
	// Network/IP specification in CIDR format.
	CIDR string `in:"query=cidr;nonzero"`

	// Alias name.
	Name string `in:"query=name;nonzero"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`
}

// POSTClusterFirewallAliases Create IP or Network Alias.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func POSTClusterFirewallAliases(c *pve.Client, req POSTClusterFirewallAliasesRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/firewall/aliases",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type PUTClusterFirewallAliasesRequest struct {
	// Network/IP specification in CIDR format.
	CIDR string `in:"query=cidr;omitempty"`

	// Name of the alias.
	Name string `in:"path=name;nonzero"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Rename an existing alias.
	Rename string `in:"query=rename;omitempty"`
}

// PUTClusterFirewallAliases Update IP or Network Alias.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func PUTClusterFirewallAliases(c *pve.Client, req PUTClusterFirewallAliasesRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// DELETEClusterFirewallAliases Delete IP or Network Alias.
//
// Parameters:
//
//   - name is the alias name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func DELETEClusterFirewallAliases(c *pve.Client, name string) (err error) {
	req := struct {
		Name string `in:"path=name"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
