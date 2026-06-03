package cluster

import (
	"github.com/iolave/go-proxmox/internal/helpers"
	"github.com/iolave/go-proxmox/pkg/pve/internal/raw_api"
)

type GetFWRuleResponse struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`

	// Rule position.
	Pos int `json:"pos"`

	// Rule type.
	//
	// 	in | out | forward | group
	Type string `json:"type"`

	// Flag to enable/disable a rule.
	Enable bool `json:"enable"`

	// Descriptive comment.
	Comment string `json:"comment"`

	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`

	//Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`

	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	DPort *string `json:"destPort"`

	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	ICMPType *string `json:"icmpType"`

	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`

	// Log level for firewall rule.
	//
	// 	emerg | alert | crit | err | warning | notice | info | debug | nolog
	Log *string `json:"log"`

	// Use predefined standard macro.
	Macro *string `json:"macro"`

	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`

	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`

	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	SPort *string `json:"srcPort"`
}

// GetFWRules returns a list of all cluster firewall rules.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWRules() ([]GetFWRuleResponse, error) {
	rules, err := rawapi.GETClusterFirewallRules(s.c)
	if err != nil {
		return nil, err
	}

	return helpers.Map(rules, func(v rawapi.GETClusterFirewallRulesResponse) GetFWRuleResponse {
		mapped := GetFWRuleResponse{
			Action:   *v.Action,
			Pos:      *v.Pos,
			Type:     *v.Type,
			Comment:  v.Comment,
			Dest:     v.Dest,
			Digest:   v.Digest,
			DPort:    v.DPort,
			ICMPType: v.ICMPType,
			Iface:    v.Iface,
			Log:      v.Log,
			Macro:    v.Macro,
			Proto:    v.Proto,
			Source:   v.Source,
			SPort:    v.SPort,
		}

		if v.Enable != nil {
			enabled := helpers.IntToBool(*v.Enable)
			mapped.Enable = enabled
		} else {
			mapped.Enable = false
		}

		return mapped
	}), nil
}

// GetFWRule returns a cluster firewall rule by it's
// position.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWRule(pos int) (GetFWRuleResponse, error) {
	res, err := rawapi.GETClusterFirewallRules_pos(s.c, pos)
	if err != nil {
		return GetFWRuleResponse{}, err
	}

	mapped := GetFWRuleResponse{
		Action:   *res.Action,
		Pos:      *res.Pos,
		Type:     *res.Type,
		Comment:  res.Comment,
		Dest:     res.Dest,
		Digest:   res.Digest,
		DPort:    res.DPort,
		ICMPType: res.ICMPType,
		Iface:    res.Iface,
		Log:      res.Log,
		Macro:    res.Macro,
		Proto:    res.Proto,
		Source:   res.Source,
		SPort:    res.SPort,
	}

	if res.Enable != nil {
		enabled := helpers.IntToBool(*res.Enable)
		mapped.Enable = enabled
	} else {
		mapped.Enable = false
	}

	return mapped, nil
}
