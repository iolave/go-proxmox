package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterSDNFabric struct {
	// Identifier for SDN fabrics
	ID string `json:"id"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `json:"protocol"`

	// OSPF area. Either a IPv4 address or a 32-bit number. Gets validated in rust.
	Area *string `json:"area"`

	// The csnp_interval property for Openfabric
	CSNPInterval *int `json:"csnp_interval"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`

	// The hello_interval property for Openfabric
	HelloInterval *int `json:"hello_interval"`

	// The IP prefix for Node IPs
	IP6Prefix *string `json:"ip6_prefix"`

	// The IP prefix for Node IPs
	IPPrefix *string `json:"ip_prefix"`

	// the token for unlocking the global SDN configuration
	LockToken *string `json:"lock-token"`
}

// ClusterSDNFabricsList SDN Fabrics Index
//
// Parameters:
//
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Only list entries where you have 'SDN.Audit' or 'SDN.Allocate' permissions on '/sdn/fabrics/<fabric>'
func ClusterSDNFabricsList(c *pve.Client, pending, running int) (res []ClusterSDNFabric, err error) {
	req := struct {
		Pending int `in:"query=pending"`
		Running int `in:"query=running"`
	}{
		Pending: pending,
		Running: running,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/fabric",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNFabricsNewFabricRequest struct {
	// Identifier for SDN fabrics
	ID string `in:"query=id;omitempty"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `in:"query=protocol;omitempty"`

	// The csnp_interval property for Openfabric
	CSNPInterval *int `in:"query=csnp_interval;omitempty"`

	// OSPF area. Either a IPv4 address or a 32-bit number. Gets validated in rust.
	Area string `in:"query=area;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// The hello_interval property for Openfabric
	HelloInterval *int `in:"query=hello_interval;omitempty"`

	// The IP prefix for Node IPs
	IP6Prefix string `in:"query=ip6_prefix;omitempty"`

	// The IP prefix for Node IPs
	IPPrefix string `in:"query=ip_prefix;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`
}

