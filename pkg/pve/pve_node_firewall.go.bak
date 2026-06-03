package pve

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type PVENodeFirewallService struct {
	api *Client
}

func newPVENodeFirewallService(api *Client) *PVENodeFirewallService {
	service := new(PVENodeFirewallService)
	service.api = api
	return service
}

type GetNodeFirewallRuleResponse[Position interface{ int | string }] struct {
	ID              string
	Action          string           `json:"action"`
	Comment         string           `json:"comment"`
	Destination     string           `json:"dest"`
	DestinationPort string           `json:"dport"`
	Enable          int              `json:"enable"`
	ICMPType        string           `json:"icmp-type"`
	Interface       string           `json:"iface"`
	IPVersion       int              `json:"ipversion"`
	LogLevel        FirewallLogLevel `json:"log"`
	Macro           string           `json:"macro"`
	Pos             Position         `json:"pos"`
	Proto           string           `json:"proto"`
	Source          string           `json:"source"`
	Sport           string           `json:"sport"`
	Type            string           `json:"type"`
}

func (r *GetNodeFirewallRuleResponse[any]) setIdFromComment() {
	idxStart := strings.IndexRune(r.Comment, '[')
	if idxStart != 0 {
		return
	}
	idxEnd := strings.IndexRune(r.Comment, ']')
	if idxEnd == -1 {
		return
	}
	substr := r.Comment[idxStart+1 : idxEnd]
	splitted := strings.Split(substr, "=")
	if len(splitted) != 2 {
		return
	}
	if splitted[0] != "id" {
		return
	}
	if err := uuid.Validate(splitted[1]); err != nil {
		return
	}
	r.ID = splitted[1]
}

// GetRules retrieves node's firewall rules.
func (s *PVENodeFirewallService) GetRules(node string) ([]GetNodeFirewallRuleResponse[int], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules")

	res := []GetNodeFirewallRuleResponse[int]{}

	err := s.api.client.sendReq(method, path, nil, &res)

	for i := 0; i < len(res); i++ {
		res[i].setIdFromComment()
	}

	return res, err
}

// GetRulesByPos Retrieves a single node's firewall rule using rule's position (pos) as an index.
func (s *PVENodeFirewallService) GetRulesByPos(node string, pos int) (GetNodeFirewallRuleResponse[string], error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/rules", strconv.Itoa(pos))

	res := GetNodeFirewallRuleResponse[string]{}
	err := s.api.client.sendReq(method, path, nil, &res)
	res.setIdFromComment()

	return res, err
}

// GetRule finds a single node's firewall rule using it's custom id within the comment field. If the rule is not found, both result and error will be nil.
//
// GET /nodes/{node}/firewall/rules requires the "Sys.Audit" permission.
func (s *PVENodeFirewallService) GetRule(node string, id string) (*GetNodeFirewallRuleResponse[int], error) {
	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	rules, err := s.GetRules(node)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if id != rule.ID {
			continue
		}
		return &rule, nil
	}

	return nil, nil
}

type CreateNodeFirewallRuleRequest struct {
	Action          string           `in:"nonzero;form=action"`      // Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name. Format: [A-Za-z][A-Za-z0-9\-\_]+.
	Node            string           `in:"nonzero;path=node"`        // The cluster node name.
	Type            string           `in:"nonzero;form=type"`        // Rule type. Values: in | out | forward | group.
	Comment         string           `in:"omitempty;form=comment"`   // Optional. Descriptive comment.
	Destination     string           `in:"omitempty;form=dest"`      // Optional. Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Digest          string           `in:"omitempty;form=digest"`    // Optional. Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	DestinationPort string           `in:"omitempty;form=dport"`     // Optional. Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Enable          int              `in:"omitempty;form=enable"`    // Optional. Flag to enable/disable a rule. Values: (0 - N).
	ICMPType        string           `in:"omitempty;form=icmp-type"` // Optional. Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	Interface       string           `in:"omitempty;form=iface"`     // Optional. Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	LogLevel        FirewallLogLevel `in:"omitempty;form=log"`       // Optional. Log level for firewall rule.
	Macro           string           `in:"omitempty;form=macro"`     // Optional. Use predefined standard macro.
	Pos             int              `in:"omitempty;form=pos"`       // Optional. Update rule at position <pos>.
	Proto           string           `in:"omitempty;form=proto"`     // Optional. IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Source          string           `in:"omitempty;form=source"`    // Optional. Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Sport           string           `in:"omitempty;form=sport"`     // Optional. Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
}

// NewRule creates a new node firewall rule. It adds metadata within the proxmox comment field with the format [id=uuid].
//
// POST /nodes/{node}/firewall/rules requires the "Sys.Modify" permission.
func (s *PVENodeFirewallService) NewRule(req CreateNodeFirewallRuleRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/firewall/rules"
	uuid := uuid.New().String()
	req.Comment = fmt.Sprintf("[id=%s] %s", uuid, req.Comment)
	if err := s.api.client.sendReq2(method, path, &req, nil); err != nil {
		return "", err
	}
	return uuid, nil
}

// DeleteRuleByPos deletes a node's firewall rule using it's position (pos).
//
//   - TODO: add digest support.
//
// DELETE /nodes/{node}/firewall/rules/{pos} requires the "Sys.Modify" permission.
func (s *PVENodeFirewallService) DeleteRuleByPos(node string, pos int) error {
	type request struct {
		Node     string `in:"nonzero;path=node"`
		Position int    `in:"path=pos"`
	}
	req := request{Node: node, Position: pos}
	method := http.MethodDelete
	path := "/nodes/{node}/firewall/rules/{pos}"

	err := s.api.client.sendReq2(method, path, &req, nil)
	return err
}

// DeleteRule deletes a node's firewall rule using it's id.
//
//   - TODO: add digest support.
//
// DELETE /nodes/{node}/firewall/rules/{pos} requires the "Sys.Modify" permission.
func (s *PVENodeFirewallService) DeleteRule(node, id string) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}

	rules, err := s.GetRules(node)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		if rule.ID == id {
			return s.DeleteRuleByPos(node, rule.Pos)
		}
	}
	return nil
}

// ReadLog Retrieves node's firewall log entries.
//
// TODO: Add missing limit, since, start, until parameters shown in [docs].
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log
func (s *PVENodeFirewallService) ReadLog(node string) ([]FirewallLogEntry, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/firewall/log")

	res := &[]FirewallLogEntry{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}
