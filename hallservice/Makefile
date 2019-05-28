
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/hall/hall.proto

.PHONY: build
build: proto
	go build -o hallservice-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover
