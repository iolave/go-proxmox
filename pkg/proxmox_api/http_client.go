package proxmoxapi

import (
	"crypto/tls"
	"net/http"
)

type proxmoxHttpClient struct {
	host     string
	port     int
	insecure bool
	client   *http.Client
}

type httpClient = http.Client

func newHttpClient(insecureSkipVerify bool) *httpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	}

	return &http.Client{Transport: transport}
}
