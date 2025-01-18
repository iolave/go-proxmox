package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/iolave/go-proxmox/pkg/errors"
)

func (s *server) sendPVERequest(sr *http.Request) (*http.Response, *errors.HTTPError) {
	pveUrl, err := url.Parse(fmt.Sprintf(
		"https://%s:%d%s",
		s.cfg.PVEHost,
		s.cfg.PVEPort,
		sr.URL.Path,
	))
	b, err := io.ReadAll(sr.Body)
	if err != nil {
		return nil, &errors.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Name:       "internal_server_error",
			Message:    "unable to read request body",
			Original:   err,
		}
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, &errors.HTTPError{
			StatusCode: http.StatusBadRequest,
			Name:       "bad_request_error",
			Message:    "unable to parse url query",
			Original:   err,
		}
	}

	cr, err := http.NewRequest(
		sr.Method,
		pveUrl.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, &errors.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Name:       "internal_server_error",
			Message:    "unable to generate new request",
			Original:   err,
		}
	}

	for k, v := range sr.Header {
		cr.Header.Set(k, strings.Join(v, ","))
	}

	res, err := s.c.Do(cr)

	if err != nil {
		return res, &errors.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Name:       "internal_server_error",
			Message:    "unable to send request to proxmox",
			Original:   err,
		}
	}

	return res, nil
}
