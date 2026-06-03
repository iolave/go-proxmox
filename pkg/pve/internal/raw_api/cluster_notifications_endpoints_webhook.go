package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterNotificationsListWebhookEndpoints Returns a list of all Webhook endpoints
//
// Required permissions:
//
//	Check: ["perm","/mapping/notifications",["Mapping.Audit"]]
func ClusterNotificationsListWebhookEndpoints(c *pve.Client) (res []struct {
	Method  string   `json:"method"`
	Name    string   `json:"name"`
	Origin  string   `json:"origin"`
	URL     string   `json:"url"`
	Body    *string  `json:"body"`
	Comment string   `json:"comment"`
	Disable *int     `json:"disable"`
	Header  []string `json:"header"`
	Secret  []string `json:"secret"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/notifications/endpoints/webhook",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterNotificationsNewWebhookEndpointRequest struct {
	// HTTP method
	//	GET | POST | PUT
	Method string `in:"query=method;omitempty"`

	// The name of the endpoint.
	Name string `in:"query=name;omitempty"`

	// Server URL
	URL string `in:"query=url;omitempty"`

	// HTTP body, base64 encoded
	Body string `in:"query=body;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// Disable this target.
	Disable *int `in:"query=disable;omitempty"`

	// HTTP headers to set. These have to be formatted as a property string in the format name=<name>,value=<base64 of value>
	Header string `in:"query=header;omitempty"`

	// Secrets to set. These have to be formatted as a property string in the format name=<name>,value=<base64 of value>
	Secret string `in:"query=secret;omitempty"`
}

// ClusterNotificationsNewWebhookEndpoint Create a new Webhook endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func ClusterNotificationsNewWebhookEndpoint(c *pve.Client, req ClusterNotificationsNewWebhookEndpointRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/webhook",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsGetWebhookEndpoint Return a specific Webhook endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsGetWebhookEndpoint(c *pve.Client, name string) (res struct {
	Method  string   `json:"method"`
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Body    *string  `json:"body"`
	Comment string   `json:"comment"`
	Digest  *string  `json:"digest"`
	Disable *int     `json:"disable"`
	Header  []string `json:"header"`
	Secret  []string `json:"secret"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/webhook/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterNotificationsUpdateWebhookEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"path=name;nonzero"`

	// HTTP body, base64 encoded
	Body string `in:"query=body;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Disable this target.
	Disable *int `in:"query=disable;omitempty"`

	// HTTP headers to set. These have to be formatted as a property string in the format name=<name>,value=<base64 of value>
	Header string `in:"query=header;omitempty"`

	// HTTP method
	//	GET | POST | PUT
	Method string `in:"query=method;omitempty"`

	// Secrets to set. These have to be formatted as a property string in the format name=<name>,value=<base64 of value>
	Secret string `in:"query=secret;omitempty"`

	// Server URL
	URL string `in:"query=url;omitempty"`
}

// ClusterNotificationsUpdateWebhookEndpoint Update existing Webhook endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func ClusterNotificationsUpdateWebhookEndpoint(c *pve.Client, req ClusterNotificationsUpdateWebhookEndpointRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/webhook/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsDeleteWebhookEndpoint Remove Webhook endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsDeleteWebhookEndpoint(c *pve.Client, name string) (err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/webhook/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
