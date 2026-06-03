package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterSDNZone struct {
	// Type of the zone.
	//	evpn | faucet | qinq | simple | vlan | vxlan
	Type string `json:"type"`

	// Name of the zone.
	Zone string `json:"zone"`

	// Advertise IP prefixes (Type-5 routes) instead of MAC/IP pairs (Type-2 routes). EVPN zone only.
	AdvertiseSubnets *int `json:"advertise-subnets"`

	// the bridge for which VLANs should be managed. VLAN & QinQ zone only.
	Bridge *string `json:"bridge"`

	// Disable auto mac learning. VLAN zone only.
	BridgeDisableMACLearning *int `json:"bridge-disable-mac-learning"`

	// ID of the controller for this zone. EVPN zone only.
	Controller *string `json:"controller"`

	// Name of DHCP server backend for this zone.
	DHCP *string `json:"dhcp"`

	// Digest of the controller section.
	Digest *string `json:"digest"`

	// Suppress IPv4 ARP && IPv6 Neighbour Discovery messages. EVPN zone only.
	DisableARPNDSuppression *int `json:"disable-arp-nd-suppression"`

	// ID of the DNS server for this zone.
	DNS *string `json:"dns"`

	// Domain name for this zone.
	DNSZone *string `json:"dnszone"`

	// List of PVE Nodes that should act as exit node for this zone. EVPN zone only.
	ExitNodes *string `json:"exitnodes"`

	// Create routes on the exit nodes, so they can connect to EVPN guests. EVPN zone only.
	ExitNodesLocalRouting *int `json:"exitnodes-local-routing"`

	// Force traffic through this exitnode first. EVPN zone only.
	ExitNodesPrimary *string `json:"exitnodes-primary"`

	// ID of the IPAM for this zone.
	IPAM *string `json:"ipam"`

	// MAC address of the anycast router for this zone.
	MAC *string `json:"mac"`

	// MTU of the zone, will be used for the created VNet bridges.
	MTU *int `json:"mtu"`

	// Nodes where this zone should be created.
	Nodes *string `json:"nodes"`

	// Comma-separated list of peers, that are part of the VXLAN zone. Usually the IPs of the nodes. VXLAN zone only.
	Peers *string `json:"peers"`

	// Changes that have not yet been applied to the running configuration.
	Pending struct {
		AdvertiseSubnets         *int    `json:"advertise-subnets"`
		Bridge                   *string `json:"bridge"`
		BridgeDisableMACLearning *int    `json:"bridge-disable-mac-learning"`
		Controller               *string `json:"controller"`
		DHCP                     *string `json:"dhcp"`
		DisableARPNDSuppression  *int    `json:"disable-arp-nd-suppression"`
		DNS                      *string `json:"dns"`
		DNSZone                  *string `json:"dnszone"`
		ExitNodes                *string `json:"exitnodes"`
		ExitNodesLocalRouting    *int    `json:"exitnodes-local-routing"`
		ExitNodesPrimary         *string `json:"exitnodes-primary"`
		IPAM                     *string `json:"ipam"`
		MAC                      *string `json:"mac"`
		MTU                      *int    `json:"mtu"`
		Nodes                    *string `json:"nodes"`
		Peers                    *string `json:"peers"`
		ReverseDNS               *string `json:"reversedns"`
		RTImport                 *string `json:"rt-import"`
	} `json:"pending"`

	// ID of the reverse DNS server for this zone.
	ReverseDNS *string `json:"reversedns"`

	// Route-Targets that should be imported into the VRF of this zone via BGP. EVPN zone only.
	RTImport *string `json:"rt-import"`

	// State of the SDN configuration object.
	//	new | changed | deleted
	State *string `json:"state"`

	// Service-VLAN Tag (outer VLAN). QinQ zone only
	Tag *int `json:"tag"`

	// VLAN protocol for the creation of the QinQ zone. QinQ zone only.
	//	802.1q | 802.1ad
	VlanProtocol *string `json:"vlan-protocol"`

	// VNI for the zone VRF. EVPN zone only.
	//	1 - 16777215
	VRFVxLan *int `json:"vrf-vxlan"`

	// VXLAN port for the zone VRF. VXLAN zone only.
	//	1 - 65536
	VxLanPort *int `json:"vxlan-port"`
}

