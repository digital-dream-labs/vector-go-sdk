PROTO_DIR = ./proto/

protos: $(PROTO_DIR)
	for dir in $^ ; do protoc \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I ${GOPATH}/src/github.com/golang/protobuf/ptypes/struct \
		-I $${dir} \
		--proto_path=. \
		--go_out=plugins=grpc,paths=source_relative:./pkg/vectorpb \
		$${dir}/*.proto \
	; done