package pve

import (
	"os"
	"testing"
)

func TestNewEnvCredsErrorWhenNoUsername(t *testing.T) {
	os.Setenv("PROXMOX_USERNAME", "")

	_, err := NewEnvCreds()

	if err == nil {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", CREDENTIALS_NOT_DETECTED_ERROR, err)
	}

	if err.Error() != CREDENTIALS_NOT_DETECTED_ERROR {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", CREDENTIALS_NOT_DETECTED_ERROR, err)
	}
}

func TestNewEnvCredsSuccessWithToken(t *testing.T) {
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "token")

	_, err := NewEnvCreds()

	if err != nil {
		t.Fatalf("NewEnvCreds(), expected err='<nil>', got err='%v'", err)
	}
}

func TestNewEnvCredsErrorCuzMissingToken(t *testing.T) {
	expectedErr := CREDENTIALS_NOT_DETECTED_ERROR
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "token_name")
	os.Setenv("PROXMOX_TOKEN", "")

	_, err := NewEnvCreds()

	if err == nil {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", expectedErr, err)
	}

	if err.Error() != expectedErr {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", expectedErr, err)
	}
}

func TestNewEnvCredsErrorCuzMissingTokenName(t *testing.T) {
	expectedErr := CREDENTIALS_NOT_DETECTED_ERROR
	os.Setenv("PROXMOX_USERNAME", "username")
	os.Setenv("PROXMOX_TOKEN_NAME", "")
	os.Setenv("PROXMOX_TOKEN", "token")

	_, err := NewEnvCreds()

	if err == nil {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", expectedErr, err)
	}

	if err.Error() != expectedErr {
		t.Fatalf("NewEnvCreds(), expected err='%s', got err='%v'", expectedErr, err)
	}
}
