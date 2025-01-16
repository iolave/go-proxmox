package server

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func getHandler(s *Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "started")
		res, err := s.sendPVERequest(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed sending request to proxmox api", err)
			return
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			log.Println("error reading response body")
			return
		}

		for k, v := range res.Header {
			w.Header().Set(k, strings.Join(v, ","))
		}
		w.WriteHeader(res.StatusCode)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "finished with status", res.StatusCode)
		return
	})
}
