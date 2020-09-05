BINARY=dokkery
VERSION := "0.1.0"
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "ziwon/dokkery"

get-tools:
	brew install golangci/tap/golangci-lint
	brew upgrade golangci/tap/golangci-lint
	brew install goreleaser/tap/goreleaser
	brew install goreleaser
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $GOPATH/bin latest

get-deps:
	go mod tidy

build:
	@echo "building ${BINARY} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	GO111MODULE=on go build -o ${BINARY} -ldflags="-w -X 'main.version=${VERSION}'" ./cmd/${BINARY}

build-linux:
	@echo "building ${BINARY} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	GO111MODULE=on go build -o ${BINARY} -ldflags="-w -linkmode external -extldflags -static -X 'main.version=${VERSION}'" ./cmd/${BINARY}

verify:
	golangci-lint run
	gosec -severity high --confidence medium -exclude G204 -quiet ./...

dist:
	@echo "building image ${BINARY} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} -t $(IMAGE_NAME):latest .

release:
	goreleaser release --rm-dist

clean:
	@test ! -e ${BINARY} || rm ${BINARY}

test:
	go test ./...


.PHONY: get-tools get-deps build build-linux dist clean test
