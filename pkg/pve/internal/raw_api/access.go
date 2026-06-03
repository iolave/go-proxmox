package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type GETAccessPermisionsRequest struct {
	UserID string `in:"omitempty;query=userid"` // User ID or full API token ID.
}

type GETAccessPermisionsResponse map[string]map[string]int

// GetPermissions retrieve effective permissions of given user/token.
//
// GET /access/permissions: each user/token is allowed to dump their own
// permissions (or that of owned tokens). A user can dump the permissions
// of another user or their tokens if they have 'Sys.Audit' permission
// on /access.
func GETAccessPermissions(c *pve.Client, req GETAccessPermisionsRequest) (
	res GETAccessPermisionsResponse,
	err error,
) {
	err = c.Do(pve.Request{
		Method:  http.MethodGet,
		Path:    "/api2/json/access/permissions",
		Payload: &req,
		Result:  &res,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
