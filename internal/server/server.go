package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type server struct {
	c   *http.Client
	s   *http.Server
	cfg serverConfig
}

type serverConfig struct {
	Host    string
	Port    int
	TLSKey  string
	TLSCrt  string
	PVEHost string
	PVEPort int
}

func New(cfg serverConfig) *server {
	server := new(server)
	server.cfg = cfg

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	server.c = c

	return server
}

func (s *server) Start() error {
	log.Println("starting server")

	mux := http.NewServeMux()

	// Proxmox api routes
	mux.HandleFunc("/{any...}", getHandler(s))

	s.s = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: mux,
	}
	defer s.s.Close()
	return s.s.ListenAndServeTLS(s.cfg.TLSCrt, s.cfg.TLSKey)
}
