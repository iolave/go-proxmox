package pve

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"slices"
	"strconv"

	"github.com/iolave/go-proxmox/internal/api_def"
	"github.com/iolave/go-proxmox/internal/models"
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
	Name     string
	Bridge   string // Bridge identifier (vmbr1).
	Firewall bool
	GW       string // GatewayIPv4.
	GW6      string // GatewayIPv6.
	HWAddr   string // XX:XX:XX:XX:XX:XX
	IP       string // IPv4/CIDR | dhcp | manual
	IP6      string // IPv6/CIDR | auto | dhcp | manua
	LinkDown bool
	MTU      int
	Rate     int // bmps
	Tag      int
}

func (n *LxcNet) String() string {
	linkDown := helpers.BoolToInt(n.LinkDown)
	fw := helpers.BoolToInt(n.Firewall)

	s := fmt.Sprintf("name=%s", n.Name)
	if n.Bridge != "" {
		s = fmt.Sprintf("%s,bridge=%s", s, n.Bridge)
	}
	if fw != 0 {
		s = fmt.Sprintf("%s,firewall=%d", s, fw)
	}
	if n.GW != "" {
		s = fmt.Sprintf("%s,gw=%s", s, n.GW)
	}
	if n.GW6 != "" {
		s = fmt.Sprintf("%s,gw6=%s", s, n.GW6)
	}
	if n.HWAddr != "" {
		s = fmt.Sprintf("%s,hwaddr=%s", s, n.HWAddr)
	}
	if n.IP != "" {
		s = fmt.Sprintf("%s,ip=%s", s, n.IP)
	}
	if n.IP6 != "" {
		s = fmt.Sprintf("%s,ip6=%s", s, n.IP6)
	}
	if linkDown != 0 {
		s = fmt.Sprintf("%s,link_down=%d", s, linkDown)
	}
	if n.MTU != 0 {
		s = fmt.Sprintf("%s,mtu=%d", s, n.MTU)
	}
	if n.Rate != 0 {
		s = fmt.Sprintf("%s,rate=%d", s, n.Rate)
	}
	if n.Tag != 0 {
		s = fmt.Sprintf("%s,tag=%d", s, n.Tag)
	}

	return s
}

// TODO: Add mount support
// mount
type LXCFeatures struct {
	ForceRWSys *bool
	Fuse       *bool
	KeyCTL     *bool
	MKNod      *bool
	Nesting    *bool
	// TODO:   Mount
}

