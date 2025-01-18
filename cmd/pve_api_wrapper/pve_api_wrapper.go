package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/iolave/go-proxmox/internal/server"
)

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

	app.Start()
}
