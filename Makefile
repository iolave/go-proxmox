GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage install-docs-dependencies

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

install-docs-dependencies:
	./scripts/install-docs-deps.sh

.PHONY: coverage-check
coverage-check: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

coverage-report: 
	go tool cover -html=cover.out -o=cover.html

coverage:
	$(MAKE) $(MAKEFLAGS) coverage-check; rc=$$? \
        ; $(MAKE) $(MAKEFLAGS) coverage-report \
        ; exit $$rc

generate-docs: install-docs-dependencies
	source /tmp/venv/go-proxmox/bin/activate; \
	go run ./cmd/gomarkdoc/main.go; \
	mkdocs build

preview-docs: install-docs-dependencies
	source /tmp/venv/go-proxmox/bin/activate; \
	mkdocs serve

build:
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = amd64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = 386)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go
	
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = arm64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = amd64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/proxmox-cli-$($@GOOS)-$($@GOARCH)" cmd/proxmox_cli/main.go
test:
	go test -v ./...

