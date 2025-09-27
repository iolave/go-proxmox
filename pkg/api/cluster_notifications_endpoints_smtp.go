package api

import (
	"net/http"
)

// ClusterNotificationsListSMTPEndpoints Returns a list of all SMTP endpoints
//
// Required permissions:
//
//	Check: ["perm","/mapping/notifications",["Mapping.Audit"]]
func (c API) ClusterNotificationsListSMTPEndpoints() (res []struct {
	FromAddress string   `json:"from-address"`
	Name        string   `json:"name"`
	Origin      string   `json:"origin"`
	Server      string   `json:"server"`
	Author      *string  `json:"author"`
	Comment     string   `json:"comment"`
	Disable     *int     `json:"disable"`
	MailTo      []string `json:"mailto"`
	MailToUser  []string `json:"mailto-user"`
	Mode        *string  `json:"mode"`
	Port        *int     `json:"port"`
	Username    *string  `json:"username"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/notifications/endpoints/smtp",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterNotificationsNewSMTPEndpointRequest struct {
	// `From` address for the mail
	FromAddress string `in:"query=from-address;omitempty"`

	// The name of the endpoint.
	Name string `in:"query=name;omitempty"`

	// The address of the SMTP server
	Server string `in:"query=server;omitempty"`

	// Author of the mail. Defaults to 'Proxmox VE'.
	Author string `in:"query=author;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// Disable this target.
	Disable *int `in:"query=disable;omitempty"`

	// List of email recipients.
	MailTo string `in:"query=mailto;omitempty"`

	// List of users.
	MailToUser string `in:"query=mailto-user;omitempty"`

	// Determine which encryption method shall be used for the connection.
	//
	//	insecure | starttls | tls
	Mode string `in:"query=mode;omitempty"`

	// Password for SMTP authentication
	Password string `in:"query=password;omitempty"`

	// The port to be used. Defaults to 465 for TLS based connections, 587 for STARTTLS based connections and port 25 for insecure plain-text connections.
	Port *int `in:"query=port;omitempty"`

	// Username for SMTP authentication
	Username string `in:"query=username;omitempty"`
}

// ClusterNotificationsNewSMTPEndpoint Create a new SMTP endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func (c API) ClusterNotificationsNewSMTPEndpoint(req ClusterNotificationsNewSMTPEndpointRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/smtp",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsGetSMTPEndpoint Return a specific SMTP endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (c API) ClusterNotificationsGetSMTPEndpoint(name string) (res struct {
	FromAddress string   `json:"from-address"`
	Name        string   `json:"name"`
	Server      string   `json:"server"`
	Author      *string  `json:"author"`
	Comment     string   `json:"comment"`
	Digest      *string  `json:"digest"`
	Disable     *int     `json:"disable"`
	MailTo      []string `json:"mailto"`
	MailToUser  []string `json:"mailto-user"`
	Mode        *string  `json:"mode"`
	Port        *int     `json:"port"`
	Username    *string  `json:"username"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/smtp/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterNotificationsUpdateSMTPEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"path=name;nonzero"`

	// Author of the mail. Defaults to 'Proxmox VE'.
	Author string `in:"query=author;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Disable this target.
	Disable *int `in:"query=disable;omitempty"`

	// `From` address for the mail
	FromAddress string `in:"query=from-address;omitempty"`

	// List of email recipients.
	MailTo string `in:"query=mailto;omitempty"`

	// List of users.
	MailToUser string `in:"query=mailto-user;omitempty"`

	// Determine which encryption method shall be used for the connection.
	//
	//	insecure | starttls | tls
	Mode string `in:"query=mode;omitempty"`

	// Password for SMTP authentication
	Password string `in:"query=password;omitempty"`

	// The port to be used. Defaults to 465 for TLS based connections, 587 for STARTTLS based connections and port 25 for insecure plain-text connections.
	Port *int `in:"query=port;omitempty"`

	// The address of the SMTP server
	Server string `in:"query=server;omitempty"`

	// Username for SMTP authentication
	Username string `in:"query=username;omitempty"`
}

// ClusterNotificationsUpdateSMTPEndpoint Update existing SMTP endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func (c API) ClusterNotificationsUpdateSMTPEndpoint(req ClusterNotificationsUpdateSMTPEndpointRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/smtp/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsDeleteSMTPEndpoint Remove SMTP endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (c API) ClusterNotificationsDeleteSMTPEndpoint(name string) (err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/smtp/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
