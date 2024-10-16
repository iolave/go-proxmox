package proxmoxapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	myerrors "github.com/iolave/go-proxmox/errors"
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

type apiResponse[T any] struct {
	Data T `json:"data"`
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

func sendGetRequest[T any](api *ProxmoxAPI, urlPath string) (T, error) {
	url := api.buildHttpRequestUrl(urlPath)
	result := &apiResponse[T]{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return result.Data, err
	}

	if api.creds.credType == CREDENTIALS_TOKEN {
		auth := fmt.Sprintf("PVEAPIToken=%s!%s=%s", api.creds.username, api.creds.tokenName, api.creds.token)
		req.Header.Add("Authorization", auth)
	} else {
		return result.Data, errors.New("only token credentials are supported at the moment")
	}

	res, err := api.httpClient.Do(req)

	if err != nil {
		return result.Data, err
	}

	if res.StatusCode != http.StatusOK {
		return result.Data, myerrors.NewHTTPErrorFromReponse(res)
	}

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return result.Data, err
	}

	err = json.Unmarshal(b, result)

	if err != nil {
		return result.Data, err
	}

	return result.Data, nil
}
