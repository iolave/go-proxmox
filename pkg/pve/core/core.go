package core

import apiclient "github.com/iolave/go-proxmox/internal/api_client"

type Service struct {
	c *apiclient.HTTPClient
}

func New(httpclient *apiclient.HTTPClient) Service {
	return Service{
		c: httpclient,
	}
}
