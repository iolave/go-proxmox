package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterHAResourcesListRequest struct {
	// Only list resources of specific type
	//
	// 	ct | vm
	Type string `in:"query=type;omitempty"`
}

// ClusterHAREsourcesList List HA resources.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHAResourcesList(c *pve.Client, req ClusterHAResourcesListRequest) (res []struct {
	SID string `json:"sid"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

type ClusterHAResourcesPostRequest struct {
	// HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	//
	// 	<type>:<name>
	SID string `in:"query=sid;omitempty"`

	// Description.
	Comment string `in:"query=comment;omitempty"`

	// Automatically migrate HA resource to the node with the highest priority according to their node affinity  rules, if a node with a higher priority than the current node comes online.
	Failback *int `in:"query=failback;omitempty"`

	// The HA group identifier.
	Group string `in:"query=group;omitempty"`

	// Maximal number of service relocate tries when a service failes to start.
	MaxRelocate *int `in:"query=max_relocate;omitempty"`

	// Maximal number of tries to restart the service on a node after its start failed.
	MaxRestart *int `in:"query=max_restart;omitempty"`

	// Requested resource state. The CRM reads this state and acts accordingly.
	// Please note that `enabled` is just an alias for `started`.
	//
	// `started`;;
	//
	// The CRM tries to start the resource. Service state is
	// set to `started` after successful start. On node failures, or when start
	// fails, it tries to recover the resource.  If everything fails, service
	// state it set to `error`.
	//
	// `stopped`;;
	//
	// The CRM tries to keep the resource in `stopped` state, but it
	// still tries to relocate the resources on node failures.
	//
	// `disabled`;;
	//
	// The CRM tries to put the resource in `stopped` state, but does not try
	// to relocate the resources on node failures. The main purpose of this
	// state is error recovery, because it is the only way to move a resource out
	// of the `error` state.
	//
	// `ignored`;;
	//
	// The resource gets removed from the manager status and so the CRM and the LRM do
	// not touch the resource anymore. All {pve} API calls affecting this resource
	// will be executed, directly bypassing the HA stack. CRM commands will be thrown
	// away while there source is in this state. The resource will not get relocated
	// on node failures.
	//
	//   	started | stopped | enabled | disabled | ignored
	State string `in:"query=state;omitempty"`

	// Resource type.
	//
	// 	vm | ct
	Type string `in:"query=type;omitempty"`
}

// ClusterHAResourcesPost Create new HA resource.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHAResourcesPost(c *pve.Client, req ClusterHAResourcesPostRequest) (err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/ha/resources",
		Method: http.MethodPost,
	})

	return err
}

// ClusterHAResourcesGet Read resource configuration.
//
// Parameters:
//
//   - sid is the HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterHAResourcesGet(c *pve.Client, sid string) (res struct {
	Digest      string  `json:"digest"`
	SID         string  `json:"sid"`
	Type        string  `json:"type"`
	Comment     string  `json:"comment"`
	Failback    *int    `json:"failback"`
	Group       *string `json:"group"`
	MaxRelocate *int    `json:"max_relocate"`
	MaxRestart  *int    `json:"max_restart"`
	State       *string `json:"state"`
}, err error) {
	req := struct {
		SID string `in:"path=sid"`
	}{
		SID: sid,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources/{sid}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterHAResourcesUpdateRequest struct {
	// HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	SID string `in:"path=sid;nonzero"`

	// Description.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Automatically migrate HA resource to the node with the highest priority according to their node affinity  rules, if a node with a higher priority than the current node comes online.
	Failback *int `in:"query=failback;omitempty"`

	// The HA group identifier.
	Group string `in:"query=group;omitempty"`

	// Maximal number of service relocate tries when a service failes to start.
	MaxRelocate *int `in:"query=max_relocate;omitempty"`

	// Maximal number of tries to restart the service on a node after its start failed.
	MaxRestart *int `in:"query=max_restart;omitempty"`

	// Requested resource state. The CRM reads this state and acts accordingly.
	// Please note that `enabled` is just an alias for `started`.
	//
	// `started`;;
	//
	// The CRM tries to start the resource. Service state is
	// set to `started` after successful start. On node failures, or when start
	// fails, it tries to recover the resource.  If everything fails, service
	// state it set to `error`.
	//
	// `stopped`;;
	//
	// The CRM tries to keep the resource in `stopped` state, but it
	// still tries to relocate the resources on node failures.
	//
	// `disabled`;;
	//
	// The CRM tries to put the resource in `stopped` state, but does not try
	// to relocate the resources on node failures. The main purpose of this
	// state is error recovery, because it is the only way to move a resource out
	// of the `error` state.
	//
	// `ignored`;;
	//
	// The resource gets removed from the manager status and so the CRM and the LRM do
	// not touch the resource anymore. All {pve} API calls affecting this resource
	// will be executed, directly bypassing the HA stack. CRM commands will be thrown
	// away while there source is in this state. The resource will not get relocated
	// on node failures.
	//
	//   	started | stopped | enabled | disabled | ignored
	State string `in:"query=state;omitempty"`
}

// ClusterHAResourcesUpdate Update HA resource.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHAResourcesUpdate(c *pve.Client, req ClusterHAResourcesUpdateRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources/{sid}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterHAResourcesDelete Delete HA resource.
//
// Parameters:
//
//   - sid is the HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHAResourcesDelete(c *pve.Client, sid string) (err error) {
	req := struct {
		SID string `in:"path=sid"`
	}{
		SID: sid,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources/{sid}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

// ClusterHAResourcesMigrate Request resource migration (online) to another node.
//
// Parameters:
//
//   - sid is the HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
//   - node is the target node name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHAResourcesMigrate(c *pve.Client, sid, node string) (res struct {
	RequestedNode     string `json:"requested-node"`
	SID               string `json:"sid"`
	BlockingResources []struct {
		Cause string `json:"cause"`
		SID   string `json:"sid"`
	} `json:"blocking-resources"`
	CoMigratedResources []string `json:"comigrated-resources"`
}, err error) {
	req := struct {
		SID  string `in:"path=sid"`
		Node string `in:"query=node,omitempty"`
	}{
		SID:  sid,
		Node: node,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources/{sid}/migrate",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterHAResourcesRelocate Request resource relocatzion to another node. This stops the service on the old node, and restarts it on the target node.

// Parameters:
//
//   - sid is the HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
//   - node is the target node name.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func ClusterHAResourcesRelocate(c *pve.Client, sid, node string) (res struct {
	RequestedNode     string `json:"requested-node"`
	SID               string `json:"sid"`
	BlockingResources []struct {
		Cause string `json:"cause"`
		SID   string `json:"sid"`
	} `json:"blocking-resources"`
	CoMigratedResources []string `json:"comigrated-resources"`
}, err error) {
	req := struct {
		SID  string `in:"path=sid"`
		Node string `in:"query=node,omitempty"`
	}{
		SID:  sid,
		Node: node,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/ha/resources/{sid}/relocate",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}
