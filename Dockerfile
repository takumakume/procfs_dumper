FROM golang:latest
WORKDIR /go/src

ENV GO111MODULE on

RUN apt-get update && apt-get install -y \
  jq
RUN go get github.com/goreleaser/goreleaser
