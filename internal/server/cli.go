package server

import (
	"fmt"
	"os"
)

const (
	NAME = "pve-api-wrapper"
	// TODO: read and implement this https://akrabat.com/setting-the-version-of-a-go-application-when-building/
	VERSION = "v0.5.1"
)

type App struct {
	Version    *bool  `arg:"--version" help:"display the program version"`
	PVEHost    string `arg:"--pve-host,env:PVE_HOST" help:"Proxmox virtual environment host" default:"localhost"`
	PVEPort    int    `arg:"--pve-port,env:PVE_PORT" help:"Proxmox virtual environment port" default:"8006"`
	Host       string `arg:"--host,env:WRAPPER_HOST" help:"API wrapper host" default:"localhost"`
	Port       int    `arg:"--port,env:WRAPPER_PORT" help:"API wrapper port" default:"8443"`
	TLSCrtPath string `arg:"--crt" help:"API wrapper tls crt path" default:"/etc/pve/local/pve-ssl.pem"`
	TLSKeyPath string `arg:"--key" help:"API wrapper tls key path" default:"/etc/pve/local/pve-ssl.key"`
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
