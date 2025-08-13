package core

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
// Any error returned is of type [errors].*HTTPError.
//
// [errors]: https://pkg.go.dev/github.com/iolave/go-errors
func (s Service) GetVersion() (GetVersionResponse, error) {
	res, err := s.c.CoreGetVersion()
	if err != nil {
		return GetVersionResponse{}, err
	}

	return GetVersionResponse{
		Release: res.Release,
		RepoID:  res.RepoID,
		Version: res.Version,
	}, nil
}
