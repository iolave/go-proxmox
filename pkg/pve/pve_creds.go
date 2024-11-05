package pve

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Proxmox api client available credential types
type CredentialType int

// TODO: Add CREDENTIALS_PASSWORD support
const (
	CREDENTIALS_TOKEN CredentialType = iota
	CREDENTIALS_PASSWORD
)

// Credentials error messages.
const (
	CREDENTIALS_NOT_DETECTED_ERROR    = "credentials could not be detected from env"
	CREDENTIALS_NOT_SUPPORTED_ERROR   = "credentials type not supported yet"
	CREDENTIALS_MISSING_REQUEST_ERROR = "*http.Request parameter is nil"
)

// Credentials store proxmox api credentials.
type Credentials struct {
	credType  CredentialType
	username  string
	tokenName string
	token     string
}

// Set adds the corresponding PVE authorization headers
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
func (c *Credentials) Set(req *http.Request) error {
	if req == nil {
		return errors.New(CREDENTIALS_MISSING_REQUEST_ERROR)
	}

	if c.credType == CREDENTIALS_TOKEN {
		auth := fmt.Sprintf("PVEAPIToken=%s!%s=%s", c.username, c.tokenName, c.token)
		req.Header.Add("Authorization", auth)
		return nil
	}

	return errors.New(CREDENTIALS_NOT_SUPPORTED_ERROR)
}

// NewTokenCreds returns a struct containing proxmox token
// credentials that can be passed to a [pve api constructor].
//
// To create a pve token, read the [docs].
//
// [pve api constructor]: https://TODO:add-the-proper-ref
// [docs]: https://pve.proxmox.com/wiki/Proxmox_VE_API#API_Tokens
func NewTokenCreds(user, tokenName, token string) *Credentials {
	return &Credentials{
		credType:  CREDENTIALS_TOKEN,
		username:  user,
		tokenName: tokenName,
		token:     token,
	}
}

// NewEnvCreds get [environment variables] values and detects
// the type of credentials based on which envs are configured.
//
// It returns an error when a credential type is not detected.
//
// [environment variables]: https://go-proxmox.iolave.com/getting-started/environment-variables
func NewEnvCreds() (*Credentials, error) {
	username := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	tokenName := os.Getenv("PROXMOX_TOKEN_NAME")
	token := os.Getenv("PROXMOX_TOKEN")

	// If no username is found
	if username == "" {
		return nil, errors.New(CREDENTIALS_NOT_DETECTED_ERROR)
	}

	if tokenName != "" && token != "" {
		return NewTokenCreds(username, tokenName, token), nil
	}

	if password != "" {
		return nil, errors.New(CREDENTIALS_NOT_SUPPORTED_ERROR)
	}

	return nil, errors.New(CREDENTIALS_NOT_DETECTED_ERROR)
}
