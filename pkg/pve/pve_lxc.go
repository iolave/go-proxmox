package pve

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/iolave/go-proxmox/pkg/helpers"
)

type PVELxcService struct {
	api *PVE
}

func newPVELxcService(api *PVE) *PVELxcService {
	service := new(PVELxcService)
	service.api = api
	return service
}

type LxcStatus string

const (
	LXC_STATUS_STOPPED LxcStatus = "stopped"
	LXC_STATUS_RUNNING LxcStatus = "running"
)

type LxcArch string

const (
	LXC_ARCH_AMD64   LxcArch = "amd64"
	LXC_ARCH_I386    LxcArch = "i386"
	LXC_ARCH_ARM64   LxcArch = "arm64"
	LXC_ARCH_ARMHF   LxcArch = "armhf"
	LXC_ARCH_RISCV32 LxcArch = "riscv32"
	LXC_ARCH_RISCV64 LxcArch = "riscv64"
)

type LxcConsoleMode string

const (
	LXC_CONSOLE_MODE_SHELL   LxcConsoleMode = "shell"
	LXC_CONSOLE_MODE_CONSOLE LxcConsoleMode = "console"
	LXC_CONSOLE_MODE_TTY     LxcConsoleMode = "tty"
)

type LxcLock string

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

// TODO: Add support for [trunks] (vlans).
//
// [trunks]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
type LxcNet struct {
	name      string
	bridge    *string // Bridge identifier (vmbr1).
	firewall  *bool
	gw        *string // GatewayIPv4.
	gw6       *string // GatewayIPv6.
	hwaddr    *string // XX:XX:XX:XX:XX:XX
	ip        *string // IPv4/CIDR | dhcp | manual
	ip6       *string // IPv6/CIDR | auto | dhcp | manua
	link_down *bool
	mtu       *int
	rate      *int // bmps
	tag       *int
}

func (n *LxcNet) String() string {
	s := fmt.Sprintf("name=%s", n.name)
	if n.bridge != nil {
		s = fmt.Sprintf("%s,bridge=%s", s, *n.bridge)
	}
	if n.firewall != nil {
		var v int

		if *n.firewall == true {
			v = 1
		} else {
			v = 0
		}

		s = fmt.Sprintf("%s,firewall=%d", s, v)
	}
	if n.gw != nil {
		s = fmt.Sprintf("%s,gw=%s", s, *n.gw)
	}
	if n.gw6 != nil {
		s = fmt.Sprintf("%s,gw6=%s", s, *n.gw6)
	}
	if n.hwaddr != nil {
		s = fmt.Sprintf("%s,hwaddr=%s", s, *n.gw)
	}
	if n.ip != nil {
		s = fmt.Sprintf("%s,ip=%s", s, *n.ip)
	}
	if n.ip6 != nil {
		s = fmt.Sprintf("%s,ip6=%s", s, *n.ip6)
	}
	if n.link_down != nil {
		var v int

		if *n.link_down == true {
			v = 1
		} else {
			v = 0
		}
		s = fmt.Sprintf("%s,link_down=%d", s, v)
	}
	if n.mtu != nil {
		s = fmt.Sprintf("%s,mtu=%d", s, *n.mtu)
	}
	if n.rate != nil {
		s = fmt.Sprintf("%s,rate=%d", s, *n.rate)
	}
	if n.tag != nil {
		s = fmt.Sprintf("%s,tag=%d", s, *n.tag)
	}

	return s
}

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

// GetAll returns node's lxc index per node.
func (s *PVELxcService) GetAll(node string) ([]GetNodeLxcsResponse, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/lxc")

	res := &[]GetNodeLxcsResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// Get returns node's lxc index.
//
// GET /nodes/:node/lxc/:vmid accessible by all authenticated users.
func (s *PVELxcService) Get(node string, vmid int) (
	res []struct {
		Subdir string `json:"subdir"`
	},
	err error,
) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/lxc", strconv.Itoa(vmid))

	err = s.api.client.sendReq(method, path, nil, &res)

	return res, err
}

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
	// unused[n] // Reference to unused volumes. This is used internally, and should not be modified manually.
	// dev[n] string Device to pass through to the container
	//mp Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.

}

type CreateLxcResponse struct {
	VMID int // LXC container id within proxmox.
}

