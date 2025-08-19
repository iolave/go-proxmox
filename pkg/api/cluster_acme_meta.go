package api

import "net/http"

type ClusterGetACMEMetaRequest struct {
	// URL of ACME CA directory endpoint. Defaults to https://acme-v02.api.letsencrypt.org/directory
	//
	// ^https?://.*
	Directory string `in:"query=directory;omitempty"`
}

// ClusterACMEGetMeta Retrieve ACME Directory Meta Information
//
// Required permissions:
//
//	Check: ["perm","/nodes/{node}",["Sys.Audit"]]
func (c API) ClusterACMEGetMeta(req ClusterGetACMEMetaRequest) (res struct {
	CAAIdentities           []string          `json:"caaIdentities"`
	ExternalAccountRequired *bool             `json:"externalAccountRequired"`
	TermsOfService          *string           `json:"termsOfService"`
	Website                 *string           `json:"website"`
	Profiles                map[string]string `json:"profiles"`
}, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/acme/meta",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}
