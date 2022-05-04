.PHONY: build test clean prepare update docker

GO = CGO_ENABLED=1 GO111MODULE=on go

MICROSERVICES=cmd/device-idt-lora-mqtt
.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_idt_lora_mqtt
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION)
GIT_SHA=$(shell git rev-parse HEAD)

GOFLAGS=-ldflags "-X github.com/TIBCOSoftware/labs-air/edgexfoundry/device-idt-lora-mqtt.Version=$(VERSION)"

build: $(MICROSERVICES)

cmd/device-idt-lora-mqtt:
	go mod tidy
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -cover

clean:
	rm -f $(MICROSERVICES)

run:
	cd bin && ./edgex-launch.sh

docker: $(DOCKERS)

docker_device_idt_lora_mqtt:
	docker build --no-cache \
		--label "git_sha=$(GIT_SHA)" \
		-t tibcosoftware/labs-air-edgex-device-idt-lora-mqtt:$(VERSION) \
		.

dockerarm64:
	docker build --no-cache \
		--label "git_sha=$(GIT_SHA)" \
		-t tibcosoftware/labs-air-edgex-device-idt-lora-mqtt-arm64:$(VERSION) \
		.

dockerbuildxarm64:
	docker buildx build \
		--platform linux/arm64 \
		--label "git_sha=$(GIT_SHA)" \
		-t alanlee/docker-device-idt-lora-mqtt-arm64:$(VERSION) \
		.

