PROTO := $(shell find ./ -name *.proto)

#--------------golang generator-------------------
GO_OUT := ../gen/go
GO_MODULE := $(shell go list -m)
GO_FILES := $(shell find ../ -name *.go)

gen-go:
	protoc -I. --experimental_allow_proto3_optional \
	--go_opt=paths=source_relative --go_out=$(GO_OUT) \
	--go-grpc_opt=paths=source_relative --go-grpc_out=$(GO_OUT) \
	$(PROTO)

clean-go:
	rm -rf $(GO_OUT)/**

#---------------------------------------------------

#--------------js generator-------------------------
JS_OUT := ./../gen/js

gen-js: 
	protoc -I=. $(PROTO) \
	--js_out=import_style=commonjs,binary:$(JS_OUT) \
	 --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:$(JS_OUT)

clean-js:
	rm -rf $(JS_OUT)/**

#--------------------general------------------------
clean: clean-go clean-js

gen: gen-go, gen-js

all:
#---------------------------------------------------