package api

import "net/http"

// ClusterNotificationsGetMacherFieldValues Returns known notification metadata fields and their known values
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (s API) ClusterNotificationsGetMacherFieldValues() (res []struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/notifications/matcher-field-values",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterNotificationsGetMacherFields Returns known notification metadata fields
//
// Required permissions:
//
//	Check: ["or",["perm","/mapping/notifications",["Mapping.Modify"]],["perm","/mapping/notifications",["Mapping.Audit"]]]
func (s API) ClusterNotificationsGetMacherFields() (res []struct {
	Name string `json:"name"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/notifications/matcher-fields",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}
