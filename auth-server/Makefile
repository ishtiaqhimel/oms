REGISTRY ?= ishtiaq99
BUILD_PLATFORM ?= linux/amd64,linux/arm64

fmt:
	go fmt ./...

.PHONY: swagger
swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

build:
	docker buildx build \
  	--platform $(BUILD_PLATFORM) \
  	-t $(REGISTRY)/auth-server:latest \
  	--push .