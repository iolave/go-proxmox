package proxmoxapi

import (
	"fmt"
)

type ProxmoxAPI struct {
	creds *credentials
}

// TODO: To test credentials, do a proxmox version
// query to ensure credentials are valid
func New() (*ProxmoxAPI, error) {
	creds, err := newCredentialsFromEnv()

	if err != nil {
		return nil, err
	}

	return &ProxmoxAPI{creds}, nil
}

// TODO: To test credentials, do a proxmox version
// query to ensure credentials are valid
func NewWithCredentials(creds *credentials) *ProxmoxAPI {
	return &ProxmoxAPI{creds}
}

func (api *ProxmoxAPI) SayHello() {
	fmt.Println("Hello from API")
}
