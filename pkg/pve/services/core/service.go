package core

import "github.com/iolave/go-proxmox/pkg/pve"

type Service struct {
	c *pve.Client
}

func New(c *pve.Client) Service {
	return Service{
		c: c,
	}
}