func (f *LXCFeatures) String() string {
	s := ""
	var intbool int

	if f.ForceRWSys != nil {
		intbool = helpers.BoolToInt(*f.ForceRWSys)
		s = fmt.Sprintf("%s,force_rw_sys=%d", s, intbool)
	}

	if f.Fuse != nil {
		intbool = helpers.BoolToInt(*f.Fuse)
		s = fmt.Sprintf("%s,fuse=%d", s, intbool)
	}

	if f.KeyCTL != nil {
		intbool = helpers.BoolToInt(*f.KeyCTL)
		s = fmt.Sprintf("%s,keyctl=%d", s, intbool)
	}

	if f.MKNod != nil {
		intbool = helpers.BoolToInt(*f.MKNod)
		s = fmt.Sprintf("%s,mknod=%d", s, intbool)
	}

	if f.Nesting != nil {
		intbool = helpers.BoolToInt(*f.Nesting)
		s = fmt.Sprintf("%s,nesting=%d", s, intbool)
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

// GetAll returns node's lxcs.
//
// GET /nodes/{node}/lxc only list CTs where you have VM.Audit permission on.
func (s *PVELxcService) GetAll(node string) ([]GetNodeLxcsResponse, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/lxc")

	res := &[]GetNodeLxcsResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// GetByID returns node's lxc.
//
//   - If the lxc is not found both "lxc" and "err" are going to be nil.
//
// GET /nodes/{node}/lxc only list CTs where you have VM.Audit permission on.
func (s *PVELxcService) GetByID(node string, vmid int) (lxc *GetNodeLxcsResponse, err error) {
	res, err := s.GetAll(node)
	if err != nil {
		return nil, err
	}

	resFiltered := helpers.FilterSlice(res, func(lxc GetNodeLxcsResponse) bool {
		return lxc.VMID == vmid
	})

	if len(resFiltered) == 0 {
		return nil, nil
	}

	return &resFiltered[0], nil
}

// Get returns node's lxc index.
//
// GET /nodes/{node}/lxc/{vmid} accessible by all authenticated users.
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

type pveCreateLxcRequest struct {
	Node               string `in:"nonzero;path=node"`
	OSTemplate         string `in:"nonzero;form=ostemplate"`
	VMID               int    `in:"nonzero;form=vmid"`
	Arch               string `in:"omitempty;form=arch"`
	BWLimit            int    `in:"omitempty;form=bwlimit"`
	CMode              string `in:"omitempty;form=cmode"`
	Console            int    `in:"omitempty;form=console"` // bool
	Cores              int    `in:"omitempty;form=cores"`
	CPULimit           int    `in:"omitempty;form=cpulimit"`
	CPUUnits           int    `in:"omitempty;form=cpuunits"`
	Debug              int    `in:"omitempty;form=debug"` // bool
	Desc               string `in:"omitempty;form=description"`
	Features           string `in:"omitempty;form=features"`
	Force              int    `in:"omitempty;form=force"` // bool
	Hookscript         string `in:"omitempty;form=hookscript"`
	Hostname           string `in:"omitempty;form=hostname"`
	IgnoreUnpackErrors int    `in:"omitempty;form=ignore-unpack-errors"` // bool
	Lock               string `in:"omitempty;form=lock"`
	Memory             int    `in:"omitempty;form=memory"`
	Nameserver         string `in:"omitempty;form=nameserver"`
	OnBoot             int    `in:"omitempty;form=onboot"` // bool
	OSType             string `in:"omitempty;form=ostype"`
	Password           string `in:"omitempty;form=password"`
	Pool               string `in:"omitempty;form=pool"`
	Protection         int    `in:"omitempty;form=protection"` // bool
	Restore            int    `in:"omitempty;form=restore"`    // bool
	RootFS             string `in:"omitempty;form=rootfs;default=local-lvm:8"`
	Searchdomain       string `in:"omitempty;form=searchdomain"`
	SSHPublicKeys      string `in:"omitempty;form=ssh-public-keys"`
	Startup            string `in:"omitempty;form=startup"`
	Storage            string `in:"omitempty;form=storage"`
	Swap               int    `in:"omitempty;form=swap"`
	Tags               string `in:"omitempty;form=tags"`
	Template           int    `in:"omitempty;form=template"` // bool
	Timezone           string `in:"omitempty;form=timezone"`
	TTY                int    `in:"omitempty;form=tty"`
	Unique             int    `in:"omitempty;form=unique"`       // bool
	Unprivileged       int    `in:"omitempty;form=unprivileged"` // bool
	// unused[n] // Reference to unused volumes. This is used internally, and should not be modified manually.
	// dev[n] string Device to pass through to the container
	//mp Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
}

type CreateLxcRequest struct {
	Node               string         // The cluster node name.
	OSTemplate         string         // The OS template or backup file (in format "{STORAGE_ID}:{TYPE}/{TEMPLATE_NAME}", i.e. "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst")
	VMID               int            // The (unique) ID of the VM.
	Arch               LxcArch        // OS architecture type.
	BWLimit            int            // Override I/O bandwidth limit (in KiB/s).
	CMode              LxcConsoleMode // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Console            bool           // Attach a console device (/dev/console) to the container.
	Cores              int            // The number of cores assigned to the container. A container can use all available cores by default.
	CPULimit           int            // Limit of CPU usage. NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	CPUUnits           int            // CPU weight for a container. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this container gets. Number is relative to the weights of all the other running guests.
	Debug              bool           // Try to be more verbose. For now this only enables debug log-level on start.
	Desc               string         // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Features           LXCFeatures    // Allow containers access to advanced features.
	Force              bool           // Allow to overwrite existing container.
	Hookscript         string         // Script that will be exectued during various steps in the containers lifetime.
	Hostname           string         // Set a host name for the container.
	IgnoreUnpackErrors bool           // Ignore errors when extracting the template.
	Lock               LxcLock        // Lock/unlock the container.
	Memory             int            // Amount of RAM for the container in MB.
	Nameserver         string         // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Net                []LxcNet       // Specifies network interfaces for the container.
	OnBoot             bool           // Specifies whether a container will be started during system bootup.
	OSType             string         // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup. debian | devuan | ubuntu | centos | fedora | opensuse | archlinux | alpine | gentoo | nixos | unmanaged
	Password           string         // Sets root password inside container.
	Pool               string         // Add the VM to the specified pool.
	Protection         bool           // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Restore            bool           // Mark this as restore task.
	RootFS             string         // Use volume as container root (in format "{STORAGE_ID}:{SIZE_IN_GIGS}", i.e. "local-lvm:8", if value not specified it defaults to "local-lvm:8", TODO: make this a struct).
	Searchdomain       string         // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	SSHPublicKeys      string         // Setup public SSH keys (one key per line, OpenSSH format).
	Startup            string         // make this a struct Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Storage            string         // Default Storage.
	Swap               int            // Amount of SWAP for the container in MB.
	Tags               string         // Tags of the Container. This is only meta information.
	Template           bool           // Enable/disable Template.
	Timezone           string         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	TTY                int            // Specify the number of tty available to the container.
	Unique             bool           // Assign a unique random ethernet address.
	Unprivileged       bool           // Makes the container run as unprivileged user. (Should not be modified manually.)
	// unused[n] // Reference to unused volumes. This is used internally, and should not be modified manually.
	// dev[n] string Device to pass through to the container
	//mp Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
}

// Create creates an LXC container and return useful information to interact with it after it's creation. If VMID property is lower than 100 a VMID will be generated automatically.
//
// POST /nodes/{node}/lxc requires the 'VM.Allocate' permission on /vms/{vmid} or on the VM pool /pool/{pool}. For restore, it is enough if the user has 'VM.Backup' permission and the VM already exists. You also need 'Datastore.AllocateSpace' permissions on the storage.
func (s *PVELxcService) Create(req CreateLxcRequest) (vmid int, err error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc"

	vmid = req.VMID
	if req.VMID < 100 {
		vmid, err = s.api.Cluster.GetNextVMID()
		if err != nil {
			return 0, err
		}
	}
	req.VMID = vmid

	// convert bool to int
	console := helpers.BoolToInt(req.Console)
	debug := helpers.BoolToInt(req.Debug)
	force := helpers.BoolToInt(req.Force)
	ignoreUnpackErrors := helpers.BoolToInt(req.IgnoreUnpackErrors)
	onBoot := helpers.BoolToInt(req.OnBoot)
	protection := helpers.BoolToInt(req.Protection)
	restore := helpers.BoolToInt(req.Restore)
	template := helpers.BoolToInt(req.Template)
	unique := helpers.BoolToInt(req.Unique)
	unprivileged := helpers.BoolToInt(req.Unprivileged)

	payload := pveCreateLxcRequest{
		Node:               req.Node,
		OSTemplate:         req.OSTemplate,
		VMID:               req.VMID,
		Arch:               string(req.Arch),
		BWLimit:            req.BWLimit,
		CMode:              string(req.CMode),
		Console:            console,
		Cores:              req.Cores,
		CPULimit:           req.CPULimit,
		CPUUnits:           req.CPUUnits,
		Debug:              debug,
		Desc:               req.Desc,
		Features:           req.Features.String(),
		Force:              force,
		Hookscript:         req.Hookscript,
		Hostname:           req.Hostname,
		IgnoreUnpackErrors: ignoreUnpackErrors,
		Lock:               string(req.Lock),
		Memory:             req.Memory,
		Nameserver:         req.Nameserver,
		OnBoot:             onBoot,
		OSType:             req.OSType,
		Password:           req.Password,
		Pool:               req.Pool,
		Protection:         protection,
		Restore:            restore,
		RootFS:             req.RootFS,
		Searchdomain:       req.Searchdomain,
		SSHPublicKeys:      req.SSHPublicKeys,
		Startup:            req.Startup,
		Storage:            req.Storage,
		Swap:               req.Swap,
		Tags:               req.Tags,
		Template:           template,
		Timezone:           req.Timezone,
		TTY:                req.TTY,
		Unique:             unique,
		Unprivileged:       unprivileged,
	}

	netValues := map[string]string{}
	for i, net := range req.Net {
		netValues[fmt.Sprintf("net%d", i)] = net.String()
	}

	err = s.api.client.sendReq3(method, path, &payload, netValues, nil)

	if err != nil {
		return 0, err
	}

	return vmid, nil
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
// DELETE /nodes/{node}/lxc/{vmid} requires the "VM.Allocate" permission.
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

// Exec executes a command inside an lxc.
//
//   - If the lxc is not found an error will be returned.
//   - If the client fails to execute the command, an error
//     will be returned.
//
// This is part of the custom features the go proxmox api wrapper
// provides. It ONLY works if the api wrapper is installed in
// a proxmox node instance.
//
// POST /custom-api/v1/lxc/{id}/exec requires the "VM.Console" permission.
func (s *PVELxcService) Exec(id int, shell string, cmd string) (out string, exitCode int, err error) {
	method := http.MethodPost
	path := fmt.Sprintf("/custom-api/v1/lxc/%d/exec", id)

	req := apidef.PostLXCExecRequest{CMD: cmd, Shell: shell}
	res := apidef.PostLXCExecResponse{}
	if err := s.api.client.sendCustomAPIRequest(method, path, req, &res); err != nil {
		return "", 0, err
	}

	return res.Output, res.ExitCode, nil
}

// ExecAsync executes a command inside an lxc asynchronously.
//
// This is part of the custom features the go proxmox api wrapper
// provides. It ONLY works if the api wrapper is installed in
// a proxmox node instance.
//
// POST /custom-api/v1/lxc/{id}/exec-async requires the "VM.Console" permission.
func (s *PVELxcService) ExecAsync(id int, shell string, cmd string) (execId string, err error) {
	method := http.MethodPost
	path := fmt.Sprintf("/custom-api/v1/lxc/%d/exec-async", id)

	req := apidef.PostLXCExecRequest{CMD: cmd, Shell: shell}
	res := apidef.PostLXCExecAsyncResponse{}
	if err := s.api.client.sendCustomAPIRequest(method, path, req, &res); err != nil {
		return "", err
	}

	return res.ID, nil
}

// GetCMDResult retrieves an async command result.
//
// This is part of the custom features the go proxmox api wrapper
// provides. It ONLY works if the api wrapper is installed in
// a proxmox node instance.
//
// POST /custom-api/v1/cmd/{id} requires the "VM.Audit" permission.
func (s *PVELxcService) GetCMDResult(id string) (result models.CMDExecution, err error) {
	method := http.MethodGet
	path := fmt.Sprintf("/custom-api/v1/cmd/%s", id)

	res := models.CMDExecution{}
	if err := s.api.client.sendCustomAPIRequest(method, path, nil, &res); err != nil {
		return res, err
	}

	return res, nil
}

type GetLxcInterfaceResponse struct {
	Name      string `json:"name"`
	HWAddress string `json:"hwaddr"`
	IPv4      string `json:"inet"`
	IPv6      string `json:"inet6"`
}

// GetInterfaces gets all lxc interfaces. If the lxc status
// is stopped or it doesnt exist, both res and err will be nil.
//
// GET /nodes/{node}/lxc/{id}/interfaces requires the "VM.Audit" permission.
func (s *PVELxcService) GetInterfaces(node string, id int) (res []GetLxcInterfaceResponse, err error) {
	method := http.MethodGet
	path := "/nodes/{node}/lxc/{id}/interfaces"

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

// GetInterfaceByName gets an specific lxc interface by name.
// If the lxc status is stopped, both res and err will be nil.
// If the interface name is not found, an error will be returned.
//
// GET /nodes/{node}/lxc/{id}/interfaces requires the "VM.Audit" permission.
func (s *PVELxcService) GetInterfaceByName(node string, id int, name string) (res GetLxcInterfaceResponse, err error) {
	ifaces, err := s.api.LXC.GetInterfaces(node, id)
	if err != nil {
		return res, err
	}

	for _, iface := range ifaces {
		if iface.Name != name {
			continue
		}
		return iface, nil

	}
	return res, fmt.Errorf("vmid '%d' or interface '%s' not found", id, name)
}

type pveUpdateLxcRequest struct {
	Node     string `in:"nonzero;path=node"`
	VMID     int    `in:"nonzero;path=vmid"`
	Arch     string `in:"omitempty;form=arch"`
	CMode    string `in:"omitempty;form=cmode"`
	Console  int    `in:"omitempty;form=console"` // bool
	Cores    int    `in:"omitempty;form=cores"`
	CPULimit int    `in:"omitempty;form=cpulimit"`
	CPUUnits int    `in:"omitempty;form=cpuunits"`
	Debug    int    `in:"omitempty;form=debug"` // bool
	Desc     string `in:"omitempty;form=description"`
	// dev[n] string Device to pass through to the container
	Features   string `in:"omitempty;form=features"`
	Hookscript string `in:"omitempty;form=hookscript"`
	Hostname   string `in:"omitempty;form=hostname"`
	Lock       string `in:"omitempty;form=lock"`
	Memory     int    `in:"omitempty;form=memory"`
	//mp Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	Nameserver   string `in:"omitempty;form=nameserver"`
	OnBoot       int    `in:"omitempty;form=onboot"` // bool
	OSType       string `in:"omitempty;form=ostype"`
	Protection   int    `in:"omitempty;form=protection"` // bool
	RootFS       string `in:"omitempty;form=rootfs"`
	Searchdomain string `in:"omitempty;form=searchdomain"`
	Startup      string `in:"omitempty;form=startup"`
	Swap         int    `in:"omitempty;form=swap"`
	Tags         string `in:"omitempty;form=tags"`
	Template     int    `in:"omitempty;form=template"` // bool
	Timezone     string `in:"omitempty;form=timezone"`
	TTY          int    `in:"omitempty;form=tty"`
	Unprivileged int    `in:"omitempty;form=unprivileged"` // bool
	// unused[n] // Reference to unused volumes. This is used internally, and should not be modified manually.
}

type UpdateLxcRequest struct {
	Node         string         // The cluster node name.
	VMID         int            // The (unique) ID of the VM.
	Arch         LxcArch        // OS architecture type.
	CMode        LxcConsoleMode // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Console      bool           // Attach a console device (/dev/console) to the container.
	Cores        int            // The number of cores assigned to the container. A container can use all available cores by default.
	CPULimit     int            // Limit of CPU usage. NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	CPUUnits     int            // CPU weight for a container. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this container gets. Number is relative to the weights of all the other running guests.
	Debug        bool           // Try to be more verbose. For now this only enables debug log-level on start.
	Desc         string         // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Features     LXCFeatures    // Allow containers access to advanced features.
	Hookscript   string         // Script that will be exectued during various steps in the containers lifetime.
	Hostname     string         // Set a host name for the container.
	Lock         LxcLock        // Lock/unlock the container.
	Memory       int            // Amount of RAM for the container in MB.
	Nameserver   string         // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Net          []LxcNet       // Specifies network interfaces for the container.
	OnBoot       bool           // Specifies whether a container will be started during system bootup.
	OSType       string         // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup. debian | devuan | ubuntu | centos | fedora | opensuse | archlinux | alpine | gentoo | nixos | unmanaged
	Protection   bool           // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	RootFS       string         // Use volume as container root (in format "{STORAGE_ID}:{SIZE_IN_GIGS}", i.e. "local-lvm:8", if value not specified it defaults to "local-lvm:8", TODO: make this a struct).
	Searchdomain string         // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Startup      string         // make this a struct Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Swap         int            // Amount of SWAP for the container in MB.
	Tags         string         // Tags of the Container. This is only meta information.
	Template     bool           // Enable/disable Template.
	Timezone     string         // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	TTY          int            // Specify the number of tty available to the container.
	Unprivileged bool           // Makes the container run as unprivileged user. (Should not be modified manually.)
	// unused[n] // Reference to unused volumes. This is used internally, and should not be modified manually.
	// dev[n] string Device to pass through to the container
	//mp Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
}

// Update updates an existing LXC container.
//
// PUT /nodes/{node}/lxc/{vmid}/config requires VM.Config.Disk, VM.Config.CPU, VM.Config.Memory, VM.Config.Network, VM.Config.Options permissions.
func (s *PVELxcService) Update(req UpdateLxcRequest) (err error) {
	method := http.MethodPut
	path := "/nodes/{node}/lxc/{vmid}/config"

	// convert bool to int
	console := helpers.BoolToInt(req.Console)
	debug := helpers.BoolToInt(req.Debug)
	onBoot := helpers.BoolToInt(req.OnBoot)
	protection := helpers.BoolToInt(req.Protection)
	template := helpers.BoolToInt(req.Template)
	unprivileged := helpers.BoolToInt(req.Unprivileged)

	payload := pveUpdateLxcRequest{
		Node:         req.Node,
		VMID:         req.VMID,
		Arch:         string(req.Arch),
		CMode:        string(req.CMode),
		Console:      console,
		Cores:        req.Cores,
		CPULimit:     req.CPULimit,
		CPUUnits:     req.CPUUnits,
		Debug:        debug,
		Desc:         req.Desc,
		Features:     req.Features.String(),
		Hookscript:   req.Hookscript,
		Hostname:     req.Hostname,
		Lock:         string(req.Lock),
		Memory:       req.Memory,
		Nameserver:   req.Nameserver,
		OnBoot:       onBoot,
		OSType:       req.OSType,
		Protection:   protection,
		RootFS:       req.RootFS,
		Searchdomain: req.Searchdomain,
		Startup:      req.Startup,
		Swap:         req.Swap,
		Tags:         req.Tags,
		Template:     template,
		Timezone:     req.Timezone,
		TTY:          req.TTY,
		Unprivileged: unprivileged,
	}

	netValues := map[string]string{}
	for i, net := range req.Net {
		netValues[fmt.Sprintf("net%d", i)] = net.String()
	}

	err = s.api.client.sendReq3(method, path, &payload, netValues, nil)

	if err != nil {
		return err
	}

	return nil
}

type pveLXCCloneRequest struct {
	Node     string `in:"nonzero;path=node"`
	VMID     int    `in:"nonzero;path=vmid"`
	NewVMID  int    `in:"nonzero;form=newid"`
	BWLimit  int    `in:"omitempty;form=bwlimit"`
	Desc     string `in:"omitempty;form=description"`
	Full     int    `in:"omitempty;form=full"` // bool
	Hostname string `in:"omitempty;form=hostname"`
	Pool     string `in:"omitempty;form=pool"`
	Snapname string `in:"omitempty;form=snapname"`
	Storage  string `in:"omitempty;form=storage"`
	Target   string `in:"omitempty;form=target"`
}

type CloneLxcRequest struct {
	Node     string // The cluster node name.
	VMID     int    // The (unique) ID of the source VM.
	NewVMID  int    // The (unique) ID of the target VM (if not set, the next available VMID will be used).
	BWLimit  int    // Override I/O bandwidth limit (in KiB/s).
	Desc     string // Description for the Container.
	Full     bool   // Create a full copy of all disks. This is always done when you clone a normal CT. For CT templates, we try to create a linked clone by default.
	Hostname string // Set a host name for the container.
	Pool     string // Add the VM to the specified pool.
	Snapname string // The name of the snapshot.
	Storage  string // Target storage for full clone.
	Target   string // Target node. Only allowed if the original VM is on shared storage.
}

// Clone creates a clone or copy of an existing LXC container.
//
// POST /nodes/{node}/lxc/{vmid}/clone needs 'VM.Clone' permissions
// on /vms/{vmid}, and 'VM.Allocate' permissions on /vms/{newid}
// (or on the VM pool /pool/{pool}). You also need
// 'Datastore.AllocateSpace' on any used storage, and 'SDN.Use'
// on any bridge.
func (s *PVELxcService) Clone(req CloneLxcRequest) (newVMID int, err error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{vmid}/clone"

	if req.VMID == 0 {
		return 0, fmt.Errorf("rep.VMID is required")
	}

	if req.NewVMID == 0 {
		if req.NewVMID, err = s.api.Cluster.GetNextVMID(); err != nil {
			return 0, err
		}
	}

	// convert bool to int
	full := helpers.BoolToInt(req.Full)

	payload := pveLXCCloneRequest{
		Node:     req.Node,
		VMID:     req.VMID,
		NewVMID:  req.NewVMID,
		BWLimit:  req.BWLimit,
		Desc:     req.Desc,
		Full:     full,
		Hostname: req.Hostname,
		Pool:     req.Pool,
		Snapname: req.Snapname,
		Storage:  req.Storage,
		Target:   req.Target,
	}
	if err = s.api.client.sendReq3(method, path, &payload, nil, nil); err != nil {
		return 0, err
	}

	return req.NewVMID, nil
}

func (s *PVELxcService) CreateTemplate(node string, vmid int) (err error) {
	method := http.MethodPost
	path := "/nodes/{node}/lxc/{vmid}/template"
	payload := struct {
		Node string `in:"nonzero;path=node"`
		VMID int    `in:"nonzero;path=vmid"`
	}{
		Node: node, VMID: vmid,
	}
	return s.api.client.sendReq3(method, path, &payload, nil, nil)
}
