
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	go install -v github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto

.PHONY: build
build: proto

	go build -o reservationservice-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t reservationservice-srv:latest
