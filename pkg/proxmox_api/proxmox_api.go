package proxmoxapi

import "fmt"

type ProxmoxAPI struct{}

func New() *ProxmoxAPI {
	return &ProxmoxAPI{}
}

func (api *ProxmoxAPI) SayHello() {
	fmt.Println("Hello from API")
}
