package proxmoxapi

import (
	"os"
	"testing"
)

func TestNewCredentialsFromEnvErrorWhenNoUsername(t *testing.T) {
	os.Setenv("PROXMOX_USERNAME", "")

	_, err := newCredentialsFromEnv()

	if err == nil {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", CREDENTIALS_NOT_DETECTED_ERROR, err)
	}

	if err.Error() != CREDENTIALS_NOT_DETECTED_ERROR {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", CREDENTIALS_NOT_DETECTED_ERROR, err)
	}
}

func TestNewCredentialsFromEnvSuccessWithToken(t *testing.T) {
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")

	_, err := newCredentialsFromEnv()

	if err != nil {
		t.Fatalf("newCredentialsFromEnv(), expected err='<nil>', got err='%v'", err)
	}
}

func TestNewCredentialsFromEnvErrorCuzMissingToken(t *testing.T) {
	expectedErr := CREDENTIALS_NOT_DETECTED_ERROR
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "")

	_, err := newCredentialsFromEnv()

	if err == nil {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", expectedErr, err)
	}

	if err.Error() != expectedErr {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", expectedErr, err)
	}
}

func TestNewCredentialsFromEnvErrorCuzMissingTokenName(t *testing.T) {
	expectedErr := CREDENTIALS_NOT_DETECTED_ERROR
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "")
	os.Setenv("PROXMOX_TOKEN", "token")

	_, err := newCredentialsFromEnv()

	if err == nil {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", expectedErr, err)
	}

	if err.Error() != expectedErr {
		t.Fatalf("newCredentialsFromEnv(), expected err='%s', got err='%v'", expectedErr, err)
	}
}
