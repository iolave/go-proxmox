package pve

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func buildGetVersionSuccessResponse() pveResponse[GetVersionResponse] {
	return pveResponse[GetVersionResponse]{
		Data: GetVersionResponse{
			Release: "Release",
			Version: "Version",
			RepoID:  "RepoID",
		},
	}
}

func TestGetVersionSuccess(t *testing.T) {
	apiRes := buildGetVersionSuccessResponse()

	data, _ := json.Marshal(apiRes)
	server := newProxmoxApiTestServer(data, http.StatusOK)
	defer server.Close()
	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())

	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")

	api, err := New(Config{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	if err != nil {
		t.Logf("should create api, but error returned")
	}

	_, err = api.GetVersion()
	if err != nil {
		t.Logf("should not return an error")
	}

}
