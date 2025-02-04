package pveutils

import (
	"errors"
	"fmt"
	"os/exec"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

// ExecLXCCmd executes a command inside a proxmox lxc.
//   - If the lxc is not found err will be nil and exitCode
//     will be -1.
//   - The command output is stored in the out varaible.
//   - The command exit code is stored in the exitCode varaible.
//   - If the err variable is not nil it means that something failed
//     while executing the cmd and it DOES NOT correspond to the
//     cmd error itself.
func ExecLXCCmd(id int, shell string, cmd string) (out string, exitCode int, err error) {
	if !IsPVEHost() {
		return "", -1, errors.New("host is not a pve node")
	}

	supportedShells := []string{"bash"}
	if !slices.Contains(supportedShells, shell) {
		return "", -1, errors.New("shell not supported")
	}

	strid := fmt.Sprintf("%d", id)
	execId := uuid.New().String()
	// executes cmd and store its result
	c := exec.Command(
		"pct",
		"exec",
		strid,
		"--",
		shell,
		"-c",
		fmt.Sprintf(
			`result=$(%s 2>&1);echo $? > /tmp/exec-%s;echo "$result" >> /tmp/exec-%s`,
			cmd,
			execId,
			execId,
		),
	)

	if c == nil {
		return "", -1, errors.New("unable to create exec command")
	}

	_, err = c.Output()
	if err != nil {
		if !strings.Contains(err.Error(), "does not exist") {
			return "", -1, nil
		}
		exitErr := err.(*exec.ExitError)
		return "", -1, errors.New(string(exitErr.Stderr))
	}

	// retrieves cmd result
	resCmd := exec.Command(
		"pct",
		"exec",
		strid,
		"cat",
		fmt.Sprintf("/tmp/exec-%s", execId),
	)

	if resCmd == nil {
		return "", -1, errors.New("unable to create result command")
	}

	b, err := resCmd.Output()
	if err != nil {
		exitErr := err.(*exec.ExitError)
		return "", -1, errors.New(string(exitErr.Stderr))
	}

	execFile := string(b)
	firstNewLineIdx := strings.IndexRune(execFile, '\n')
	if firstNewLineIdx == -1 {
		return "", -1, errors.New("invalid execution file content")
	}

	exitCodeLine := execFile[0:firstNewLineIdx]
	exitCodeLine = strings.ReplaceAll(exitCodeLine, " ", "")

	exitCode, err = strconv.Atoi(exitCodeLine)
	if err != nil {
		return "", -1, err
	}
	out = execFile[firstNewLineIdx+1 : len(execFile)-1]

	// removes result log
	rmCmd := exec.Command(
		"pct",
		"exec",
		strid,
		"rm",
		fmt.Sprintf("/tmp/exec-%s", execId),
	)
	if rmCmd == nil {
		return "", -1, errors.New("unable to create rm command")
	}
	_, _ = rmCmd.Output()

	// returns result
	return out, exitCode, nil
}
