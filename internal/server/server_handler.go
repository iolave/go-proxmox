package server

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/iolave/go-proxmox/pkg/errors"
)

func getHandler(s *server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "started")
		res, httpErr := s.sendPVERequest(r)
		if httpErr != nil {
			httpErr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "error sending request to proxmox api:", httpErr.Error())
			return
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			httpErr := errors.NewHTTPError(res.StatusCode, "unable to read proxmox response body", err)
			httpErr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "error reading response body")
			return
		}

		for k, v := range res.Header {
			w.Header().Set(k, strings.Join(v, ","))
		}

		if res.StatusCode != http.StatusOK {
			msg := strings.Join(strings.Split(res.Status, " ")[1:], " ")
			httpErr := errors.NewHTTPError(res.StatusCode, msg, string(b))
			httpErr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed with status", res.StatusCode)
			return
		}

		w.WriteHeader(res.StatusCode)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}
