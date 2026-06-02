.PHONY: all clean dependencies protoc-gen lib slib

all: dependencies lib dist

clean:
	rm -rf dist

dependencies:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

lib: lib/lib.pb.go lib/lib_grpc.pb.go

lib/lib.pb.go lib/lib_grpc.pb.go: lib/lib.proto
	protoc --go_out=. lib/lib.proto
	protoc --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false lib/lib.proto
#https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc#future-proofing-services

dist: dist/server dist/gen dist/uut-go slib dist/uut-c

dist/server: lib/lib.pb.go lib/lib_grpc.pb.go server.go
	go build -a -o dist/server server.go

dist/gen: lib/lib.pb.go lib/lib_grpc.pb.go gen.go
	go build -a -o dist/gen gen.go

dist/uut-go: lib/lib.pb.go lib/lib_grpc.pb.go client/client.go uut.go
	go build -a -o dist/uut-go uut.go

slib: dist/libgrpc-go.so

dist/libgrpc-go.so dist/libgrpc-go.h: client.export.go client/client.go
	go build -a -o dist/libgrpc-go.so -buildmode=c-shared client.export.go

dist/uut-c: dist/libgrpc-go.so dist/libgrpc-go.h uut.c
	gcc -o dist/uut-c uut.c -Idist -Ldist -lgrpc-go
