BINARY=dokkery
VERSION := "0.1.0"
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "ziwon/dokkery"

get-tools:
	brew install golangci/tap/golangci-lint
	brew upgrade golangci/tap/golangci-lint

get-deps:
	go mod tidy

build:
	@echo "building ${BINARY} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	GO111MODULE=on go build -o ${BINARY} -ldflags="-w -X 'main.version=${VERSION}'" ./cmd/${BINARY}

build-alpine:
	@echo "building ${BINARY} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	GO111MODULE=on go build -o ${BINARY} -ldflags="-w -linkmode external -extldflags -static -X 'main.version=${VERSION}'" ./cmd/${BINARY}

dist:
	@echo "building image ${BINARY} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} -t $(IMAGE_NAME):latest .

clean:
	@test ! -e ${BINARY} || rm ${BINARY}

test:
	go test ./...


.PHONY: get-tools get-deps build build-alpine dist clean test
