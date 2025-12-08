package api

import (
	"net/http"
)

type ClusterSDNVNet struct {
	// The SDN vnet object identifier.
	VNet string `json:"vnet"`

	// zone id
	Zone string `json:"zone"`

	// alias name of the vnet
	Alias string `json:"alias"`

	// If true, sets the isolated property for all members of this VNet
	IsolatePorts *int `json:"isolate-ports"`

	// the token for unlocking the global SDN configuration
	LockToken *string `json:"lock-token"`

	// vlan or vxlan id
	Tag *int `json:"tag"`

	// Type
	Type *string `json:"type"`

	// Allow vm VLANs to pass through this vnet.
	VLanAware *int `json:"vlan-aware"`
}

// ClusterSDNVNetsList SDN vnets index.
//
// Parameters:
//
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Only list entries where you have 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/zones/<zone>/<vnet>'
func (c API) ClusterSDNVNetsList(pending, running int) (res []ClusterSDNVNet, err error) {
	req := struct {
		Pending int `in:"query=pending"`
		Running int `in:"query=running"`
	}{
		Pending: pending,
		Running: running,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNVNetsNewVNetRequest struct {
	// The SDN vnet object identifier.
	VNet string `in:"query=vnet;omitempty"`

	// zone id
	Zone string `in:"query=zone;omitempty"`

	// alias name of the vnet
	Alias string `in:"query=alias;omitempty"`

	// If true, sets the isolated property for all members of this VNet
	IsolatePorts *int `in:"query=isolate-ports;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// vlan or vxlan id
	Tag *int `json:"tag"`

	// Type
	Type string `in:"query=type;omitempty"`

	// Allow vm VLANs to pass through this vnet.
	VLanAware *int `in:"query=vlanaware;omitempty"`
}

// ClusterSDNVNetsNewVNet Create a new sdn vnet object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}",["SDN.Allocate"]]
func (c API) ClusterSDNVNetsNewVNet(req ClusterSDNVNetsNewVNetRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNVNetsGetVNet Read sdn vnet configuration.
//
// Parameters:
//
//   - id: The SDN vnet object identifier.
//
// Required permissions:
//
//	Require 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/zones/<zone>/<vnet>'
func (c API) ClusterSDNVNetsGetVNet(id string) (res ClusterSDNVNet, err error) {
	req := struct {
		VNet string `in:"path=vnet;nonzero"`
	}{
		VNet: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNVNetsUpdateVNetRequest struct {
	// The SDN vnet object identifier.
	VNet string `in:"path=vnet;nonzero"`

	// alias name of the vnet
	Alias string `in:"query=alias;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// If true, sets the isolated property for all members of this VNet
	IsolatePorts *int `in:"query=isolate-ports;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// vlan or vxlan id
	Tag *int `json:"tag"`

	// Allow vm VLANs to pass through this vnet.
	VLanAware *int `in:"query=vlanaware;omitempty"`

	// zone id
	Zone string `in:"query=zone;omitempty"`
}

// ClusterSDNVNetsUpdateVNet Update sdn vnet configuration.
//
// Required permissions:
//
//	Require 'SDN.Allocate' permission on '/sdn/zones/<zone>/<vnet>'
func (c API) ClusterSDNVNetsUpdateVNet(req ClusterSDNVNetsUpdateVNetRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNVNetsDeleteVNet Delete sdn vnet configuration.
//
// Parameters:
//
//   - id: The SDN vnet object identifier.
//
// Required permissions:
//
//	Require 'SDN.Allocate' permission on '/sdn/zones/<zone>/<vnet>'
func (c API) ClusterSDNVNetsDeleteVNet(id string) (err error) {
	req := struct {
		VNet string `in:"path=vnet;nonzero"`
	}{
		VNet: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

type ClusterSDNVNetsNewIPMappingRequest struct {
	// The IP address to associate with the given MAC address
	IP string `in:"query=ip;omitempty"`

	// The SDN vnet object identifier.
	VNet string `in:"path=vnet;omitempty"`

	// The SDN zone object identifier.
	Zone string `in:"query=zone;omitempty"`

	// A common MAC address with the I/G (Individual/Group) bit not set.
	MAC string `in:"query=mac;omitempty"`
}

// ClusterSDNVNetsNewIPMapping Create a new sdn vnet ip mapping object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}/{vnet}",["SDN.Allocate"]]
func (c API) ClusterSDNVNetsNewIPMapping(req ClusterSDNVNetsNewIPMappingRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}/ips",
		Payload: &req,
	})

	return err
}

type ClusterSDNVNetsUpdateIPMappingRequest struct {
	// The IP address to associate with the given MAC address
	IP string `in:"query=ip;omitempty"`

	// The SDN vnet object identifier.
	VNet string `in:"path=vnet;omitempty"`

	// The SDN zone object identifier.
	Zone string `in:"query=zone;omitempty"`

	// A common MAC address with the I/G (Individual/Group) bit not set.
	MAC string `in:"query=mac;omitempty"`

	// The (unique) ID of the VM.
	VMID *int `in:"query=vmid;omitempty"`
}

type ClusterSDNVNetsDeleteIPMappingRequest struct {
	// The IP address to associate with the given MAC address
	IP string `in:"query=ip;omitempty"`

	// The SDN vnet object identifier.
	VNet string `in:"path=vnet;omitempty"`

	// The SDN zone object identifier.
	Zone string `in:"query=zone;omitempty"`

	// A common MAC address with the I/G (Individual/Group) bit not set.
	MAC string `in:"query=mac;omitempty"`
}

// ClusterSDNVNetsUpdateIPMapping Update sdn vnet ip mapping configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}/{vnet}",["SDN.Allocate"]]
func (c API) ClusterSDNVNetsUpdateIPMapping(req ClusterSDNVNetsUpdateIPMappingRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}/ips",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNVNetsDeleteIPMapping Delete sdn vnet ip mapping configuration.
//
// Parameters:
//
//   - id: The SDN vnet object identifier.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}/{vnet}",["SDN.Allocate"]]
func (c API) ClusterSDNVNetsDeleteIPMapping(id string) (err error) {
	req := struct {
		VNet string `in:"path=vnet;nonzero"`
	}{
		VNet: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/vnets/{vnet}/ips",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
