package core

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type GetVersionResponse struct {
	// The current Proxmox VE point release in `x.y` format.
	Release string `json:"release"`

	// The short git revision from which this version was build.
	RepoID string `json:"repoId"`

	// The full pve-manager package version of this node.
	Version string `json:"version"`
}

// GetVersion returns API version details, including some
// parts of the global datacenter config.
//
// Proxmox permissions required:
//
//	Available for all users.
func (s Service) GetVersion() (res GetVersionResponse, err error) {
	err = s.c.Do(pve.Request{
		Path:   "/api2/json/version",
		Method: http.MethodGet,
		Result: &res,
	})
	if err != nil {
		return res, err
	}

	return GetVersionResponse{
		Release: res.Release,
		RepoID:  res.RepoID,
		Version: res.Version,
	}, nil
}
