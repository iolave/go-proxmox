package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	apidef "github.com/iolave/go-proxmox/internal/api_wrapper/api_def"
	"github.com/iolave/go-proxmox/internal/api_wrapper/pve_utils"
	"github.com/iolave/go-proxmox/pkg/errors"
)

func addCustomRoutes(m *http.ServeMux, s *server) {
	m.HandleFunc("GET /custom-api/v1/lxc/{id}/ip", getLXCIPHandlerV1(s))
	m.HandleFunc("POST /custom-api/v1/lxc/{id}/exec", postLXCCMDHandlerV1(s))
	m.HandleFunc("POST /custom-api/v1/lxc/{id}/exec-async", postLXCCMDAsyncHandlerV1(s))
	m.HandleFunc("GET /custom-api/v1/cmd/{id}", getCMDResultHandlerV1(s))
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			return
		}

		if !authorized {
			httperr := errors.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
				err,
			)
			httperr.WriteResponse(w)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		ip, err := pveutils.GetLXCIPv4(id)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"failed to retrieve ip",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", fmt.Sprintf("ip=%s", ip), fmt.Sprintf("err=%s", string(b)))
			httperr.WriteResponse(w)
			return
		}

		if ip == "" {
			httperr := errors.NewHTTPError(
				http.StatusNotFound,
				"lxc not found",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		res := apidef.GetLXCIPResponse{IP: ip}
		b, _ := json.Marshal(res)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}

// POST LXC CMD godoc
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		if !authorized {
			httperr := errors.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"missing request body",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}
		in := apidef.PostLXCExecRequest{}
		if err = json.Unmarshal(b, &in); err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"invalid request body",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		if in.CMD == "" {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"found empty or missing .cmd property",
				nil,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		out, exitCode, err := pveutils.ExecLXCCmd(id, in.Shell, in.CMD)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"failed to execute command",
				map[string]string{"error": err.Error()},
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		if exitCode == -1 {
			httperr := errors.NewHTTPError(
				http.StatusNotFound,
				"lxc not found",
				nil,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		res := apidef.PostLXCExecResponse{Output: out, ExitCode: exitCode}
		b, _ = json.Marshal(res)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}

// POST LXC CMD Async godoc
//
//	@Tags		LXC
//	@Summary	Executes a command inside a proxmox lxc and returns an id used to query the result.
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
//	@Router		/lxc/{id}/exec-async [post]
func postLXCCMDAsyncHandlerV1(s *server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "started")
		authorized, err := s.IsUserAuthorized(r, "VMS", "VM.Console")
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"unable to authorize request",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		if !authorized {
			httperr := errors.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"missing request body",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}
		in := apidef.PostLXCExecRequest{}
		if err = json.Unmarshal(b, &in); err != nil {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"invalid request body",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		if in.CMD == "" {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"found empty or missing .cmd property",
				nil,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		execId, err := s.models.CMDExecution.New()
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"unable to store execution id",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		res := apidef.PostLXCExecAsyncResponse{ID: execId}
		b, _ = json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		go func() {
			// async proccess
			out, exitCode, err := pveutils.ExecLXCCmd(id, in.Shell, in.CMD)
			if err != nil {
				if err := s.models.CMDExecution.SetFailed(execId, err); err != nil {
					log.Println(r.Method, r.URL.Path, "failed_async", err.Error())
				}
				return
			}

			if exitCode == -1 {
				httperr := errors.NewHTTPError(
					http.StatusNotFound,
					"lxc not found",
					nil,
				)
				if err := s.models.CMDExecution.SetFailed(execId, httperr); err != nil {
					log.Println(r.Method, r.URL.Path, "failed_async", err.Error())
				}
				return
			}

			if err := s.models.CMDExecution.SetSucceeded(execId, out, exitCode); err != nil {
				log.Println(r.Method, r.URL.Path, "failed_async", err.Error())
				return
			}
			log.Println(r.Method, r.URL.Path, "succeeded_async")
			return
		}()
	})
}

// Get CMD execution result godoc
//
//	@Tags		CMD
//	@Summary	Get cmd execution result.
//	@Description	Requires VM.Audit scope. Gets a cmd execution result using the id returned by an exec-async endpoint.
//	@Param		id	path	string	true	"execution id"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	models.CMDExecution
//	@Failure	400	{object}	errors.HTTPError
//	@Failure	401	{object}	errors.HTTPError
//	@Failure	404	{object}	errors.HTTPError
//	@Failure	500	{object}	errors.HTTPError
//	@Router		/cmd/{id} [get]
func getCMDResultHandlerV1(s *server) http.HandlerFunc {
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
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			return
		}

		if !authorized {
			httperr := errors.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
				err,
			)
			httperr.WriteResponse(w)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			return
		}

		id := r.PathValue("id")
		if id == "" {
			httperr := errors.NewHTTPError(
				http.StatusBadRequest,
				"id property is empty",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", string(b))
			httperr.WriteResponse(w)
			return
		}

		result, err := s.models.CMDExecution.Get(id)
		if err != nil {
			httperr := errors.NewHTTPError(
				http.StatusInternalServerError,
				"failed to retrieve cmd result",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", fmt.Sprintf("id=%s", id), fmt.Sprintf("err=%s", string(b)))
			httperr.WriteResponse(w)
			return
		}

		if result == nil {
			httperr := errors.NewHTTPError(
				http.StatusNotFound,
				"cmd execution result not found",
				err,
			)
			b, _ := httperr.Marshall()
			log.Println(r.Method, r.URL.Path, "failed", fmt.Sprintf("id=%s", id), fmt.Sprintf("err=%s", string(b)))
			httperr.WriteResponse(w)
			return
		}

		b, _ := json.Marshal(result)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		log.Println(r.Method, r.URL.Path, "succeeded")
		return
	})
}
