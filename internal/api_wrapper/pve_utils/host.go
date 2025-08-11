package pveutils

import "os/exec"

func IsPVEHost() bool {
	cmd := exec.Command("/usr/bin/which", "pvedaemon")
	if cmd == nil {
		return false
	}

	_, err := cmd.Output()
	if err != nil {
		return false
	}

	return true
}
