# Build Stage
FROM golang:1.13 AS builder
LABEL app="dokkery"
LABEL REPO="https://github.com/ziwon/dokkery"

ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ADD . /go/src/github.com/ziwon/dokkery
WORKDIR /go/src/github.com/ziwon/dokkery
RUN make build-alpine

# Dist Stage
FROM alpine:3.12.0
RUN apk add openssh
COPY --from=builder /go/src/github.com/ziwon/dokkery/dokkery /app/dokkery
COPY config.example.yml /app/config.yml
WORKDIR /app
CMD ["./dokkery", "serve"]
