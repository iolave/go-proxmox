# go-proxmox

> [!WARNING]
> All versions released prior to `v1.0.0` are to be considered [breaking changes](https://semver.org/#how-do-i-know-when-to-release-100)

A fully-featured set of tools for Proxmox Virtual Environment.
This repository provides three layers:

- Low-level 1:1 API implementation: minimal abstractions, closely matching the upstream Proxmox REST API. To see a list of all implemented api endpoints and more information, click [here].
- [API Wrapper] (binary): exposes additional features and allows original api calls to be passed through.
- High-level, ergonomic client – a set of well-documented Go interfaces that smooths over quirks of the raw API and integrates extensions provided by the wrapper.

[here]: https://go-proxmox.iolave.com/go-client/
[API Wrapper]: https://go-proxmox.iolave.com/api-wrapper/

## Features
### 1:1 Proxmox API Implementation
- Direct mapping of Proxmox endpoints, parameters, and responses.
- As close to upstream behavior as possible.
- Useful for users who need full control or raw input/output.
- Unclear or inconsistent API behaviors are intentionally preserved (mirrors the original).

### API Wrapper
Small binary that exposes:

- Original API.
- Custom API with missing functionality from the native one.

It is designed to be deployed on a Proxmox node.

### High-level Client
- Human-friendly Go API.
- Strong typing and validation where the raw API is ambiguous.
- Includes wrapper-only features

## Project Structure
```text
.
├── pkg/                  # Packages that can be imported by other projects
│   │
│   ├── api/              # Proxmox API 1:1 implementation
│   ├── helpers/          # General purpose helpers
│   └── pve/              # High-level Proxmox API client
│
├── cmd/                  # Command line tools
│   │
│   ├── pve-api-wrapper/  # Proxmox API wrapper
│   └── gomarkdoc/        # Documentation generator
│
├── internal/             # API Wrapper internals
├── scripts/              # Shell scripts, mostly for development and Makefile
├── docs/                 # MKDocs documentation
├── README.md             # This file
├── Makefile              # Makefile for building, testing and documentation generation
├── LICENSE               # License
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
└── mkdocs.yml            # MKDocs configuration
```

## Documentation
- The go client documentation is available [here].
- The API wrapper documentation is available [here](https://go-proxmox.iolave.com/api-wrapper/).

## Contributing
Contributions are welcome! 

Please:

- Open an issue before large changes.
- Keep PRs focused and well‑scoped.
- Include relevant tests.

[here]: https://go-proxmox.iolave.com/go-client/
[API Wrapper]: https://go-proxmox.iolave.com/api-wrapper/
