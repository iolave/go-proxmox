package main

import (
	"fmt"
	"os"

	"github.com/iolave/go-proxmox/internal/server"
)

func main() {
	s := server.New("", "", "localhost", 8006)

	err := s.Start()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}

}
