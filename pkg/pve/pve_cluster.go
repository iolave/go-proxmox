package pve

import (
	"math/rand/v2"
	"net/http"
	"path"
	"strconv"
)

type PVEClusterService struct {
	api      *Client
	Firewall *PVEClusterFirewallService
}

func newPVEClusterService(api *Client) *PVEClusterService {
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

// GetRandomVMID chooses a random vmid and checks wether it's
// available or not.
//
// GET /cluster/resources is accessible by all authenticated users.
func (s *PVEClusterService) GetRandomVMID() (id int, err error) {
	// getting all vmids
	ids, err := s.api.Cluster.GetVMIDs()
	if err != nil {
		return 0, err
	}

	// choosing a number between 100 and vmid max
	for id <= 100 {
		id = rand.IntN(999999999)
	}

	// checking if vmid is avaialble
	for _, remoteId := range ids {
		if remoteId == id {
			return s.api.Cluster.GetRandomVMID()
		}
	}

	return id, err
}

// GetVMIDs retrieves all vmids in the cluster.
//
// GET /cluster/resources is accessible by all authenticated users.
func (s *PVEClusterService) GetVMIDs() (ids []int, err error) {
	method := http.MethodGet
	path := "/cluster/resources"

	type Request struct {
		Type string `in:"query=type"`
	}
	type Response struct {
		ID int `json:"vmid"`
	}

	req := Request{Type: "vm"}
	res := []Response{}
	err = s.api.client.sendReq2(method, path, &req, &res)
	if err != nil {
		return ids, err
	}
	for _, singleRes := range res {
		ids = append(ids, singleRes.ID)
	}

	return ids, nil
}

// IsVMIDAvailable checks if a vmid is available or not.
//
// GET /cluster/resources is accessible by all authenticated users.
func (s *PVEClusterService) IsVMIDAvailable(id int) (avaialble bool, err error) {
	ids, err := s.api.Cluster.GetVMIDs()
	if err != nil {
		return false, err
	}

	for _, remoteId := range ids {
		if remoteId == id {
			return false, nil
		}
	}

	return true, nil
}
