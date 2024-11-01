# proxmoxapi

```go
import "github.com/iolave/go-proxmox/pkg/proxmox_api"
```





## Constants

<a name="CREDENTIALS_NOT_DETECTED_ERROR"></a>Credentials error messages.

```go
const (
    CREDENTIALS_NOT_DETECTED_ERROR  = "credentials could not be detected from env"
    CREDENTIALS_NOT_SUPPORTED_ERROR = "credentials type not supported yet"
)
```

<a name="Credentials"></a>
## type [Credentials](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/api_credentials.go#L22-L27>)

Credentials is the struct that stores proxmox api credentials.

```go
type Credentials struct {
    // contains filtered or unexported fields
}
```

<a name="NewTokenCredentials"></a>
### func [NewTokenCredentials](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/api_credentials.go#L32>)

```go
func NewTokenCredentials(user, tokenName, token string) *Credentials
```

NewTokenCredentials returns a struct containing proxmox token credentials that can be passed to the [NewWithCredentials](<https://go-proxmox.iolave.com/reference/pkg/proxmox_api/#func-newwithcredentials>) method.

<a name="FirewallLogEntry"></a>
## type [FirewallLogEntry](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/firewall.go#L19-L22>)

Proxmox firewall log entry.

```go
type FirewallLogEntry struct {
    Id   int    `json:"n"`
    Text string `json:"t"`
}
```

<a name="FirewallLogLevel"></a>
## type [FirewallLogLevel](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/firewall.go#L4>)

Promox firewall availabe log levels.

```go
type FirewallLogLevel string
```

<a name="FIREWALL_LOG_LEVEL_EMERG"></a>

```go
const (
    FIREWALL_LOG_LEVEL_EMERG  FirewallLogLevel = "emerg"
    FIREWALL_LOG_LEVEL_ALERT  FirewallLogLevel = "alert"
    FIREWALL_LOG_LEVEL_CRIT   FirewallLogLevel = "crit"
    FIREWALL_LOG_LEVEL_ERR    FirewallLogLevel = "err"
    FIREWALL_LOG_LEVEL_WARN   FirewallLogLevel = "warning"
    FIREWALL_LOG_LEVEL_NOTICE FirewallLogLevel = "notice"
    FIREWALL_LOG_LEVEL_INFO   FirewallLogLevel = "info"
    FIREWALL_LOG_LEVEL_DEBUG  FirewallLogLevel = "debug"
    FIREWALL_LOG_LEVEL_NOLOG  FirewallLogLevel = "nolog"
)
```

<a name="GetAliasResponse"></a>
## type [GetAliasResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L9-L14>)



```go
type GetAliasResponse struct {
    CIDR    string `json:"cidr"`
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetIPSetResponse"></a>
## type [GetIPSetResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L80-L84>)



```go
type GetIPSetResponse struct {
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetNodeDatastoreContentResponse"></a>
## type [GetNodeDatastoreContentResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L31-L42>)

TODO: Add missing verification property from [docs](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content>).

```go
type GetNodeDatastoreContentResponse struct {
    Format    string  `json:"format"`    // Format identifier ('raw', 'qcow2', 'subvol', 'iso', 'tgz' ...)
    Size      int     `json:"size"`      // Volume size in bytes.
    VolumeID  string  `json:"volid"`     // Volume identifier.
    CreatedAt *int    `json:"ctime"`     // Creation time (seconds since the UNIX Epoch).
    Encrypted *string `json:"encrypted"` // If whole backup is encrypted, value is the fingerprint or '1'  if encrypted. Only useful for the Proxmox Backup Server storage type.
    Notes     *string `json:"notes"`     // Optional notes. If they contain multiple lines, only the first one is returned here.
    Parent    *string `json:"parent"`    // Volume identifier of parent (for linked cloned).
    Protected *bool   `json:"protected"` // Protection status. Currently only supported for backups.
    Used      *int    `json:"used"`      // Used space. Please note that most storage plugins do not report anything useful here.
    VmID      *int    `json:"vmid"`      // Associated Owner VMID.
}
```

<a name="GetNodeFirewallRulesResponse"></a>
## type [GetNodeFirewallRulesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_firewall.go#L9-L25>)



```go
type GetNodeFirewallRulesResponse[Position interface{ int | string }] struct {
    Action          string           `json:"action"`
    Comment         string           `json:"comment"`
    Destination     string           `json:"dest"`
    DestinationPort string           `json:"dport"`
    Enable          int              `json:"enable"`
    ICMPType        string           `json:"icmp-type"`
    Interface       string           `json:"iface"`
    IPVersion       int              `json:"ipversion"`
    LogLevel        FirewallLogLevel `json:"log"`
    Macro           string           `json:"macro"`
    Pos             Position         `json:"pos"`
    Proto           string           `json:"proto"`
    Source          string           `json:"source"`
    Sport           string           `json:"sport"`
    Type            string           `json:"type"`
}
```

<a name="GetNodeLxcsResponse"></a>
## type [GetNodeLxcsResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_lxc.go#L8-L19>)



```go
type GetNodeLxcsResponse struct {
    Status  LxcStatus `json:"status"`
    VmID    int       `json:"vmid"`
    Cpus    *int      `json:"cpus"`
    Lock    *string   `json:"lock"`
    MaxDisk *int      `json:"maxdisk"`
    MaxMem  *int      `json:"maxmem"`
    MaxSwap *int      `json:"maxswap"`
    Name    *string   `json:"name"`
    Tags    *string   `json:"tags"`
    Uptime  *int      `json:"uptime"`
}
```

<a name="GetNodeStoragesResponse"></a>
## type [GetNodeStoragesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L9-L20>)



```go
type GetNodeStoragesResponse struct {
    Content      string   `json:"content"`       // Allowed storage content types.
    Storage      string   `json:"storage"`       // The storage identifier.
    Type         string   `json:"type"`          // Storage type.
    Active       *bool    `json:"active"`        // Set when storage is accessible.
    Available    *int     `json:"avail"`         // Available storage space in bytes.
    Enabled      *bool    `json:"enabled"`       // Set when storage is enabled (not disabled).
    Shared       *bool    `json:"shared"`        // Shared flag from storage configuration.
    TotalSpace   *int     `json:"total"`         // Total storage space in bytes.
    UsedSpace    *int     `json:"used"`          // Total storage space in bytes.
    UsedFraction *float64 `json:"used_fraction"` // Used fraction (used/total).
}
```

<a name="GetNodesResponse"></a>
## type [GetNodesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/nodes.go#L14-L24>)



```go
type GetNodesResponse struct {
    Node           string     `json:"node"`
    Status         NodeStatus `json:"status"`
    CPU            float64    `json:"cpu"`
    Level          string     `json:"level"`
    MaxCpu         int        `json:"maxcpu"`
    MaxMem         int        `json:"maxmem"`
    Mem            int        `json:"mem"`
    SSLFingerprint string     `json:"ssl_fingerprint"`
    Uptime         int        `json:"uptime"`
}
```

<a name="GetRulesResponse"></a>
## type [GetRulesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L91-L93>)



```go
type GetRulesResponse struct {
    Pos int `json:"pos"`
}
```

<a name="GetVersionResponse"></a>
## type [GetVersionResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/version.go#L5-L9>)



```go
type GetVersionResponse struct {
    Release string `json:"release"`
    Version string `json:"version"`
    RepoID  string `json:"repoid"`
}
```

<a name="LxcStatus"></a>
## type [LxcStatus](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/lxc.go#L3>)



```go
type LxcStatus string
```

<a name="LXC_STATUS_STOPPED"></a>

```go
const (
    LXC_STATUS_STOPPED LxcStatus = "stopped"
    LXC_STATUS_RUNNING LxcStatus = "running"
)
```

<a name="NodeStatus"></a>
## type [NodeStatus](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/nodes.go#L6>)

Proxmox availabe node statuses

```go
type NodeStatus string
```

<a name="NODE_STATUS_ONLINE"></a>

```go
const (
    NODE_STATUS_ONLINE  NodeStatus = "online"
    NODE_STATUS_OFFLINE NodeStatus = "offline"
    NODE_STATUS_UNKNOWN NodeStatus = "unknown"
)
```

<a name="ProxmoxAPI"></a>
## type [ProxmoxAPI](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/proxmox_api.go#L18-L22>)



```go
type ProxmoxAPI struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/proxmox_api.go#L24>)

```go
func New(config ProxmoxAPIConfig) (*ProxmoxAPI, error)
```



<a name="NewWithCredentials"></a>
### func [NewWithCredentials](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/proxmox_api.go#L48>)

```go
func NewWithCredentials(config ProxmoxAPIConfig, creds *Credentials) (*ProxmoxAPI, error)
```

TODO: To test credentials, do a proxmox version query to ensure credentials are valid

<a name="ProxmoxAPI.CreateClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) [CreateClusterFirewallAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L28>)

```go
func (api *ProxmoxAPI) CreateClusterFirewallAlias(name, cidr string, comment *string) error
```

CreateClusterFirewallAlias creates a cluster firewall IP or Network Alias.

<a name="ProxmoxAPI.DeleteClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) [DeleteClusterFirewallAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L68>)

```go
func (api *ProxmoxAPI) DeleteClusterFirewallAlias(name string, digest *string) error
```

DeleteClusterFirewallAlias removes a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

<a name="ProxmoxAPI.DownloadISOToNodeDatastore"></a>
### func \(\*ProxmoxAPI\) [DownloadISOToNodeDatastore](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L59>)

```go
func (api *ProxmoxAPI) DownloadISOToNodeDatastore(node, storageId, fileName, URL string) error
```

DownloadISOToNodeDatastore downloads an iso from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="ProxmoxAPI.DownloadVZTemplateToNodeDatastore"></a>
### func \(\*ProxmoxAPI\) [DownloadVZTemplateToNodeDatastore](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L75>)

```go
func (api *ProxmoxAPI) DownloadVZTemplateToNodeDatastore(node, storageId, fileName, URL string) error
```

DownloadVZTemplateToNodeDatastore downloads a vztemplate from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="ProxmoxAPI.GetClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) [GetClusterFirewallAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L22>)

```go
func (api *ProxmoxAPI) GetClusterFirewallAlias(name string) (GetAliasResponse, error)
```

GetClusterFirewallAlias retrieves cluster firewall alias by it's name.

<a name="ProxmoxAPI.GetClusterFirewallAliases"></a>
### func \(\*ProxmoxAPI\) [GetClusterFirewallAliases](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L17>)

```go
func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasResponse, error)
```

GetClusterFirewallAliases retrieves all cluster firewall aliases.

<a name="ProxmoxAPI.GetClusterFirewallIPSet"></a>
### func \(\*ProxmoxAPI\) [GetClusterFirewallIPSet](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L87>)

```go
func (api *ProxmoxAPI) GetClusterFirewallIPSet() ([]GetIPSetResponse, error)
```

GetClusterFirewallIPSet retrieves all cluster firewall IPSets.

<a name="ProxmoxAPI.GetClusterFirewallRules"></a>
### func \(\*ProxmoxAPI\) [GetClusterFirewallRules](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L96>)

```go
func (api *ProxmoxAPI) GetClusterFirewallRules() ([]GetRulesResponse, error)
```

GetClusterFirewallRules retrieves all cluster firewall rules.

<a name="ProxmoxAPI.GetNodeDatastoreContent"></a>
### func \(\*ProxmoxAPI\) [GetNodeDatastoreContent](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L49>)

```go
func (api *ProxmoxAPI) GetNodeDatastoreContent(node, storageId string) ([]GetNodeDatastoreContentResponse, error)
```

GetNodeDatastoreContent retrieves node's datastores info.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content>).

<a name="ProxmoxAPI.GetNodeDatastores"></a>
### func \(\*ProxmoxAPI\) [GetNodeDatastores](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_storage.go#L23>)

```go
func (api *ProxmoxAPI) GetNodeDatastores(node string) ([]GetNodeStoragesResponse, error)
```

GetNodeDatastores retrieves node's datastores info.

<a name="ProxmoxAPI.GetNodeFirewallRules"></a>
### func \(\*ProxmoxAPI\) [GetNodeFirewallRules](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_firewall.go#L28>)

```go
func (api *ProxmoxAPI) GetNodeFirewallRules(node string) ([]GetNodeFirewallRulesResponse[int], error)
```

GetNodeFirewallRules retrieves node's firewall rules.

<a name="ProxmoxAPI.GetNodeFirewallRulesByPos"></a>
### func \(\*ProxmoxAPI\) [GetNodeFirewallRulesByPos](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_firewall.go#L33>)

```go
func (api *ProxmoxAPI) GetNodeFirewallRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error)
```

GetNodeFirewallRulesByPos Retrieves a single node's firewall rule using rule's position \(pos\) as an index.

<a name="ProxmoxAPI.GetNodeLxcs"></a>
### func \(\*ProxmoxAPI\) [GetNodeLxcs](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_lxc.go#L22>)

```go
func (api *ProxmoxAPI) GetNodeLxcs(node string) ([]GetNodeLxcsResponse, error)
```

GetNodeLxcs returns node's lxc index per node.

<a name="ProxmoxAPI.GetNodes"></a>
### func \(\*ProxmoxAPI\) [GetNodes](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/nodes.go#L27>)

```go
func (api *ProxmoxAPI) GetNodes() ([]GetNodesResponse, error)
```

GetNodes retrieves nodes.

<a name="ProxmoxAPI.GetVersion"></a>
### func \(\*ProxmoxAPI\) [GetVersion](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/version.go#L12>)

```go
func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error)
```

GetVersion retrieves proxmox version.

<a name="ProxmoxAPI.ReadNodeFirewallLog"></a>
### func \(\*ProxmoxAPI\) [ReadNodeFirewallLog](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/node_firewall.go#L42>)

```go
func (api *ProxmoxAPI) ReadNodeFirewallLog(node string) ([]FirewallLogEntry, error)
```

ReadNodeFirewallLog Retrieves node's firewall log entries.

TODO: Add missing limit, since, start, until parameters shown in [docs](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log>).

<a name="ProxmoxAPI.UpdateClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) [UpdateClusterFirewallAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/cluster_firewall.go#L44>)

```go
func (api *ProxmoxAPI) UpdateClusterFirewallAlias(name, cidr string, comment *string, digest *string, rename *string) error
```

UpdateClusterFirewallAlias updates a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

<a name="ProxmoxAPIConfig"></a>
## type [ProxmoxAPIConfig](<https://github.com/iolave/go-proxmox/blob/master/pkg/proxmox_api/proxmox_api.go#L11-L16>)



```go
type ProxmoxAPIConfig struct {
    Host               string
    Port               int
    InsecureSkipVerify bool
    CfServiceToken     *cloudflare.CloudflareServiceToken
}
```