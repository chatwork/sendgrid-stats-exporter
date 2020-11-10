default: build

.PHONY: build
build:
	go build -o exporter .