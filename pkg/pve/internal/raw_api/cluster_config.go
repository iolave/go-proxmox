package rawapi

import (
	"net/http"
	"strconv"

	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterConfigPostConfigRequest struct {
	// The name of the cluster.
	ClusterName string `in:"query=clustername;nonzero"`

	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	//
	// 	[address=]<IP> [,priority=<integer>]
	Links []string `in:"query_n=link;omitempty"`

	// Node id for this node.
	NodeID *int `in:"query=nodeid;omitempty"`

	// Number of votes for this node.
	Votes *int `in:"query=votes;omitempty"`
}

// ClusterConfigPostConfig Generate new cluster configuration. If no links given, default to local IP address as link0.
//
// Required permissions:
//
//	Root only.
func ClusterConfigPostConfig(c *pve.Client, req ClusterConfigPostConfigRequest) (res string, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/config",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterConfigGetAPIVersion Return the version of the cluster join API available on this node.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterConfigGetAPIVersion(c *pve.Client) (int, error) {
	res := ""

	err := c.Do(pve.Request{
		Path:   "/api2/json/cluster/config/apiversion",
		Method: http.MethodGet,
		Result: &res,
	})

	if err != nil {
		return 0, err
	}

	version, err := strconv.Atoi(res)
	if err != nil {
		return 0, errors.NewInternalServerError(
			"failed to parse response",
			err,
		)
	}

	return version, err
}

type ClusterConfigGetJoinInfoRequest struct {
	// The node for which the joinee gets the nodeinfo.
	Node string `in:"query=node;omitempty"`
}

// ClusterConfigGetJoinInfo Get information needed to join this cluster over the connected node.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterConfigGetJoinInfo(c *pve.Client, req ClusterConfigGetJoinInfoRequest) (res struct {
	ConfigDigest string `json:"config_digest"`
	NodeList     []struct {
		Name        string `json:"name"`
		NodeID      int    `json:"nodeid"`
		PVEAddr     string `json:"pve_addr"`
		PVEFP       string `json:"pve_fp"`
		QuorumVotes int    `json:"quorum_votes"`
		// TODO: add proper type
		Ring0Addr string `json:"ring0_addr"`
	} `json:"nodelist"`
	PreferredNode string `json:"preferred_node"`
	// TODO: add proper type
	Totem any `json:"totem"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/config/joininfo",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterConfigJoinClusterRequest struct {
	// Certificate SHA 256 fingerprint.
	//
	// 	([A-Fa-f0-9]{2}:){31}[A-Fa-f0-9]{2}
	Fingerprint string `in:"query=fingerprint;nonzero"`

	// Hostname (or IP) of an existing cluster member.
	Hostname string `in:"query=hostname;nonzero"`

	// Superuser (root) password of peer node.
	Password string `in:"query=password;nonzero"`

	// Do not throw error if node already exists.
	Force *int `in:"query=force;omitempty"`

	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	//
	// 	[address=]<IP> [,priority=<integer>]
	Links []string `in:"query_n=link;omitempty"`

	// Node id for this node.
	Nodeid *int `in:"query=nodeid;omitempty"`

	// Number of votes for this node
	Votes *int `in:"query=votes;omitempty"`
}

// ClusterConfigJoinCluster Joins this node into an existing cluster. If no links are given, default to IP resolved by node's hostname on single link (fallback fails for clusters with multiple links).
//
// Required permissions:
//
//	Root only.
func ClusterConfigJoinCluster(c *pve.Client, req ClusterConfigJoinClusterRequest) (res string, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/config/join",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterConfigGetQDevice Get QDevice status
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// TODO: add proper response type
func ClusterConfigGetQDevice(c *pve.Client) (res any, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/config/qdevice",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterConfigGetTotem Get corosync totem protocol settings.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// TODO: add proper response type
func ClusterConfigGetTotem(c *pve.Client) (res any, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/config/totem",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterConfigGetNodes Corosync node list.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterConfigGetNodes(c *pve.Client) (res []struct {
	Node string `json:"node"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/config/nodes",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterConfigAddNodeRequest struct {
	// The cluster node name.
	Node string `in:"path=node;nonzero"`

	// The JOIN_API_VERSION of the new node.
	APIVersion *int `in:"query=apiversion;omitempty"`

	// Do not throw error if node already exists.
	Force *int `in:"query=force;omitempty"`

	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	//
	// 	[address=]<IP> [,priority=<integer>]
	Links []string `in:"query_n=link;omitempty"`

	// IP Address of node to add. Used as fallback if no links are given.
	NewNodeIP string `in:"query=new_node_ip;omitempty"`

	// Node id for this node.
	NodeID *int `in:"query=nodeid;omitempty"`

	// Number of votes for this node
	Votes *int `in:"query=votes;omitempty"`
}

// ClusterConfigAddNode Adds a node to the cluster configuration. This call is for internal use.
//
// Required permissions:
//
//	Root only.
func ClusterConfigAddNode(c *pve.Client, req ClusterConfigAddNodeRequest) (res struct {
	CoroSyncAuthKey string   `json:"corosync_authkey"`
	CoroSyncConf    string   `json:"corosync_conf"`
	Warnings        []string `json:"warnings"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/config/nodes/{node}",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterConfigDeleteNode Removes a node from the cluster configuration.
//
// Parameters:
//
//   - node is the cluster node name.
//
// Required permissions:
//
//	Root only.
func ClusterConfigDeleteNode(c *pve.Client, node string) (err error) {
	req := struct {
		Node string `in:"path=node;nonzero"`
	}{
		Node: node,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/config/nodes/{node}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
