GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-docs-dependencies
install-docs-dependencies:
	./scripts/install-docs-deps.sh

.phony: install-dependencies
install-dependencies: install-docs-dependencies
	go mod tidy

.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	./scripts/coverage.sh

.PHONY: generate-docs
generate-docs: install-docs-dependencies
	bash ./scripts/generate-docs.sh

.PHONY: preview-docs
preview-docs: install-docs-dependencies generate-docs
	bash ./scripts/preview-docs.sh

.PHONY: build
build: install-dependencies
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = amd64)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = 386)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go
	
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm64)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = arm64)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = amd64)
	CGO_ENABLED=0 GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -ldflags="-extldflags=-static" -o "bin/pve-api-wrapper-$($@GOOS)-$($@GOARCH)" ./cmd/pve_api_wrapper/pve_api_wrapper.go

