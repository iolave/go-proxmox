package pve

import "net/http"

type GetVersionResponse struct {
	Release string `json:"release"`
	Version string `json:"version"`
	RepoID  string `json:"repoid"`
}

// GetVersion retrieves proxmox version.
func (api *PVE) GetVersion() (GetVersionResponse, error) {
	path := "/version"
	method := http.MethodPost

	res := &GetVersionResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}
