package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/iolave/go-proxmox/internal/pve_utils"
	"github.com/iolave/go-proxmox/pkg/errors"
)

func addCustomRoutes(m *http.ServeMux, s *server) {
	m.HandleFunc("GET /custom-api/v1/lxc/{id}/ip", getLXCIPHandlerV1(s))
}

func getLXCIPHandlerV1(s *server) http.HandlerFunc {
	// lxc-info -i -n 100 | head -n 1 | awk '{print $2}'
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "started")
		authorized, err := s.IsUserAuthorized(r, "VMS", "VM.Audit")
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"unable to authorize request",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		if !authorized {
			httperr := errors.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		rawId := r.PathValue("id")
		id, err := strconv.Atoi(rawId)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"id property from path is not a number",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		ip, err := pveutils.GetLXCIPv4(id)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"failed to retrieve ip",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		if ip == "" {
			httperr := errors.NewHTTPError(
				http.StatusNotFound,
				"lxc not found",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		res := struct {
			IP string `json:"ip"`
		}{IP: ip}
		b, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}
