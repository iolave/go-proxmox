# pve

```go
import "github.com/iolave/go-proxmox/pkg/pve"
```





## Constants

<a name="CREDENTIALS_NOT_DETECTED_ERROR"></a>Credentials error messages.

```go
const (
    CREDENTIALS_NOT_DETECTED_ERROR    = "credentials could not be detected from env"
    CREDENTIALS_NOT_SUPPORTED_ERROR   = "credentials type not supported yet"
    CREDENTIALS_MISSING_REQUEST_ERROR = "*http.Request parameter is nil"
)
```

<a name="APTRepoInfoError"></a>
## type [APTRepoInfoError](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L59-L62>)



```go
type APTRepoInfoError struct {
    Error string `json:"error"` // The error message.
    Path  string `json:"path"`  // Path to the problematic file.
}
```

<a name="APTRepoInfoFile"></a>
## type [APTRepoInfoFile](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L80-L85>)



```go
type APTRepoInfoFile struct {
    Digest       []int                 `json:"digest"`       // Digest of the file as bytes.
    FileType     string                `json:"file-type"`    // Format of the file ("list", "sources").
    Path         string                `json:"path"`         // Path to the problematic file.
    Repositories []APTRepoInfoFileRepo `json:"repositories"` // The parsed repositories.
}
```

<a name="APTRepoInfoFileRepo"></a>
## type [APTRepoInfoFileRepo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L69-L78>)



```go
type APTRepoInfoFileRepo struct {
    Comment    string                   // Associated comment.
    Components []string                 // List of repository components
    Enabled    bool                     // Whether the repository is enabled or not.
    FileType   string                   // Format of the defining file ("list", "sources").
    Options    []APTRepoInfoFileRepoOpt // Additional options.
    Suites     []string                 // List of package distribuitions
    Types      []string                 // List of package types ("deb", "deb-src").
    URIs       []string                 // List of repository URIs.
}
```

<a name="APTRepoInfoFileRepoOpt"></a>
## type [APTRepoInfoFileRepoOpt](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L64-L67>)



```go
type APTRepoInfoFileRepoOpt struct {
    Key    string
    Values []string
}
```

<a name="APTRepoInfoInfos"></a>
## type [APTRepoInfoInfos](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L87-L93>)



```go
type APTRepoInfoInfos struct {
    Index    string `json:"index"`    // Index of the associated repository within the file.
    Kind     string `json:"kind"`     // Kind of the information (e.g. warning).
    Message  string `json:"message"`  // Information message.
    Path     string `json:"path"`     // Path to the associated file.
    Property string `json:"property"` // Property from which the info originates.
}
```

<a name="APTRepoInfoStdRepo"></a>
## type [APTRepoInfoStdRepo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L95-L99>)



```go
type APTRepoInfoStdRepo struct {
    Handle string `json:"handle"` // Handle to identify the repository.
    Name   string `json:"name"`   // Full name of the repository.
    Status *bool  `json:"status"` // Indicating enabled/disabled status, if the repository is configured.
}
```

<a name="Config"></a>
## type [Config](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve.go#L9-L14>)



```go
type Config struct {
    Host               string
    Port               int
    InsecureSkipVerify bool
    CfServiceToken     *cloudflare.ServiceToken
}
```

<a name="CreateLxcRequest"></a>
## type [CreateLxcRequest](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L178-L223>)



