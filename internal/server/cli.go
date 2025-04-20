package server

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

const (
	NAME = "pve-api-wrapper"
	// TODO: read and implement this https://akrabat.com/setting-the-version-of-a-go-application-when-building/
	VERSION = "v0.7.0"
)

//go:embed systemd.service
var systemdService []byte

type ServiceInstallCMD struct{}

func (s ServiceInstallCMD) Install() {
	cmd := exec.Command("which", "systemctl")
	if err := cmd.Run(); err != nil {
		fmt.Println("error: your system is not supported yet")
		os.Exit(1)
	}

	// create systemd service file
	f, err := os.Create("/etc/systemd/system/pve-api-wrapper.service")
	if err != nil {
		fmt.Println("error: unable to create systemd service file")
		os.Exit(1)
	}

	if _, err := f.Write(systemdService); err != nil {
		fmt.Println("error: unable to write systemd service file")
		os.Exit(1)
	}

	f.Close()
	fmt.Println("systemd service file created")
}

type ServiceCMD struct {
	Install *ServiceInstallCMD `arg:"subcommand:install" help:"install api-wrapper as a service"`
}

type App struct {
	Version    *bool       `arg:"--version" help:"display the program version"`
	PVEHost    string      `arg:"--pve-host,env:PVE_HOST" help:"Proxmox virtual environment host" default:"localhost"`
	PVEPort    int         `arg:"--pve-port,env:PVE_PORT" help:"Proxmox virtual environment port" default:"8006"`
	Host       string      `arg:"--host,env:WRAPPER_HOST" help:"API wrapper host" default:"localhost"`
	Port       int         `arg:"--port,env:WRAPPER_PORT" help:"API wrapper port" default:"8443"`
	TLSCrtPath string      `arg:"--crt" help:"API wrapper tls crt path" default:"/etc/pve/local/pve-ssl.pem"`
	TLSKeyPath string      `arg:"--key" help:"API wrapper tls key path" default:"/etc/pve/local/pve-ssl.key"`
	Service    *ServiceCMD `arg:"subcommand:service" help:"API wrapper service tools"`
}

func (a App) Start() {
	if a.Version != nil && *a.Version == true {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	cfg := serverConfig{
		PVEHost: a.PVEHost,
		PVEPort: a.PVEPort,
		Host:    a.Host,
		Port:    a.Port,
		TLSKey:  a.TLSKeyPath,
		TLSCrt:  a.TLSCrtPath,
	}

	s := New(cfg)

	if err := s.Start(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
