# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [unreleased]
### Changed
- `ProxmoxAPI.NewWithCredentials` method can now return an error.
- `ProxmoxAPI.New` and `ProxmoxAPI.NewWithCredentials` methods now sends a Get Version request to the remote proxmox api to check for valid credentials.


## [v0.1.0]

### Added
- Proxmox api token credentials support.
- Proxmox api version endpoint.

[unreleased]: https://github.com/iolave/go-proxmox/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.1.0
