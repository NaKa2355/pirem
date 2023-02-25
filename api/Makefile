OUT := ../
DEPENDS := ./apirem/v1 ./plugin/v1
PROTO := $(shell find ./ -name *.proto)

all:
	protoc -I. --go_opt=module=github.com/NaKa2355/pirem --go-grpc_opt=module=github.com/NaKa2355/pirem --experimental_allow_proto3_optional --go_out=$(OUT) --go-grpc_out=$(OUT) $(PROTO)