package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterSDNIPAMS struct {
	IPAM string `json:"ipam"`
	Type string `json:"type"`

	// Certificate SHA 256 fingerprint.
	Fingerprint string `json:"fingerprint"`

	// the token for unlocking the global SDN configuration
	LockToken string `json:"lock-token"`

	Section string `json:"section"`
	Token   string `json:"token"`
	URL     string `json:"url"`
}

// ClusterSDNIPAMSList SDN ipams index.
//
// Parameters:
//
//   - t ("netbox" | "phpipam" | "pve" | ""): Only list sdn ipams of specific type
//
// Required permissions:
//
//	Only list entries where you have 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/ipams/<ipam>'
func ClusterSDNIPAMSList(c *pve.Client, t string) (res []ClusterSDNIPAMS, err error) {
	req := struct {
		Type string `in:"query=type;omitempty"`
	}{
		Type: t,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/ipams",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNIPAMSNewIPAMRequest struct {
	// The SDN ipam object identifier.
	IPAM string `in:"query=ipam;omitempty"`

	// Plugin type.
	//	netbox | phpipam | pve
	Type string `in:"query=type;omitempty"`

	// Certificate SHA 256 fingerprint.
	Fingerprint string `in:"query=fingerprint;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	Section string `in:"query=section;omitempty"`
	Token   string `in:"query=token;omitempty"`
	URL     string `in:"query=url;omitempty"`
}

// ClusterSDNIPAMSNewIPAM Create a new sdn ipam object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/ipams",["SDN.Allocate"]]
func ClusterSDNIPAMSNewIPAM(c *pve.Client, req ClusterSDNIPAMSNewIPAMRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/ipams",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNIPAMSGetIPAM Read sdn ipam configuration.
//
// Parameters:
//
//   - id: The SDN ipam object identifier.
//
// Required permissions:
//
//	Check: ["perm","/sdn/ipams/{ipam}",["SDN.Allocate"]]
func ClusterSDNIPAMSGetIPAM(c *pve.Client, id string) (res ClusterSDNIPAMS, err error) {
	req := struct {
		IPAM string `in:"path=ipam;nonzero"`
	}{
		IPAM: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/ipams/{ipam}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNIPAMSUpdateIPAMRequest struct {
	// The SDN ipam object identifier.
	IPAM string `in:"query=ipam;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Certificate SHA 256 fingerprint.
	Fingerprint string `in:"query=fingerprint;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	Section string `in:"query=section;omitempty"`
	Token   string `in:"query=token;omitempty"`
	URL     string `in:"query=url;omitempty"`
}

// ClusterSDNIPAMSUpdateIPAM Update sdn ipam configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/ipams/{ipam}",["SDN.Allocate"]]
func ClusterSDNIPAMSUpdateIPAM(c *pve.Client, req ClusterSDNIPAMSUpdateIPAMRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/ipams/{ipam}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNIPAMSDeleteIPAM Delete sdn ipam configuration.
//
// Parameters:
//
//   - id: The SDN ipam object identifier.
//
// Required permissions:
//
//	Check: ["perm","/sdn/ipams",["SDN.Allocate"]]
func ClusterSDNIPAMSDeleteIPAM(c *pve.Client, id string) (err error) {
	req := struct {
		IPAM string `in:"path=ipam;nonzero"`
	}{
		IPAM: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/ipams/{ipam}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

// FIXME: /api2/json/cluster/sdn/ipams/{ipam}/status not in docs(?)
