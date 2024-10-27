package proxmoxapi

import "net/http"

type GetVersionResponse struct {
	Release string `json:"release"`
	Version string `json:"version"`
	RepoID  string `json:"repoid"`
}

// GetVersion retrieves proxmox version.
func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error) {
	return sendRequest[GetVersionResponse](http.MethodGet, api, "/version", nil)
}
