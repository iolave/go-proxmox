package cluster

import (
	"math/rand/v2"

	"github.com/iolave/go-proxmox/pkg/pve/internal/raw_api"
)

// GetNextVMID returns the next available VMID.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (s Service) GetNextVMID() (vmid int, err error) {
	return rawapi.ClusterGetNextID(
		s.c,
		rawapi.ClusterGetNextIDRequest{},
	)
}

// GetRandomVMID returns a random VMID.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (s Service) GetRandomVMID() (vmid int, err error) {
	// choosing a number between 100 and vmid max
	for vmid <= 100 {
		vmid = rand.IntN(999999999)
	}

	available, err := s.IsVMIDAvailable(vmid)
	if err != nil {
		return 0, err
	}

	if !available {
		return s.GetRandomVMID()
	}

	return vmid, nil
}

// IsVMIDAvailable checks if a vmid is available or not.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func (s Service) IsVMIDAvailable(id int) (bool, error) {
	resources, err := rawapi.ClusterGetResources(
		s.c,
		rawapi.ClusterGetResourcesRequest{
			Type: "vm",
		},
	)
	if err != nil {
		return false, err
	}

	for _, resource := range resources {
		if *resource.VMID == id {
			return false, nil
		}
	}

	return true, nil
}
