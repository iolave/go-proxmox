package pve

import (
	"net/http"
	"net/url"
	"path"
)

type GetNodeDatastoreResponse struct {
	Content      string   `json:"content"`       // Allowed storage content types.
	Storage      string   `json:"storage"`       // The storage identifier.
	Type         string   `json:"type"`          // Storage type.
	Active       *bool    `json:"active"`        // Set when storage is accessible.
	Available    *int     `json:"avail"`         // Available storage space in bytes.
	Enabled      *bool    `json:"enabled"`       // Set when storage is enabled (not disabled).
	Shared       *bool    `json:"shared"`        // Shared flag from storage configuration.
	TotalSpace   *int     `json:"total"`         // Total storage space in bytes.
	UsedSpace    *int     `json:"used"`          // Total storage space in bytes.
	UsedFraction *float64 `json:"used_fraction"` // Used fraction (used/total).
}

// GetNodeDatastores retrieves node's datastores info.
func (api *PVE) GetNodeDatastores(node string) ([]GetNodeDatastoreResponse, error) {
	path := path.Join("/nodes", node, "/storage")
	method := http.MethodGet

	res := &[]GetNodeDatastoreResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// TODO: Add missing verification property from [docs].
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content
type GetNodeDatastoreContentResponse struct {
	Format    string  `json:"format"`    // Format identifier ('raw', 'qcow2', 'subvol', 'iso', 'tgz' ...)
	Size      int     `json:"size"`      // Volume size in bytes.
	VolumeID  string  `json:"volid"`     // Volume identifier.
	CreatedAt *int    `json:"ctime"`     // Creation time (seconds since the UNIX Epoch).
	Encrypted *string `json:"encrypted"` // If whole backup is encrypted, value is the fingerprint or '1'  if encrypted. Only useful for the Proxmox Backup Server storage type.
	Notes     *string `json:"notes"`     // Optional notes. If they contain multiple lines, only the first one is returned here.
	Parent    *string `json:"parent"`    // Volume identifier of parent (for linked cloned).
	Protected *bool   `json:"protected"` // Protection status. Currently only supported for backups.
	Used      *int    `json:"used"`      // Used space. Please note that most storage plugins do not report anything useful here.
	VmID      *int    `json:"vmid"`      // Associated Owner VMID.
}

// GetNodeDatastoreContent retrieves node's datastores info.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content
func (api *PVE) GetNodeDatastoreContent(node, storageId string) ([]GetNodeDatastoreContentResponse, error) {
	path := path.Join("/nodes", node, "/storage", storageId, "/content")
	method := http.MethodGet

	res := &[]GetNodeDatastoreContentResponse{}
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// DownloadISOToNodeDatastore downloads an iso from an url into a node's datastore.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url
func (api *PVE) DownloadISOToNodeDatastore(node, storageId, fileName, URL string) error {
	path := path.Join("/nodes", node, "/storage", storageId, "/download-url")
	method := http.MethodPost

	payload := &url.Values{}
	payload.Add("url", URL)
	payload.Add("content", "iso")
	payload.Add("filename", fileName)

	err := api.httpClient.sendReq(method, path, payload, nil)
	return err
}

// DownloadVZTemplateToNodeDatastore downloads a vztemplate from an url into a node's datastore.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url
func (api *PVE) DownloadVZTemplateToNodeDatastore(node, storageId, fileName, URL string) error {
	path := path.Join("/nodes", node, "/storage", storageId, "/download-url")
	method := http.MethodPost

	payload := &url.Values{}
	payload.Add("url", URL)
	payload.Add("content", "vztmpl")
	payload.Add("filename", fileName)

	err := api.httpClient.sendReq(method, path, payload, nil)
	return err
}
