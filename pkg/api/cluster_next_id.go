package api

import (
	"net/http"
	"strconv"

	"github.com/iolave/go-errors"
)

type ClusterGetNextIDRequest struct {
	VMID int `in:"query=vmid;omitempty"`
}

// ClusterGetNextID Get next free VMID. Pass a VMID
// to assert that its free (at time of check).
func (c API) ClusterGetNextID(req ClusterGetNextIDRequest) (int, error) {
	res := ""

	err := c.SendPVERequest(PVERequest{
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
