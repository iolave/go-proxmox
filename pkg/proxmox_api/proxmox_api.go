package proxmoxapi

import (
	"fmt"
	"net/http"
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
}

// TODO: To test credentials, do a proxmox version
// query to ensure credentials are valid
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

	return api, nil
}

// TODO: To test credentials, do a proxmox version
// query to ensure credentials are valid
func NewWithCredentials(config ProxmoxAPIConfig, creds *credentials) *ProxmoxAPI {
	api := &ProxmoxAPI{
		creds:      creds,
		httpClient: newHttpClient(config.InsecureSkipVerify),
	}

	return api
}

type apiResponse[T any] struct {
	Data T `json:"data"`
}

func (api *ProxmoxAPI) buildHttpRequestUrl(path string) string {
	if path[0] == '/' {
		path = path[1:]
	}

	return fmt.Sprintf("https://%s:%d/api2/json/%s", api.config.Host, api.config.Port, path)

}
