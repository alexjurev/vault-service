VERSION ?= $(shell git describe --abbrev=4 --dirty --always --tags)
GOPATH ?= $(HOME)/go
BIN_DIR = $(GOPATH)/bin
TMPDIR ?= $(shell dirname $$(mktemp -u))
DOCKER_IMG="vault:develop"
LDFLAGS := -X main.version=$(VERSION)

# Go 1.13 uses Google's proxy by default. That's why we have to:
# - indicate packages that should not be accessed via proxy;
# - set default proxy to the company-managed one.
# https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules
export DOCKER_BUILDKIT = 1

# Project specific variables

PACKAGE = vault-service
APP_NAME ?= $(PACKAGE)

CONFIG_LOGGER_LEVEL ?= debug

# Main targets

all: test build
.DEFAULT_GOAL := all

.PHONY: build
build: ## Build the project binary
	go build \
		-ldflags '-X main.version=$(VERSION)' \
		-o build/vault-service \
		./cmd/main.go

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f Dockerfile .

run-img: build-img
	docker run \
    --rm \
	-it \
	-p 8040:8040 \
	-p 3110:3110 \
    $(DOCKER_IMG)

.PHONY: generate-server
generate-server: ## Generate rest from swagger
	@docker run --rm -v ${PWD}:${PWD} -w ${PWD} -u 1000:1000 quay.io/goswagger/swagger:v0.30.3 \
		generate server -f ./docs/swagger.yaml  \
		-C ./swagger-gen/default-server.yml \
		--template-dir ./swagger-gen/templates \
		--target ./pkg/infrastructure \
		--name ${PACKAGE}
