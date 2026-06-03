package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterGetACMEMetaRequest struct {
	// URL of ACME CA directory endpoint. Defaults to https://acme-v02.api.letsencrypt.org/directory
	//
	// ^https?://.*
	Directory string `in:"query=directory;omitempty"`
}

// ClusterACMEGetMeta Retrieve ACME Directory Meta Information
//
// Required permissions:
//
//	Check: ["perm","/nodes/{node}",["Sys.Audit"]]
func ClusterACMEGetMeta(c *pve.Client, req ClusterGetACMEMetaRequest) (res struct {
	CAAIdentities           []string          `json:"caaIdentities"`
	ExternalAccountRequired *bool             `json:"externalAccountRequired"`
	TermsOfService          *string           `json:"termsOfService"`
	Website                 *string           `json:"website"`
	Profiles                map[string]string `json:"profiles"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/acme/meta",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterACMEGetDirectories Get named known ACME directory endpoints.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterACMEGetDirectories(c *pve.Client) (res []struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/acme/directories",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterACMEGetChallSchema Get schema of ACME challenge types.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterACMEGetChallSchema(c *pve.Client) (res []struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Schema struct {
		Name   string `json:"name,omitempty"`
		Fields map[string]struct {
			Type string `json:"type"`
			Desc string `json:"description"`
		} `json:"fields,omitempty"`
	} `json:"schema"`
	Type string `json:"type"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/acme/challenge-schema",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
