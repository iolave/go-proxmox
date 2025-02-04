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

type getLXCIPResponse struct {
	IP string `json:"ip" example:"10.10.10.10"`
}

// Get LXC IP godoc
//
//	@Tags			LXC
//	@Summary		Get lxc assigned ip.
//	@Description	Gets the assigned ip of an lxc by running the lxc-info command in the host machine.
//	@Param			id	path	int	true	"vmid"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	server.getLXCIPResponse
//	@Failure		400	{object}	errors.HTTPError
//	@Failure		401	{object}	errors.HTTPError
//	@Failure		404	{object}	errors.HTTPError
//	@Failure		500	{object}	errors.HTTPError
//	@Router			/lxc/{id}/ip [get]
func getLXCIPHandlerV1(s *server) http.HandlerFunc {
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

		res := getLXCIPResponse{IP: ip}
		b, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}
