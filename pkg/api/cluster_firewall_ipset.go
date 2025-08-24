package api

import "net/http"

// ClusterFirewallIPSetList List IPSets
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallIPSetList() (res []struct {
	Name    string `json:"name"`
	Digest  string `json:"digest"`
	Comment string `json:"comment"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/ipset",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterFirewallIPSetGet List IPSet content
//
// Parameters:
//
//   - name is the IPSet name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallIPSetGet(name string) (res []struct {
	CIDR    string `json:"cidr"`
	Digest  string `json:"digest"`
	Comment string `json:"comment"`
	NoMatch *int   `json:"nomatch"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterFirewallIPSetPostRequest struct {
	// IP set name.
	Name string `in:"query=name;nonzero"`

	// IP set comment.
	Comment string `in:"query=comment;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Rename string `in:"query=rename;omitempty"`
}

// ClusterFirewallIPSetPost Create new IPSet
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallIPSetPost(req ClusterFirewallIPSetPostRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterFirewallIPSetAddCIDRRequest struct {
	// IP set name.
	Name string `in:"path=name;nonzero"`

	// Network/IP specification in CIDR format.
	CIDR string `in:"query=cidr;nonzero"`

	// IP set comment.
	Comment string `in:"query=comment;omitempty"`

	// TODO: add proper comment
	NoMatch *int `in:"query=nomatch;omitempty"`
}

// ClusterFirewallIPSetAddCIDR Add IP or Network to IPSet.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallIPSetAddCIDR(req ClusterFirewallIPSetAddCIDRRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset/{name}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterFirewallIPSetUpdateCIDRRequest struct {
	Name    string `in:"path=name;nonzero"`
	CIDR    string `in:"path=cidr;nonzero"`
	Comment string `in:"query=comment;omitempty"`
	Digest  string `in:"query=digest;omitempty"`
	NoMatch *int   `in:"query=nomatch;omitempty"`
}

// ClusterFirewallIPSetUpdateCIDR Update IP or Network settings
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallIPSetUpdateCIDR(req ClusterFirewallIPSetUpdateCIDRRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset/{name}/{cidr}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterFirewallIPSetDeleteCIDR Remove IP or Network from IPSet.
//
// Parameters:
//
//   - name is the IPSet name.
//   - cidr is the IP or Network to remove.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallIPSetDeleteCIDR(name string, cidr string) (err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
		CIDR string `in:"path=cidr;nonzero"`
	}{
		Name: name,
		CIDR: cidr,
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset/{name}/{cidr}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

// ClusterFirewallIPSetDelete Delete IPSet.
//
// Parameters:
//
//   - name is the IPSet name.
//   - force delete all members of the IPSet, if there are any.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallIPSetDelete(name string, force bool) (err error) {
	req := struct {
		Name  string `in:"path=name;nonzero"`
		Force *int   `in:"query=force;omitempty"`
	}{
		Name: name,
	}

	if force {
		req.Force = Int(1)
	} else {
		req.Force = Int(0)
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/ipset/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
