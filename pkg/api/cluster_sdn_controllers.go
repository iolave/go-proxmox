package api

import (
	"net/http"
)

// ClusterSDNControllersList List SDN controllers
//
// Parameters:
//
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//   - t: Only list sdn controllers of specific type (bgp | evpn | faucet | isis)
//
// Required permissions:
//
//	Check: ["perm","/sdn",["SDN.Allocate"]]
func (c API) ClusterSDNControllersList(pending, running int, t string) (res []struct {
	Controller string  `json:"controller"`
	Type       string  `json:"type"`
	Pending    *int    `json:"pending"`
	State      *string `json:"state"`
}, err error) {
	req := struct {
		Pending int    `in:"query=pending"`
		Running int    `in:"query=running"`
		T       string `in:"query=type"`
	}{
		Pending: pending,
		Running: running,
		T:       t,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/controllers",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNControllersNewControllerRequest struct {
	// The SDN controller object identifier.
	Controller string `in:"query=controller;omitempty"`

	// Plugin type.
	//	bgp | evpn | faucet | isis
	Type string `in:"query=type;omitempty"`

	// autonomous system number
	ASN *int `in:"query=asn;omitempty"`

	BGPMultipathAsPathRelax *int `in:"query=bgp-multipath-as-path-relax;omitempty"`

	// Enable ebgp. (remote-as external)
	EBGP *int `in:"query=ebgp;omitempty"`

	EBGPMultihop *int `in:"query=ebgp-multihop;omitempty"`

	// SDN fabric to use as underlay for this EVPN controller.
	Fabric string `in:"query=fabric;omitempty"`

	// ISIS domain.
	ISISDomain string `in:"query=isis-domain;omitempty"`

	// ISIS interface.
	ISISIfaces string `in:"query=isis-ifaces;omitempty"`

	// ISIS network entity title.
	ISISNet string `in:"query=isis-net;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// source loopback interface.
	Loopback string `in:"query=loopback;omitempty"`

	// The cluster node name.
	Node string `in:"query=node;omitempty"`

	// peers address list.
	Peers string `in:"query=peers;omitempty"`
}

// ClusterSDNControllersNewController Create a new sdn controller object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/controllers",["SDN.Allocate"]]
func (c API) ClusterSDNControllersNewController(req ClusterSDNControllersNewControllerRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/controllers",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNControllersGetController Read sdn controller configuration.
//
// Parameters:
//
//   - controller: The SDN controller object identifier.
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Check: ["perm","/sdn/controllers/{controller}",["SDN.Allocate"]]
func (c API) ClusterSDNControllersGetController(controller string, pending, running int) (res struct {
	Controller string  `json:"controller"`
	Type       string  `json:"type"`
	Pending    *int    `json:"pending"`
	State      *string `json:"state"`
}, err error) {
	req := struct {
		Controller string `in:"path=controller;nonzero"`
		Pending    int    `in:"query=pending"`
		Running    int    `in:"query=running"`
	}{
		Controller: controller,
		Pending:    pending,
		Running:    running,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/controllers/{controller}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNControllersUpdateControllerRequest struct {
	// The SDN controller object identifier.
	Controller string `in:"path=controller;omitempty"`

	// autonomous system number
	ASN *int `in:"query=asn;omitempty"`

	BGPMultipathAsPathRelax *int `in:"query=bgp-multipath-as-path-relax;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Enable ebgp. (remote-as external)
	EBGP *int `in:"query=ebgp;omitempty"`

	EBGPMultihop *int `in:"query=ebgp-multihop;omitempty"`

	// SDN fabric to use as underlay for this EVPN controller.
	Fabric string `in:"query=fabric;omitempty"`

	// ISIS domain.
	ISISDomain string `in:"query=isis-domain;omitempty"`

	// ISIS interface.
	ISISIfaces string `in:"query=isis-ifaces;omitempty"`

	// ISIS network entity title.
	ISISNet string `in:"query=isis-net;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// source loopback interface.
	Loopback string `in:"query=loopback;omitempty"`

	// The cluster node name.
	Node string `in:"query=node;omitempty"`

	// peers address list.
	Peers string `in:"query=peers;omitempty"`
}

// ClusterSDNControllersUpdateController Update sdn controller configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/controllers/{controller}",["SDN.Allocate"]]
func (c API) ClusterSDNControllersUpdateController(req ClusterSDNControllersUpdateControllerRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/controllers/{controller}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNControllersDeleteController Delete sdn controller object configuration.
// Parameters:
//
//   - controller: The SDN controller object identifier.
//   - lockToken (optional): the token for unlocking the global SDN configuration
//
// Required permissions:
//
//	Check: ["perm","/sdn/controllers",["SDN.Allocate"]]
func (c API) ClusterSDNControllersDeleteController(controller, lockToken string) (err error) {
	req := struct {
		Controller string `in:"path=controller;nonzero"`
		LockToken  string `in:"query=lock-token;omitempty"`
	}{
		Controller: controller,
		LockToken:  lockToken,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/sdn/controllers/{controller}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
