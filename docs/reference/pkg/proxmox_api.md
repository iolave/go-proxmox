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

<a name="CreateLxcRequest"></a>
## type CreateLxcRequest



```go
type CreateLxcRequest struct {
    Node               string          // The cluster node name.
    OSTemplate         string          // The OS template or backup file.
    VMID               *int            // The (unique) ID of the VM.
    Arch               *LxcArch        // OS architecture type.
    BWLimit            *int            // Override I/O bandwidth limit (in KiB/s).
    CMode              *LxcConsoleMode // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
    Console            *bool           // Attach a console device (/dev/console) to the container.
    Cores              *int            // The number of cores assigned to the container. A container can use all available cores by default.
    CPULimit           *int            // Limit of CPU usage. NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
    CPUUnits           *int            // CPU weight for a container. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this container gets. Number is relative to the weights of all the other running guests.
    Debug              *bool           // Try to be more verbose. For now this only enables debug log-level on start.
    Desc               *string         // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
    Features           *string         // Allow containers access to advanced features.
    Force              *bool           // Allow to overwrite existing container.
    Hookscript         *string         // Script that will be exectued during various steps in the containers lifetime.
    Hostname           *string         // Set a host name for the container.
    IgnoreUnpackErrors *bool           // Ignore errors when extracting the template.
    Lock               *LxcLock        // Lock/unlock the container.
    Memory             *int            // Amount of RAM for the container in MB.
    Nameserver         *string         // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
    Net                *[]LxcNet       // Specifies network interfaces for the container.
    OnBoot             *bool           // Specifies whether a container will be started during system bootup.
    OSType             *string         // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup. debian | devuan | ubuntu | centos | fedora | opensuse | archlinux | alpine | gentoo | nixos | unmanaged
    Password           *string         // Sets root password inside container.
    Pool               *string         // Add the VM to the specified pool.
    Protection         *bool           // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
    Restore            *bool           // Mark this as restore task.
    RootFS             *string         // make this a struct Use volume as container root.
    Searchdomain       *string         // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
    SSHPublicKeys      *string         // Setup public SSH keys (one key per line, OpenSSH format).
    Start              *bool           // Start the CT after its creation finished successfully.
    Startup            *string         // make this a struct Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
    Storage            *string         // Default Storage.
    Swap               *int            // Amount of SWAP for the container in MB.
    Tags               *string         // Tags of the Container. This is only meta information.
    Template           *bool           // Enable/disable Template.
    Timezone           *string         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
    TTY                *int            // Specify the number of tty available to the container.
    Unique             *bool           // Assign a unique random ethernet address.
    Unprivileged       *bool           // Makes the container run as unprivileged user. (Should not be modified manually.)

}
```

<a name="Credentials"></a>
## type Credentials

Credentials is the struct that stores proxmox api credentials.

```go
type Credentials struct {
    // contains filtered or unexported fields
}
```

<a name="NewTokenCredentials"></a>
### func NewTokenCredentials

```go
func NewTokenCredentials(user, tokenName, token string) *Credentials
```

