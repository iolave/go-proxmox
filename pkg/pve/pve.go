package pve

import (
	"crypto/tls"
	"net/http"

	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/pkg/pve/internal/helpers"
)

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

type Client struct {
	// cfg is the proxmox api configuration.
	cfg Config

	// httpc is the underlying http client used
	// to send requests to the proxmox api.
	httpc *http.Client
}

// New returns a new go-proxmox client which can be used to
// create go-proxmox services.
//
// It returns an error when if the config is invalid.
//
// - Any error returned is of type [errors].Error.
// - It also initializes custom httpin directives.
//
// [errors]: https://pkg.go.dev/github.com/iolave/go-errors
func New(cfg Config) (*Client, error) {
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

	c := &Client{
		httpc: httpc,
		cfg:   cfg,
	}

	if err := helpers.Validate.Struct(c); err != nil {
		return nil, errors.NewWithNameAndErr(
			"validation_error",
			"config contains invalid values",
			err,
		)
	}

	return c, nil
}
