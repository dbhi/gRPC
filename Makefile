.PHONY: all clean dependencies protoc-gen lib slib

all: dependencies lib dist

clean:
	rm -rf dist

GOLIBS=lib/lib.pb.go lib/lib_grpc.pb.go
CLIBS=dist/libgrpc-go.so dist/libgrpc-go.h
CLIENT=client/client.go

dependencies:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

lib: $(GOLIBS)

lib/lib.pb.go: lib/lib.proto
	protoc --go_out=. $<

lib/lib_grpc.pb.go: lib/lib.proto
	protoc --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false $<
#https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc#future-proofing-services

dist: dist/server dist/gen dist/uut-go slib dist/uut-c

dist/server: server.go $(GOLIBS)
	go build -a -o $@ $<

dist/gen: gen.go $(CLIENT) $(GOLIBS)
	go build -a -o $@ $<

dist/uut-go: uut.go $(CLIENT) $(GOLIBS)
	go build -a -o $@ $<

slib: $(CLIBS)

$(CLIBS): client.export.go $(CLIENT) $(GOLIBS)
	go build -a -o dist/libgrpc-go.so -buildmode=c-shared $<

dist/uut-c: uut.c $(CLIBS)
	gcc -o $@ $< -Idist -Ldist -lgrpc-go
