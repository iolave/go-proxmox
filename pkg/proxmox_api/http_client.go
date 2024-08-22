package proxmoxapi

import (
	"crypto/tls"
	"net/http"
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
