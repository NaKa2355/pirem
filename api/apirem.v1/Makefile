TARGET := ../../pkg

all:
	protoc -I. --experimental_allow_proto3_optional --go_out=$(TARGET) --go-grpc_out=$(TARGET) *.proto