GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage install-docs-dependencies

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

install-docs-dependencies:
	./scripts/install-docs-deps.sh

install-dependencies: install-docs-dependencies
	go mod tidy


.PHONY: coverage-check
coverage-check: install-go-test-coverage
	go test -v ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

coverage-report: 
	go tool cover -html=cover.out -o=cover.html

coverage:
	$(MAKE) $(MAKEFLAGS) coverage-check; rc=$$? \
        ; $(MAKE) $(MAKEFLAGS) coverage-report \
        ; exit $$rc

generate-docs: install-docs-dependencies
	bash ./scripts/generate-docs.sh

preview-docs: install-docs-dependencies generate-docs
	bash ./scripts/preview-docs.sh

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
test:
	go test -v ./...

