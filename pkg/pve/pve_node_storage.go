package pve

import (
	"net/http"
	"net/url"
	"path"
)

type PVENodeStorageService struct {
	api *PVE
}

func newPVENodeStorageService(api *PVE) *PVENodeStorageService {
	service := new(PVENodeStorageService)
	service.api = api
	return service
}

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

// GetDatastores retrieves node's datastores info.
func (s *PVENodeStorageService) GetDatastores(node string) ([]GetNodeDatastoreResponse, error) {
	path := path.Join("/nodes", node, "/storage")
	method := http.MethodGet

	res := &[]GetNodeDatastoreResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

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

// GetDatastoreContent retrieves node's datastores info.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content
func (s *PVENodeStorageService) GetDatastoreContent(node, storageId string) ([]GetNodeDatastoreContentResponse, error) {
	path := path.Join("/nodes", node, "/storage", storageId, "/content")
	method := http.MethodGet

	res := &[]GetNodeDatastoreContentResponse{}
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// DownloadISOToDatastore downloads an iso from an url into a node's datastore.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url
func (s *PVENodeStorageService) DownloadISOToDatastore(node, storageId, fileName, URL string) error {
	path := path.Join("/nodes", node, "/storage", storageId, "/download-url")
	method := http.MethodPost

	payload := &url.Values{}
	payload.Add("url", URL)
	payload.Add("content", "iso")
	payload.Add("filename", fileName)

	err := s.api.client.sendReq(method, path, payload, nil)
	return err
}

// DownloadVZTemplateToDatastore downloads a vztemplate from an url into a node's datastore.
//
// TODO: Add optional [parameters].
//
// [parameters]: https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/download-url
func (s *PVENodeStorageService) DownloadVZTemplateToDatastore(node, storageId, fileName, URL string) error {
	path := path.Join("/nodes", node, "/storage", storageId, "/download-url")
	method := http.MethodPost

	payload := &url.Values{}
	payload.Add("url", URL)
	payload.Add("content", "vztmpl")
	payload.Add("filename", fileName)

	err := s.api.client.sendReq(method, path, payload, nil)
	return err
}
