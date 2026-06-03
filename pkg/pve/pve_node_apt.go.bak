package pve

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/iolave/go-proxmox/pkg/helpers"
)

type PVENodeAPTService struct {
	api *Client
}

func newPVENodeAPTService(api *Client) *PVENodeAPTService {
	service := new(PVENodeAPTService)
	service.api = api
	return service
}

type NodeAPTIndex struct {
	ID string `json:"id"`
}

// GetIndex returns node's directory index for apt (Advanced Package Tool).
//
// GET /nodes/:node/apt accessible by all authenticated users.
func (s *PVENodeAPTService) GetIndex(node string) ([]NodeAPTIndex, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt")

	res := new([]NodeAPTIndex)
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// GetChangelog returns the changelog for a given pacakge name. If version is nil, the latest version available will be considered and otherwise, it will return the changelog found for the given version.
//
// GET /nodes/:node/apt/changelog requires the "Sys.Audit" permission.
func (s *PVENodeAPTService) GetChangelog(node, name string, version *string) (string, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/changelog")

	payload := &url.Values{}

	payload.Add("name", name)
	if version != nil {
		payload.Add("version", *version)
	}

	res := helpers.NewStr("")
	err := s.api.client.sendReq(method, path, payload, res)

	return *res, err
}

type APTRepoInfoError struct {
	Error string `json:"error"` // The error message.
	Path  string `json:"path"`  // Path to the problematic file.
}

type APTRepoInfoFileRepoOpt struct {
	Key    string
	Values []string
}

type APTRepoInfoFileRepo struct {
	Comment    string                   // Associated comment.
	Components []string                 // List of repository components
	Enabled    bool                     // Whether the repository is enabled or not.
	FileType   string                   // Format of the defining file ("list", "sources").
	Options    []APTRepoInfoFileRepoOpt // Additional options.
	Suites     []string                 // List of package distribuitions
	Types      []string                 // List of package types ("deb", "deb-src").
	URIs       []string                 // List of repository URIs.
}

type APTRepoInfoFile struct {
	Digest       []int                 `json:"digest"`       // Digest of the file as bytes.
	FileType     string                `json:"file-type"`    // Format of the file ("list", "sources").
	Path         string                `json:"path"`         // Path to the problematic file.
	Repositories []APTRepoInfoFileRepo `json:"repositories"` // The parsed repositories.
}

type APTRepoInfoInfos struct {
	Index    string `json:"index"`    // Index of the associated repository within the file.
	Kind     string `json:"kind"`     // Kind of the information (e.g. warning).
	Message  string `json:"message"`  // Information message.
	Path     string `json:"path"`     // Path to the associated file.
	Property string `json:"property"` // Property from which the info originates.
}

type APTRepoInfoStdRepo struct {
	Handle string `json:"handle"` // Handle to identify the repository.
	Name   string `json:"name"`   // Full name of the repository.
	Status *bool  `json:"status"` // Indicating enabled/disabled status, if the repository is configured.
}

type GetNodeAPTRepoInfo struct {
	Digest   string               `json:"digest"`         // Common digest of all files.
	Errors   []APTRepoInfoError   `json:"errors"`         // List of problematic repository files.
	Files    []APTRepoInfoFile    `json:"files"`          // List of parsed repository files.
	Infos    []APTRepoInfoInfos   `json:"infos"`          // Additional information/warnings for APT repositories.
	StdRepos []APTRepoInfoStdRepo `json:"standard-repos"` // List of standard repositories and their configuration status.
}

// GetRepoInfo returns APT repository information.
//
// GET /nodes/:node/apt/repositories requires the "Sys.Audit" permission.
func (s *PVENodeAPTService) GetRepoInfo(node string) (GetNodeAPTRepoInfo, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/repositories")

	res := new(GetNodeAPTRepoInfo)
	err := s.api.client.sendReq(method, path, nil, res)

	return *res, err
}

// SetRepoProps changes the properties of a repository (currently only allows enabling/disabling).
//   - "index": Index within the file (starting from 0).
//   - "node": Cluster node name.
//   - "filePath": Path to the containing file.
//   - "digest": Digest to detect modifications.
//   - "enabled": Whether the repository should be enabled or not.
//
// POST /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.
func (s *PVENodeAPTService) SetRepoProps(index int, node, filePath string, digest *string, enabled *bool) error {
	method := http.MethodPost
	path := path.Join("/nodes", node, "/apt/repositories")

	payload := &url.Values{}

	payload.Add("index", fmt.Sprintf("%d", index))
	payload.Add("path", filePath)
	if digest != nil {
		payload.Add("digest", *digest)
	}
	if enabled != nil {
		payload.Add("enabled", fmt.Sprintf("%d", helpers.BoolToInt(*enabled)))
	}

	err := s.api.client.sendReq(method, path, nil, nil)

	return err
}

// AddStdRepo adds a standard repository to the configuration.
//   - node: Cluster node name.
//   - handle: Handle that identifies a repository.
//   - digest: Digest to detect modifications.
//
// PUT /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.
func (s *PVENodeAPTService) AddStdRepo(node, handle string, digest *string) error {
	method := http.MethodPut
	path := path.Join("/nodes", node, "/apt/repositories")

	payload := &url.Values{}

	if digest != nil {
		payload.Add("digest", *digest)
	}

	err := s.api.client.sendReq(method, path, nil, nil)

	return err
}

// ListUpdates list available updates.
//
// TODO: [docs] lacks of response definition (map it).
//   - node: Cluster node name.
//
// GET /nodes/:node/apt/update requires the "Sys.Modify" permission.
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/update
func (s *PVENodeAPTService) ListUpdates(node string) (interface{}, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/update")

	var res *interface{}
	err := s.api.client.sendReq(method, path, nil, res)

	return res, err
}

// UpdateIndex this is used to resynchronize the package index files from their sources (apt-get update).
//
//   - node: Cluster node name.
//   - notify: Send notification about new packages.
//   - quiet: Only produces output suitable for logging, omitting progress indicators.
//
// POST /nodes/:node/apt/update requires the "Sys.Modify" permission.
func (s *PVENodeAPTService) UpdateIndex(node string, notify, quiet bool) (string, error) {
	method := http.MethodPost
	path := path.Join("/nodes", node, "/apt/update")

	res := ""
	err := s.api.client.sendReq(method, path, nil, &res)

	return res, err
}

// GetPVEInfo get package information for important Proxmox packages.
//
// TODO: [docs] lacks of response definition (map it).
//   - node: Cluster node name.
//
// GET /nodes/:node/apt/versions requires the "Sys.Audit" permission.
//
// [docs]: https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/apt/versions
func (s *PVENodeAPTService) GetPVEInfo(node string) (interface{}, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/versions")

	var res *interface{}
	err := s.api.client.sendReq(method, path, nil, res)

	return res, err
}