```go
type CreateLxcRequest struct {
    Node               string          // The cluster node name.
    OSTemplate         string          // The OS template or backup file (in format "{STORAGE_ID}:{TYPE}/{TEMPLATE_NAME}", i.e. "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst")
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
    RootFS             *string         // Use volume as container root (in format "{STORAGE_ID}:{SIZE_IN_GIGS}", i.e. "local-lvm:8", if value not specified it defaults to "local-lvm:8", TODO: make this a struct).
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

<a name="CreateLxcResponse"></a>
## type [CreateLxcResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L225-L227>)



```go
type CreateLxcResponse struct {
    VMID int // LXC container id within proxmox.
}
```

<a name="CredentialType"></a>
## type [CredentialType](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_creds.go#L11>)

Proxmox api client available credential types

```go
type CredentialType int
```

<a name="CREDENTIALS_TOKEN"></a>TODO: Add CREDENTIALS\_PASSWORD support

```go
const (
    CREDENTIALS_TOKEN CredentialType = iota
    CREDENTIALS_PASSWORD
)
```

<a name="Credentials"></a>
## type [Credentials](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_creds.go#L27-L32>)

Credentials store proxmox api credentials.

```go
type Credentials struct {
    // contains filtered or unexported fields
}
```

<a name="NewEnvCreds"></a>
### func [NewEnvCreds](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_creds.go#L82>)

```go
func NewEnvCreds() (*Credentials, error)
```

NewEnvCreds get [environment variables](<https://go-proxmox.iolave.com/getting-started/#enviroment-variables>) values and detects the type of credentials based on which envs are configured.

It returns an error when a credential type is not detected.

<a name="NewTokenCreds"></a>
### func [NewTokenCreds](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_creds.go#L67>)

```go
func NewTokenCreds(user, tokenName, token string) *Credentials
```

NewTokenCreds returns a struct containing proxmox token credentials that can be passed to a [pve api constructor](<https://TODO:add-the-proper-ref>).

To create a pve token, read the [docs](<https://pve.proxmox.com/wiki/Proxmox_VE_API#API_Tokens>).

<a name="Credentials.Set"></a>
### func \(\*Credentials\) [Set](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_creds.go#L46>)

```go
func (c *Credentials) Set(req *http.Request) error
```

Set adds the corresponding PVE authorization headers to the req parameter.

\* It returns an error with the [CREDENTIALS\\\_MISSING\\\_REQUEST\\\_ERROR](<https://go-proxmox.iolave.com/reference/pkg/pve#constants>) message when nil is passed to the req parameter.

\* It returns an error with the [CREDENTIALS\\\_NOT\\\_SUPPORTED\\\_ERROR](<https://go-proxmox.iolave.com/reference/pkg/pve#constants>) message when [CredentialType](<https://go-proxmox.iolave.com/reference/pkg/pve#type-credentialtype>) is not supported.

<a name="DeleteLXCOptions"></a>
## type [DeleteLXCOptions](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L311-L315>)



```go
type DeleteLXCOptions struct {
    DestroyUnreferencedDisks *bool // If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config.
    Force                    *bool // Force destroy, even if running.
    Purge                    *bool // Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed.
}
```

<a name="FirewallLogEntry"></a>
## type [FirewallLogEntry](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_firewall.go#L19-L22>)

Proxmox firewall log entry.

```go
type FirewallLogEntry struct {
    Id   int    `json:"n"`
    Text string `json:"t"`
}
```

<a name="FirewallLogLevel"></a>
## type [FirewallLogLevel](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_firewall.go#L4>)

Promox firewall available log levels.

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

<a name="GetClusterFirewallAliasesResponse"></a>
## type [GetClusterFirewallAliasesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L19-L24>)



```go
type GetClusterFirewallAliasesResponse struct {
    CIDR    string `json:"cidr"`
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetClusterFirewallIPSetResponse"></a>
## type [GetClusterFirewallIPSetResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L108-L112>)



```go
type GetClusterFirewallIPSetResponse struct {
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    Comment string `json:"comment"`
}
```

<a name="GetClusterFirewallRulesResponse"></a>
## type [GetClusterFirewallRulesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L125-L127>)



```go
type GetClusterFirewallRulesResponse struct {
    Pos int `json:"pos"`
}
```

<a name="GetNodeAPTRepoInfo"></a>
## type [GetNodeAPTRepoInfo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L101-L107>)



```go
type GetNodeAPTRepoInfo struct {
    Digest   string               `json:"digest"`         // Common digest of all files.
    Errors   []APTRepoInfoError   `json:"errors"`         // List of problematic repository files.
    Files    []APTRepoInfoFile    `json:"files"`          // List of parsed repository files.
    Infos    []APTRepoInfoInfos   `json:"infos"`          // Additional information/warnings for APT repositories.
    StdRepos []APTRepoInfoStdRepo `json:"standard-repos"` // List of standard repositories and their configuration status.
}
```

<a name="GetNodeDatastoreContentResponse"></a>
## type [GetNodeDatastoreContentResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L46-L57>)

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

<a name="GetNodeDatastoreResponse"></a>
## type [GetNodeDatastoreResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L19-L30>)



```go
type GetNodeDatastoreResponse struct {
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

<a name="GetNodeFirewallRulesResponse"></a>
## type [GetNodeFirewallRulesResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_firewall.go#L19-L35>)



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
## type [GetNodeLxcsResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L137-L148>)



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

<a name="GetNodeResponse"></a>
## type [GetNodeResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node.go#L33-L43>)



```go
type GetNodeResponse struct {
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

<a name="GetVersionResponse"></a>
## type [GetVersionResponse](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_version.go#L5-L9>)



```go
type GetVersionResponse struct {
    Release string `json:"release"`
    Version string `json:"version"`
    RepoID  string `json:"repoid"`
}
```

<a name="LxcArch"></a>
## type [LxcArch](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L31>)



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
## type [LxcConsoleMode](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L42>)



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
## type [LxcLock](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L50>)



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
## type [LxcNet](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L68-L81>)

TODO: Add support for [trunks](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc>) \(vlans\).

```go
type LxcNet struct {
    // contains filtered or unexported fields
}
```

<a name="LxcNet.String"></a>
### func \(\*LxcNet\) [String](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L83>)

```go
func (n *LxcNet) String() string
```



<a name="LxcStatus"></a>
## type [LxcStatus](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L24>)



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

<a name="NodeAPTIndex"></a>
## type [NodeAPTIndex](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L22-L24>)



```go
type NodeAPTIndex struct {
    ID string `json:"id"`
}
```

<a name="NodeStatus"></a>
## type [NodeStatus](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node.go#L25>)

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

<a name="PVE"></a>
## type [PVE](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve.go#L16-L25>)



```go
type PVE struct {

    // PVE API implementations
    Node    *PVENodeService
    Cluster *PVEClusterService
    LXC     *PVELxcService
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve.go#L27>)

```go
func New(config Config) (*PVE, error)
```



<a name="NewWithCredentials"></a>
### func [NewWithCredentials](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve.go#L57>)

```go
func NewWithCredentials(config Config, creds *Credentials) (*PVE, error)
```



<a name="PVE.GetVersion"></a>
### func \(\*PVE\) [GetVersion](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_version.go#L12>)

```go
func (api *PVE) GetVersion() (GetVersionResponse, error)
```

GetVersion retrieves proxmox version.

<a name="PVEClusterFirewallService"></a>
## type [PVEClusterFirewallService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L9-L11>)



```go
type PVEClusterFirewallService struct {
    // contains filtered or unexported fields
}
```

<a name="PVEClusterFirewallService.CreateAlias"></a>
### func \(\*PVEClusterFirewallService\) [CreateAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L49>)

```go
func (s *PVEClusterFirewallService) CreateAlias(name, cidr string, comment *string) error
```

CreateAlias creates a cluster firewall IP or Network Alias.

<a name="PVEClusterFirewallService.DeleteAlias"></a>
### func \(\*PVEClusterFirewallService\) [DeleteAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L94>)

```go
func (s *PVEClusterFirewallService) DeleteAlias(name string, digest *string) error
```

DeleteAlias removes a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

<a name="PVEClusterFirewallService.GetAlias"></a>
### func \(\*PVEClusterFirewallService\) [GetAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L38>)

```go
func (s *PVEClusterFirewallService) GetAlias(name string) (GetClusterFirewallAliasesResponse, error)
```

GetAlias retrieves cluster firewall alias by it's name.

<a name="PVEClusterFirewallService.GetAliases"></a>
### func \(\*PVEClusterFirewallService\) [GetAliases](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L27>)

```go
func (s *PVEClusterFirewallService) GetAliases() ([]GetClusterFirewallAliasesResponse, error)
```

GetAliases retrieves all cluster firewall aliases.

<a name="PVEClusterFirewallService.GetIPSet"></a>
### func \(\*PVEClusterFirewallService\) [GetIPSet](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L115>)

```go
func (s *PVEClusterFirewallService) GetIPSet() ([]GetClusterFirewallIPSetResponse, error)
```

GetIPSet retrieves all cluster firewall IPSets.

<a name="PVEClusterFirewallService.GetRules"></a>
### func \(\*PVEClusterFirewallService\) [GetRules](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L130>)

```go
func (s *PVEClusterFirewallService) GetRules() ([]GetClusterFirewallRulesResponse, error)
```

GetRules retrieves all cluster firewall rules.

<a name="PVEClusterFirewallService.UpdateAlias"></a>
### func \(\*PVEClusterFirewallService\) [UpdateAlias](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster_firewall.go#L68>)

```go
func (s *PVEClusterFirewallService) UpdateAlias(name, cidr string, comment *string, digest *string, rename *string) error
```

UpdateAlias updates a cluster firewall IP or Network alias.

Digest prevents changes if current configuration file has a different digest. This can be used to prevent concurrent modifications.

<a name="PVEClusterService"></a>
## type [PVEClusterService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster.go#L9-L12>)



```go
type PVEClusterService struct {
    Firewall *PVEClusterFirewallService
    // contains filtered or unexported fields
}
```

<a name="PVEClusterService.GetNextVMID"></a>
### func \(\*PVEClusterService\) [GetNextVMID](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_cluster.go#L22>)

```go
func (s *PVEClusterService) GetNextVMID() (int, error)
```

GetNextVMID returns the next available VMID.

<a name="PVELxcService"></a>
## type [PVELxcService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L14-L16>)



```go
type PVELxcService struct {
    // contains filtered or unexported fields
}
```

<a name="PVELxcService.Create"></a>
### func \(\*PVELxcService\) [Create](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L230>)

```go
func (s *PVELxcService) Create(req CreateLxcRequest) (CreateLxcResponse, error)
```

Create creates an LXC container and return useful information to interact with it after it's creation.

<a name="PVELxcService.Delete"></a>
### func \(\*PVELxcService\) [Delete](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L323>)

```go
func (s *PVELxcService) Delete(node string, vmid int, opts *DeleteLXCOptions) (res string, err error)
```

Delete destroy the container \(also delete all uses files\).

- opts.Force default value is false.
- opts.DestroyUnreferencedDisks default value is false.

DELETE /nodes/:node/lxc/:vmid requires the "VM.Allocate" permission.

<a name="PVELxcService.Get"></a>
### func \(\*PVELxcService\) [Get](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L164-L169>)

```go
func (s *PVELxcService) Get(node string, vmid int) (res []struct {
    Subdir string `json:"subdir"`
}, err error)
```

Get returns node's lxc index.

GET /nodes/:node/lxc/:vmid accessible by all authenticated users.

<a name="PVELxcService.GetAll"></a>
### func \(\*PVELxcService\) [GetAll](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_lxc.go#L151>)

```go
func (s *PVELxcService) GetAll(node string) ([]GetNodeLxcsResponse, error)
```

GetAll returns node's lxc index per node.

<a name="PVENodeAPTService"></a>
## type [PVENodeAPTService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L12-L14>)



```go
type PVENodeAPTService struct {
    // contains filtered or unexported fields
}
```

<a name="PVENodeAPTService.AddStdRepo"></a>
### func \(\*PVENodeAPTService\) [AddStdRepo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L156>)

```go
func (s *PVENodeAPTService) AddStdRepo(node, handle string, digest *string) error
```

AddStdRepo adds a standard repository to the configuration.

- node: Cluster node name.
- handle: Handle that identifies a repository.
- digest: Digest to detect modifications.

PUT /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.

<a name="PVENodeAPTService.GetChangelog"></a>
### func \(\*PVENodeAPTService\) [GetChangelog](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L42>)

```go
func (s *PVENodeAPTService) GetChangelog(node, name string, version *string) (string, error)
```

GetChangelog returns the changelog for a given pacakge name. If version is nil, the latest version available will be considered and otherwise, it will return the changelog found for the given version.

GET /nodes/:node/apt/changelog requires the "Sys.Audit" permission.

<a name="PVENodeAPTService.GetIndex"></a>
### func \(\*PVENodeAPTService\) [GetIndex](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L29>)

```go
func (s *PVENodeAPTService) GetIndex(node string) ([]NodeAPTIndex, error)
```

GetIndex returns node's directory index for apt \(Advanced Package Tool\).

GET /nodes/:node/apt accessible by all authenticated users.

<a name="PVENodeAPTService.GetPVEInfo"></a>
### func \(\*PVENodeAPTService\) [GetPVEInfo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L214>)

```go
func (s *PVENodeAPTService) GetPVEInfo(node string) (interface{}, error)
```

GetPVEInfo get package information for important Proxmox packages.

TODO: [docs](<https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/versions>) lacks of response definition \(map it\).

- node: Cluster node name.

GET /nodes/:node/apt/versions requires the "Sys.Audit" permission.

<a name="PVENodeAPTService.GetRepoInfo"></a>
### func \(\*PVENodeAPTService\) [GetRepoInfo](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L112>)

```go
func (s *PVENodeAPTService) GetRepoInfo(node string) (GetNodeAPTRepoInfo, error)
```

GetRepoInfo returns APT repository information.

GET /nodes/:node/apt/repositories requires the "Sys.Audit" permission.

<a name="PVENodeAPTService.ListUpdates"></a>
### func \(\*PVENodeAPTService\) [ListUpdates](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L179>)

```go
func (s *PVENodeAPTService) ListUpdates(node string) (interface{}, error)
```

ListUpdates list available updates.

TODO: [docs](<https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/update>) lacks of response definition \(map it\).

- node: Cluster node name.

GET /nodes/:node/apt/update requires the "Sys.Modify" permission.

<a name="PVENodeAPTService.SetRepoProps"></a>
### func \(\*PVENodeAPTService\) [SetRepoProps](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L130>)

```go
func (s *PVENodeAPTService) SetRepoProps(index int, node, filePath string, digest *string, enabled *bool) error
```

SetRepoProps changes the properties of a repository \(currently only allows enabling/disabling\).

- "index": Index within the file \(starting from 0\).
- "node": Cluster node name.
- "filePath": Path to the containing file.
- "digest": Digest to detect modifications.
- "enabled": Whether the repository should be enabled or not.

POST /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.

<a name="PVENodeAPTService.UpdateIndex"></a>
### func \(\*PVENodeAPTService\) [UpdateIndex](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_apt.go#L196>)

```go
func (s *PVENodeAPTService) UpdateIndex(node string, notify, quiet bool) (string, error)
```

UpdateIndex this is used to resynchronize the package index files from their sources \(apt\-get update\).

- node: Cluster node name.
- notify: Send notification about new packages.
- quiet: Only produces output suitable for logging, omitting progress indicators.

POST /nodes/:node/apt/update requires the "Sys.Modify" permission.

<a name="PVENodeFirewallService"></a>
## type [PVENodeFirewallService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_firewall.go#L9-L11>)



```go
type PVENodeFirewallService struct {
    // contains filtered or unexported fields
}
```

<a name="PVENodeFirewallService.GetRules"></a>
### func \(\*PVENodeFirewallService\) [GetRules](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_firewall.go#L38>)

```go
func (s *PVENodeFirewallService) GetRules(node string) ([]GetNodeFirewallRulesResponse[int], error)
```

GetRules retrieves node's firewall rules.

<a name="PVENodeFirewallService.GetRulesByPos"></a>
### func \(\*PVENodeFirewallService\) [GetRulesByPos](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_firewall.go#L49>)

```go
func (s *PVENodeFirewallService) GetRulesByPos(node string, pos int) (GetNodeFirewallRulesResponse[string], error)
```

GetRulesByPos Retrieves a single node's firewall rule using rule's position \(pos\) as an index.

<a name="PVENodeFirewallService.ReadLog"></a>
### func \(\*PVENodeFirewallService\) [ReadLog](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_firewall.go#L64>)

```go
func (s *PVENodeFirewallService) ReadLog(node string) ([]FirewallLogEntry, error)
```

ReadLog Retrieves node's firewall log entries.

TODO: Add missing limit, since, start, until parameters shown in [docs](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log>).

<a name="PVENodeService"></a>
## type [PVENodeService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node.go#L8-L13>)



```go
type PVENodeService struct {
    APT      *PVENodeAPTService
    Firewall *PVENodeFirewallService
    Storage  *PVENodeStorageService
    // contains filtered or unexported fields
}
```

<a name="PVENodeService.Get"></a>
### func \(\*PVENodeService\) [Get](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node.go#L57>)

```go
func (service *PVENodeService) Get(node string) (GetNodeResponse, error)
```

Get retrieves a single nodes.

<a name="PVENodeService.GetAll"></a>
### func \(\*PVENodeService\) [GetAll](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node.go#L46>)

```go
func (service *PVENodeService) GetAll() ([]GetNodeResponse, error)
```

GetAll retrieves all nodes.

<a name="PVENodeStorageService"></a>
## type [PVENodeStorageService](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L9-L11>)



```go
type PVENodeStorageService struct {
    // contains filtered or unexported fields
}
```

<a name="PVENodeStorageService.DownloadISOToDatastore"></a>
### func \(\*PVENodeStorageService\) [DownloadISOToDatastore](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L79>)

```go
func (s *PVENodeStorageService) DownloadISOToDatastore(node, storageId, fileName, URL string) error
```

DownloadISOToDatastore downloads an iso from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="PVENodeStorageService.DownloadVZTemplateToDatastore"></a>
### func \(\*PVENodeStorageService\) [DownloadVZTemplateToDatastore](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L97>)

```go
func (s *PVENodeStorageService) DownloadVZTemplateToDatastore(node, storageId, fileName, URL string) error
```

DownloadVZTemplateToDatastore downloads a vztemplate from an url into a node's datastore.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url>).

<a name="PVENodeStorageService.GetDatastoreContent"></a>
### func \(\*PVENodeStorageService\) [GetDatastoreContent](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L64>)

```go
func (s *PVENodeStorageService) GetDatastoreContent(node, storageId string) ([]GetNodeDatastoreContentResponse, error)
```

GetDatastoreContent retrieves node's datastores info.

TODO: Add optional [parameters](<https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content>).

<a name="PVENodeStorageService.GetDatastores"></a>
### func \(\*PVENodeStorageService\) [GetDatastores](<https://github.com/iolave/go-proxmox/blob/master/pkg/pve/pve_node_storage.go#L33>)

```go
func (s *PVENodeStorageService) GetDatastores(node string) ([]GetNodeDatastoreResponse, error)
```

GetDatastores retrieves node's datastores info.