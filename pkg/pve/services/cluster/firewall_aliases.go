package cluster

import (
	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/internal/helpers"
	rawapi "github.com/iolave/go-proxmox/pkg/pve/internal/raw_api"
)

type GetFWAliasResponse struct {
	// Alias name/id.
	Name string `json:"name"`

	// Network/IP specification in CIDR format.
	CIDR string `json:"cidr"`

	// Prevent changes if current configuration
	// file has a different digest. This can be
	// used to prevent concurrent modifications.
	Digest string `json:"digest"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// GetFWAlias returns a cluster firewall alias by it's
// name/id.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWAlias(name string) (GetFWAliasResponse, error) {
	res, err := rawapi.GETClusterFirewallAliases_name(s.c, name)
	if err != nil {
		return GetFWAliasResponse{}, err
	}

	return GetFWAliasResponse{
		Name:    res.Name,
		CIDR:    res.CIDR,
		Digest:  res.Digest,
		Comment: res.Comment,
	}, nil
}

// GetFWAliases returns a list of all cluster
// firewall aliases in the cluster.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWAliases() ([]GetFWAliasResponse, error) {
	res, err := rawapi.GETClusterFirewallAliases(s.c)
	if err != nil {
		return nil, err
	}

	mapFn := func(v rawapi.GETClusterFirewallAliasesResponse) GetFWAliasResponse {
		return GetFWAliasResponse{
			Name:    v.Name,
			CIDR:    v.CIDR,
			Digest:  v.Digest,
			Comment: v.Comment,
		}
	}

	return helpers.Map(res, mapFn), nil
}

type NewFWAliasRequest struct {
	// Alias name/id. It must be unique
	// and non-empty.
	Name string `json:"name" validate:"required"`

	// Network/IP specification in CIDR format.
	CIDR string `json:"cidr" validate:"required"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// NewFWAlias creates a new cluster firewall alias.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s Service) NewFWAlias(req NewFWAliasRequest) error {
	err := helpers.Validate.Struct(req)
	if err != nil {
		return errors.NewBadRequestError(
			"validation_error",
			err,
		)
	}

	err = rawapi.POSTClusterFirewallAliases(s.c, rawapi.POSTClusterFirewallAliasesRequest{
		CIDR:    req.CIDR,
		Name:    req.Name,
		Comment: req.Comment,
	})
	return err
}

type UpdateFWAliasRequest struct {
	// Alias name/id.
	Name string `json:"name" validate:"required"`

	// Network/IP specification in CIDR format.
	CIDR *string `json:"cidr"`

	// Prevent changes if current configuration
	// file has a different digest. This can be
	// used to prevent concurrent modifications.
	Digest *string `json:"digest"`

	// Descriptive comment.
	Comment *string `json:"comment"`

	// If specified, the alias will be renamed
	// to the given name/id.
	Rename *string `json:"rename"`
}

// UpdateFWAlias updates an existing cluster firewall alias.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s Service) UpdateFWAlias(req UpdateFWAliasRequest) error {
	err := helpers.Validate.Struct(req)
	if err != nil {
		return errors.NewBadRequestError(
			"validation_error",
			err,
		)
	}

	err = rawapi.PUTClusterFirewallAliases(s.c, rawapi.PUTClusterFirewallAliasesRequest{
		Name:    req.Name,
		CIDR:    helpers.StringFromPtr(req.CIDR),
		Comment: helpers.StringFromPtr(req.Comment),
		Digest:  helpers.StringFromPtr(req.Digest),
		Rename:  helpers.StringFromPtr(req.Rename),
	})
	return err
}
