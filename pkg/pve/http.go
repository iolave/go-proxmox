package pve

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ggicci/httpin"
	"github.com/iolave/go-errors"
	strutils "github.com/iolave/go-proxmox/internal/str_utils"
)

// SendPVERequest sends a request to the proxmox api. It returns
// an error when the request fails.
//
// Any error returned is of type [errors].*HTTPError.
//
// [errors]: https://pkg.go.dev/github.com/iolave/go-errors
func (c Client) Do(req Request) error {
	base := fmt.Sprintf("%s://%s:%d", c.cfg.Proto, c.cfg.Host, c.cfg.Port)
	url, err := url.JoinPath(base, req.Path)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to build request url",
			err,
		)
	}
	if req.Payload == nil {
		req.Payload = struct{}{}
	}
	inreq, err := httpin.NewRequest(
		req.Method,
		url,
		req.Payload,
		httpin.Option.WithNestedDirectivesEnabled(true),
	)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to create request",
			err,
		)
	}

	// Add the custom headers to the request
	for k, v := range c.cfg.CustomHeaders {
		inreq.Header[k] = v
	}

	auth, err := c.cfg.Credentials.getAuthorization()
	if err != nil {
		return errors.NewInternalServerError(
			"failed to get authorization header",
			err,
		)
	}
	inreq.Header.Set("Authorization", auth)

	// If the request has additional payload,
	// a clone of the request is created in
	// order to parse the form data and populate
	// the additional payload. Then, a new reader
	// is created from the populated form data
	// and the original request body is set to it.
	if req.AdditionalPayload != nil {
		ctx := inreq.Context()
		reqClone := inreq.Clone(ctx)
		err := reqClone.ParseForm()
		if err != nil {
			return errors.NewInternalServerError(
				"failed to parse form data to add additional payload",
				err,
			)
		}

		for k, v := range req.AdditionalPayload {
			reqClone.Form.Add(k, v)
		}

		newBody := io.NopCloser(strings.NewReader(reqClone.Form.Encode()))
		inreq.Body = newBody
	}

	// Send the request
	res, err := c.httpc.Do(inreq)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to send request",
			err,
		)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to read response body",
			err,
		)
	}

	if res.StatusCode != http.StatusOK {
		return errors.NewHTTPError(
			res.StatusCode,
			fmt.Sprintf("%s_error", strutils.ToSnakeCase(http.StatusText(res.StatusCode))),
			string(b),
			nil,
		)
	}

	if req.Result == nil {
		return nil
	}

	pveres := struct {
		Data any `json:"data"`
	}{}
	err = json.Unmarshal(b, &pveres)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to unmarshal response body",
			err,
		)
	}

	b, _ = json.Marshal(pveres.Data)
	json.Unmarshal(b, &req.Result)

	return nil
}

// Request is the proxmox api request struct.
type Request struct {
	// Method is the http method.
	//
	// eg: GET, POST, PUT, DELETE
	Method string

	// Path is the http path.
	//
	// eg: /api2/json/nodes/{node}
	Path string

	// Payload is the request payload.
	//
	// It uses the httpin [Struct Tags] feature
	// to easily encode the form request.
	//
	// [Struct Tags]: https://github.com/ggicci/httpin/tree/62858140ae3d12b723a7ad8fa7bbf17c50a46d62?tab=readme-ov-file#add-httpin-directives-by-tagging-the-struct-fields-with-in
	Payload any

	// AdditionalPayload is an additional payload
	// that will be added to the request form data.
	AdditionalPayload map[string]string

	// Result is a pointer to a variable that will
	// be populated with the response. Passing nil
	// will prevent the response from being stored.
	Result any
}
