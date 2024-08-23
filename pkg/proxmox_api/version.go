package proxmoxapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetVersionResponse struct {
	Release string `json:"release"`
	Version string `json:"version"`
	RepoID  string `json:"repoid"`
}

func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error) {
	url := api.buildHttpRequestUrl("/version")

	result := &apiResponse[GetVersionResponse]{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return result.Data, err
	}

	if api.creds.credType == CREDENTIALS_TOKEN {
		auth := fmt.Sprintf("PVEAPIToken=%s!%s=%s", api.creds.username, api.creds.tokenName, api.creds.token)
		req.Header.Add("Authorization", auth)
	} else {
		return result.Data, errors.New("TODO: Change error: only token credentials supported")
	}

	res, err := api.httpClient.Do(req)

	if err != nil {
		return result.Data, err
	}

	if res.StatusCode != http.StatusOK {
		return result.Data, errors.New("TODO: Change error: status code is not 200")
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
