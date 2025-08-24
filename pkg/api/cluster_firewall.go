package api

import "net/http"

type ClusterFirewallGetRefsRequest struct {
	// Only list references of specified type.
	//
	// 	alias | ipset
	Type string `in:"query=type;omitempty"`
}

// ClusterFirewallGetRefs Lists possible IPSet/Alias reference which are allowed in source/dest properties.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallGetRefs(req ClusterFirewallGetRefsRequest) (res []struct {
	Name    string `json:"name"`
	Ref     string `json:"ref"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/refs",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterFirewallGetOptions Get Firewall options.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterFirewallGetOptions() (res struct {
	EBTables      *int    `json:"ebtables"`
	Enable        *int    `json:"enable"`
	LogRateLimit  *string `json:"log_ratelimit"`
	PolicyForward *int    `json:"policy_forward"`
	PolicyIn      *int    `json:"policy_in"`
	PolicyOut     *int    `json:"policy_out"`
	Digest        *string `json:"digest"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/options",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterFirewallPutOptionsRequest struct {
	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Enable ebtables rules cluster wide.
	EBTables *int `in:"query=ebtables;omitempty"`

	// Enable or disable the firewall cluster wide.
	Enable *int `in:"query=enable;omitempty"`

	// Log ratelimiting settings
	//
	// 	[enable=]<1|0> [,burst=<integer>] [,rate=<rate>]
	LogRateLimit string `in:"query=log_ratelimit;omitempty"`

	// Forward policy.
	//
	// 	ACCEPT | DROP
	PolicyForward string `in:"query=policy_forward;omitempty"`

	// Input policy.
	//
	// 	ACCEPT | REJECT | DROP
	PolicyIn string `in:"query=policy_in;omitempty"`

	// Output policy.
	//
	// 	ACCEPT | REJECT | DROP
	PolicyOut string `in:"query=policy_out;omitempty"`
}

// ClusterFirewallPutOptions Set Firewall options.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterFirewallPutOptions(req ClusterFirewallPutOptionsRequest) (err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/firewall/options",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterFirewallGetMacros List available macros
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (s API) ClusterFirewallGetMacros() (res []struct {
	Description string `json:"descr"`
	Macro       string `json:"macro"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/firewall/macros",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