// Create creates an LXC container and return useful information to interact with it after it's creation.
func (s *PVELxcService) Create(req CreateLxcRequest) (CreateLxcResponse, error) {
	method := http.MethodPost
	path := path.Join("/nodes", req.Node, "/lxc")

	var err error
	if req.Node == "" {
		return CreateLxcResponse{}, errors.New("missing 'Node' parameter")
	}

	p := &url.Values{}

	addPayloadValue(p, "ostemplate", &req.OSTemplate, nil)

	var vmid int
	if req.VMID == nil {
		vmid, err = s.api.Cluster.GetNextVMID()
		if err != nil {
			return CreateLxcResponse{}, err
		}
	} else {
		vmid = *req.VMID
	}
	addPayloadValue(p, "vmid", &vmid, nil)
	//addPayloadValue(p, "arch", req.Arch)
	addPayloadValue(p, "bwlimit", req.BWLimit, nil)
	//addPayloadValue(p, "cmode", req.CMode)
	addPayloadValue(p, "console", req.Console, nil)
	addPayloadValue(p, "cores", req.Cores, nil)
	addPayloadValue(p, "cpulimit", req.CPULimit, nil)
	addPayloadValue(p, "cpuunits", req.CPUUnits, nil)
	addPayloadValue(p, "debug", req.Debug, nil)
	addPayloadValue(p, "description", req.Desc, nil)
	addPayloadValue(p, "features", req.Features, nil)
	addPayloadValue(p, "force", req.Force, nil)
	addPayloadValue(p, "hookscript", req.Hookscript, nil)
	addPayloadValue(p, "ignore-unpack-errors", req.IgnoreUnpackErrors, nil)
	//addPayloadValue(p, "lock", req.Lock, nil)
	addPayloadValue(p, "memory", req.Memory, nil)
	addPayloadValue(p, "hostname", req.Nameserver, nil)
	if req.Net != nil {
		for i := 0; i < len(*req.Net); i++ {
			content := (*req.Net)[i].String()
			addPayloadValue(p, fmt.Sprintf("net%d", i), &content, nil)
		}
	}
	addPayloadValue(p, "onboot", req.OnBoot, nil)
	addPayloadValue(p, "ostype", req.OSType, nil)
	addPayloadValue(p, "password", req.Password, nil)
	addPayloadValue(p, "pool", req.Pool, nil)
	addPayloadValue(p, "protection", req.Protection, nil)
	addPayloadValue(p, "restore", req.Restore, nil)
	addPayloadValue(p, "rootfs", req.RootFS, helpers.NewStr("local-lvm:8"))
	addPayloadValue(p, "searchdomain", req.Searchdomain, nil)
	addPayloadValue(p, "ssh-public-keys", req.SSHPublicKeys, nil)
	if req.Start != nil && *req.Start == true {
		// addPayloadValue(p, "start", req.Start)
		// TODO: Manually start using the lxc status start endpoint
	}
	addPayloadValue(p, "startup", req.Startup, nil)
	addPayloadValue(p, "storage", req.Storage, nil)
	addPayloadValue(p, "swap", req.Swap, nil)
	addPayloadValue(p, "tags", req.Tags, nil)
	addPayloadValue(p, "template", req.Template, nil)
	addPayloadValue(p, "timezone", req.Timezone, nil)
	addPayloadValue(p, "tty", req.TTY, nil)
	addPayloadValue(p, "unique", req.Unique, nil)
	addPayloadValue(p, "unprivileged", req.Unprivileged, nil)

	//addPayloadValue(p, "unused[n]", req.Unuseds)
	//addPayloadValue(p, "dev[n]", req.Devs)
	//addPayloadValue(p, "mp[n]", req.MPs)

	err = s.api.client.sendReq(method, path, p, nil)

	if err != nil {
		return CreateLxcResponse{}, nil
	}

	return CreateLxcResponse{VMID: vmid}, nil
}

type DeleteLXCOptions struct {
	DestroyUnreferencedDisks *bool // If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config.
	Force                    *bool // Force destroy, even if running.
	Purge                    *bool // Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed.
}

// Delete destroy the container (also delete all uses files).
//
//   - if opts is set, the request might fail.
//   - opts.Force default value is false.
//   - opts.DestroyUnreferencedDisks default value is false.
//
// DELETE /nodes/:node/lxc/:vmid requires the "VM.Allocate" permission.
func (s *PVELxcService) Delete(node string, vmid int, opts *DeleteLXCOptions) (res string, err error) {
	method := http.MethodDelete
	path := path.Join("/nodes", node, "/lxc", strconv.Itoa(vmid))

	payload := &url.Values{}

	if opts != nil {
		addPayloadValue(payload, "force", opts.Force, nil)
		addPayloadValue(payload, "purge", opts.Purge, nil)
		addPayloadValue(payload, "destroy-unreferenced-disks", opts.DestroyUnreferencedDisks, nil)
	}

	err = s.api.client.sendReq(method, path, payload, &res)

	return res, err
}

type LXCStartRequest struct {
	Node     string `in:"nonzero;path=node"`       // The cluster node name.
	ID       int    `in:"nonzero;path=id"`         // The (unique) ID of the VM.
	Debug    int    `in:"omitempty;form=debug"`    // If set, enables very verbose debug log-level on start. Defaults to 0 (false).
	SkipLock int    `in:"omitempty;form=skiplock"` // Ignore locks - only root is allowed to use this option.
}

