package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iolave/go-proxmox/pkg/cloudflare"
	"github.com/iolave/go-proxmox/pkg/helpers"
	"github.com/iolave/go-proxmox/pkg/pve"
)

func main() {
	fmt.Println("proxmox-cli")

	creds := pve.NewTokenCreds("root@pam", "test", "41e70ba5-3edb-4d9a-8ca6-84f297d9c2bc")

	config := pve.Config{
		Host:               "pve-prd.pingolabs.cl",
		Port:               443,
		InsecureSkipVerify: true,
		CfServiceToken:     cloudflare.NewServiceToken("ec0c353083ecb09a4682de4638464996.access", "55fb2b3eb6f8294a8e56df1f8fcb60165847040844a162d22c016f9a0a4fdb64"),
	}

	api, err := pve.NewWithCredentials(config, creds)

	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	node := "pve-prd-1"
	p1(api.GetLxcs(node))
	p1(api.GetNextVMID())

	p1(api.CreateLxc(pve.CreateLxcRequest{
		Node:       node,
		OSTemplate: "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst",
		//RootFS:     helpers.NewStr("local-lvm:8"),
		//Start:      helpers.NewBool(true),
		VMID: helpers.NewInt(100),
	}))
}

func p1(res any, err error) {
	if err != nil {
		fmt.Println("error:", err.Error(), err)
		return
	}

	b, err := json.Marshal(res)

	if err != nil {
		fmt.Println("error:", err.Error(), err)
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
