package rawapi

import (
	"net/http"
	"strconv"

	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/pkg/pve"
)

type GETClusterLogRequest struct {
	// Maximum number of entries (1 - N)
	MaxEntries *int `in:"query=max;omitempty"`
}

// GETClusterLog Read cluster log.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func GETClusterLog(c *pve.Client, req GETClusterLogRequest) (res []struct {
	ID      string `json:"id"`
	Message string `json:"msg"`
	UID     string `json:"uid"`
	Time    int    `json:"time"`
	Node    string `json:"node"`
	PID     int    `json:"pid"`
	Pri     int    `json:"pri"`
	User    string `json:"user"`
	Tag     string `json:"tag"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/log",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterGetTasks List recent tasks (cluster wide).
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterGetTasks(c *pve.Client) (res []struct {
	UpID      string  `json:"upid"`
	Node      *string `json:"node"`
	Status    *string `json:"status"`
	ID        *string `json:"id"`
	StartTime *int    `json:"starttime"`
	Saved     *string `json:"saved"`
	User      *string `json:"user"`
	EndTime   *int    `json:"endtime"`
	Type      *string `json:"type"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/tasks",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterGetStatus Get cluster status information.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func ClusterGetStatus(c *pve.Client) (res []struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	IP      *string `json:"ip"`
	Level   *string `json:"level"`
	Local   *int    `json:"local"`
	NodeID  *int    `json:"nodeid"`
	Nodes   *int    `json:"nodes"`
	Online  *int    `json:"online"`
	QuoRate *int    `json:"quorate"`
	Version *int    `json:"version"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/status",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterGetResourcesRequest struct {
	// Resource type.
	//
	// 	vm | storage | node | sdn
	Type string `in:"query=type"`
}

// ClusterGetResources Resources index (cluster wide).
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterGetResources(c *pve.Client, req ClusterGetResourcesRequest) (res []struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	CGroupMode *int    `json:"cgroup-mode"`
	Content    *string `json:"content"`
	CPU        *int    `json:"cpu"`
	Disk       *int    `json:"disk"`
	DiskRead   *int    `json:"diskread"`
	DiskWrite  *int    `json:"diskwrite"`
	HAState    *string `json:"hastate"`
	Level      *string `json:"level"`
	Lock       *string `json:"lock"`
	MaxCPU     *int    `json:"maxcpu"`
	MaxDisk    *int    `json:"maxdisk"`
	MaxMem     *int    `json:"maxmem"`
	Mem        *int    `json:"mem"`
	MemHost    *int    `json:"memhost"`
	Name       *string `json:"name"`
	NetIn      *int    `json:"netin"`
	NetOut     *int    `json:"netout"`
	Node       *string `json:"node"`
	PluginType *string `json:"plugintype"`
	Pool       *string `json:"pool"`
	Status     *string `json:"status"`
	Storage    *string `json:"storage"`
	Tags       *string `json:"tags"`
	Template   *int    `json:"template"`
	Uptime     *int    `json:"uptime"`
	VMID       *int    `json:"vmid"`
}, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/resources",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

// ClusterGetOptions Get datacenter options. Without 'Sys.Audit'
// on '/' not all options are returned.
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterGetOptions(c *pve.Client) (res struct {
	MigrationUnsecure *int      `json:"migration_unsecure"`
	Keyboard          *string   `json:"keyboard"`
	Description       *string   `json:"description"`
	RegisteredTags    *[]string `json:"registered-tags"`
	AllowedTags       *[]string `json:"allowed-tags"`
	BWLimit           *string   `json:"bwlimit"`
	Fencing           *string   `json:"fencing"`
	EmailFrom         *string   `json:"email_from"`
	Language          *string   `json:"language"`
	HTTPProxy         *string   `json:"http_proxy"`
	MacPrefix         *string   `json:"mac_prefix"`
	Console           *string   `json:"console"`
	MaxWorkers        *string   `json:"max_workers"`
	TagStyle          struct {
		CaseSensitive *int    `json:"case-sensitive"`
		Ordering      *string `json:"ordering"`
		Shape         *string `json:"shape"`
	} `json:"tag-style"`
	WebAuthn struct {
		ID              *string `json:"id"`
		RP              *string `json:"rp"`
		AllowSubdomains *int    `json:"allow-subdomains"`
		Origin          *string `json:"origin"`
	} `json:"webauthn"`
	CRS struct {
		HARebalanceOnStart *int    `json:"ha-rebalance-on-start"`
		HA                 *string `json:"ha"`
	} `json:"crs"`
	HA struct {
		ShutdownPolicy *string `json:"shutdown_policy"`
	} `json:"ha"`
	Notify struct {
		TargetPackageUpdates *string `json:"target-package-updates"`
		Fencing              *string `json:"fencing"`
		TargetFencing        *string `json:"target-fencing"`
		PackageUpdates       *string `json:"package-updates"`
		Replication          *string `json:"replication"`
		TargetReplication    *string `json:"target-replication"`
	} `json:"notify"`
	U2F struct {
		Origin *string `json:"origin"`
		AppID  *string `json:"appid"`
	}
	Migration struct {
		Type    *string `json:"type"`
		Network *string `json:"network"`
	} `json:"migration"`
	NextID struct {
		Lower *string `json:"lower"`
		Upper *string `json:"upper"`
	} `json:"next-id"`
	UserTagAccess struct {
		UserAllowList *[]string `json:"user-allow-list"`
		UserAllow     *string   `json:"user-allow"`
	} `json:"user-tag-access"`
}, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/options",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterPutOptionsRequest struct {
	// Set I/O bandwidth limit for various operations (in KiB/s).
	//
	// 	[clone=<LIMIT>] [,default=<LIMIT>] [,migration=<LIMIT>] [,move=<LIMIT>] [,restore=<LIMIT>]
	BWLimit string `in:"query=bwlimit;omitempty"`

	// docs says this is available in API but...
	// ¡NOT AVAILABLE IN API!
	// ConsentText       string `in:"form;omitempty,form=consent-text"`

	// Select the default Console viewer. You can either use the builtin java applet (VNC; deprecated and maps to html5), an external virt-viewer comtatible application (SPICE), an HTML5 based vnc viewer (noVNC), or an HTML5 based console client (xtermjs). If the selected viewer is not available (e.g. SPICE not activated for the VM), the fallback is noVNC.
	//
	// 	applet | vv | html5 | xtermjs
	Console string `in:"query=console;omitempty"`

	// Cluster resource scheduling settings.
	//
	// 	[ha=<basic|static>] [,ha-rebalance-on-start=<1|0>]
	CRS string `in:"query=crs;omitempty"`

	// A list of settings you want to delete.
	//
	// 	eg: bwlimit,console,crs
	Delete string `in:"query=delete;omitempty"`

	// Datacenter description. Shown in the web-interface datacenter notes panel. This is saved as comment inside the configuration file.
	Description string `in:"query=description;omitempty"`

	// Specify email address to send notification from (default is root@$hostname)
	EmailFrom string `in:"query=email_from;omitempty"`

	// Set the fencing mode of the HA cluster. Hardware mode needs a valid configuration of fence devices in /etc/pve/ha/fence.cfg. With both all two modes are used.
	//
	// WARNING: 'hardware' and 'both' are EXPERIMENTAL & WIP
	// 	watchdog | hardware | both
	Fencing string `in:"query=fencing;omitempty"`

	// Cluster wide HA settings.
	//
	// 	shutdown_policy=<enum>
	HA string `in:"query=ha;omitempty"`

	// Specify external http proxy which is used for downloads (example: 'http://username:password@host:port/')
	HTTPProxy string `in:"query=http_proxy;omitempty"`

	// Default keybord layout for vnc server.
	//
	// 	de | de-ch | da | en-gb | en-us | es | fi | fr | fr-be | fr-ca | fr-ch | hu | is | it | ja | lt | mk | nl | no | pl | pt | pt-br | sv | sl | tr
	Keyboard string `in:"query=keyboard;omitempty"`

	// Default GUI language.
	//
	// 	ar | ca | da | de | en | es | eu | fa | fr | hr | he | it | ja | ka | kr | nb | nl | nn | pl | pt_BR | ru | sl | sv | tr | ukr | zh_CN | zh_TW
	Language string `in:"query=language;omitempty"`

	// Prefix for the auto-generated MAC addresses of virtual guests. The default `BC:24:11` is the Organizationally Unique Identifier (OUI) assigned by the IEEE to Proxmox Server Solutions GmbH for a MAC Address Block Large (MA-L). You're allowed to use this in local networks, i.e., those not directly reachable by the public (e.g., in a LAN or NAT/Masquerading).
	//
	// Note that when you run multiple cluster that (partially) share the networks of their virtual guests, it's highly recommended that you extend the default MAC prefix, or generate a custom (valid) one, to reduce the chance of MAC collisions. For example, add a separate extra hexadecimal to the Proxmox OUI for each cluster, like `BC:24:11:0` for the first, `BC:24:11:1` for the second, and so on.
	// Alternatively, you can also separate the networks of the guests logically, e.g., by using VLANs.
	//
	// For publicly accessible guests it's recommended that you get your own https://standards.ieee.org/products-programs/regauth/[OUI from the IEEE] registered or coordinate with your, or your hosting providers, network admins.
	MacPrefix string `in:"query=mac_prefix;omitempty"`

	// Defines how many workers (per node) are maximal started  on actions like 'stopall VMs' or task from the ha-manager.
	MaxWorkers int `in:"query=max_workers;omitempty"`

	// For cluster wide migration settings.
	//
	// 	[type=]<secure|insecure> [,network=<CIDR>]
	Migration string `in:"query=migration;omitempty"`

	// Migration is secure using SSH tunnel by default. For secure private networks you can disable it to speed up migration. Deprecated, use the 'migration' property instead!
	MigrationUnsecure string `in:"query=migration_unsecure;omitempty"`

	// Control the range for the free VMID auto-selection pool.
	//
	// 	[lower=<integer>] [,upper=<integer>]
	NextID string `in:"query=next-id;omitempty"`

	// Cluster-wide notification settings.
	//
	// 	[fencing=<always|never>] [,package-updates=<auto|always|never>] [,replication=<always|never>] [,target-fencing=<TARGET>] [,target-package-updates=<TARGET>] [,target-replication=<TARGET>]
	Notify string `in:"query=notify;omitempty"`

	// A list of tags that require a `Sys.Modify` on '/' to set and delete. Tags set here that are also in 'user-tag-access' also require `Sys.Modify`.
	//
	// 	<tag>[;<tag>...]
	RegisteredTags string `in:"query=registered-tags;omitempty"`

	// docs says this is available in API but...
	// ¡NOT AVAILABLE IN API!
	// Replication       string `in:"form;omitempty,form=replication"`

	//Tag style options.
	//
	// 	[case-sensitive=<1|0>] [,color-map=<tag>:<hex-color>[:<hex-color-for-text>][;<tag>=...]] [,ordering=<config|alphabetical>] [,shape=<enum>]
	TagStyle string `in:"query=tag-style;omitempty"`

	// u2f
	//
	// 	[appid=<APPID>] [,origin=<URL>]
	U2F string `in:"query=u2f;omitempty"`

	// Privilege options for user-settable tags
	//
	// 	[user-allow=<enum>] [,user-allow-list=<tag>[;<tag>...]]
	UserTagAccess string `in:"query=user-tag-access;omitempty"`

	// webauthn configuration
	//
	// 	[allow-subdomains=<1|0>] [,id=<DOMAINNAME>] [,origin=<URL>] [,rp=<RELYING_PARTY>]
	WebAuthn string `in:"query=webauthn;omitempty"`
}

// ClusterPutOptions Set datacenter options.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func ClusterPutOptions(c *pve.Client, req ClusterPutOptionsRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/options",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

type ClusterGetNextIDRequest struct {
	// The (unique) ID of the VM.
	VMID int `in:"query=vmid;omitempty"`
}

// ClusterGetNextID Get next free VMID. Pass a VMID
// to assert that its free (at time of check).
//
// Required permissions:
//
//	Accessible by all authenticated users.
func ClusterGetNextID(c *pve.Client, req ClusterGetNextIDRequest) (int, error) {
	res := ""

	err := c.Do(pve.Request{
		Path:    "/api2/json/cluster/nextid",
		Method:  http.MethodGet,
		Payload: &req,
		Result:  &res,
	})

	vmid, err := strconv.Atoi(res)
	if err != nil {
		return 0, errors.NewInternalServerError(
			"failed to parse response",
			err,
		)
	}

	return vmid, err
}
