package main

import (
	"fmt"

	proxmoxapi "github.com/iolave/go-proxmox/pkg/proxmox_api"
)

func main() {
	proxmoxApi := proxmoxapi.New()
	fmt.Println("proxmox-cli")
	proxmoxApi.SayHello()

}
