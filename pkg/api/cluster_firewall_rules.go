package api

import (
	"net/http"
	"strconv"

	"github.com/iolave/go-errors"
)

type ClusterFirewallRulesGetResponse struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`

	// Rule type.
	//
	// 	in | out | forward | group
	Type *string `json:"type"`

	// Descriptive comment.
	Comment string `json:"comment"`

	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`

	//Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`

	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	DPort *string `json:"dport"`

	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`

	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	ICMPType *string `json:"icmp-type"`

	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`

	// Log level for firewall rule.
	//
	// 	emerg | alert | crit | err | warning | notice | info | debug | nolog
	Log *string `json:"log"`

	// Use predefined standard macro.
	Macro *string `json:"macro"`

	// rule at position <pos>.
	Pos *int `json:"pos"`

	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`

	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`

	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	SPort *string `json:"sport"`
}

// ClusterFirewallRulesGetAll List rules.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallRulesGetAll() (res []ClusterFirewallRulesGetResponse, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/rules",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterFirewallRulesGet Get single rule data.
//
// Parameters:
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallRulesGet(pos int) (ClusterFirewallRulesGetResponse, error) {
	res := struct {
		ClusterFirewallRulesGetResponse
		Pos string `json:"pos"`
	}{}

	req := struct {
		Pos int `in:"path=pos"`
	}{
		Pos: pos,
	}

	err := s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/rules/{pos}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})
	if err != nil {
		return ClusterFirewallRulesGetResponse{}, err
	}

	pos, err = strconv.Atoi(res.Pos)
	if err != nil {
		return ClusterFirewallRulesGetResponse{}, errors.NewInternalServerError(
			"failed to parse response",
			err,
		)
	}
	res.ClusterFirewallRulesGetResponse.Pos = Int(pos)

	return res.ClusterFirewallRulesGetResponse, err
}

type ClusterFirewallRulesPostRequest struct {
	Action   string `in:"query=action;omitempty"`
	Type     string `in:"query=type;omitempty"`
	Comment  string `in:"query=comment;omitempty"`
	Dest     string `in:"query=dest;omitempty"`
	Digest   string `in:"query=digest;omitempty"`
	DPort    string `in:"query=dport;omitempty"`
	Enable   *int   `in:"query=enable;omitempty"`
	ICMPType string `in:"query=icmp-type;omitempty"`
	Iface    string `in:"query=iface;omitempty"`
	Log      string `in:"query=log;omitempty"`
	Macro    string `in:"query=macro;omitempty"`
	Pos      *int   `in:"query=pos;omitempty"`
	Proto    string `in:"query=proto;omitempty"`
	Source   string `in:"query=source;omitempty"`
	SPort    string `in:"query=sport;omitempty"`
}

// ClusterFirewallRulesPost Add rule.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallRulesPost(req ClusterFirewallRulesPostRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/rules",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterFirewallRulesUpdateRequest struct {
	ClusterFirewallRulesPostRequest

	// Rule position.
	Pos *int `in:"path=pos"`

	// Move rule to new position. Other arguments are ignored.
	MoveTo *int `in:"query=moveto;omitempty"`
}

// ClusterFirewallRulesUpdate Modify rule data.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallRulesUpdate(req ClusterFirewallRulesUpdateRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/rules/{pos}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

type ClusterFirewallRulesDeleteRequest struct {
	// Rule position.
	Pos *int `in:"path=pos"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`
}

// ClusterFirewallRulesDelete Delete rule.
//
// Parameters:
//
//   - pos is the position of a rule.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallRulesDelete(req ClusterFirewallRulesDeleteRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/rules/{pos}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
