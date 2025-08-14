package pve

import (
	"fmt"

	apiclient "github.com/iolave/go-proxmox/internal/api_client"
	"github.com/iolave/go-proxmox/pkg/cloudflare"
	"github.com/iolave/go-proxmox/pkg/pve/core"
)

type Config struct {
	Host               string
	Port               int
	InsecureSkipVerify bool
	CfServiceToken     *cloudflare.ServiceToken
	APIWrapper         bool
}

type Client struct {
	// httpc is the underlying http client used
	// to send requests to the proxmox api.
	APIClient *apiclient.HTTPClient
	httpc     *apiclient.HTTPClient

	config Config
	creds  *Credentials
	client *httpClient

	// PVE API implementations
	Access  *PVEAccessService
	Node    *PVENodeService
	Cluster *PVEClusterService
	LXC     *PVELxcService

	// v1.0.0 API implementations
	Core core.Service
}

func New(config Config) (*Client, error) {
	creds, err := NewEnvCreds()
	if err != nil {
		return nil, err
	}
	return NewWithCredentials(config, creds)

}

func NewWithCredentials(config Config, creds *Credentials) (*Client, error) {
	httpc, err := apiclient.NewHTTPClient(
		"https",
		config.Host,
		config.Port,
		config.InsecureSkipVerify,
	)
	if err != nil {
		return nil, err
	}

	auth, err := creds.getAuthorization()
	if err != nil {
		return nil, err
	}
	httpc.CustomHeaders.Set("Authorization", auth)

	if config.CfServiceToken != nil {
		httpc.CustomHeaders.Set("CF-Access-Client-Id", config.CfServiceToken.ClientId)
		httpc.CustomHeaders.Set("CF-Access-Client-Secret", config.CfServiceToken.ClientSecret)
	}

	api := &Client{
		httpc:     httpc,
		APIClient: httpc,
		config:    config,
		creds:     creds,
		client: newHttpClient(
			creds,
			config.CfServiceToken,
			config.Host,
			config.Port,
			config.InsecureSkipVerify,
			config.APIWrapper,
		),
	}

	initializeServices(api)

	api.Core = core.New(api.httpc)

	_, err = api.Core.GetVersion()
	if err != nil {
		return nil, fmt.Errorf("Unable to comunicate with proxmox api, %v\n", err)
	}

	return api, nil
}

func initializeServices(api *Client) {
	api.Access = newPVEAccessService(api)
	api.Node = newPVENodeService(api)
	api.Cluster = newPVEClusterService(api)
	api.LXC = newPVELxcService(api)
}
