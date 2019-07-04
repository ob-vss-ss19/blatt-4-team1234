
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --gogoslick_out=. proto/user/user.proto

.PHONY: build
build: proto

	go build main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover
