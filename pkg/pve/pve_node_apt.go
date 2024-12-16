package pve

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/iolave/go-proxmox/pkg/helpers"
)

type NodeAPTIndex struct {
	ID string `json:"id"`
}

// GetNodeAPTIndex returns node's directory index for apt (Advanced Package Tool).
//
// GET /nodes/:node/apt accessible by all authenticated users.
func (api *PVE) GetNodeAPTIndex(node string) ([]NodeAPTIndex, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt")

	res := new([]NodeAPTIndex)
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// GetNodeAPTChangelog returns the changelog for a given pacakge name. If version is nil, the latest version available will be considered and otherwise, it will return the changelog found for the given version.
//
// GET /nodes/:node/apt/changelog requires the "Sys.Audit" permission.
func (api *PVE) GetNodeAPTChangelog(node, name string, version *string) (string, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/changelog")

	payload := &url.Values{}

	payload.Add("name", name)
	if version != nil {
		payload.Add("version", *version)
	}

	res := helpers.NewStr("")
	err := api.httpClient.sendReq(method, path, payload, res)

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

// GetNodeAPTRepoInfo returns APT repository information.
//
// GET /nodes/:node/apt/repositories requires the "Sys.Audit" permission.
func (api *PVE) GetNodeAPTRepoInfo(node string) (GetNodeAPTRepoInfo, error) {
	method := http.MethodGet
	path := path.Join("/nodes", node, "/apt/repositories")

	res := new(GetNodeAPTRepoInfo)
	err := api.httpClient.sendReq(method, path, nil, res)

	return *res, err
}

// SetNodeAPTRepoProps changes the properties of a repository (currently only allows enabling/disabling).
//   - "index": Index within the file (starting from 0).
//   - "node": Cluster node name.
//   - "filePath": Path to the containing file.
//   - "digest": Digest to detect modifications.
//   - "enabled": Whether the repository should be enabled or not.
//
// POST /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.
func (api *PVE) SetNodeAPTRepoProps(index int, node, filePath string, digest *string, enabled *bool) error {
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

	err := api.httpClient.sendReq(method, path, nil, nil)

	return err
}

// AddNodeAPTStdRepo adds a standard repository to the configuration.
//   - node: Cluster node name.
//   - handle: Handle that identifies a repository.
//   - digest: Digest to detect modifications.
//
// PUT /nodes/:node/apt/repositoreies requires the "Sys.Modify" permission.
func (api *PVE) AddNodeAPTStdRepo(node, handle string, digest *string) error {
	method := http.MethodPut
	path := path.Join("/nodes", node, "/apt/repositories")

	payload := &url.Values{}

	if digest != nil {
		payload.Add("digest", *digest)
	}

	err := api.httpClient.sendReq(method, path, nil, nil)

	return err
}
