package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/iolave/go-proxmox/pkg/helpers"
	"github.com/iolave/go-proxmox/pkg/pve"
)

const authRegex = `^PVEAPIToken=(?P<user>[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+)!(?P<tokenName>[a-zA-Z0-9_-]+)=(?P<token>[a-zA-Z0-9_-]+)$`

func (s *server) IsUserAuthorized(r *http.Request, path, scope string) (bool, error) {
	auth := r.Header.Get("authorization")
	params := helpers.GetRegexpParams(authRegex, auth)

	creds := pve.NewTokenCreds(
		params["user"],
		params["tokenName"],
		params["token"],
	)

	client, err := pve.NewWithCredentials(pve.Config{
		Host:               s.cfg.PVEHost,
		Port:               s.cfg.PVEPort,
		InsecureSkipVerify: true,
	}, creds)
	if err != nil {
		if strings.Contains(err.Error(), "auth") {
			return false, nil
		}
		return false, err
	}

	perms, err := client.Access.GetPermissions(pve.GetAccessPermisionsRequest{})
	if err != nil {
		return false, err
	}
	b, _ := json.Marshal(perms)
	permsMap := map[string]map[string]bool{}
	json.Unmarshal(b, &permsMap)

	permsScope := permsMap[path][scope]
	return permsScope, nil
}
