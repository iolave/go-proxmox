package proxmoxapi

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	myerrors "github.com/iolave/go-proxmox/errors"
)

type httpClient = http.Client

func newHttpClient(insecureSkipVerify bool) *httpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	}

	return &http.Client{Transport: transport}
}

type apiResponse[T any] struct {
	Data T `json:"data"`
}

func sendRequest[Response any](method string, api *ProxmoxAPI, urlPath string, payload *url.Values) (Response, error) {
	url := api.buildHttpRequestUrl(urlPath)
	result := &apiResponse[Response]{}

	var req *http.Request
	var err error

	if payload != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(payload.Encode()))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return result.Data, err
	}

	if api.creds.credType == CREDENTIALS_TOKEN {
		auth := fmt.Sprintf("PVEAPIToken=%s!%s=%s", api.creds.username, api.creds.tokenName, api.creds.token)
		req.Header.Add("Authorization", auth)
	} else {
		return result.Data, errors.New("only token credentials are supported at the moment")
	}

	if api.config.CfServiceToken != nil {
		req.Header.Add("CF-Access-Client-Id", api.config.CfServiceToken.ClientId)
		req.Header.Add("CF-Access-Client-Secret", api.config.CfServiceToken.ClientSecret)
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
