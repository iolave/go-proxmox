package proxmoxapi

import (
	"errors"
	"os"
)

type credentialEnum = int

// Supported credetials.
const (
	CREDENTIALS_TOKEN credentialEnum = iota
)

// Credentials error messages.
const (
	CREDENTIALS_NOT_DETECTED_ERROR  = "credentials could not be detected from env"
	CREDENTIALS_NOT_SUPPORTED_ERROR = "credentials type not supported yet"
)

// Credentials is the struct that stores proxmox api credentials.
type Credentials struct {
	credType  credentialEnum
	username  string
	tokenName string
	token     string
}

// NewTokenCredentials returns a struct containing proxmox token credentials that can be passed to the [NewWithCredentials] method.
//
// [NewWithCredentials]: https://go-proxmox.iolave.com/reference/pkg/proxmox_api/#func-newwithcredentials
func NewTokenCredentials(user, tokenName, token string) *Credentials {
	return &Credentials{
		credType:  CREDENTIALS_TOKEN,
		username:  user,
		tokenName: tokenName,
		token:     token,
	}
}

// newCredentialsFromEnv() will return an error only when
// credentials are not detected from environment variables.
func newCredentialsFromEnv() (*Credentials, error) {
	username := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	tokenName := os.Getenv("PROXMOX_TOKEN_NAME")
	token := os.Getenv("PROXMOX_TOKEN")

	// If no username is found
	if username == "" {
		return nil, errors.New(CREDENTIALS_NOT_DETECTED_ERROR)
	}

	if tokenName != "" && token != "" {
		return NewTokenCredentials(username, tokenName, token), nil
	}

	if password != "" {
		return nil, errors.New(CREDENTIALS_NOT_SUPPORTED_ERROR)
	}

	return nil, errors.New(CREDENTIALS_NOT_DETECTED_ERROR)
}
