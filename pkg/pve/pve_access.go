package pve

import (
	"net/http"
)

type PVEAccessService struct {
	api *Client
}

func newPVEAccessService(api *Client) *PVEAccessService {
	service := new(PVEAccessService)
	service.api = api
	return service
}

type GetAccessPermisionsRequest struct {
	UserID string `in:"omitempty;query=userid"` // User ID or full API token ID.
}

type GetAccessPermisionsResponse struct {
	Access       map[string]bool
	AccessGroups map[string]bool
	Nodes        map[string]bool
	Pool         map[string]bool
	SDN          map[string]bool
	Storage      map[string]bool
	VMS          map[string]bool
}

// GetPermissions retrieve effective permissions of given user/token.
//
// GET /access/permissions: each user/token is allowed to dump their own
// permissions (or that of owned tokens). A user can dump the permissions
// of another user or their tokens if they have 'Sys.Audit' permission
// on /access.
func (s *PVEAccessService) GetPermissions(req GetAccessPermisionsRequest) (
	res GetAccessPermisionsResponse,
	err error,
) {
	method := http.MethodGet
	path := "/access/permissions"

	r := struct {
		Access       map[string]int `json:"/access"`
		AccessGroups map[string]int `json:"/access/groups"`
		Nodes        map[string]int `json:"/nodes"`
		Pool         map[string]int `json:"/pool"`
		SDN          map[string]int `json:"/sdn"`
		Storage      map[string]int `json:"/storage"`
		VMS          map[string]int `json:"/vms"`
	}{}

	if err = s.api.client.sendReq2(method, path, &req, &r); err != nil {
		return res, err
	}

	res.Access = convertIntMapToBool(r.Access)
	res.AccessGroups = convertIntMapToBool(r.AccessGroups)
	res.Nodes = convertIntMapToBool(r.Nodes)
	res.Pool = convertIntMapToBool(r.Pool)
	res.SDN = convertIntMapToBool(r.SDN)
	res.Storage = convertIntMapToBool(r.Storage)
	res.VMS = convertIntMapToBool(r.VMS)

	return res, nil
}

func convertIntMapToBool(in map[string]int) (out map[string]bool) {
	out = map[string]bool{}
	for k, v := range in {
		boolv := false
		if v == 1 {
			boolv = true
		}
		out[k] = boolv
	}

	return out
}
