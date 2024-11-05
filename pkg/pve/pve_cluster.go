package pve

import (
	"errors"
	"net/http"
	"path"
	"strconv"
)

// GetNextVMID returns the next available VMID.
func (api *PVE) GetNextVMID() (int, error) {
	method := http.MethodGet
	path := path.Join("/cluster", "/nextid")

	res := new(string)
	err := api.httpClient.sendReq(method, path, nil, res)

	if err != nil {
		return 0, err
	}

	if res == nil {
		errors.New("GetNextVMID is trying to access a nil reference")
	}

	return strconv.Atoi(*res)
}
