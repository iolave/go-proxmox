package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	_ "github.com/iolave/go-proxmox/docs/api-wrapper"
	"github.com/iolave/go-proxmox/internal/api_wrapper/server"
)

// @title			Proxmox API Wrapper
// @version			1.0
// @description			Proxmox api wrapper that provides custom features.
//
// @contact.name		Ignacio Olave
// @contact.url			http://www.github.com/iolave/go-proxmox/issues
// @contact.email		contact@iolave.com
// @license.name		Apache 2.0
// @license.url			http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8443
// @BasePath			/custom-api/v1
// @externalDocs.description	go-promox docs
// @externalDocs.url		https://go-proxmox.iolave.com
func main() {
	var app server.App

	p, err := arg.NewParser(arg.Config{Program: server.NAME}, &app)
	if err != nil {
		log.Fatalf("there was an error in the definition of the Go struct: %v\n", err)
	}
	err = p.Parse(os.Args[1:])
	switch {
	case err == arg.ErrHelp: // found "--help" on command line
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	case err != nil:
		fmt.Printf("error: %v\n", err)
		p.WriteUsage(os.Stdout)
		os.Exit(1)
	}

	switch {
	case app.Service != nil:
		switch {
		case app.Service.Install != nil:
			app.Service.Install.Install()
		}
	default:
		app.Start()
	}
}
