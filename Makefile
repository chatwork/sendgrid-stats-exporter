USER     := ada-u
REPO     := sendgrid-stats-exporter
GIT_TAG  := $(shell git tag --points-at HEAD)
GIT_HASH := $(shell git rev-parse HEAD)
VERSION  := $(shell if [ -n "$(GIT_TAG)" ]; then echo "$(GIT_TAG)"; else echo "$(GIT_HASH)"; fi)

DIST_DIR := $(shell if [ -n "$(GOOS)$(GOARCH)" ]; then echo "./dist/$(GOOS)-$(GOARCH)"; else echo "./dist"; fi)

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
DOCKER_BUILD_PLATFORMS ?= linux/amd64,linux/arm64
DOCKER_BUILDX_ARGS ?= --push

default: build

.PHONY: build
build:
	@echo "version: $(VERSION) hash: $(GIT_HASH) tag: $(GIT_TAG)"
	go build -ldflags "-s -w -X main.version=$(VERSION) -X main.gitCommit=$(GIT_HASH)" -o $(DIST_DIR)/exporter .

.PHONY: build-image
build-image:
	docker build -t chatwork/"$(REPO)" .
	docker tag chatwork/"$(REPO)":latest chatwork/"$(REPO)":"$(VERSION)"

.PHONY: push-image
push-image:
	docker push chatwork/"$(REPO)"

build-image-multi:
	docker buildx build -t chatwork/"$(REPO)":"$(VERSION)" --platform=$(DOCKER_BUILD_PLATFORMS) .

push-image-multi:
	docker buildx build $(DOCKER_BUILDX_ARGS) -t chatwork/"$(REPO)":"$(VERSION)" --platform=$(DOCKER_BUILD_PLATFORMS) .