// Start starts an lxc container.
//
// POST /nodes/{node}/lxc/{id}/status/start requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Start(req LXCStartRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/start"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type LXCStopRequest struct {
	Node             string `in:"nonzero;path=node"`                // The cluster node name.
	ID               int    `in:"nonzero;path=id"`                  // The (unique) ID of the VM.
	OverruleShutdown int    `in:"omitempty;form=overrule-shutdown"` // Try to abort active 'vzshutdown' tasks before stopping. Defaults to 0 (false).
	SkipLock         int    `in:"omitempty;form=skiplock"`          // Ignore locks - only root is allowed to use this option.
}

// Stop stops an lxc container. This will abruptly stop all processes running in the container.
//
// POST /nodes/{node}/lxc/{id}/status/stop requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Stop(req LXCStopRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/stop"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type LXCSuspendRequest struct {
	Node string `in:"nonzero;path=node"` // The cluster node name.
	ID   int    `in:"nonzero;path=id"`   // The (unique) ID of the VM.
}

// Suspend suspends an lxc. This is experimental.
//
// POST /nodes/{node}/lxc/{id}/status/suspend requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Suspend(req LXCSuspendRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/suspend"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type LXCShutdownRequest struct {
	Node    string `in:"nonzero;path=node"`        // The cluster node name.
	ID      int    `in:"nonzero;path=id"`          // The (unique) ID of the VM.
	Force   int    `in:"omitempty;form=forceStop"` // Make sure the Container stops. Defaults to 0 (false).
	Timeout int    `in:"omitempty;path=timeout"`   // Wait maximal timeout seconds. Defaults to 60.
}

// Shutdown shutdowns an lxc. This will trigger a clean shutdown of the container, see lxc-stop(1) for details.
//
// POST /nodes/{node}/lxc/{id}/status/shutdown requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Shutdown(req LXCShutdownRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/shutdown"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type LXCResumeRequest struct {
	Node string `in:"nonzero;path=node"` // The cluster node name.
	ID   int    `in:"nonzero;path=id"`   // The (unique) ID of the VM.
}

// Resume resumes an lxc.
//
// POST /nodes/{node}/lxc/{id}/status/resume requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Resume(req LXCResumeRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/resume"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type LXCRebootRequest struct {
	Node    string `in:"nonzero;path=node"`      // The cluster node name.
	ID      int    `in:"nonzero;path=id"`        // The (unique) ID of the VM.
	Timeout int    `in:"omitempty;path=timeout"` // Wait maximal timeout seconds.
}

// Reboot reboots an lxc by shutting it down, and starting it again. Applies pending changes.
//
// POST /nodes/{node}/lxc/{id}/status/reboot requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) Reboot(req LXCRebootRequest) (string, error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{id}/status/reboot"

	res := ""
	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return "", err
	}

	return res, nil
}

type GetLxcStatusResponse struct {
	CPUs      int `json:"cpus"`      // Maximum usable CPUs.
	Disk      int `json:"disk"`      // Root disk image space-usage in bytes.
	DiskRead  int `json:"diskread"`  // The amount of bytes the guest read from it's block devices since the guest was started. (Note: This info is not available for all storage types.)
	DiskWrite int `json:"diskwrite"` // The amount of bytes the guest wrote from it's block devices since the guest was started. (Note: This info is not available for all storage types.)
	// TODO: add ha object support
	//ha        object  // HA manager service status.
	Lock     string `json:"lock"`     // The current config lock, if any.
	MaxDisk  int    `json:"maxdisk"`  // Root disk image size in bytes.
	MaxMem   int    `json:"maxmem"`   // Maximum memory in bytes.
	MaxSwap  int    `json:"maxswap"`  // Maximum SWAP memory in bytes.
	Name     string `json:"name"`     // Container name.
	NetIn    int    `json:"netin"`    // The amount of traffic in bytes that was sent to the guest over the network since it was started.
	NetOut   int    `json:"netout"`   // The amount of traffic in bytes that was sent from the guest over the network since it was started.
	Status   string `json:"status"`   // LXC Container status.
	Tags     string `json:"tags"`     // The current configured tags, if any.
	Template int    `json:"template"` // Determines if the guest is a template.
	Uptime   int    `json:"uptime"`   // Uptime in seconds.
	ID       int    `json:"vmid"`     // The (unique) ID of the VM.
}

// GetStatus gets an lxc status.
//
// POST /nodes/{node}/lxc/{id}/status/current requires the "VM.PowerMgmt" permission.
func (s *PVELxcService) GetStatus(node string, id int) (res GetLxcStatusResponse, err error) {
	method := http.MethodGet
	path := "/nodes/{node}/lxc/{id}/status/current"

	req := struct {
		Node string `in:"nonzero;path=node"`
		ID   int    `in:"nonzero;path=id"`
	}{
		Node: node,
		ID:   id,
	}

	if err := s.api.client.sendReq2(method, path, &req, &res); err != nil {
		return res, err
	}

	return res, nil
}
