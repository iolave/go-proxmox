package apiclient

import "net/http"

func (c HTTPClient) ClusterGetOptions() (res struct {
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
	err = c.SendPVERequest(PVERequest{
		Path:   "/api2/json/cluster/options",
		Method: http.MethodGet,
		Result: &res,
	})

	return res, err
}

type ClusterPutOptionsRequest struct {
	BWLimit string `in:"query=bwlimit;omitempty"`

	// ¡NOT AVAILABLE IN API!
	// ConsentText       string `in:"form;omitempty,form=consent-text"`
	Console           string `in:"query=console;omitempty"`
	CRS               string `in:"query=crs;omitempty"`
	Delete            string `in:"query=delete;omitempty"`
	Description       string `in:"query=description;omitempty"`
	EmailFrom         string `in:"query=email_from;omitempty"`
	Fencing           string `in:"query=fencing;omitempty"`
	HA                string `in:"query=ha;omitempty"`
	HTTPProxy         string `in:"query=http_proxy;omitempty"`
	Keyboard          string `in:"query=keyboard;omitempty"`
	Language          string `in:"query=language;omitempty"`
	MacPrefix         string `in:"query=mac_prefix;omitempty"`
	MaxWorkers        int    `in:"query=max_workers;omitempty"`
	Migration         string `in:"query=migration;omitempty"`
	MigrationUnsecure string `in:"query=migration_unsecure;omitempty"`
	NextID            string `in:"query=next-id;omitempty"`
	Notify            string `in:"query=notify;omitempty"`
	RegisteredTags    string `in:"query=registered-tags;omitempty"`
	// ¡NOT AVAILABLE IN API!
	// Replication       string `in:"form;omitempty,form=replication"`
	TagStyle      string `in:"query=tag-style;omitempty"`
	U2F           string `in:"query=u2f;omitempty"`
	UserTagAccess string `in:"query=user-tag-access;omitempty"`
	WebAuthn      string `in:"query=webauthn;omitempty"`
}

func (c HTTPClient) ClusterPutOptions(req ClusterPutOptionsRequest) (err error) {
	err = c.SendPVERequest(PVERequest{
		Path:    "/api2/json/cluster/options",
		Method:  http.MethodPut,
		Payload: &req,
	})

	return err
}
