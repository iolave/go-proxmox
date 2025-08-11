GOBIN ?= $$(go env GOPATH)/bin

.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	./scripts/coverage.sh

.PHONY: docs
docs:
	./scripts/generate-docs.sh

.PHONY: preview-docs
preview-docs: 
	./scripts/generate-docs.sh -p

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

