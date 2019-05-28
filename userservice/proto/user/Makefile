regenerate-docker:
	docker run --rm -it \
		-v `pwd`:/user \
		-w /user \
		obraun/vss-protoactor-jenkins \
		make regenerate

regenerate:
	go install -v github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=plugins=grpc:. user.proto