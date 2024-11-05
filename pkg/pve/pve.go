package pve

import (
	"fmt"

	"github.com/iolave/go-proxmox/pkg/cloudflare"
)

type Config struct {
	Host               string
	Port               int
	InsecureSkipVerify bool
	CfServiceToken     *cloudflare.ServiceToken
}

type PVE struct {
	config     Config
	creds      *Credentials
	httpClient *httpClient
}

func New(config Config) (*PVE, error) {
	creds, err := NewEnvCreds()

	if err != nil {
		return nil, err
	}

	api := &PVE{
		config: config,
		creds:  creds,
		httpClient: newHttpClient(
			creds,
			config.CfServiceToken,
			config.Host,
			config.Port,
			config.InsecureSkipVerify,
		),
	}

	_, err = api.GetVersion()

	if err != nil {
		return nil, fmt.Errorf("Unable to comunicate with proxmox api, %v\n", err)
	}

	return api, nil
}

func NewWithCredentials(config Config, creds *Credentials) (*PVE, error) {
	api := &PVE{
		config: config,
		creds:  creds,
		httpClient: newHttpClient(
			creds,
			config.CfServiceToken,
			config.Host,
			config.Port,
			config.InsecureSkipVerify,
		),
	}

	_, err := api.GetVersion()

	if err != nil {
		return nil, fmt.Errorf("Unable to comunicate with proxmox api, %v\n", err)
	}

	return api, nil
}
