package rawapi

import (
	"net/http"

	"github.com/iolave/go-proxmox/pkg/pve"
)

type ClusterSDNApplyRequest struct {
	// The token for unlocking the global SDN configuration
	LockToken string `in:"query=lock-token;omitempty"`

	// When lock-token has been provided and configuration successfully commited, release the lock automatically afterwards
	ReleaseLock *int `in:"query=release-lock;omitempty"`
}

// ClusterSDNApply Apply sdn controller changes && reload.
//
// Required permissions:
//
//	Check: ["perm","/sdn",["SDN.Allocate"]]
func ClusterSDNApply(c *pve.Client, req ClusterSDNApplyRequest) (res string, err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn",
		Method:  http.MethodPost,
		Payload: &req,
		Result:  &res,
	})

	return res, err
}

// ClusterSDNLock Acquire global lock for SDN configuration
//
// Parameters:
//
//   - allowPending (0|1): if true, allow acquiring lock even though there are pending changes
//
// Required permissions:
//
//	Check: ["perm","/sdn",["SDN.Allocate"]]
func ClusterSDNLock(c *pve.Client, allowPending int) (res string, err error) {
	err = c.Do(pve.Request{
		Path:   "/api2/json/cluster/sdn/lock",
		Method: http.MethodPost,
		Result: &res,
	})

	return res, err
}

// ClusterSDNUnlock Release global lock for SDN configuration
//
// Parameters:
//
//   - force (0|1): if true, allow releasing lock without providing the token.
//   - lockToken: the token for unlocking the global SDN configuration.
//
// Required permissions:
//
//	Check: ["perm","/sdn",["SDN.Allocate"]]
func ClusterSDNUnlock(c *pve.Client, force int, lockToken string) (err error) {
	req := struct {
		Force     int    `in:"query=force"`
		LockToken string `in:"query=lock-token"`
	}{
		Force:     force,
		LockToken: lockToken,
	}

	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/lock",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}

// ClusterSDNRollback Rollback pending changes to SDN configuration
//
// Required permissions:
//
//	Check: ["perm","/sdn",["SDN.Allocate"]]
func ClusterSDNRollback(c *pve.Client, req ClusterSDNApplyRequest) (err error) {
	err = c.Do(pve.Request{
		Path:    "/api2/json/cluster/sdn/rollback",
		Method:  http.MethodPost,
		Payload: &req,
	})

	return err
}
