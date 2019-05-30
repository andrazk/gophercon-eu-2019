GOOS?=linux
GOARCH?=amd64

PROJECT?=github.com/andrazk/tenerife
BUILD_PATH?=cmd/tenerife
APP?=tenerife

# Current version
RELEASE?=0.0.1

# Parameters to push images and release app to Kubernetes or try it with Docker
REGISTRY?=docker.io/andrazk
NAMESPACE?=andrazk
CONTAINER_NAME?=${NAMESPACE}-${APP}
CONTAINER_IMAGE?=${REGISTRY}/${CONTAINER_NAME}
VALUES?=values

build:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: build
	docker push $(CONTAINER_IMAGE):$(RELEASE)
