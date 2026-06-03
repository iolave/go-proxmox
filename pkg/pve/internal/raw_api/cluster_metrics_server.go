package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

// ClusterMetricsServerList List configured metric servers.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterMetricsServerList(c *pve.Client) (res []struct {
	Disable int    `json:"disable"`
	ID      string `json:"id"`
	Port    int    `json:"port"`
	Server  string `json:"server"`
	Type    string `json:"type"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/metrics/server",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterMetricsServerGet Read metric server configuration.
//
// Parameters:
//
//   - id is the metric server ID.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
//
// TODO: add response type, [proxmox docs] don't specify
//
// [proxmox docs] https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/metrics/server/{id}
func ClusterMetricsServerGet(c *pve.Client, id string) (res any, err error) {
	req := struct {
		ID string `in:"path=id"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/metrics/server/{id}",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

type ClusterMetricsServerPostRequest struct {
	// The ID of the entry.
	ID string `in:"path=id;nonzero"`

	// server network port
	Port int `in:"query=port;omitempty"`

	// server dns name or IP address
	Server string `in:"query=server;omitempty"`

	// Plugin type.
	//   graphite | influxdb | opentelemetry
	Type string `in:"query=type;omitempty"`

	// An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	APIPathPrefix string `in:"query=api-path-prefix;omitempty"`

	// The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Bucket string `in:"query=bucket;omitempty"`

	// Flag to disable the plugin.
	Disable *int `in:"query=disable;omitempty"`

	// The InfluxDB protocol.
	// 	udp | http | https
	InfluxDBProto string `in:"query=influxdbproto;omitempty"`

	// InfluxDB max-body-size in bytes. Requests are batched up to this size.
	MaxBodySize *int `in:"query=max-body-size;omitempty"`

	// MTU for metrics transmission over UDP
	MTU *int `in:"query=mtu;omitempty"`

	// The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Organization string `in:"query=organization;omitempty"`

	// Compression algorithm for requests
	// 	none | gzip
	OTELCompression string `in:"query=otel-compression;omitempty"`

	// Custom HTTP headers (JSON format, base64 encoded)
	OTELHeaders string `in:"query=otel-headers;omitempty"`

	// Maximum request body size in bytes
	OTELMaxBodySize *int `in:"query=otel-max-body-size;omitempty"`

	// OTLP endpoint path
	OTELPath string `in:"query=otel-path;omitempty"`

	// HTTP protocol
	// 	http | https
	OTELProtocol string `in:"query=otel-protocol;omitempty"`

	// Additional resource attributes as JSON, base64 encoded
	OTELResourceAttributes string `in:"query=otel-resource-attributes;omitempty"`

	// HTTP request timeout in seconds
	OTELTimeout *int `in:"query=otel-timeout;omitempty"`

	// Verify SSL certificates
	OTELVerifySSL *int `in:"query=otel-verify-ssl;omitempty"`

	// root graphite path (ex: proxmox.mycluster.mykey)
	Path string `in:"query=path;omitempty"`

	// Protocol to send graphite data. TCP or UDP (default)
	//   tcp | udp
	Proto string `in:"query=proto;omitempty"`

	// graphite TCP socket timeout (default=1)
	Timeout *int `in:"query=timeout;omitempty"`

	// The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
	Token string `in:"query=token;omitempty"`

	// Set to 0 to disable certificate verification for https endpoints.
	VerifyCertificate *int `in:"query=verify-certificate;omitempty"`
}

// ClusterMetricsServerPost Create a new external metric server config
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func ClusterMetricsServerPost(c *pve.Client, req ClusterMetricsServerPostRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/metrics/server/{id}",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}

type ClusterMetricsServerUpdateRequest struct {
	// The ID of the entry.
	ID string `in:"path=id;nonzero"`

	// server network port
	Port int `in:"query=port;omitempty"`

	// server dns name or IP address
	Server string `in:"query=server;omitempty"`

	// An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	APIPathPrefix string `in:"query=api-path-prefix;omitempty"`

	// The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Bucket string `in:"query=bucket;omitempty"`

	// A list of settings you want to delete.
	Delete string `in:"query=delete;omitempty"`

	// prevent changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.
	Digest string `in:"query=digest;omitempty"`

	// Flag to disable the plugin.
	Disable *int `in:"query=disable;omitempty"`

	// The InfluxDB protocol.
	// 	udp | http | https
	InfluxDBProto string `in:"query=influxdbproto;omitempty"`

	// InfluxDB max-body-size in bytes. Requests are batched up to this size.
	MaxBodySize *int `in:"query=max-body-size;omitempty"`

	// MTU for metrics transmission over UDP
	MTU *int `in:"query=mtu;omitempty"`

	// The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Organization string `in:"query=organization;omitempty"`

	// Compression algorithm for requests
	// 	none | gzip
	OTELCompression string `in:"query=otel-compression;omitempty"`

	// Custom HTTP headers (JSON format, base64 encoded)
	OTELHeaders string `in:"query=otel-headers;omitempty"`

	// Maximum request body size in bytes
	OTELMaxBodySize *int `in:"query=otel-max-body-size;omitempty"`

	// OTLP endpoint path
	OTELPath string `in:"query=otel-path;omitempty"`

	// HTTP protocol
	// 	http | https
	OTELProtocol string `in:"query=otel-protocol;omitempty"`

	// Additional resource attributes as JSON, base64 encoded
	OTELResourceAttributes string `in:"query=otel-resource-attributes;omitempty"`

	// HTTP request timeout in seconds
	OTELTimeout *int `in:"query=otel-timeout;omitempty"`

	// Verify SSL certificates
	OTELVerifySSL *int `in:"query=otel-verify-ssl;omitempty"`

	// root graphite path (ex: proxmox.mycluster.mykey)
	Path string `in:"query=path;omitempty"`

	// Protocol to send graphite data. TCP or UDP (default)
	//   tcp | udp
	Proto string `in:"query=proto;omitempty"`

	// graphite TCP socket timeout (default=1)
	Timeout *int `in:"query=timeout;omitempty"`

	// The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
	Token string `in:"query=token;omitempty"`

	// Set to 0 to disable certificate verification for https endpoints.
	VerifyCertificate *int `in:"query=verify-certificate;omitempty"`
}

// ClusterMetricsServerUpdate Update metric server configuration.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func ClusterMetricsServerUpdate(c *pve.Client, req ClusterMetricsServerUpdateRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/metrics/server/{id}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterMetricsServerDelete Remove Metric server.
//
// Parameters:
//
//   - id is the metric server ID.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func ClusterMetricsServerDelete(c *pve.Client, id string) (err error) {
	req := struct {
		ID string `in:"path=id;nonzero"`
	}{
		ID: id,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/metrics/server/{id}",
		Method:  http.MethodDelete,
		Payload: &req,
	})

	return err
}
