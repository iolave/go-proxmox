package api

import "net/http"

// ClusterFirewallGroupsList List security groups.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (c API) ClusterFirewallGroupsList() (res []struct {
	Group   string `json:"group"`
	Comment string `json:"comment"`
	Digest  string `json:"digest"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/groups",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterFirewallGroupsPostRequest struct {
	// Security Group name.
	Name string `in:"query=group;omitempty"`

	// Security Group comment.
	Comment string `in:"query=comment;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Rename/update an existing security group. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing group.
	Rename string `in:"query=rename;omitempty"`
}

// ClusterFirewallGroupsPost Create new security group.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallGroupsPost(req ClusterFirewallGroupsPostRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterFirewallGroupsGetResponse struct {
	Action    string  `json:"action"`
	Pos       int     `json:"pos"`
	Type      string  `json:"type"`
	Comment   string  `json:"comment"`
	Dest      *string `json:"dest"`
	DPort     *string `json:"dport"`
	Enable    *int    `json:"enable"`
	ICMPType  *string `json:"icmp-type"`
	Iface     *string `json:"iface"`
	IPVersion *int    `json:"ipversion"`
	Log       *string `json:"log"`
	Macro     *string `json:"macro"`
	Protocol  *string `json:"proto"`
	Source    *string `json:"source"`
	SPort     *string `json:"sport"`
}

// ClusterFirewallGroupsGet List rules.
//
// Parameters:
//
//   - name is the security group name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterFirewallGroupsGet(name string) (res []ClusterFirewallGroupsGetResponse, err error) {
	req := struct {
		Group string `in:"path=group"`
	}{
		Group: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterFirewallGroupsNewRuleRequest struct {
	// Security Group name.
	Name string `in:"path=group;nonzero"`

	// Rule type.
	//
	// 	in | out | forward | group
	Type string `in:"query=type;nonzero"`

	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `in:"query=action;nonzero"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`

	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest string `in:"query=dest;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport string `in:"query=dport;omitempty"`

	// Flag to enable/disable a rule.
	Enable *int `in:"query=enable;omitempty"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	ICMPType string `in:"query=icmp-type;omitempty"`

	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface string `in:"query=iface;omitempty"`

	// Log level for firewall rule.
	//
	// 	emerg | alert | crit | err | warning | notice | info | debug | nolog
	Log string `in:"query=log;omitempty"`

	// Use predefined standard macro.
	Macro string `in:"query=macro;omitempty"`

	// Update rule at position <pos>.
	Pos *int `in:"query=pos;omitempty"`

	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto string `in:"query=proto;omitempty"`

	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source string `in:"query=source;omitempty"`

	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport string `in:"query=sport;omitempty"`
}

// ClusterFirewallGroupsNewRule Create new rule.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallGroupsNewRule(req ClusterFirewallGroupsNewRuleRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterFirewallGroupsDelete Delete security group.
//
// Parameters:
//
//   - name is the security group name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallGroupsDelete(name string) (err error) {
	req := struct {
		Group string `in:"path=group"`
	}{
		Group: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

// ClusterFirewallGroupsGetRule Get single rule data.
//
// Parameters:
//
//   - name is the security group name.
//   - pos is the rule position.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterFirewallGroupsGetRule(name string, pos int) (res ClusterFirewallGroupsGetResponse, err error) {
	req := struct {
		Group string `in:"path=group"`
		Pos   int    `in:"path=pos"`
	}{
		Group: name,
		Pos:   pos,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}/{pos}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterFirewallGroupsUpdateRuleRequest struct {
	// Security Group name.
	Name string `in:"path=group;nonzero"`

	// Rule position.
	Pos int `in:"path=pos"`

	// Move rule to new position <moveto>. Other arguments are ignored.
	MoveTo *int `in:"query=moveto;omitempty"`

	// Name and Pos will be ignored from this property.
	ClusterFirewallGroupsNewRuleRequest
}

// ClusterFirewallGroupsUpdateRule Modify rule data.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallGroupsUpdateRule(req ClusterFirewallGroupsUpdateRuleRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}/{pos}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterFirewallGroupsDeleteRule Delete rule.
//
// Parameters:
//
//   - name is the security group name.
//   - pos is the rule position.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterFirewallGroupsDeleteRule(name string, pos int) (err error) {
	req := struct {
		Group string `in:"path=group"`
		Pos   int    `in:"path=pos"`
	}{
		Group: name,
		Pos:   pos,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/groups/{group}/{pos}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