type ClusterSDNZoneListRequest struct {
	// Display pending config.
	Pending int `in:"query=pending"`

	// Display running config.
	Running int `in:"query=running"`

	// Only list SDN zones of specific type
	//	evpn | faucet | qinq | simple | vlan | vxlan
	Type string `in:"query=type"`
}

// ClusterSDNZonesList SDN zones index.
//
// Required permissions:
//
//	Only list entries where you have 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/zones/<zone>'
func ClusterSDNZonesList(c *pve.Client, req ClusterSDNZoneListRequest) (res []ClusterSDNZone, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/zones",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNZonesNewZoneRequest struct {
	// Plugin type.
	//	evpn | faucet | qinq | simple | vlan | vxlan
	Type string `in:"query=type;omitempty"`
	// The SDN zone object identifier.
	Zone string `in:"query=zone;omitempty"`

	// Advertise IP prefixes (Type-5 routes) instead of MAC/IP pairs (Type-2 routes).
	AdvertiseSubnets *int `in:"query=advertise-subnets;omitempty"`

	// The bridge for which VLANs should be managed.
	Bridge string `in:"query=bridge;omitempty"`

	// Disable auto mac learning.
	BridgeDisableMACLearning *int `in:"query=bridge-disable-mac-learning;omitempty"`

	// Controller for this zone.
	Controller string `in:"query=controller;omitempty"`

	// Type of the DHCP backend for this zone
	DHCP string `in:"query=dhcp;omitempty"`

	// Suppress IPv4 ARP && IPv6 Neighbour Discovery messages.
	DisableARPNDSuppression *int `in:"query=disable-arp-nd-suppression;omitempty"`

	// dns api server
	DNS string `in:"query=dns;omitempty"`

	// dns domain zone
	//	ex: mydomain.com
	DNSZone string `in:"query=dnszone;omitempty"`

	// Faucet dataplane id
	DPID *int `in:"query=dp-id;omitempty"`

	// List of cluster node names.
	ExitNodes string `in:"query=exitnodes;omitempty"`

	// Allow exitnodes to connect to EVPN guests.
	ExitNodesLocalRouting *int `in:"query=exitnodes-local-routing;omitempty"`

	// Force traffic through this exitnode first.
	ExitNodesPrimary string `in:"query=exitnodes-primary;omitempty"`

	// SDN fabric to use as underlay for this VXLAN zone.
	Fabric string `in:"query=fabric;omitempty"`

	// use a specific ipam
	IPAM string `in:"query=ipam;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// Anycast logical router mac address.
	MAC string `in:"query=mac;omitempty"`

	// MTU of the zone, will be used for the created VNet bridges.
	MTU *int `in:"query=mtu;omitempty"`

	// List of cluster node names.
	Nodes string `in:"query=nodes;omitempty"`

	// Comma-separated list of peers, that are part of the VXLAN zone. Usually the IPs of the nodes.
	Peers string `in:"query=peers;omitempty"`

	// reverse dns api server
	ReverseDNS string `in:"query=reversedns;omitempty"`

	// List of Route Targets that should be imported into the VRF of the zone.
	RTImport string `in:"query=rt-import;omitempty"`

	// service-VLAN Tag (outer VLAN)
	//	0 - N
	Tag *int `in:"query=tag;omitempty"`

	// Which VLAN protocol should be used for the creation of the QinQ zone.
	//	802.1q | 802.1ad
	VlanProtocol string `in:"query=vlan-protocol;omitempty"`

	// VNI for the zone VRF.
	//	1 - 16777215
	VRFVxLan *int `in:"query=vrf-vxlan;omitempty"`

	// UDP port that should be used for the VXLAN tunnel (default 4789).
	//	1 - 65536
	VxLanPort *int `in:"query=vxlan-port;omitempty"`
}

// ClusterSDNZonesNewZone Create a new sdn zone object.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones",["SDN.Allocate"]]
func ClusterSDNZonesNewZone(c *pve.Client, req ClusterSDNZonesNewZoneRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/zones",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNZonesGetZone Read sdn zone configuration.
//
// Parameters:
//
//   - zone: The SDN zone object identifier.
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}",["SDN.Allocate"]]
func ClusterSDNZonesGetZone(c *pve.Client, zone string, pending, running int) (res ClusterSDNZone, err error) {
	req := struct {
		Zone    string `in:"path=zone;nonzero"`
		Pending int    `in:"query=pending"`
		Running int    `in:"query=running"`
	}{
		Zone:    zone,
		Pending: pending,
		Running: running,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/zones/{zone}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNZonesUpdateZoneRequest struct {
	// The SDN zone object identifier.
	Zone string `in:"path=zone;nonzero"`

	// Advertise IP prefixes (Type-5 routes) instead of MAC/IP pairs (Type-2 routes).
	AdvertiseSubnets *int `in:"query=advertise-subnets;omitempty"`

	// The bridge for which VLANs should be managed.
	Bridge string `in:"query=bridge;omitempty"`

	// Disable auto mac learning.
	BridgeDisableMACLearning *int `in:"query=bridge-disable-mac-learning;omitempty"`

	// Controller for this zone.
	Controller string `in:"query=controller;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Type of the DHCP backend for this zone
	DHCP string `in:"query=dhcp;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Suppress IPv4 ARP && IPv6 Neighbour Discovery messages.
	DisableARPNDSuppression *int `in:"query=disable-arp-nd-suppression;omitempty"`

	// dns api server
	DNS string `in:"query=dns;omitempty"`

	// dns domain zone
	//	ex: mydomain.com
	DNSZone string `in:"query=dnszone;omitempty"`

	// Faucet dataplane id
	DPID *int `in:"query=dp-id;omitempty"`

	// List of cluster node names.
	ExitNodes string `in:"query=exitnodes;omitempty"`

	// Allow exitnodes to connect to EVPN guests.
	ExitNodesLocalRouting *int `in:"query=exitnodes-local-routing;omitempty"`

	// Force traffic through this exitnode first.
	ExitNodesPrimary string `in:"query=exitnodes-primary;omitempty"`

	// SDN fabric to use as underlay for this VXLAN zone.
	Fabric string `in:"query=fabric;omitempty"`

	// use a specific ipam
	IPAM string `in:"query=ipam;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// Anycast logical router mac address.
	MAC string `in:"query=mac;omitempty"`

	// MTU of the zone, will be used for the created VNet bridges.
	MTU *int `in:"query=mtu;omitempty"`

	// List of cluster node names.
	Nodes string `in:"query=nodes;omitempty"`

	// Comma-separated list of peers, that are part of the VXLAN zone. Usually the IPs of the nodes.
	Peers string `in:"query=peers;omitempty"`

	// reverse dns api server
	ReverseDNS string `in:"query=reversedns;omitempty"`

	// List of Route Targets that should be imported into the VRF of the zone.
	RTImport string `in:"query=rt-import;omitempty"`

	// service-VLAN Tag (outer VLAN)
	//	0 - N
	Tag *int `in:"query=tag;omitempty"`

	// Which VLAN protocol should be used for the creation of the QinQ zone.
	//	802.1q | 802.1ad
	VlanProtocol string `in:"query=vlan-protocol;omitempty"`

	// VNI for the zone VRF.
	//	1 - 16777215
	VRFVxLan *int `in:"query=vrf-vxlan;omitempty"`

	// UDP port that should be used for the VXLAN tunnel (default 4789).
	//	1 - 65536
	VxLanPort *int `in:"query=vxlan-port;omitempty"`
}

// ClusterSDNZonesUpdateZone Update sdn zone configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}",["SDN.Allocate"]]
func ClusterSDNZonesUpdateZone(c *pve.Client, req ClusterSDNZonesUpdateZoneRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/zones/{zone}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNZonesDeleteZone Delete sdn zone configuration.
//
// Parameters:
//
//   - zone: The SDN zone object identifier.
//   - lockToken (optional): the token for unlocking the global SDN configuration
//
// Required permissions:
//
//	Check: ["perm","/sdn/zones/{zone}",["SDN.Allocate"]]
func ClusterSDNZonesDeleteZone(c *pve.Client, zone string) (err error) {
	req := struct {
		Zone string `in:"path=zone;nonzero"`
	}{
		Zone: zone,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/zones/{zone}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
