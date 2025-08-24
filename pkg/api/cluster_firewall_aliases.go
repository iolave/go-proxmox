package api

import "net/http"

type ClusterFirewallAliasesGetResponse struct {
	// Network/IP specification in CIDR format.
	CIDR string `json:"cidr"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `json:"digest"`

	// Alias name.
	Name string `json:"name"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// ClusterFirewallAliasesList List aliases.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterFirewallAliasesList() (res []ClusterFirewallAliasesGetResponse, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/aliases",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterFirewallAliasesGet Get IP or Network Alias.
//
// Parameters:
//
//   - name is the alias name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterFirewallAliasesGet(name string) (res ClusterFirewallAliasesGetResponse, err error) {
	req := struct {
		Name string `in:"path=name"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterFirewallAliasesNewRequest struct {
	// Network/IP specification in CIDR format.
	CIDR string `in:"query=cidr;nonzero"`

	// Alias name.
	Name string `in:"query=name;nonzero"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`
}

// ClusterFirewallAliasesNew Create IP or Network Alias.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallAliasesNew(req ClusterFirewallAliasesNewRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/aliases",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterFirewallAliasesUpdateRequest struct {
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

// ClusterFirewallAliasesUpdate Update IP or Network Alias.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallAliasesUpdate(req ClusterFirewallAliasesUpdateRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterFirewallAliasesDelete Delete IP or Network Alias.
//
// Parameters:
//
//   - name is the alias name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallAliasesDelete(name string) (err error) {
	req := struct {
		Name string `in:"path=name"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/aliases/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
