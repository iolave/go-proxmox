# proxmoxapi

```go
import "github.com/iolave/go-proxmox/pkg/proxmox_api"
```





## Constants

<a name="CREDENTIALS_NOT_DETECTED_ERROR"></a>

```go
const (
    CREDENTIALS_NOT_DETECTED_ERROR  = "credentials could not be detected from env"
    CREDENTIALS_NOT_SUPPORTED_ERROR = "credentials type not supported yet"
)
```

<a name="FirewallLogEntry"></a>
## type FirewallLogEntry



```go
type FirewallLogEntry struct {
    Id   int    `json:"n"`
    Text string `json:"t"`
}
```

<a name="FirewallLogLevel"></a>
## type FirewallLogLevel



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
## type GetAliasResponse



```go
type GetAliasResponse struct {
    CIDR    string `json:"cidr"`
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetIPSetResponse"></a>
## type GetIPSetResponse



```go
type GetIPSetResponse struct {
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetNodeFirewallRulesResponse"></a>
## type GetNodeFirewallRulesResponse



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

<a name="GetNodesResponse"></a>
## type GetNodesResponse



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
## type GetRulesResponse



```go
type GetRulesResponse struct {
    Pos int `json:"pos"`
}
```

<a name="GetVersionResponse"></a>
## type GetVersionResponse



```go
type GetVersionResponse struct {
    Release string `json:"release"`
    Version string `json:"version"`
    RepoID  string `json:"repoid"`
}
```

<a name="NodeStatus"></a>
## type NodeStatus



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
## type ProxmoxAPI



```go
type ProxmoxAPI struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func New

```go
func New(config ProxmoxAPIConfig) (*ProxmoxAPI, error)
```



<a name="NewWithCredentials"></a>
### func NewWithCredentials

```go
func NewWithCredentials(config ProxmoxAPIConfig, creds *credentials) (*ProxmoxAPI, error)
```

TODO: To test credentials, do a proxmox version query to ensure credentials are valid

<a name="ProxmoxAPI.CreateClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) CreateClusterFirewallAlias

```go
func (api *ProxmoxAPI) CreateClusterFirewallAlias(name, cidr string, comment *string) error
```



<a name="ProxmoxAPI.DeleteClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) DeleteClusterFirewallAlias

```go
func (api *ProxmoxAPI) DeleteClusterFirewallAlias(name string, digest *string) error
```



<a name="ProxmoxAPI.GetClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallAlias

```go
func (api *ProxmoxAPI) GetClusterFirewallAlias(name string) (GetAliasResponse, error)
```



<a name="ProxmoxAPI.GetClusterFirewallAliases"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallAliases

```go
func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasResponse, error)
```



<a name="ProxmoxAPI.GetClusterFirewallIPSet"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallIPSet

```go
func (api *ProxmoxAPI) GetClusterFirewallIPSet() ([]GetIPSetResponse, error)
```



<a name="ProxmoxAPI.GetClusterFirewallRules"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallRules

```go
func (api *ProxmoxAPI) GetClusterFirewallRules() ([]GetRulesResponse, error)
```



<a name="ProxmoxAPI.GetNodeFirewallRules"></a>
### func \(\*ProxmoxAPI\) GetNodeFirewallRules

```go
func (api *ProxmoxAPI) GetNodeFirewallRules(node string) ([]GetNodeFirewallRulesResponse[int], error)
```



<a name="ProxmoxAPI.GetNodeFirewallRulesByPos"></a>
### func \(\*ProxmoxAPI\) GetNodeFirewallRulesByPos

```go
func (api *ProxmoxAPI) GetNodeFirewallRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error)
```



<a name="ProxmoxAPI.GetNodes"></a>
### func \(\*ProxmoxAPI\) GetNodes

```go
func (api *ProxmoxAPI) GetNodes() ([]GetNodesResponse, error)
```



<a name="ProxmoxAPI.GetVersion"></a>
### func \(\*ProxmoxAPI\) GetVersion

```go
func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error)
```



<a name="ProxmoxAPI.ReadNodeFirewallLog"></a>
### func \(\*ProxmoxAPI\) ReadNodeFirewallLog

```go
func (api *ProxmoxAPI) ReadNodeFirewallLog(node string) ([]FirewallLogEntry, error)
```



<a name="ProxmoxAPIConfig"></a>
## type ProxmoxAPIConfig



```go
type ProxmoxAPIConfig struct {
    Host               string
    Port               int
    InsecureSkipVerify bool
    CfServiceToken     *cloudflare.CloudflareServiceToken
}
```