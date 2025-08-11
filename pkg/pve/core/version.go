package core

import (
	"net/http"

	apiclient "github.com/iolave/go-proxmox/internal/api_client"
)

type GetVersionResponse struct {
	// The current Proxmox VE point release in `x.y` format.
	Release string `json:"release"`

	// The short git revision from which this version was build.
	RepoID string `json:"repoid"`

	// The full pve-manager package version of this node.
	Version string `json:"version"`
}

// GetVersion returns API version details, including some
// parts of the global datacenter config.
func (s Service) GetVersion() (GetVersionResponse, error) {
	res := GetVersionResponse{}
	err := s.httpc.SendPVERequest(apiclient.PVERequest{
		Path:   "/api2/json/version",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
