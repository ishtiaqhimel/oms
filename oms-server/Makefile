REGISTRY ?= ishtiaq99
BUILD_PLATFORM ?= linux/amd64,linux/arm64

fmt:
	go fmt ./...

.PHONY: swagger
swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

test:
	go test ./handlers -timeout 30s ./...

build:
	docker buildx build \
  	--platform $(BUILD_PLATFORM) \
  	-t $(REGISTRY)/oms-server:latest \
  	--push .