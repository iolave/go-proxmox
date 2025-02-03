The pve api wrapper is an http server ment to be installed on the proxmox host server that wraps the proxmox api and adds missing features to it.

_pve-api-wrapper can definitely be installed somewhere else but some functionalities might not work._

## Features
- Proper error responses on failed requests,
- lxc exec command endpoint (not available yet).

## Installation
### Latest release
The installation script installs the `pve-api-wrapper` binary into `/usr/local/bin`.

```bash
curl https://raw.githubusercontent.com/iolave/go-proxmox/refs/tags/latest/scripts/install.sh | bash
```

_Inspect the installation script code [here]._
### Build from source
```
# Download the source code
curl -L -o go-proxmox-latest.tar.gz https://github.com/iolave/go-proxmox/archive/refs/tags/latest.tar.gz
tar -xvzf go-proxmox-latest.tar.gz

# Build
cd go-proxmox-latest
# This step will create a folder `bin` with the pve-api-wrapper binaries.
make build
```

## Usage
```bash
pve-api-wrapper [--version] [--pve-host PVE-HOST] [--pve-port PVE-PORT] [--host HOST] [--port PORT] [--crt CRT] [--key KEY]
```
### Options
- `--version` displays the program version.
- `--pve-host` proxmox virtual environment host (env:`PVE_HOST`, default:`localhost`).
- `--pve-port` proxmox virtual environment port (env:`PVE_PORT`, default:`8006`).
- `--host` api wrapper host (env:`WRAPPER_HOST`, default:`localhost`).
- `--port` api wrapper port (env:`WRAPPER_PORT`, default:`8443`).
- `--crt` api wrapper tls crt path (default:`/etc/pve/local/pve-ssl.pem`).
- `--key` api wrapper tls key path (default:`/etc/pve/local/pve-ssl.key`).

[here]: https://github.com/iolave/go-proxmox/blob/latest/scripts/install.sh
<!--
    TODO: host the shell script within the docs https://github.com/squidfunk/mkdocs-material/discussions/3458
-->
