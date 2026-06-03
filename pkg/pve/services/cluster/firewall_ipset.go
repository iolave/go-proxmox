package cluster

import (
	"github.com/iolave/go-errors"
	"github.com/iolave/go-proxmox/internal/helpers"
	"github.com/iolave/go-proxmox/pkg/pve/internal/raw_api"
)

type IPSet struct {
	// IP set name.
	Name string `json:"name" validate:"required"`

	// Digest of the configuration. Used to prevent
	// concurrent modifications.
	Digest string `json:"digest"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// GetFWIPSets returns a list of all cluster firewall IPSet
// configured.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWIPSets() ([]IPSet, error) {
	res, err := rawapi.GETClusterFirewallIPSet(s.c)
	if err != nil {
		return nil, err
	}

	ipsets := make([]IPSet, len(res))
	for i, v := range res {
		ipsets[i] = IPSet{
			Name:    v.Name,
			Digest:  v.Digest,
			Comment: v.Comment,
		}
	}

	return ipsets, nil
}

// NewFWIPSet creates a new IPSet.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Modify"]]
func (s Service) NewFWIPSet(cfg IPSet) error {
	err := helpers.Validate.Struct(cfg)
	if err != nil {
		return errors.NewBadRequestError(
			"validation_error",
			err,
		)
	}

	err = rawapi.POSTClusterFirewallIPSet(s.c, rawapi.POSTClusterFirewallIPSetRequest{
		Name:    cfg.Name,
		Comment: cfg.Comment,
	})

	return err
}

type IPSetEntry struct {
	// Network/IP specification in CIDR format.
	CIDR string `json:"cidr"`

	// Digest of the configuration. Used to prevent
	// concurrent modifications.
	Digest string `json:"digest"`

	// TODO: add proper documentation (proxmox api does not
	// provide any)
	NoMatch bool `json:"noMatch"`

	// Descriptive comment.
	Comment string `json:"comment"`
}

// GetFWIPSetEntries returns the IPSet entries by it's name.
//
// Required proxmox permissions:
//
//	Check: ["perm","/",["Sys.Audit"]]
func (s Service) GetFWIPSetEntries(name string) ([]IPSetEntry, error) {
	if name == "" {
		return nil, errors.NewBadRequestError(
			"name cannot be empty",
			nil,
		)
	}

	res, err := rawapi.GETClusterFirewallIPSet_name(s.c, name)
	if err != nil {
		return nil, err
	}

	entries := make([]IPSetEntry, len(res))
	for i, v := range res {
		entries[i] = IPSetEntry{
			CIDR:    v.CIDR,
			Digest:  v.Digest,
			NoMatch: helpers.IntToBool(helpers.IntFromPtr(v.NoMatch)),
			Comment: v.Comment,
		}
	}

	return entries, nil
}