NewTokenCredentials returns a struct containing proxmox token credentials that can be passed to the [NewWithCredentials](<https://go-proxmox.iolave.com/reference/pkg/proxmox_api/#func-newwithcredentials>) method.

<a name="FirewallLogEntry"></a>
## type FirewallLogEntry

Proxmox firewall log entry.

```go
type FirewallLogEntry struct {
    Id   int    `json:"n"`
    Text string `json:"t"`
}
```

<a name="FirewallLogLevel"></a>
## type FirewallLogLevel

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

<a name="GetNodeDatastoreContentResponse"></a>
## type GetNodeDatastoreContentResponse

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

<a name="GetNodeLxcsResponse"></a>
## type GetNodeLxcsResponse



```go
type GetNodeLxcsResponse struct {
    Status  LxcStatus `json:"status"`
    VMID    int       `json:"vmid"`
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
## type GetNodeStoragesResponse



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

<a name="LxcArch"></a>
## type LxcArch



```go
type LxcArch string
```

<a name="LXC_ARCH_AMD64"></a>

```go
const (
    LXC_ARCH_AMD64   LxcArch = "amd64"
    LXC_ARCH_I386    LxcArch = "i386"
    LXC_ARCH_ARM64   LxcArch = "arm64"
    LXC_ARCH_ARMHF   LxcArch = "armhf"
    LXC_ARCH_RISCV32 LxcArch = "riscv32"
    LXC_ARCH_RISCV64 LxcArch = "riscv64"
)
```

<a name="LxcConsoleMode"></a>
## type LxcConsoleMode



```go
type LxcConsoleMode string
```

<a name="LXC_CONSOLE_MODE_SHELL"></a>

```go
const (
    LXC_CONSOLE_MODE_SHELL   LxcConsoleMode = "shell"
    LXC_CONSOLE_MODE_CONSOLE LxcConsoleMode = "console"
    LXC_CONSOLE_MODE_TTY     LxcConsoleMode = "tty"
)
```

<a name="LxcLock"></a>
## type LxcLock



```go
type LxcLock string
```

<a name="LXC_LOCK_BACKUP"></a>

```go
const (
    LXC_LOCK_BACKUP          LxcLock = "backup"
    LXC_LOCK_CREATE          LxcLock = "create"
    LXC_LOCK_DESTROYED       LxcLock = "destroyed"
    LXC_LOCK_DISK            LxcLock = "disk"
    LXC_LOCK_FSTRIM          LxcLock = "fstrim"
    LXC_LOCK_MIGRATE         LxcLock = "migrate"
    LXC_LOCK_MOUNTED         LxcLock = "mounted"
    LXC_LOCK_ROLLBACK        LxcLock = "rollback"
    LXC_LOCK_SNAPSHOT        LxcLock = "snapshot"
    LXC_LOCK_SNAPSHOT_DELETE LxcLock = "snapshot-delete"
)
```

<a name="LxcNet"></a>
## type LxcNet

TODO: Add support for [trunks](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc>) \(vlans\).

```go
type LxcNet struct {
    // contains filtered or unexported fields
}
```

<a name="LxcNet.String"></a>
### func \(\*LxcNet\) String

```go
func (n *LxcNet) String() string
```



<a name="LxcStatus"></a>
## type LxcStatus



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
## type NodeStatus

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
func NewWithCredentials(config ProxmoxAPIConfig, creds *Credentials) (*ProxmoxAPI, error)
```

TODO: To test credentials, do a proxmox version query to ensure credentials are valid

<a name="ProxmoxAPI.CreateClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) CreateClusterFirewallAlias

```go
func (api *ProxmoxAPI) CreateClusterFirewallAlias(name, cidr string, comment *string) error
```

CreateClusterFirewallAlias creates a cluster firewall IP or Network Alias.

<a name="ProxmoxAPI.CreateLxc"></a>
### func \(\*ProxmoxAPI\) CreateLxc

```go
func (api *ProxmoxAPI) CreateLxc(req CreateLxcRequest) (string, error)
```



<a name="ProxmoxAPI.DeleteClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) DeleteClusterFirewallAlias

```go
func (api *ProxmoxAPI) DeleteClusterFirewallAlias(name string, digest *string) error
```

DeleteClusterFirewallAlias removes a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

<a name="ProxmoxAPI.DownloadISOToNodeDatastore"></a>
### func \(\*ProxmoxAPI\) DownloadISOToNodeDatastore

```go
func (api *ProxmoxAPI) DownloadISOToNodeDatastore(node, storageId, fileName, URL string) error
```

DownloadISOToNodeDatastore downloads an iso from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="ProxmoxAPI.DownloadVZTemplateToNodeDatastore"></a>
### func \(\*ProxmoxAPI\) DownloadVZTemplateToNodeDatastore

```go
func (api *ProxmoxAPI) DownloadVZTemplateToNodeDatastore(node, storageId, fileName, URL string) error
```

DownloadVZTemplateToNodeDatastore downloads a vztemplate from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="ProxmoxAPI.GetClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallAlias

```go
func (api *ProxmoxAPI) GetClusterFirewallAlias(name string) (GetAliasResponse, error)
```

GetClusterFirewallAlias retrieves cluster firewall alias by it's name.

<a name="ProxmoxAPI.GetClusterFirewallAliases"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallAliases

```go
func (api *ProxmoxAPI) GetClusterFirewallAliases() ([]GetAliasResponse, error)
```

GetClusterFirewallAliases retrieves all cluster firewall aliases.

<a name="ProxmoxAPI.GetClusterFirewallIPSet"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallIPSet

```go
func (api *ProxmoxAPI) GetClusterFirewallIPSet() ([]GetIPSetResponse, error)
```

GetClusterFirewallIPSet retrieves all cluster firewall IPSets.

<a name="ProxmoxAPI.GetClusterFirewallRules"></a>
### func \(\*ProxmoxAPI\) GetClusterFirewallRules

```go
func (api *ProxmoxAPI) GetClusterFirewallRules() ([]GetRulesResponse, error)
```

GetClusterFirewallRules retrieves all cluster firewall rules.

<a name="ProxmoxAPI.GetLxcs"></a>
### func \(\*ProxmoxAPI\) GetLxcs

```go
func (api *ProxmoxAPI) GetLxcs(node string) ([]GetNodeLxcsResponse, error)
```

GetLxcs returns node's lxc index per node.

<a name="ProxmoxAPI.GetNextVMID"></a>
### func \(\*ProxmoxAPI\) GetNextVMID

```go
func (api *ProxmoxAPI) GetNextVMID() (int, error)
```

GetNextVMID returns the next available VMID.

<a name="ProxmoxAPI.GetNodeDatastoreContent"></a>
### func \(\*ProxmoxAPI\) GetNodeDatastoreContent

```go
func (api *ProxmoxAPI) GetNodeDatastoreContent(node, storageId string) ([]GetNodeDatastoreContentResponse, error)
```

GetNodeDatastoreContent retrieves node's datastores info.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content>).

<a name="ProxmoxAPI.GetNodeDatastores"></a>
### func \(\*ProxmoxAPI\) GetNodeDatastores

```go
func (api *ProxmoxAPI) GetNodeDatastores(node string) ([]GetNodeStoragesResponse, error)
```

GetNodeDatastores retrieves node's datastores info.

<a name="ProxmoxAPI.GetNodeFirewallRules"></a>
### func \(\*ProxmoxAPI\) GetNodeFirewallRules

```go
func (api *ProxmoxAPI) GetNodeFirewallRules(node string) ([]GetNodeFirewallRulesResponse[int], error)
```

GetNodeFirewallRules retrieves node's firewall rules.

<a name="ProxmoxAPI.GetNodeFirewallRulesByPos"></a>
### func \(\*ProxmoxAPI\) GetNodeFirewallRulesByPos

```go
func (api *ProxmoxAPI) GetNodeFirewallRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error)
```

GetNodeFirewallRulesByPos Retrieves a single node's firewall rule using rule's position \(pos\) as an index.

<a name="ProxmoxAPI.GetNodes"></a>
### func \(\*ProxmoxAPI\) GetNodes

```go
func (api *ProxmoxAPI) GetNodes() ([]GetNodesResponse, error)
```

GetNodes retrieves nodes.

<a name="ProxmoxAPI.GetVersion"></a>
### func \(\*ProxmoxAPI\) GetVersion

```go
func (api *ProxmoxAPI) GetVersion() (GetVersionResponse, error)
```

GetVersion retrieves proxmox version.

<a name="ProxmoxAPI.ReadNodeFirewallLog"></a>
### func \(\*ProxmoxAPI\) ReadNodeFirewallLog

```go
func (api *ProxmoxAPI) ReadNodeFirewallLog(node string) ([]FirewallLogEntry, error)
```

ReadNodeFirewallLog Retrieves node's firewall log entries.

TODO: Add missing limit, since, start, until parameters shown in [docs](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log>).

<a name="ProxmoxAPI.UpdateClusterFirewallAlias"></a>
### func \(\*ProxmoxAPI\) UpdateClusterFirewallAlias

```go
func (api *ProxmoxAPI) UpdateClusterFirewallAlias(name, cidr string, comment *string, digest *string, rename *string) error
```

UpdateClusterFirewallAlias updates a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

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