# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.7.1]
### PVE API client
#### Fixed
- Removed unused imports.

## [v0.7.0]
### PVE API wrapper
#### Added
- `service install` command to install a services that runs the server automatically.
- `POST /custom-api/v1/lxc/{id}/exec-async` endpoint.
- `GET  /custom-api/v1/cmd/{id}` endpoint.

#### Fixed
- custom handlers header content type.
- missing custom handlers write response call on errors.

### PVE API client
#### Added
- `POST   /custom-api/v1/lxc/{id}/exec-async` implementation as `PVE.LXC.ExecAsync`.
- `GET    /custom-api/v1/cmd/{id}` implementation as `PVE.LXC.GetCMDResult`.
- `GET    /nodes/{node}/lxc/{id}/config` implementation as `PVE.LXC.Update`.
- `POST   /nodes/{node}/lxc/{vmid}/clone` implementation as `PVE.LXC.Clone`.
- `POST   /nodes/{node}/lxc/{vmid}/template` implementation as `PVE.LXC.CreateTemplate`.

## [v0.6.1]
### PVE API client
#### Fixed
- Fixed app version.
- LXC features properties are now pointes.

## [v0.6.0]
### PVE API wrapper
#### Added
- `GET  /custom-api/v1/lxc/{id}/ip` endpoint.
- `POST /custom-api/v1/lxc/{id}/exec` endpoint.

#### Fixed
- pve api wrapper host option is now `--host`.
- pve api wrapper port option is now `--port`.

### PVE API client
#### Added
- `GET    /access/permissions` implementation as `PVE.Access.GetPermissions`.
- `GET    /cluster/resources` implementation as `PVE.Cluster.GetVMIDs`, `PVE.Cluster.GetRandomVMID` and `PVE.Cluster.IsVMIDAvailable`.
- `GET    /nodes/{node}/firewall/rules` implementation as `PVE.Node.Firewall.GetRule`.
- `POST   /nodes/{node}/firewall/rules` implementation as `PVE.Node.Firewall.NewRule`.
- `DELETE /nodes/{node}/firewall/rules/{pos}` implementation as `PVE.Node.Firewall.DeleteRule` and `PVE.Node.Firewall.DeleteRuleByPos`.
- `GET    /nodes/{node}/lxc/{id}/interfaces` implementation as `PVE.LXC.GetInterfaces` and `PVE.LXC.GetInterface`.
- `GET    /nodes/{node}/lxc/{id}/status/current` implementation as `PVE.LXC.GetStatus`.
- `POST   /nodes/{node}/lxc/{id}/status/reboot` implementation as `PVE.LXC.Reboot`.
- `POST   /nodes/{node}/lxc/{id}/status/resume` implementation as `PVE.LXC.Resume`.
- `POST   /nodes/{node}/lxc/{id}/status/shutdown` implementation as `PVE.LXC.Shutdown`.
- `POST   /nodes/{node}/lxc/{id}/status/start` implementation as `PVE.LXC.Start`.
- `POST   /nodes/{node}/lxc/{id}/status/stop` implementation as `PVE.LXC.Stop`.
- `POST   /custom-api/v1/lxc/{id}/exec` implementation as `PVE.LXC.Exec`.

#### Fixed
- `PVE.LXC.Create` hostname assignment.
- `PVE.LXC.Create` net assignment.
- `PVE.LXC.Create` multiple networks support.

#### Changed
- `PVE.LXC.Create` request's `Features` property is now a struct.

## [v0.5.0]
### PVE API wrapper
#### Added
- new http server that redirects http requests to the desired proxmox instance.
- proxmox error response wrapping into `pkg/errors.HTTPError`.  

### PVE API client
#### Added
- `/nodes/:node/lxc/:vmid` implementations.

#### Fixed
- pve api result assignment.
- `PVE.version` http method.
- `PVE.Lxc.GetAll` empty result.

## [v0.4.0]
### Changed
- Cluster endpoints implementations are now located under `PVE.Cluster` field.
- Cluster Firewall endpoints implementations are now located under `PVE.Cluster.Firewall` field.
- Node endpoints implementations are now located under `PVE.Node` field.
- Node Firewall endpoints implementations are now located under `PVE.Node.Firewall` field.
- Node APT endpoints implementations are now located under `PVE.Node.APT` field.
- Node Storage endpoints implementations are now located under `PVE.Node.Storage` field.
- LXC endpoints implementations are now located under `PVE.LXC` field.

## [v0.3.0]
### Added
- Node's APT endpoints implementations.

## [v0.2.1]
### Fixed
- `GetClusterFirewallAliases` http method is now get.
- `GetClusterFirewallIPSet` http method is now get.
- `GetClusterFirewallRules` url path is now correct.

## [v0.2.0]
### Added
- Helpers:
    - `BoolToInt` converts `true` to `1` and `false` to `0`.
    - `NewStr`, `NewInt` and `NewBool` methods that converts the primitive value to a pointer to it.
- Cloudflare Zero Trust support with Service Tokens.
- Proxmox Cluster:
    - `GetClusterFirewallAliases` retrieves all cluster firewall aliases.
    - `GetClusterFirewallAlias` retrieves cluster firewall alias by it's name.
    - `CreateClusterFirewallAlias` creates a cluster firewall IP or Network Alias.
    - `UpdateClusterFirewallAlias` updates a cluster firewall IP or Network alias.
    - `DeleteClusterFirewallAlias` removes a cluster firewall IP or Network alias.
    - `GetClusterFirewallIPSet` retrieves all cluster firewall IPSets.
    - `GetClusterFirewallRules` retrieves all cluster firewall rules.
- Proxmox nodes:
    - `GetAll` retrieves all nodes.
    - `Get` retrieves a single nodes.
    - `GetNodeRules` retrieves node's firewall rules.
    - `GetNodeRulesByPos` Retrieves a single node's firewall rule using rule's position (pos) as an index.
    - `ReadNodeLog` Retrieves node's firewall log entries.
    - `GetNodeDatastores` retrieves node's datastores info.
    - `GetNodeDatastoreContent` retrieves node's datastores info.
    - `DownloadISOToNodeDatastore` downloads an iso from an url into a node's datastore.
    - `DownloadVZTemplateToNodeDatastore` downloads a vztemplate from an url into a node's datastore.
- Proxmox LXC:
    - `GetLxcs` returns node's lxc index per node.
    - `CreateLxc` creates an LXC container and return useful information to interact with it after it's creation.

### Changed
- Main package was renamed from proxmoxapi to pve.
- `pve.NewEnvCreds` method can now return an error.
- `pve.New` and `pve.NewWithCredentials` methods now sends a Get Version request to the remote proxmox api to check for valid credentials.
- `pve.Credentials` now have a Set method that adds the authorization header to the given http request.


## [v0.1.0]

### Added
- Proxmox api token credentials support.
- Proxmox api version endpoint.

[unreleased]: https://github.com/iolave/go-proxmox/compare/v0.7.1...staging
[v0.7.1]: https://github.com/iolave/go-proxmox/releases/tag/v0.7.1
[v0.7.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.7.0
[v0.6.1]: https://github.com/iolave/go-proxmox/releases/tag/v0.6.1
[v0.6.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.6.0
[v0.5.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.5.0
[v0.4.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.4.0
[v0.3.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.3.0
[v0.2.1]: https://github.com/iolave/go-proxmox/releases/tag/v0.2.1
[v0.2.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.2.0
[v0.1.0]: https://github.com/iolave/go-proxmox/releases/tag/v0.1.0
