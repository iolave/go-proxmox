package cloudflare

import (
	"errors"
	"net/http"
)

type ServiceToken struct {
	ClientId     string // Client id created when creating a new cloudflare service.
	ClientSecret string // Client secret created when creating a new cloudflare service.
}

// NewServiceToken generates a pointer to a ServiceToken struct with it's
// properties initialized.
func NewServiceToken(clientId, secret string) *ServiceToken {
	return &ServiceToken{
		ClientId:     clientId,
		ClientSecret: secret,
	}
}

// Set adds the corresponding Cloudflare access client headers
// to the given request.
//
// It returns an error only when nil is passed to the request parameter.
func (t *ServiceToken) Set(req *http.Request) error {
	if req == nil {
		return errors.New("req parameter is missing")
	}

	req.Header.Add("CF-Access-Client-Id", t.ClientId)
	req.Header.Add("CF-Access-Client-Secret", t.ClientSecret)

	return nil
}
