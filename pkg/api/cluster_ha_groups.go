package api

import "net/http"

// ClusterHAGroupsList Get HA groups. (deprecated in favor of HA rules)
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (c API) ClusterHAGroupsList() (res []struct {
	Group string `json:"group"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/ha/groups",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterHAGroupsPostRequest struct {
	// The HA group identifier.
	Group string `in:"query=group;omitempty"`

	// List of cluster node members, where a priority can be given to each node. A resource bound to a group will run on the available nodes with the highest priority. If there are more nodes in the highest priority class, the services will get distributed to those nodes. The priorities have a relative meaning only. The higher the number, the higher the priority.
	//
	// 	<node>[:<pri>]{,<node>[:<pri>]}*
	Nodes string `in:"query=nodes;omitempty"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`

	// The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	NoFailback *int `in:"query=nofailback;omitempty"`

	// Resources bound to restricted groups may only run on nodes defined by the group. The resource will be placed in the stopped state if no group node member is online. Resources on unrestricted groups may run on any cluster node if all group members are offline, but they will migrate back as soon as a group member comes online. One can implement a 'preferred node' behavior using an unrestricted group with only one member.
	Restricted *int `in:"query=restricted;omitempty"`

	// Group type.
	Type string `in:"query=type;omitempty"`
}

// ClusterHAGroupsPost Create a new HA group. (deprecated in favor of HA rules)
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func (c API) ClusterHAGroupsPost(req ClusterHAGroupsPostRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ha/groups",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterHAGroupsGet Read ha group configuration. (deprecated in favor of HA rules)
//
// Parameters:
//
//   - group is the HA group identifier.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// NOTE:Proxmox api docs does not specify response type.
func (c API) ClusterHAGroupsGet(group string) (res any, err error) {
	req := struct {
		Group string `in:"path=group"`
	}{
		Group: group,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ha/groups/{group}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterHAGroupsUpdateRequest struct {
	// The HA group identifier.
	Group string `in:"path=group;omitempty"`

	// Descriptive comment.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// List of cluster node members, where a priority can be given to each node. A resource bound to a group will run on the available nodes with the highest priority. If there are more nodes in the highest priority class, the services will get distributed to those nodes. The priorities have a relative meaning only. The higher the number, the higher the priority.
	//
	// 	<node>[:<pri>]{,<node>[:<pri>]}*
	Nodes string `in:"query=nodes;omitempty"`

	// The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	NoFailback *int `in:"query=nofailback;omitempty"`

	// Resources bound to restricted groups may only run on nodes defined by the group. The resource will be placed in the stopped state if no group node member is online. Resources on unrestricted groups may run on any cluster node if all group members are offline, but they will migrate back as soon as a group member comes online. One can implement a 'preferred node' behavior using an unrestricted group with only one member.
	Restricted *int `in:"query=restricted;omitempty"`
}

// ClusterHAGroupsUpdate Update ha group configuration. (deprecated in favor of HA rules)
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func (c API) ClusterHAGroupsUpdate(req ClusterHAGroupsUpdateRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ha/groups/{group}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterHAGroupsDelete Delete ha group configuration. (deprecated in favor of HA rules)
//
// Parameters:
//
//   - group is the HA group identifier.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Console"]]
func (c API) ClusterHAGroupsDelete(group string) (err error) {
	req := struct {
		Group string `in:"path=group"`
	}{
		Group: group,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ha/groups/{group}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
