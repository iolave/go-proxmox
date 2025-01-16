The pve api wrapper is an http server ment to be installed on the proxmox host server (but can definitely be installed somewhere else) that wraps the proxmox api and adds missing features to it.

## Features
- Proper error responses on failed requests,
- lxc exec command endpoint (not available yet).

## Installation
### Latest release
The installation script installs the `pve-api-wrapper` binary into `/usr/local/bin`.

```bash
curl https://raw.githubusercontent.com/iolave/go-proxmox/refs/tags/latest/scripts/install.sh | sh
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


## Serve
```bash
pve-api-wrapper
```

[here]: https://github.com/iolave/go-proxmox/blob/latest/scripts/install.sh
<!--
    TODO: host the shell script within the docs https://github.com/squidfunk/mkdocs-material/discussions/3458
-->
