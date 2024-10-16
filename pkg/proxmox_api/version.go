package proxmoxapi

type GetVersionResponse struct {
	Release string `json:"release"`
	Version string `json:"version"`
	RepoID  string `json:"repoid"`
}

func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error) {
	return sendGetRequest[GetVersionResponse](api, "/version")
}
