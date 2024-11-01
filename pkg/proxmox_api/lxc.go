package proxmoxapi

type LxcStatus string

const (
	LXC_STATUS_STOPPED LxcStatus = "stopped"
	LXC_STATUS_RUNNING LxcStatus = "running"
)
