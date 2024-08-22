package proxmoxapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func TestGetVersionSuccess(t *testing.T) {
	apiRes := apiResponse[GetVersionResponse]{
		Data: GetVersionResponse{
			Release: "Release",
			Version: "Version",
			RepoID:  "RepoID",
		},
	}
	_ = apiRes.Data

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(apiRes)
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	api, err := New(ProxmoxAPIConfig{
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

func TestGetVersionErrorCuzStatusCode(t *testing.T) {
	apiRes := apiResponse[GetVersionResponse]{
		Data: GetVersionResponse{
			Release: "Release",
			Version: "Version",
			RepoID:  "RepoID",
		},
	}
	_ = apiRes.Data

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(apiRes)
		w.Write(data)
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	api, err := New(ProxmoxAPIConfig{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	if err != nil {
		t.Logf("should create api, but error returned")
	}

	res, err := api.GetVersion()
	t.Log(res, err)
	if err == nil {
		t.Logf("should return an error")
	}

}
func TestGetVersionErrorCuzInvalidBody(t *testing.T) {
	apiRes := apiResponse[GetVersionResponse]{
		Data: GetVersionResponse{
			Release: "Release",
			Version: "Version",
			RepoID:  "RepoID",
		},
	}
	_ = apiRes.Data

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := "text"
		w.Write([]byte(data))
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	api, err := New(ProxmoxAPIConfig{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	if err != nil {
		t.Logf("should create api, but error returned")
	}

	res, err := api.GetVersion()
	t.Log(res, err)
	if err == nil {
		t.Logf("should return an error")
	}

}
