package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ggicci/httpin"
	"github.com/go-playground/validator/v10"
	errors "github.com/iolave/go-errors"
	strutils "github.com/iolave/go-proxmox/internal/str_utils"
)

// API is an http client used to send requests
// to the proxmox api.
type API struct {
	// cfg is the proxmox api configuration.
	cfg Config

	// httpc is the underlying http client used
	// to send requests to the proxmox api.
	httpc *http.Client
}

type Config struct {
	// CustomHeaders is a map of custom headers
	// to be sent with each request.
	CustomHeaders http.Header

	// Proto is the protocol used to send requests
	// to the proxmox api.
	Proto string `validate:"required,oneof=http https"`

	// Host is the host used to send requests
	// to the proxmox api.
	Host string `validate:"required"`

	// Port is the port used to send requests
	// to the proxmox api.
	Port int `validate:"required"`

	// Credentials is the proxmox api credentials.
	Credentials *Credentials `validate:"required"`

	// InsecureSkipVerify is a flag that indicates
	// whether the client should skip verifying the
	// server's certificate chain and host name.
	// It is used to disable SSL certificate verification.
	InsecureSkipVerify bool
}

// New returns a new Proxmox API client. It returns an
// error when the config is invalid.
//
// - Any error returned is of type [errors].Error.
// - It also initializes custom httpin directives.
//
// [errors]: https://pkg.go.dev/github.com/iolave/go-errors
func New(cfg Config) (*API, error) {
	httpinInit()

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.InsecureSkipVerify,
		},
	}
	httpc := &http.Client{Transport: transport}

	if cfg.CustomHeaders == nil {
		cfg.CustomHeaders = http.Header{}
	}

	c := &API{
		httpc: httpc,
		cfg:   cfg,
	}

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)
	if err := validate.Struct(c); err != nil {
		return nil, errors.NewWithNameAndErr(
			"validation error",
			"config contains invalid values",
			err,
		)
	}

	return c, nil
}

// PVERequest is the proxmox api request struct.
type PVERequest struct {
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

// SendPVERequest sends a request to the proxmox api. It returns
// an error when the request fails.
//
// Any error returned is of type [errors].*HTTPError.
//
// [errors]: https://pkg.go.dev/github.com/iolave/go-errors
func (c API) SendPVERequest(pvereq PVERequest) error {
	base := fmt.Sprintf("%s://%s:%d", c.cfg.Proto, c.cfg.Host, c.cfg.Port)
	url, err := url.JoinPath(base, pvereq.Path)
	if err != nil {
		return errors.NewInternalServerError(
			"failed to build request url",
			err,
		)
	}
	if pvereq.Payload == nil {
		pvereq.Payload = struct{}{}
	}
	req, err := httpin.NewRequest(
		pvereq.Method,
		url,
		pvereq.Payload,
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
		req.Header[k] = v
	}

	auth, err := c.cfg.Credentials.getAuthorization()
	if err != nil {
		return errors.NewInternalServerError(
			"failed to get authorization header",
			err,
		)
	}
	req.Header.Set("Authorization", auth)

	// If the request has additional payload,
	// a clone of the request is created in
	// order to parse the form data and populate
	// the additional payload. Then, a new reader
	// is created from the populated form data
	// and the original request body is set to it.
	if pvereq.AdditionalPayload != nil {
		ctx := req.Context()
		reqClone := req.Clone(ctx)
		err := reqClone.ParseForm()
		if err != nil {
			return errors.NewInternalServerError(
				"failed to parse form data to add additional payload",
				err,
			)
		}

		for k, v := range pvereq.AdditionalPayload {
			reqClone.Form.Add(k, v)
		}

		newBody := io.NopCloser(strings.NewReader(reqClone.Form.Encode()))
		req.Body = newBody
	}

	// Send the request
	res, err := c.httpc.Do(req)
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

	if pvereq.Result == nil {
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
	json.Unmarshal(b, &pvereq.Result)

	return nil
}
