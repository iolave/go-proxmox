package main

import (
	"encoding/json"
	"fmt"

	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/pkg/api"
	"github.com/iolave/go-proxmox/pkg/pve/core"
)

func main() {
	// fmt.Println("proxmox-cli")

	config := api.Config{
		Proto:              "https",
		Host:               "pve.internal.pingolabs.cl",
		Port:               8006,
		InsecureSkipVerify: true,
		Credentials:        api.NewTokenCreds("root@pam", "test-go-proxmox-mm", "4607e5bd-84c2-420a-84ee-4b7418618bdd"),
	}

	api, _ := api.New(config)
	cs := core.New(api)
	p1(cs.GetVersion())
	// c.CustomHeaders.Set("Authorization", "PVEAPIToken=root@pam!test-go-proxmox-mm=4607e5bd-84c2-420a-84ee-4b7418618bdd")

	// p1(c.ClusterMappingListResourceTypes())
	// p1(c.ClusterMappingListDir(api.ClusterMappingListDirRequest{}))
	// p1(c.ClusterMappingListPCI(api.ClusterMappingListPCIRequest{}))

	// p1(c.Core.GetVersion())
	//p1(c.APIClient.ClusterGetStatus())
	//p1(c.APIClient.ClusterGetResources(api.ClusterGetResourcesRequest{
	//	Type: "sdn",
	//}))
	// p1(c.APIClient.ClusterGetNextID(api.ClusterGetNextIDRequest{
	// 	// VMID: 100,
	// }))

	//_ = "pve-prd-1"
	//p1(api.Node.Firewall.NewRule(pve.CreateNodeFirewallRuleRequest{
	//	Node:        node,
	//	Type:        "in",
	//	Action:      "ACCEPT",
	//	Macro:       "SSH",
	//	Destination: "dc/pve-prd-1-local",
	//	Comment:     "API_CREATED",
	//}))
	//p1(api.Node.Firewall.GetRule(node, "59a45ebd-b66c-4dd6-9b1f-f8b2c63873b5"))
	//
	// p1(api.LXC.GetAll(node))
	// p1(api.LXC.Delete(node, 110, nil))
	// time.Sleep(time.Microsecond * 5000)
	// p1(api.LXC.GetAll(node))
	// p1(api.Cluster.GetNextVMID())
	//
	// p1(api.LXC.Create(pve.CreateLxcRequest{
	// 	Node:       node,
	// 	OSTemplate: "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst",
	// 	//RootFS:     "local-lvm:8",
	// 	Hostname: "client-2",
	// 	VMID:     110,
	// 	Features: pve.LXCFeatures{},
	// 	Net: []pve.LxcNet{
	// 		{
	// 			Name:     "eth0",
	// 			Bridge:   "vmbr7",
	// 			Firewall: true,
	// 			IP:       "10.10.0.24/24",
	// 			GW:       "10.10.0.1",
	// 		},
	// 	},
	// 	Nameserver: "10.10.0.1",
	// }))
	//p2(api.LXC.CreateTemplate(node, 105))
	//p2(api.LXC.Update(pve.UpdateLxcRequest{
	//	Node:     node,
	//	VMID:     109,
	//	Template: true,
	//}))
	// p1(api.LXC.Clone(pve.CloneLxcRequest{
	// 	Node:     node,
	// 	VMID:     105,
	// 	Hostname: "pve-api-clone",
	// }))
	//p1(api.LXC.Start(pve.LXCStartRequest{
	//	ID:   100,
	//	Node: node,
	//}))
	//p1(api.LXC.GetStatus(node, 100))
	//p1(api.Cluster.GetVMIDs())
	//p1(api.Access.GetPermissions(pve.GetAccessPermisionsRequest{}))
	//p1(api.LXC.GetIP(857388269))
	//r, _ := http.NewRequest("GET", "/", nil)
	//r.Header.Set("authorization", "PVEAPIToken=root@pam!test-go-proxmox-mm=4607e5bd-84c2-420a-84ee-4b7418618bdd")
	//fmt.Println(api.LXC.Exec(100, "bash", "ls -l /"))
	//p1(api.LXC.GetInterfaceByName(node, 1022, "eth0"))
}

func p1(res any, err error) {
	if err != nil {
		err := err.(*errors.HTTPError)
		fmt.Println(string(err.JSON()))
		return
	}

	b, err := json.Marshal(res)

	if err != nil {
		fmt.Println("json marshal error:", err.Error())
		return
	}

	fmt.Println(string(b))
}

func p2(err error) {
	if err != nil {
		p1(nil, err)
		return
	}
}
