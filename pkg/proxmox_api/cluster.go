package proxmoxapi

import (
	"net/http"
	"path"
	"strconv"
)

// GetNextVMID returns the next available VMID.
func (api *ProxmoxAPI) GetNextVMID() (int, error) {
	path := path.Join("/cluster", "/nextid")
	vmid, err := sendRequest[string](http.MethodGet, api, path, nil)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(vmid)
}
