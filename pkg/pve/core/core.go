package core

import "github.com/iolave/go-proxmox/pkg/api"

type Service struct {
	api *api.API
}

func New(api *api.API) Service {
	return Service{
		api: api,
	}
}
