package api

import "net/http"

type ClusterACMEGetPluginsRequest struct {
	// Only list ACME plugins of a specific type
	//
	// 	dns | standalone
	Type string `in:"query=type;omitempty"`
}

type ClusterACMEGetPluginResponse struct {
	Plugin          string  `json:"plugin"`
	Type            *string `json:"type"`
	Digest          *string `json:"digest"`
	Disable         *int    `json:"disable"`
	Nodes           *string `json:"nodes"`
	Data            *string `json:"data"`
	ValidationDelay *int    `json:"validation-delay"`
	API             *string `json:"api"`
}

// ClusterACMEGetPlugins ACME plugin index.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterACMEGetPlugins() (res []ClusterACMEGetPluginResponse, err error) {
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/acme/plugins",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterACMEGetPluginRequest struct {
	// Unique identifier for ACME plugin instance.
	ID string `in:"path=id;nonzero"`
}

// ClusterACMEGetPlugin Get ACME plugin configuration.
//
// Parameters:
//
//   - id is the unique identifier for ACME plugin instance.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterACMEGetPlugin(id string) (res ClusterACMEGetPluginResponse, err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/acme/plugins/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterACMEPostPluginRequest struct {
	// ACME Plugin ID name
	ID string `in:"query=id;nonzero"`

	// ACME challenge type.
	//
	// 	dns | standalone
	Type string `in:"query=type;nonzero"`

	// API plugin name
	//
	// 	1984hosting | acmedns | acmeproxy | active24 | ad | ali | alviy | anx | artfiles | arvan | aurora | autodns | aws | azion | azure | beget | bookmyname | bunny | cf | clouddns | cloudns | cn | conoha | constellix | cpanel | curanet | cyon | da | ddnss | desec | df | dgon | dnsexit | dnshome | dnsimple | dnsservices | doapi | domeneshop | dp | dpi | dreamhost | duckdns | durabledns | dyn | dynu | dynv6 | easydns | edgecenter | edgedns | euserv | exoscale | fornex | freedns | freemyip | gandi_livedns | gcloud | gcore | gd | geoscaling | googledomains | he | he_ddns | hetzner | hexonet | hostingde | huaweicloud | infoblox | infomaniak | internetbs | inwx | ionos | ionos_cloud | ipv64 | ispconfig | jd | joker | kappernet | kas | kinghost | knot | la | leaseweb | lexicon | limacity | linode | linode_v4 | loopia | lua | maradns | me | miab | mijnhost | misaka | myapi | mydevil | mydnsjp | mythic_beasts | namecheap | namecom | namesilo | nanelo | nederhost | neodigit | netcup | netlify | nic | njalla | nm | nsd | nsone | nsupdate | nw | oci | omglol | one | online | openprovider | openstack | opnsense | ovh | pdns | pleskxml | pointhq | porkbun | rackcorp | rackspace | rage4 | rcode0 | regru | scaleway | schlundtech | selectel | selfhost | servercow | simply | technitium | tele3 | tencent | timeweb | transip | udr | ultra | unoeuro | variomedia | veesp | vercel | vscale | vultr | websupport | west_cn | world4you | yandex360 | yc | zilore | zone | zoneedit | zonomi
	API string `in:"query=api;omitempty"`

	// DNS plugin data. (base64 encoded)
	Data string `in:"query=data;omitempty"`

	// Flag to disable the config.
	//
	// 	1 | 0
	Disable *int `in:"query=disable;omitempty"`

	// List of cluster node names.
	//
	// 	node1,node2,node3,...
	Nodes string `in:"query=nodes;omitempty"`

	// Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	//
	// 	0 - 172800
	ValidationDelay *int `in:"query=validation-delay;omitempty"`
}

// ClusterACMEPostPlugin Add ACME plugin configuration.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterACMEPostPlugin(req ClusterACMEPostPluginRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/acme/plugins",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

// ClusterACMEDeletePlugin Delete ACME plugin configuration.
//
// Parameters:
//
//   - id is the unique identifier for ACME plugin instance.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterACMEDeletePlugin(id string) (err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/acme/plugins/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}

type ClusterACMEUpdatePluginRequest struct {
	// ACME Plugin ID name
	ID string `in:"path=id;nonzero"`

	// API plugin name
	//
	// 	1984hosting | acmedns | acmeproxy | active24 | ad | ali | alviy | anx | artfiles | arvan | aurora | autodns | aws | azion | azure | beget | bookmyname | bunny | cf | clouddns | cloudns | cn | conoha | constellix | cpanel | curanet | cyon | da | ddnss | desec | df | dgon | dnsexit | dnshome | dnsimple | dnsservices | doapi | domeneshop | dp | dpi | dreamhost | duckdns | durabledns | dyn | dynu | dynv6 | easydns | edgecenter | edgedns | euserv | exoscale | fornex | freedns | freemyip | gandi_livedns | gcloud | gcore | gd | geoscaling | googledomains | he | he_ddns | hetzner | hexonet | hostingde | huaweicloud | infoblox | infomaniak | internetbs | inwx | ionos | ionos_cloud | ipv64 | ispconfig | jd | joker | kappernet | kas | kinghost | knot | la | leaseweb | lexicon | limacity | linode | linode_v4 | loopia | lua | maradns | me | miab | mijnhost | misaka | myapi | mydevil | mydnsjp | mythic_beasts | namecheap | namecom | namesilo | nanelo | nederhost | neodigit | netcup | netlify | nic | njalla | nm | nsd | nsone | nsupdate | nw | oci | omglol | one | online | openprovider | openstack | opnsense | ovh | pdns | pleskxml | pointhq | porkbun | rackcorp | rackspace | rage4 | rcode0 | regru | scaleway | schlundtech | selectel | selfhost | servercow | simply | technitium | tele3 | tencent | timeweb | transip | udr | ultra | unoeuro | variomedia | veesp | vercel | vscale | vultr | websupport | west_cn | world4you | yandex360 | yc | zilore | zone | zoneedit | zonomi
	API string `in:"query=api;omitempty"`

	// DNS plugin data. (base64 encoded)
	Data string `in:"query=data;omitempty"`

	// Flag to disable the config.
	//
	// 	1 | 0
	Disable *int `in:"query=disable;omitempty"`

	// List of cluster node names.
	//
	// 	node1,node2,node3,...
	Nodes string `in:"query=nodes;omitempty"`

	// Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	//
	// 	0 - 172800
	ValidationDelay *int `in:"query=validation-delay;omitempty"`

	// Prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`
}

// ClusterACMEUpdatePlugin Update ACME plugin configuration.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (c API) ClusterACMEUpdatePlugin(req ClusterACMEUpdatePluginRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/acme/plugins/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}
