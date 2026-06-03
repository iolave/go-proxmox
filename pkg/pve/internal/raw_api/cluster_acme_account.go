package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterACMEGetAccount Return existing ACME account information.
//
// Parameters:
//
//   - name is the ACME account config file name.
//
// Required permissions:
//
//	Root only.
func ClusterACMEGetAccount(c *pve.Client, name string) (res struct {
	Account   map[string]any `json:"account"`
	Directory *string        `json:"directory"`
	Location  *string        `json:"location"`
	TOS       *string        `json:"tos"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/acme/account/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterACMEPostAccountRequest struct {
	// Contact email addresses.
	Contact string `in:"query=contact;nonzero"`

	// URL of ACME CA directory endpoint.
	Directory string `in:"query=directory;omitempty"`

	// HMAC key for External Account Binding.
	EABHMACKey string `in:"query=eab-hmac-key;omitempty"`

	// Key Identifier for External Account Binding.
	EABKid string `in:"query=eab-kid;omitempty"`

	// ACME account config file name.
	Name string `in:"query=name;omitempty"`

	// URL of CA TermsOfService - setting this indicates agreement.
	TOSURL string `in:"query=tos_url;omitempty"`
}

// ClusterACMEPostAccount Register a new ACME account with CA.
func ClusterACMEPostAccount(c *pve.Client, req ClusterACMEPostAccountRequest) (res string, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/acme/account",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterACMEDeleteAccount Deactivate existing ACME account at CA.
//
// Required permissions:
//
//	Root only.
func ClusterACMEDeleteAccount(c *pve.Client, name string) (res string, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/acme/account/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterACMEUpdateAccountRequest struct {
	// ACME account config file name.
	Name string `in:"path=name;nonzero"`

	// Contact email addresses.
	Contact string `in:"query=contact;omitempty"`
}

// ClusterACMEUpdateAccount Update existing ACME account information with CA. Note: not specifying any new account information triggers a refresh.
//
// Required permissions:
//
//	Root only.
func ClusterACMEUpdateAccount(c *pve.Client, req ClusterACMEUpdateAccountRequest) (res string, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/acme/account/{name}",
		Method:  http.MethodPut,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}
