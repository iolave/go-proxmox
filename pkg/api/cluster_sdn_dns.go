package api

import (
	"net/http"
)

// ClusterSDNDNSList SDN DNS index
//
// Parameters:
//
//   - t (optional: powerdns): Only list sdn dns of specific type
//
// Required permissions:
//
//	Only list entries where you have 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/dns/<dns>'
func (c API) ClusterSDNDNSList(t string) (res []struct {
	DNS  string `json:"dns"`
	Type string `json:"type"`
}, err error) {
	req := struct {
		Type string `in:"query=type;omitempty"`
	}{
		Type: t,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/dns",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNDNSNewDNSRequest struct {
	// The SDN dns object identifier.
	DNS string `in:"query=dns;omitempty"`

	Key string `in:"query=key;omitempty"`

	// Plugin type.
	//	powerdns
	Type string `in:"query=type;omitempty"`

	URL string `in:"query=url;omitempty"`

	// Certificate SHA 256 fingerprint.
	Fingerprint string `in:"query=fingerprint;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	ReverseMaskV6 *int `in:"query=reversemaskv6;omitempty"`

	ReverseV6Mask *int `in:"query=reversev6mask;omitempty"`

	TTL *int `in:"query=ttl;omitempty"`
}

// ClusterSDNDNSNewDNS Create a new sdn dns object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/dns",["SDN.Allocate"]]
func (c API) ClusterSDNDNSNewDNS(req ClusterSDNDNSNewDNSRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/dns",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNDNSGetDNS Read sdn dns configuration.
//
// Parameters:
//
//   - dns: The SDN dns object identifier.
//
// Required permissions:
//
//	Check: ["perm","/sdn/dns/{dns}",["SDN.Allocate"]]
func (c API) ClusterSDNDNSGetDNS(dns string) (res struct {
	DNS  string `json:"dns"`
	Type string `json:"type"`
}, err error) {
	req := struct {
		DNS string `in:"path=dns;nonzero"`
	}{
		DNS: dns,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/dns/{dns}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNDNSUpdateDNSRequest struct {
	// The SDN dns object identifier.
	DNS string `in:"path=dns;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Certificate SHA 256 fingerprint.
	Fingerprint string `in:"query=fingerprint;omitempty"`

	Key string `in:"query=key;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	ReverseMaskV6 *int `in:"query=reversemaskv6;omitempty"`

	TTL *int `in:"query=ttl;omitempty"`

	URL string `in:"query=url;omitempty"`
}

// ClusterSDNDNSUpdateDNS Update sdn dns configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/dns/{dns}",["SDN.Allocate"]]
func (c API) ClusterSDNDNSUpdateDNS(req ClusterSDNDNSUpdateDNSRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/dns/{dns}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNDNSDeleteDNS Delete sdn dns object configuration.
//
// Parameters:
//
//   - dns: The SDN dns object identifier.
//   - lockToken (optional): the token for unlocking the global SDN configuration
//
// Required permissions:
//
//	Check: ["perm","/sdn/dns",["SDN.Allocate"]]
func (c API) ClusterSDNDNSDeleteDNS(dns, lockToken string) (err error) {
	req := struct {
		DNS       string `in:"path=dns;nonzero"`
		LockToken string `in:"query=lock-token;omitempty"`
	}{
		DNS:       dns,
		LockToken: lockToken,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/dns/{dns}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
