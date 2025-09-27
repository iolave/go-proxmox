package api

import (
	"net/http"
)

// ClusterNotificationsListSendmailEndpoints Returns a list of all sendmail endpoints
//
// Required permissions:
//
//	Check: ["perm","/mapping/notifications",["Mapping.Audit"]]
func (c API) ClusterNotificationsListSendmailEndpoints() (res []struct {
	Name        string   `json:"name"`
	Origin      string   `json:"origin"`
	Author      *string  `json:"author"`
	Comment     string   `json:"comment"`
	Disable     *int     `json:"disable"`
	FromAddress *string  `json:"from-address"`
	MailTo      []string `json:"mailto"`
	MailToUser  []string `json:"mailto-user"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/notifications/endpoints/sendmail",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterNotificationsNewSendmailEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"query=name;omitempty"`

	// Author of the mail
	Author string `in:"query=author;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// Disable this target
	Disable *int `in:"query=disable;omitempty"`

	// `From` address for the mail
	FromAddress string `in:"query=from-address;omitempty"`

	// List of email recipients
	MailTo string `in:"query=mailto;omitempty"`

	// List of users
	MailToUser string `in:"query=mailto-user;omitempty"`
}

// ClusterNotificationsNewSendmailEndpoint Create a new sendmail endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func (c API) ClusterNotificationsNewSendmailEndpoint(req ClusterNotificationsNewSendmailEndpointRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/sendmail",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsGetSendmailEndpoint Return a specific sendmail endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (c API) ClusterNotificationsGetSendmailEndpoint(name string) (res struct {
	Name        string   `json:"name"`
	Author      *string  `json:"author"`
	Comment     string   `json:"comment"`
	Digest      *string  `json:"digest"`
	Disable     *int     `json:"disable"`
	FromAddress *string  `json:"from-address"`
	MailTo      []string `json:"mailto"`
	MailToUser  []string `json:"mailto-user"`
}, err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/sendmail/{name}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterNotificationsUpdateSendmailEndpointRequest struct {
	// The name of the endpoint.
	Name string `in:"path=name;nonzero"`

	// Author of the mail
	Author string `in:"query=author;omitempty"`

	// Comment
	Comment string `in:"query=comment;omitempty"`

	// A list of settings you want to delete
	Delete string `in:"query=delete;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications
	Digest string `in:"query=digest;omitempty"`

	// Disable this target
	Disable *int `in:"query=disable;omitempty"`

	// `From` address for the mail
	FromAddress string `in:"query=from-address;omitempty"`

	// List of email recipients
	MailTo string `in:"query=mailto;omitempty"`

	// List of users
	MailToUser string `in:"query=mailto-user;omitempty"`
}

// ClusterNotificationsUpdateSendmailEndpoint Update existing sendmail endpoint
//
// Required permissions:
//
//	Check: ["and",["perm","/mapping/notifications",["Mapping.Modify"]],["or",["perm","/",["Sys.Audit","Sys.Modify"]],["perm","/",["Sys.AccessNetwork"]]]]
func (c API) ClusterNotificationsUpdateSendmailEndpoint(req ClusterNotificationsUpdateSendmailEndpointRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/sendmail/{name}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterNotificationsDeleteSendmailEndpoint Remove sendmail endpoint
//
// Parameters:
//
//   - name is the name of the endpoint
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (c API) ClusterNotificationsDeleteSendmailEndpoint(name string) (err error) {
	req := struct {
		Name string `in:"path=name;nonzero"`
	}{
		Name: name,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/notifications/endpoints/sendmail/{name}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
