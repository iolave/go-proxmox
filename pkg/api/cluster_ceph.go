package api

import "net/http"

// ClusterCephGetStatus Get ceph status.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit","Datastore.Audit"],"any",1]
func (s API) ClusterCephGetStatus() (res any, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/ceph/status",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterCephGetMetadataRequest struct {
	// Only return metadata for this config.
	//
	// 	all | versions
	Scope string `in:"query=scope;omitempty"`
}

// ClusterCephGetMetadata Get ceph metadata.
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit","Datastore.Audit"],"any",1]
func (s API) ClusterCephGetMetadata(req ClusterCephGetMetadataRequest) (res struct {
	MDS map[string]struct {
		Addr             string `json:"addr"`
		CephRelease      string `json:"ceph_release"`
		CephVersion      string `json:"ceph_version"`
		CephVersionShort string `json:"ceph_version_short"`
		Hostname         string `json:"hostname"`
		MemSwapKB        int    `json:"mem_swap_kb"`
		MemTotalKB       int    `json:"mem_total_kb"`
		Name             string `json:"name"`
	} `json:"mds"`
	MGR map[string]struct {
		Addr             string `json:"addr"`
		CephRelease      string `json:"ceph_release"`
		CephVersion      string `json:"ceph_version"`
		CephVersionShort string `json:"ceph_version_short"`
		Hostname         string `json:"hostname"`
		MemSwapKB        int    `json:"mem_swap_kb"`
		MemTotalKB       int    `json:"mem_total_kb"`
		Name             string `json:"name"`
	} `json:"mgr"`
	MON map[string]struct {
		Addrs            string `json:"addrs"`
		CephRelease      string `json:"ceph_release"`
		CephVersion      string `json:"ceph_version"`
		CephVersionShort string `json:"ceph_version_short"`
		Hostname         string `json:"hostname"`
		MemSwapKB        int    `json:"mem_swap_kb"`
		MemTotalKB       int    `json:"mem_total_kb"`
		Name             string `json:"name"`
	} `json:"mon"`
	Node map[string]struct {
		BuildCommit string `json:"buildcommit"`
		Version     struct {
			Parts  []string `json:"parts"`
			String string   `json:"str"`
		} `json:"version"`
	} `json:"node"`
	OSD map[string]struct {
		BackAddr         string `json:"addrs"`
		CephRelease      string `json:"ceph_release"`
		CephVersion      string `json:"ceph_version"`
		CephVersionShort string `json:"ceph_version_short"`
		DeviceID         string `json:"device_id"`
		FrontAddr        string `json:"front_addr"`
		Hostname         string `json:"hostname"`
		ID               int    `json:"id"`
		MemSwapKB        int    `json:"mem_swap_kb"`
		MemTotalKB       int    `json:"mem_total_kb"`
		OSDData          string `json:"osd_data"`
		OSDObjectStore   string `json:"osd_objectstore"`
	} `json:"osd"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ceph/metadata",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

// ClusterCephFlagsGetStatus get the status of all ceph flags
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterCephFlagsGetAllStatus() (res []struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Value       int    `json:"value"`
}, err error) {
	err = s.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/ceph/flags",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

// ClusterCephFlagsGetStatus Get the status of a specific ceph flag.
//
// Parameters:
//
// - flag is the ceph flag name.
//
//	nobackfill | nodeep-scrub | nodown | noin | noout | norebalance | norecover | noscrub | notieragent | noup | pause
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s API) ClusterCephFlagsGetStatus(flag string) (res int, err error) {
	req := struct {
		Flag string `in:"path=flag;nonzero"`
	}{
		Flag: flag,
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ceph/flags/{flag}",
		Method:  http.MethodGet,
		Result:  &res,
		Payload: &req,
	})

	return res, err
}

// ClusterCephFlagsSetStatus Set the status of a specific ceph flag.
//
// Parameters:
//
//   - flag is the ceph flag name.
//
//     nobackfill | nodeep-scrub | nodown | noin | noout | norebalance | norecover | noscrub | notieragent | noup | pause
//
//   - value is the ceph flag value.
//
//     0 | 1
//
// Required permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s API) ClusterCephFlagsSetStatus(flag string, value int) (err error) {
	req := struct {
		Flag  string `in:"path=flag;nonzero"`
		Value int    `in:"query=value"`
	}{
		Flag:  flag,
		Value: value,
	}

	err = s.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/ceph/flags/{flag}",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}
