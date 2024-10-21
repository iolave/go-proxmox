package proxmoxapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iolave/go-proxmox/pkg/cloudflare"
)

type ProxmoxAPI struct {
	config     ProxmoxAPIConfig
	creds      *credentials
	httpClient *http.Client
}

type ProxmoxAPIConfig struct {
	Host               string
	Port               int
	InsecureSkipVerify bool
	CfServiceToken     *cloudflare.CloudflareServiceToken
}

func New(config ProxmoxAPIConfig) (*ProxmoxAPI, error) {
	creds, err := newCredentialsFromEnv()

	if err != nil {
		return nil, err
	}

	api := &ProxmoxAPI{
		config:     config,
		creds:      creds,
		httpClient: newHttpClient(config.InsecureSkipVerify),
	}

	_, err = api.GetVersion()

	if err != nil {
		return nil, fmt.Errorf("Unable to comunicate with proxmox api, %v\n", err)
	}

	return api, nil
}

// TODO: To test credentials, do a proxmox version
// query to ensure credentials are valid
func NewWithCredentials(config ProxmoxAPIConfig, creds *credentials) (*ProxmoxAPI, error) {
	api := &ProxmoxAPI{
		config:     config,
		creds:      creds,
		httpClient: newHttpClient(config.InsecureSkipVerify),
	}

	_, err := api.GetVersion()

	if err != nil {
		return nil, fmt.Errorf("Unable to comunicate with proxmox api, %v\n", err)
	}

	return api, nil
}

func (api *ProxmoxAPI) buildHttpRequestUrl(path string) string {
	checkForwardSlashRune := func(r rune) bool {
		if r == '/' {
			return true
		}

		return false
	}

	path = strings.TrimFunc(path, checkForwardSlashRune)

	return fmt.Sprintf("https://%s:%d/api2/json/%s", api.config.Host, api.config.Port, path)

}
