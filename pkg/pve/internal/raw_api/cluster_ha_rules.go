package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterHARulesListRequest struct {
	// Limit the returned list to rules affecting the specified resource.
	Resource string `in:"query=resource;omitempty"`

	// Limit the returned list to the specified rule type.
	//
	//	node-affinity | resource-affinity
	Type string `in:"query=type;omitempty"`
}

// ClusterHARulesList Get HA rules.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHARulesList(c *pve.Client, req ClusterHARulesListRequest) (res []struct {
	Rule string `json:"rule"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/rules",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

type ClusterHARulesPostRequest struct {
	// List of HA resource IDs. This consists of a list of resource types followed by a resource specific name separated with a colon (example: vm:100,ct:101).
	//
	//	<type>:<name>{,<type>:<name>}*
	Resources string `in:"query=resources;nonzero"`

	//  HA rule identifier.
	Rule string `in:"query=rule;nonzero"`

	//  HA rule type.
	//
	//	node-affinity | resource-affinity
	Type string `in:"query=type;nonzero"`

	// Describes whether the node affinity rule is strict or non-strict.
	//
	// A non-strict node affinity rule makes resources prefer to be on the defined nodes.
	// If none of the defined nodes are available, the resource may run on any other node.
	//
	// A strict node affinity rule makes resources be restricted to the defined nodes. If
	// none of the defined nodes are available, the resource will be stopped.
	Strict *int `in:"query=strict;omitempty"`

	// Describes whether the HA resources are supposed to be kept on the same node ('positive'), or are supposed to be kept on separate nodes ('negative').
	//
	//	positive | negative
	Affinity string `in:"query=affinity;omitempty"`

	// HA rule comment.
	Comment string `in:"query=comment;omitempty"`

	// Whether the HA rule is disabled.
	Disable *int `in:"query=disable;omitempty"`

	// List of cluster node members, where a priority can be given to each node. A resource bound to a group will run on the available nodes with the highest priority. If there are more nodes in the highest priority class, the services will get distributed to those nodes. The priorities have a relative meaning only. The higher the number, the higher the priority.
	//
	// 	<node>[:<pri>]{,<node>[:<pri>]}*
	Nodes string `in:"query=nodes;omitempty"`
}

// ClusterHARulesPost Create HA rule.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHARulesPost(c *pve.Client, req ClusterHARulesPostRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/rules",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterHARulesGet Get HA rule.
//
// Parameters:
//
//   - rule HA rule identifier.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHARulesGet(c *pve.Client, rule string) (res struct {
	Rule string `json:"rule"`
	Type string `json:"type"`
}, err error) {
	req := struct {
		Rule string `in:"path=rule;nonzero"`
	}{
		Rule: rule,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/rules/{rule}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterHARulesPutRequest struct {
	//  HA rule identifier.
	Rule string `in:"query=path;nonzero"`

	//  HA rule type.
	//
	//	node-affinity | resource-affinity
	Type string `in:"query=type;nonzero"`

	// Describes whether the HA resources are supposed to be kept on the same node ('positive'), or are supposed to be kept on separate nodes ('negative').
	//
	//	positive | negative
	Affinity string `in:"query=affinity;omitempty"`

	// HA rule comment.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Whether the HA rule is disabled.
	Disable *int `in:"query=disable;omitempty"`

	// List of cluster node members, where a priority can be given to each node. A resource bound to a group will run on the available nodes with the highest priority. If there are more nodes in the highest priority class, the services will get distributed to those nodes. The priorities have a relative meaning only. The higher the number, the higher the priority.
	//
	// 	<node>[:<pri>]{,<node>[:<pri>]}*
	Nodes string `in:"query=nodes;omitempty"`

	// List of HA resource IDs. This consists of a list of resource types followed by a resource specific name separated with a colon (example: vm:100,ct:101).
	//
	//	<type>:<name>{,<type>:<name>}*
	Resources string `in:"query=resources;omitempty"`

	// Describes whether the node affinity rule is strict or non-strict.
	//
	// A non-strict node affinity rule makes resources prefer to be on the defined nodes.
	// If none of the defined nodes are available, the resource may run on any other node.
	//
	// A strict node affinity rule makes resources be restricted to the defined nodes. If
	// none of the defined nodes are available, the resource will be stopped.
	Strict *int `in:"query=strict;omitempty"`
}

// ClusterHARulesPut Update HA rule.
//
// Parameters:
//
//   - rule HA rule identifier.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHARulesPut(c *pve.Client, req ClusterHARulesPutRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/rules/{rule}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterHARulesDelete Delete HA rule.
//
// Parameters:
//
//   - rule HA rule identifier.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHARulesDelete(c *pve.Client, rule string) (err error) {
	req := struct {
		Rule string `in:"path=rule"`
	}{
		Rule: rule,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/rules/{rule}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
