VENV="${TMPDIR}/venv/go-proxmox"

source "${VENV}/bin/activate"
rm -rf ./docs/go-client/pkg
mkdir -p ./docs/go-client/pkg/
go run ./cmd/gomarkdoc/main.go
swag init -g ./cmd/pve_api_wrapper/pve_api_wrapper.go --parseInternal
mkdocs build
