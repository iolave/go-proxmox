package proxmoxapi

import (
	"fmt"
	"os"
	"testing"
)

func TestNewErrorWhenInvalidCredentials(t *testing.T) {
	os.Setenv("PROXMOX_USER", "")

	cfg := ProxmoxAPIConfig{
		Host: "",
		Port: 0,
	}

	api, err := New(cfg)

	if err == nil {
		t.Fatalf("New(%v), expected (<nil>, 'TODO: Change error'), got (%v, %v)", cfg, api, err)
	}

}

func TestBuildHttpRequestUrl(t *testing.T) {
	os.Setenv("PROXMOX_USER", "user")
	os.Setenv("PROXMOX_TOKEN_NAME", "token-name")
	os.Setenv("PROXMOX_USER", "token")

	cfg := ProxmoxAPIConfig{
		Host: "",
		Port: 0,
	}

	api, _ := New(cfg)

	normalizedPath := "some-path"
	expected := fmt.Sprintf("https://%s:%d/api2/json/%s", cfg.Host, cfg.Port, normalizedPath)

	testCases := []string{
		fmt.Sprintf("%s", normalizedPath),
		fmt.Sprintf("/%s", normalizedPath),
		fmt.Sprintf("//%s", normalizedPath),
	}

	for i := 0; i < len(testCases); i++ {
		result := api.buildHttpRequestUrl(testCases[i])

		if result != expected {
			t.Fatalf(`buildHttpRequestUrl("%s"), expected "%s", got "%s"`, testCases[i], expected, result)
		}

	}
}
