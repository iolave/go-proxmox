package pve

import (
	"errors"
	"net/http"
	"path"
	"strconv"
)

type PVEClusterService struct {
	api      *PVE
	Firewall *PVEClusterFirewallService
}

func newPVEClusterService(api *PVE) *PVEClusterService {
	service := new(PVEClusterService)
	service.api = api
	service.Firewall = newPVEClusterFirewallService(api)
	return service
}

// GetNextVMID returns the next available VMID.
func (s *PVEClusterService) GetNextVMID() (int, error) {
	method := http.MethodGet
	path := path.Join("/cluster", "/nextid")

	res := new(string)
	err := s.api.client.sendReq(method, path, nil, res)

	if err != nil {
		return 0, err
	}

	if res == nil {
		errors.New("GetNextVMID is trying to access a nil reference")
	}

	return strconv.Atoi(*res)
}
