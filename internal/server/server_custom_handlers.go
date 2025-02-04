package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	apidef "github.com/iolave/go-proxmox/internal/api_def"
	"github.com/iolave/go-proxmox/internal/pve_utils"
	"github.com/iolave/go-proxmox/pkg/errors"
)

func addCustomRoutes(m *http.ServeMux, s *server) {
	m.HandleFunc("GET /custom-api/v1/lxc/{id}/ip", getLXCIPHandlerV1(s))
	m.HandleFunc("POST /custom-api/v1/lxc/{id}/exec", postLXCCMDHandlerV1(s))
}

// Get LXC IP godoc
//
//	@Tags			LXC
//	@Summary		Get lxc assigned ip.
//	@Description	Gets the assigned ip of an lxc by running the lxc-info command in the host machine.
//	@Param			id	path	int	true	"vmid"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apidef.GetLXCIPResponse
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

		res := apidef.GetLXCIPResponse{IP: ip}
		b, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}

// Get LXC IP godoc
//
//	@Tags		LXC
//	@Summary	Executes a command inside a proxmox lxc.
//	@Description	Requires VM.Console scope. Run a command inside a proxmox lxc. It stores the exit code and it's output in a temprary file in the /tmp dir. Once that file is read, it is removed.
//	@Param		id	path	int	true	"vmid"
//	@Param		request body	apidef.PostLXCExecRequest true "command input"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	apidef.PostLXCExecResponse
//	@Failure	400	{object}	errors.HTTPError
//	@Failure	401	{object}	errors.HTTPError
//	@Failure	404	{object}	errors.HTTPError
//	@Failure	500	{object}	errors.HTTPError
//	@Router		/lxc/{id}/exec [post]
func postLXCCMDHandlerV1(s *server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "started")
		authorized, err := s.IsUserAuthorized(r, "VMS", "VM.Console")
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

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"missing request body",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}
		in := apidef.PostLXCExecRequest{}
		if err = json.Unmarshal(b, &in); err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"invalid request body",
				err,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		if in.CMD == "" {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"found empty or missing .cmd property",
				nil,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		out, exitCode, err := pveutils.ExecLXCCmd(id, in.Shell, in.CMD)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"failed to execute command",
				map[string]string{"error": err.Error()},
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		if exitCode == -1 {
			httperr := errors.NewHTTPError(
				http.StatusNotFound,
				"lxc not found",
				nil,
			)
			httperr.WriteResponse(w)
			log.Println(r.Method, r.URL.Path, "failed", httperr.Message)
			return
		}

		res := apidef.PostLXCExecResponse{Output: out, ExitCode: exitCode}
		b, _ = json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}
