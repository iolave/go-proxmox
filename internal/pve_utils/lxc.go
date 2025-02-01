package pveutils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/iolave/go-proxmox/pkg/helpers"
)

const ipRegex = `(?P<ip>((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))`

// GetLXCIPv4
//   - if lxc is not found, empty string is returned
//   - if host is not a pve machine, an error is returned
func GetLXCIPv4(id int) (string, error) {
	if !IsPVEHost() {
		return "", errors.New("host is not a pve node")
	}

	strid := fmt.Sprintf("%d", id)
	cmd := exec.Command(
		"/usr/bin/lxc-info", "-i", "-n", strid,
	)
	if cmd == nil {
		return "", errors.New("unable to create a command")
	}

	b, err := cmd.Output()
	if err != nil {
		exitErr := err.(*exec.ExitError)
		return "", errors.New(string(exitErr.Stderr))
	}

	outstr := string(b)
	original := outstr
	lines := strings.Split(outstr, "\n")
	firstLine := ""
	if len(lines) >= 1 {
		firstLine = lines[0]
	}

	params := helpers.GetRegexpParams(ipRegex, firstLine)
	ip := params["ip"]

	if ip == "" {
		if !strings.Contains(original, "doesn't exist") {
			return "", fmt.Errorf("got invalid data: %s", original)
		}

		return "", nil
	}

	return ip, nil
}
