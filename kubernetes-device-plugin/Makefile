NAME=coffee-device-plugin
VERSION=v0.0.1
REPO=quay.io/huyiqiu1998
K8S_YAML=coffee-ds.yaml

# Docker 镜像名称
IMAGE = $(REPO)/$(NAME):$(VERSION)

# Go 源文件
GO_SRC = main.go

BINARY = $(NAME)

GO_FLAGS = "-a -installsuffix cgo"

.PHONY: all build clean docker-build docker-push deploy

all: build

build:
	@echo "==> Building Go binary"
	CGO_ENABLED=0 GOOS=linux go build $(GO_FLAGS) -o bin/$(BINARY) $(GO_SRC)

clean:
	@echo "==> Cleaning up"
	rm -f bin/$(BINARY)

docker-build: build
	@echo "==> Building Docker image"
	docker build -t $(IMAGE) .

docker-push:
	@echo "==> Pushing Docker image"
	docker push $(IMAGE)

deploy:
	@echo "==> Deploying to Kubernetes"
	kubectl apply -f deploy/$(K8S_YAML)

all-deploy: docker-build docker-push deploy