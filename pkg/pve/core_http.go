package pve

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/iolave/go-proxmox/pkg/cloudflare"
	"github.com/iolave/go-proxmox/pkg/errors"
	"github.com/iolave/go-proxmox/pkg/helpers"
)

type httpClient struct {
	Client       *http.Client
	Creds        *Credentials
	Host         string
	Port         int
	ServiceToken *cloudflare.ServiceToken
	APIWrapper   bool
}

func newHttpClient(creds *Credentials, st *cloudflare.ServiceToken, host string, port int, insecureSkipVerify bool, wrapper bool) *httpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	}

	return &httpClient{
		Client:       &http.Client{Transport: transport},
		Creds:        creds,
		Host:         host,
		Port:         port,
		ServiceToken: st,
		APIWrapper:   wrapper,
	}
}

type pveResponse[T any] struct {
	Data T `json:"data"`
}

// buildHttpRequestUrl builds the proxmox api url based in the
// given path and the configured api host and port.
func (c *httpClient) buildRequestUrl(path string) string {
	checkForwardSlashRune := func(r rune) bool {
		if r == '/' {
			return true
		}

		return false
	}

	path = strings.TrimFunc(path, checkForwardSlashRune)

	return fmt.Sprintf("https://%s:%d/api2/json/%s", c.Host, c.Port, path)
}

// sendReq sends an http request to the configured proxmox api.
//
// It stores the response value into the result parameter only
// if no error has been returned. If an error is returned, the
// passed result parameter will be intact.
func (c *httpClient) sendReq(method, path string, payload *url.Values, result any) error {
	url := c.buildRequestUrl(path)

	var req *http.Request
	var err error

	if payload != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(payload.Encode()))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	if err := c.Creds.Set(req); err != nil {
		return err
	}

	if c.ServiceToken != nil {
		c.ServiceToken.Set(req)
	}

	res, err := c.Client.Do(req)

	if err != nil {
		return err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		if !c.APIWrapper {
			return errors.NewHTTPErrorFromResponse(res)
		}

		httpErr := new(errors.HTTPError)
		if err := json.Unmarshal(b, httpErr); err != nil {
			return err
		}

		return httpErr

	}

	switch t := reflect.TypeOf(result); t {
	case reflect.TypeFor[*string]():
		pveRes := &pveResponse[string]{}
		err = json.Unmarshal(b, pveRes)
		if err != nil {
			return err
		}
		*result.(*string) = pveRes.Data
		return nil

	case reflect.TypeFor[*int]():
		pveRes := &pveResponse[int]{}
		err = json.Unmarshal(b, pveRes)
		if err != nil {
			return err
		}
		*result.(*int) = pveRes.Data
		return nil
	default:
		pveRes := &pveResponse[any]{}
		err = json.Unmarshal(b, pveRes)
		if err != nil {
			return err
		}
		b, _ := json.Marshal(pveRes.Data)
		json.Unmarshal(b, result)
		return nil
	}
}

func addPayloadValue[T string | bool | int](p *url.Values, key string, value *T, defaultValue *T) {
	switch t := any(value).(type) {
	case *bool:
		if t == nil {
			if defaultValue == nil {
				return
			}
			v := helpers.BoolToInt(*any(defaultValue).(*bool))
			p.Set(key, fmt.Sprintf("%d", v))
			return
		}
		v := helpers.BoolToInt(*t)
		p.Set(key, fmt.Sprintf("%d", v))
		return
	case *int:
		if t == nil {
			if defaultValue == nil {
				return
			}
			p.Set(key, fmt.Sprintf("%d", *any(defaultValue).(*int)))
			return
		}
		p.Set(key, fmt.Sprintf("%d", *t))
		return
	case *string:
		if t == nil {
			if defaultValue == nil {
				return
			}
			p.Set(key, *any(defaultValue).(*string))
			return
		}
		p.Set(key, *t)
		return
	default:
		panic("'AddPayloadValue' value parameter type not supported")
	}
}
