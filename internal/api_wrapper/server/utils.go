package server

import (
	"net/http"
	"strings"

	"github.com/iolave/go-proxmox/pkg/helpers"
	"github.com/iolave/go-proxmox/pkg/pve"
	"github.com/iolave/go-proxmox/pkg/pve/services/access"
)

const authRegex = `^PVEAPIToken=(?P<user>[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+)!(?P<tokenName>[a-zA-Z0-9_-]+)=(?P<token>[a-zA-Z0-9_-]+)$`

func (s *server) IsUserAuthorized(
	r *http.Request,
	privelege access.Privilege,
	path string,
) (bool, error) {
	auth := r.Header.Get("authorization")
	params := helpers.GetRegexpParams(authRegex, auth)

	creds := pve.NewTokenCreds(
		params["user"],
		params["tokenName"],
		params["token"],
	)

	client, err := pve.New(pve.Config{
		Proto:              "https",
		Host:               s.cfg.PVEHost,
		Port:               s.cfg.PVEPort,
		InsecureSkipVerify: true,
		Credentials:        creds,
	})
	if err != nil {
		if strings.Contains(err.Error(), "auth") {
			return false, nil
		}
		return false, err
	}

	accessService := access.New(client)

	perms, err := accessService.GetPermissions()
	if err != nil {
		return false, err
	}
	if !perms[privelege][path] {
		return false, nil
	}

	return true, nil
}