// ClusterSDNFabricsNewFabric Add a fabric
//
// Required permissions:
//
//	Check: ["perm","/sdn/fabrics",["SDN.Allocate"]]
func ClusterSDNFabricsNewFabric(c *pve.Client, req ClusterSDNFabricsNewFabricRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/fabric",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNFabricsGetFabric Read sdn fabric configuration.
//
// Parameters:
//
//   - id: Identifier for SDN fabrics
//
// Required permissions:
//
//	Check: ["perm","/sdn/fabrics/{id}",["SDN.Audit","SDN.Allocate"],"any",1]
func ClusterSDNFabricsGetFabric(c *pve.Client, id string) (res ClusterSDNFabric, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/fabric/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNFabricsUpdateFabricRequest struct {
	// Identifier for SDN fabrics
	ID string `in:"query=id;omitempty"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `in:"query=protocol;omitempty"`

	// (OpenFabric only) The csnp_interval property for Openfabric
	CSNPInterval *int `in:"query=csnp_interval;omitempty"`

	// (OpenFabric only) The hello_interval property for Openfabric
	HelloInterval *int `in:"query=hello_interval;omitempty"`

	// (OSPF only) OSPF area. Either a IPv4 address or a 32-bit number. Gets validated in rust.
	Area string `in:"query=area;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// A list of settings you want to delete.
	//	- OpenFabric: [hello_interval | csnp_interval, ...]
	//	- OSPF: [area, ...]
	Delete string `in:"query=delete;omitempty"`

	// The IP prefix for Node IPs
	IP6Prefix string `in:"query=ip6_prefix;omitempty"`

	// The IP prefix for Node IPs
	IPPrefix string `in:"query=ip_prefix;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`
}

// ClusterSDNFabricsUpdateFabric Update sdn fabric configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn/fabrics/{id}",["SDN.Allocate"]]
func ClusterSDNFabricsUpdateFabric(c *pve.Client, req ClusterSDNFabricsUpdateFabricRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/fabric/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNFabricsDeleteFabric Delete sdn fabric configuration.
//
// Parameters:
//
//   - id: Identifier for SDN fabrics
//
// Required permissions:
//
//	Check: ["perm","/sdn/fabrics",["SDN.Allocate"]]
func ClusterSDNFabricsDeleteFabric(c *pve.Client, id string) (err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/fabric/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

type ClusterSDNFabricNode struct {
	// Identifier for SDN fabrics
	FabricID string `json:"fabric_id"`

	Interfaces []any

	// Identifier for nodes in an SDN fabric
	NodeID string `json:"node_id"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `json:"protocol"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`

	// IPv4 address for this node
	IP *string `json:"ip"`

	// IPv6 address for this node
	IP6 *string `json:"ip6"`

	// The token for unlocking the global SDN configuration
	LockToken *string `json:"lock-token"`
}

// ClusterSDNFabricsListNodes SDN Fabrics Index.
//
// Parameters:
//
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Only list nodes where you have 'SDN.Audit' or 'SDN.Allocate' permissions on
//
// '/sdn/fabrics/<fabric>' and 'Sys.Audit' or 'Sys.Modify' on /nodes/<node_id>
func ClusterSDNFabricsListNodes(c *pve.Client, pending, running int) (res []ClusterSDNFabricNode, err error) {
	req := struct {
		Pending int `in:"query=pending"`
		Running int `in:"query=running"`
	}{
		Pending: pending,
		Running: running,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterSDNFabricsGetNodeFabrics Read sdn fabric configuration.
//
// Parameters:
//
//   - fabricId: Identifier for SDN fabrics
//   - pending (0|1): Display pending config.
//   - running (0|1): Display running config.
//
// Required permissions:
//
//	Only returns nodes where you have 'Sys.Audit' or 'Sys.Modify' permissions.
//	Check: ["perm","/sdn/fabrics/{fabric_id}",["SDN.Audit"]]
func ClusterSDNFabricsGetFabricNodes(c *pve.Client,
	fabricId string,
	pending, running int,
) (res []ClusterSDNFabricNode, err error) {
	req := struct {
		FabricID string `in:"path=fabric_id;nonzero"`
		Pending  int    `in:"query=pending"`
		Running  int    `in:"query=running"`
	}{
		FabricID: fabricId,
		Pending:  pending,
		Running:  running,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node/{fabric_id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNFabricsAddNodeRequest struct {
	// Identifier for SDN fabrics
	FabricID string `in:"path=fabric_id;nonzero"`

	// Identifier for nodes in an SDN fabric
	NodeID string `in:"query=node_id;omitempty"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `in:"query=protocol;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// network interface
	//	openfabric: [name=<string> [,hello_multiplier=<integer>] [,ip=<string>] [,ip6=<string>], ...]
	//	ospf: [name=<string> [,ip=<string>], ...]
	Interfaces string `in:"query=interfaces;omitempty"`

	// IPv4 address for this node
	IP string `in:"query=ip;omitempty"`

	// IPv6 address for this node
	IP6 string `in:"query=ip6;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`
}

// ClusterSDNFabricsAddNode Add a node to sdn fabric configuration.
//
// Required permissions:
//
//	Check: ["and",["perm","/sdn/fabrics/{fabric_id}",["SDN.Allocate"]],["perm","/nodes/{node_id}",["Sys.Modify"]]]
func ClusterSDNFabricsAddNode(c *pve.Client, req ClusterSDNFabricsAddNodeRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node/{fabric_id}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterSDNFabricsGetFabricNode Get a node
//
// Parameters:
//
//   - fabricId: Identifier for SDN fabrics
//   - nodeId: Identifier for nodes in an SDN fabric

// Required permissions:
//
//	Check: ["perm","/sdn/fabrics/{fabric_id}",["SDN.Audit"]]
func ClusterSDNFabricsGetFabricNode(c *pve.Client, fabricId, nodeId string) (res ClusterSDNFabricNode, err error) {
	req := struct {
		FabricID string `in:"path=fabric_id;nonzero"`
		NodeID   string `in:"path=node_id;nonzero"`
	}{
		FabricID: fabricId,
		NodeID:   nodeId,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node/{fabric_id}/{node_id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterSDNFabricNodeUpdateRequest struct {
	// Identifier for SDN fabrics
	FabricID string `in:"path=fabric_id;nonzero"`

	// Identifier for nodes in an SDN fabric
	NodeID string `in:"path=node_id;nonzero"`

	// Type of configuration entry in an SDN Fabric section config
	//	openfabric | ospf
	Protocol string `in:"query=protocol;omitempty"`

	// A list of settings you want to delete.
	//	[interfaces | ip | ip6, ...]
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// network interface
	//	openfabric: [name=<string> [,hello_multiplier=<integer>] [,ip=<string>] [,ip6=<string>], ...]
	//	ospf: [name=<string> [,ip=<string>], ...]
	Interfaces string `in:"query=interfaces;omitempty"`

	// IPv4 address for this node
	IP string `in:"query=ip;omitempty"`

	// IPv6 address for this node
	IP6 string `in:"query=ip6;omitempty"`

	// the token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`
}

// ClusterSDNFabricsUpdateNode Update sdn fabric configuration.
//
// Required permissions:
//
//	Check: ["and",["perm","/sdn/fabrics/{fabric_id}",["SDN.Allocate"]],["perm","/nodes/{node_id}",["Sys.Modify"]]]
func ClusterSDNFabricsUpdateNode(c *pve.Client, req ClusterSDNFabricNodeUpdateRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node/{fabric_id}/{node_id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNFabricsDeleteNode Delete node from sdn fabric configuration.
//
// Parameters:
//
//   - fabricId: Identifier for SDN fabrics
//   - nodeId: Identifier for nodes in an SDN fabric
//
// Required permissions:
//
//	Check: ["and",["perm","/sdn/fabrics/{fabric_id}",["SDN.Allocate"]],["perm","/nodes/{node_id}",["Sys.Modify"]]]
func ClusterSDNFabricsDeleteNode(c *pve.Client, fabricId, nodeId string) (err error) {
	req := struct {
		FabricID string `in:"path=fabric_id;nonzero"`
		NodeID   string `in:"path=node_id;nonzero"`
	}{
		FabricID: fabricId,
		NodeID:   nodeId,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/fabrics/node/{fabric_id}/{node_id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
