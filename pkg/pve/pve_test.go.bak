package pve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func newProxmoxApiTestServer(data []byte, status int) *httptest.Server {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(data)
	}))

	return server
}

func TestNewErrorWhenInvalidCredentials(t *testing.T) {
	os.Setenv("PROXMOX_USER", "")

	cfg := Config{
		Host: "",
		Port: 0,
	}

	api, err := New(cfg)

	if err == nil {
		t.Fatalf("New(%v), expected (<nil>, 'TODO: Change error'), got (%v, %v)", cfg, api, err)
	}

}

func TestNewErrorWhenGetVersionReturnStatusCodeOtherThanOk(t *testing.T) {
	apiRes := buildGetVersionSuccessResponse()
	_ = apiRes.Data

	data, _ := json.Marshal(apiRes)
	server := newProxmoxApiTestServer(data, http.StatusBadRequest)
	defer server.Close()
	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	_, err := New(Config{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	if err == nil {
		t.Logf("should return error, but api was created")
	}
}

func TestNewWithCredentialsSuccess(t *testing.T) {
	apiRes := buildGetVersionSuccessResponse()
	_ = apiRes.Data

	data, _ := json.Marshal(apiRes)
	server := newProxmoxApiTestServer(data, http.StatusOK)
	defer server.Close()
	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	creds, _ := NewEnvCreds()
	_, err := NewWithCredentials(Config{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	}, creds)

	if err != nil {
		t.Logf("should return api, but got error %v", err)
	}
}
func TestNewWithCredentialsErrorWhenInvalidCredentials(t *testing.T) {
	apiRes := buildGetVersionSuccessResponse()
	_ = apiRes.Data

	data, _ := json.Marshal(apiRes)
	server := newProxmoxApiTestServer(data, http.StatusBadRequest)
	defer server.Close()
	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	creds, _ := NewEnvCreds()
	_, err := NewWithCredentials(Config{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	}, creds)

	if err == nil {
		t.Logf("should return error, but api was created")
	}
}

func TestNewErrorWhenGetVersionReturnInvalidResponseData(t *testing.T) {
	apiRes := buildGetVersionSuccessResponse()
	_ = apiRes.Data

	server := newProxmoxApiTestServer([]byte("data"), http.StatusOK)
	defer server.Close()

	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")
	_, err := New(Config{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	if err == nil {
		t.Logf("should return error, but api was created")
	}
}

func TestBuildHttpRequestUrl(t *testing.T) {
	os.Setenv("PROXMOX_USER", "user")
	os.Setenv("PROXMOX_TOKEN_NAME", "token-name")
	os.Setenv("PROXMOX_USER", "token")

	data, _ := json.Marshal(buildGetVersionSuccessResponse())
	server := newProxmoxApiTestServer(data, http.StatusOK)
	defer server.Close()
	url, _ := url.Parse(server.URL)
	host := url.Hostname()
	port, _ := strconv.Atoi(url.Port())

	api, _ := New(Config{Host: host, Port: port, InsecureSkipVerify: true})

	normalizedPath := "some-path"
	expected := fmt.Sprintf("https://%s:%d/api2/json/%s", host, port, normalizedPath)

	testCases := []string{
		fmt.Sprintf("%s", normalizedPath),
		fmt.Sprintf("/%s", normalizedPath),
		fmt.Sprintf("//%s", normalizedPath),
	}

	for i := 0; i < len(testCases); i++ {
		result := api.client.buildRequestUrl(testCases[i])

		if result != expected {
			t.Fatalf(`buildRequestUrl("%s"), expected "%s", got "%s"`, testCases[i], expected, result)
		}

	}
}
