package pve

import (
	"errors"
	"fmt"
	"net/http"
)

// go-proxmox client credential type
type CredentialType int

const (
	// Credential type for pve api token
	CREDENTIALS_TOKEN CredentialType = iota
)

// Credentials error messages.
const (
	CREDS_TYPE_NOT_SUPPORTED_ERROR = "credentials type not supported yet"
	CREDS_MISSING_REQ_PARAM_ERROR  = "*http.Request parameter is nil"
)

// Credential
type Credentials struct {
	credType  CredentialType
	username  string
	tokenName string
	token     string
}

// getAuthorization builds the authorization header based on the
// credentials type.
//
// It returns an error if the credentials type is not supported.
func (c Credentials) getAuthorization() (string, error) {
	switch c.credType {
	case CREDENTIALS_TOKEN:
		auth := fmt.Sprintf("PVEAPIToken=%s!%s=%s",
			c.username,
			c.tokenName,
			c.token,
		)
		return auth, nil
	default:
		return "", errors.New(CREDS_TYPE_NOT_SUPPORTED_ERROR)
	}
}

// set adds the corresponding PVE authorization headers
// to the req parameter.
//
// * It returns an error with the [CREDENTIALS_MISSING_REQUEST_ERROR] message
// when nil is passed to the req parameter.
//
// * It returns an error with the [CREDENTIALS_NOT_SUPPORTED_ERROR] message
// when [CredentialType] is not supported.
//
// [CREDENTIALS_MISSING_REQUEST_ERROR]: https://go-proxmox.iolave.com/reference/pkg/pve#constants
// [CREDENTIALS_NOT_SUPPORTED_ERROR]: https://go-proxmox.iolave.com/reference/pkg/pve#constants
// [CredentialType]: https://go-proxmox.iolave.com/reference/pkg/pve#type-credentialtype
func (c *Credentials) set(req *http.Request) error {
	if req == nil {
		return errors.New(CREDS_MISSING_REQ_PARAM_ERROR)
	}

	var auth string
	switch c.credType {
	case CREDENTIALS_TOKEN:
		auth = fmt.Sprintf("PVEAPIToken=%s!%s=%s", c.username, c.tokenName, c.token)
		break
	default:
		return errors.New(CREDS_TYPE_NOT_SUPPORTED_ERROR)
	}

	req.Header.Add("Authorization", auth)
	return nil
}

// NewTokenCreds returns a [Credentials] struct containing pve token
// credentials that can be passed to the [New] function.
//
// To create a pve token, read the [docs].
//
// [docs]: https://pve.proxmox.com/wiki/Proxmox_VE_API#API_Tokens
func NewTokenCreds(user, tokenName, token string) *Credentials {
	return &Credentials{
		credType:  CREDENTIALS_TOKEN,
		username:  user,
		tokenName: tokenName,
		token:     token,
	}
}
