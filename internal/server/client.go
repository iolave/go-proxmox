package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) sendPVERequest(sr *http.Request) (*http.Response, error) {
	pveUrl, err := url.Parse(fmt.Sprintf(
		"https://%s:%d%s",
		s.PVEHost,
		s.PVEPort,
		sr.URL.Path,
	))
	b, err := io.ReadAll(sr.Body)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}

	cr, err := http.NewRequest(
		sr.Method,
		pveUrl.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	for k, v := range sr.Header {
		cr.Header.Set(k, strings.Join(v, ","))
	}

	res, err := s.c.Do(cr)
	return res, err
}
