package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host    string
	port    int
	tlsKey  string
	tlsCrt  string
	c       *http.Client
	s       *http.Server
	PVEHost string
	PVEPort int
}

func New(tlsCrt, tlsKey, pveHost string, pvePort int) *Server {
	server := new(Server)
	server.PVEHost = pveHost
	server.PVEPort = pvePort
	server.host = "localhost"
	server.port = 8443
	server.tlsCrt = tlsCrt
	server.tlsKey = tlsKey

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

func (s *Server) Start() error {
	log.Println("starting server")
	s.s = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: getHandler(s),
	}
	defer s.s.Close()
	return s.s.ListenAndServeTLS(s.tlsCrt, s.tlsKey)
}
