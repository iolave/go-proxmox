package access

import rawapi "github.com/iolave/go-proxmox/pkg/pve/internal/raw_api"

// Privilege is the type of privileges that can be granted to a pve-api path.
type Privilege = string

// Privilege
const (
	PRIVILEGE_DATASTORE_ALLOCATE_TEMPLATE Privilege = "Datastore.AllocateTemplate"
	PRIVILEGE_DATASTORE_ALLOCATE          Privilege = "Datastore.Allocate"
	PRIVILEGE_DATASTORE_ALLOCATE_SPACE    Privilege = "Datastore.AllocateSpace"
	PRIVILEGE_DATASTORE_AUDIT             Privilege = "Datastore.Audit"
	PRIVILEGE_GROUP_ALLOCATE              Privilege = "Group.Allocate"
	PRIVILEGE_MAPPING_MODIFY              Privilege = "Mapping.Modify"
	PRIVILEGE_MAPPING_AUDIT               Privilege = "Mapping.Audit"
	PRIVILEGE_MAPPING_USE                 Privilege = "Mapping.Use"
	PRIVILEGE_PERMISSIONS_MODIFY          Privilege = "Permissions.Modify"
	PRIVILEGE_POOL_ALLOCATE               Privilege = "Pool.Allocate"
	PRIVILEGE_POOL_AUDIT                  Privilege = "Pool.Audit"
	PRIVILEGE_REALM_ALLOCATE              Privilege = "Realm.Allocate"
	PRIVILEGE_REALM_ALLOCATE_USER         Privilege = "Realm.AllocateUser"
	PRIVILEGE_SYS_MODIFY                  Privilege = "Sys.Modify"
	PRIVILEGE_SYS_CONSOLE                 Privilege = "Sys.Console"
	PRIVILEGE_SYS_POWERMGMT               Privilege = "Sys.PowerMgmt"
	PRIVILEGE_SYS_SYSLOG                  Privilege = "Sys.Syslog"
	PRIVILEGE_SYS_INCOMING                Privilege = "Sys.Incoming"
	PRIVILEGE_SYS_AUDIT                   Privilege = "Sys.Audit"
	PRIVILEGE_SYS_ACCESS_NETWORK          Privilege = "Sys.AccessNetwork"
	PRIVILEGE_SDN_ALLOCATE                Privilege = "SDN.Allocate"
	PRIVILEGE_SDN_AUDIT                   Privilege = "SDN.Audit"
	PRIVILEGE_SDN_USE                     Privilege = "SDN.Use"
	PRIVILEGE_USER_MODIFY                 Privilege = "User.Modify"
	PRIVILEGE_VM_CONFIG_MEMORY            Privilege = "VM.Config.Memory"
	PRIVILEGE_VM_AUDIT                    Privilege = "VM.Audit"
	PRIVILEGE_VM_MIGRATE                  Privilege = "VM.Migrate"
	PRIVILEGE_VM_CONFIG_CDROM             Privilege = "VM.Config.CDROM"
	PRIVILEGE_VM_BACKUP                   Privilege = "VM.Backup"
	PRIVILEGE_VM_CONFIG_DISK              Privilege = "VM.Config.Disk"
	PRIVILEGE_VM_ALLOCATE                 Privilege = "VM.Allocate"
	PRIVILEGE_VM_SNAPSHOT_ROLLBACK        Privilege = "VM.Snapshot.Rollback"
	PRIVILEGE_VM_CONFIG_CLOUDINIT         Privilege = "VM.Config.Cloudinit"
	PRIVILEGE_VM_SNAPSHOT                 Privilege = "VM.Snapshot"
	PRIVILEGE_VM_CONFIG_CPU               Privilege = "VM.Config.CPU"
	PRIVILEGE_VM_CLONE                    Privilege = "VM.Clone"
	PRIVILEGE_VM_CONFIG_HWTYPE            Privilege = "VM.Config.HWType"
	PRIVILEGE_VM_MONITOR                  Privilege = "VM.Monitor"
	PRIVILEGE_VM_CONFIG_NETWORK           Privilege = "VM.Config.Network"
	PRIVILEGE_VM_POWERMGMT                Privilege = "VM.PowerMgmt"
	PRIVILEGE_VM_CONSOLE                  Privilege = "VM.Console"
	PRIVILEGE_VM_CONFIG_OPTIONS           Privilege = "VM.Config.Options"
)

// GetPermissionsResponse is a map of privileges to a map of paths to
// whether the privilege is granted or not.
//
// In example:
//
//	{
//		"VM.Config.Memory": {
//			"/": true,
//			"/vms/10000": true,
//			"/vms/10001": false,
//		},
//		"VM.Audit": {
//			"/": true,
//			"/vms/10000": true,
//			"/vms/10001": false,
//		},
//	}
type GetPermisionsResponse map[Privilege]map[string]bool

// GetPermissions retrieve effective permissions of the current authenticated user.
//
// Required proxmox permissions:
//
//	GET /access/permissions: each user/token is allowed to dump their own
//	permissions (or that of owned tokens).
func (s Service) GetPermissions() (
	res GetPermisionsResponse,
	err error,
) {
	perms, err := rawapi.GETAccessPermissions(s.c, rawapi.GETAccessPermisionsRequest{})
	if err != nil {
		return res, err
	}

	for path, perms := range perms {
		for perm, intEnabled := range perms {
			enabled := false
			if intEnabled == 1 {
				enabled = true
			}

			res[Privilege(perm)][path] = enabled
		}
	}

	return res, nil
}
