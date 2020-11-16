USER     := ada-u
REPO     := sendgrid-stats-exporter
GIT_TAG  := $(shell git tag --points-at HEAD)
GIT_HASH := $(shell git rev-parse HEAD)
VERSION  := $(shell if [ -n "$(GIT_TAG)" ]; then echo "$(GIT_TAG)"; else echo "$(GIT_HASH)"; fi)

DIST_DIR := $(shell if [ -n "$(GOOS)$(GOARCH)" ]; then echo "./dist/$(GOOS)-$(GOARCH)"; else echo "./dist"; fi)

default: build

.PHONY: build
build:
	go build -ldflags "-s -w -X main.version=$(VERSION) -X main.gitCommit=$(GIT_HASH)" -o $(DIST_DIR)/exporter .

.PHONY: build-image
build-image:
		docker build -t chatwork/"$(REPO)" .;
    	@version=$$(docker inspect -f {{.Config.Labels.version}} chatwork/"$(REPO)"); \
    		if [ -n "$$version" ]; then \
    			docker tag chatwork/"$(REPO)":latest chatwork/"$(REPO)":$$version; \
    		fi

.PHONY: push-image
push-image:
	@version=$$(docker inspect -f {{.Config.Labels.version}} chatwork/"$(REPO)":latest); \
		if docker inspect --format='{{index .RepoDigests 0}}' chatwork/"$(REPO)":$$version >/dev/null 2>&1; then \
			echo "no changes"; \
		else \
			docker push chatwork/"$(REPO)"; \
		fi
