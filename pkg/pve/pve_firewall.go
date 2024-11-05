package pve

// Promox firewall available log levels.
type FirewallLogLevel string

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

// Proxmox firewall log entry.
type FirewallLogEntry struct {
	Id   int    `json:"n"`
	Text string `json:"t"`
}
