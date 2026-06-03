package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterNotificationsListGotifyEndpoints Returns a list of all gotify endpoints
//
// Required permissions:
//
//	Check: ["perm","/mapping/notifications",["Mapping.Audit"]]
func ClusterNotificationsListGotifyEndpoints(c *pve.Client) (res []struct {
	Name    string `json:"name"`
	Origin  string `json:"origin"`
	Server  string `json:"server"`
	Comment string `json:"comment"`
	Disable *int   `json:"disable"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/notifications/endpoints/gotify",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterNotificationsNewGotifyEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"query=name;omitempty"`

	// Server URL
	Server string `in:"query=server;omitempty"`

	// Secret token
	Token string `in:"query=token;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// Disable this target
	Disable *int `in:"query=disable;omitempty"`
}

// ClusterNotificationsNewGotifyEndpoint Create a new gotify endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func ClusterNotificationsNewGotifyEndpoint(c *pve.Client, req ClusterNotificationsNewGotifyEndpointRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/gotify",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsGetGotifyEndpoint Return a specific gotify endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsGetGotifyEndpoint(c *pve.Client, name string) (res struct {
	Name    string  `json:"name"`
	Server  string  `json:"server"`
	Comment string  `json:"comment"`
	Digest  *string `json:"digest"`
	Disable *int    `json:"disable"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/gotify/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterNotificationsUpdateGotifyEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"path=name;nonzero"`

	// Comment.
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Disable this target
	Disable *int `in:"query=disable;omitempty"`

	// Server URL
	Server string `in:"query=server;omitempty"`

	// Secret token
	Token string `in:"query=token;omitempty"`
}

// ClusterNotificationsUpdateGotifyEndpoint Update existing gotify endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func ClusterNotificationsUpdateGotifyEndpoint(c *pve.Client, req ClusterNotificationsUpdateGotifyEndpointRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/gotify/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsDeleteGotifyEndpoint Remove gotify endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func ClusterNotificationsDeleteGotifyEndpoint(c *pve.Client, name string) (err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/notifications/endpoints/gotify/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
