# configure make
SHELL := /bin/bash
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables

# set directories for artifacts
BIN_DIR := $(CURDIR)/bin
DIST_DIR := $(CURDIR)/dist

# evaluate variables for versioning
COMMIT_SHA1 ?=
RELEASE_TAG ?=
VERSION_TAG ?=  $(if $(RELEASE_TAG),$(RELEASE_TAG),$(if $(COMMIT_SHA1),build-$(COMMIT_SHA1),user-$(USER)))

# configure the go toolchain
GOTEST_RACE ?= true
LDFLAGS_META := github.com/olliebun/qrcode/internal/meta
GOOS ?= linux
GOARCH ?= amd64

.DEFAULT_GOAL := help
.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 | "sort"}' $(MAKEFILE_LIST)

.PHONY: test
test:
	go test -count=1 -race=$(GOTEST_RACE) ./...

.PHONY: build
build:
	CGO_ENABLED="0" go build \
		-ldflags="-X '$(LDFLAGS_META).Version=$(VERSION_TAG)'" \
		-o="$(BIN_DIR)/qrcode" -trimpath

.PHONY: dist
dist:
dist: build
	@mkdir -p "$(DIST_DIR)"
	tar \
		--create --gzip \
		--directory "$(BIN_DIR)" \
		--file "$(DIST_DIR)/qrcode-$(RELEASE_TAG)-$(GOOS)-$(GOARCH).tgz" \
		qrcode

.PHONY: publish
publish:
	gh release upload \
		--clobber "$(RELEASE_TAG)" \
		"$(DIST_DIR)/qrcode-$(RELEASE_TAG)-$(GOOS)-$(GOARCH).tgz"
