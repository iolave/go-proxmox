package proxmoxapi

import (
	"errors"
	"os"
)

type credentialEnum = int

const (
	CREDENTIALS_TOKEN credentialEnum = iota
)

const (
	CREDENTIALS_NOT_DETECTED_ERROR  = "credentials could not be detected from env"
	CREDENTIALS_NOT_SUPPORTED_ERROR = "credentials type not supported yet"
)

type credentials struct {
	credType  credentialEnum
	username  string
	tokenName string
	token     string
}

func NewTokenCredentials(user, tokenName, token string) *credentials {
	return &credentials{
		credType:  CREDENTIALS_TOKEN,
		username:  user,
		tokenName: tokenName,
		token:     token,
	}
}

// newCredentialsFromEnv() will return an error only when
// credentials are not detected from environment variables.
func newCredentialsFromEnv() (*credentials, error) {
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
