
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --gogoslick_out=. proto/show/show.proto

.PHONY: build
build: proto

	go build -o showservice-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover
