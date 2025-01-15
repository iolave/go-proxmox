package pve

import (
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

	result := ""
	err := s.api.client.sendReq(method, path, nil, &result)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}
