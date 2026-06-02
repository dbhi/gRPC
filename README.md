# [DBHI] Dynamic Binary Hardware Injection: gRPC server and clients

This repository contains resources to provide inter-process communication through [gRPC Remote Procedure Calls](https://grpc.io/) in the context of Dynamic Binary Hardware Injection (DBHI).
gRPC supports cross-platform client and server bindings for many languages.
We used [go(lang)](https://golang.org/) to write both a buffered channel (FIFO) server and two example clients.
The server provides a go channel for each identifier added to the list through the API.
Then, read and write methods are available for clients to interchange information with FIFO-alike interfaces.
Moreover, helper functions for clients can be built to a shared library and it can be used in third-party applications (probably written in C/C++).

- `lib/`: definition of the API.
  Language-specific sources generated with `protoc` (`protoc-gen-go` and `protoc-gen-go-grpc`).
- `server.go`: server that implements the API defined in `lib/lib.proto`, and provides channels as FIFO interfaces.
- `client/`: helper functions that depend on `lib`, to write client applications.
- `client.export.go`: functions from package `client` to be exported to C.
- `gen.go`: client that waits until two channels exist in the server, then pushes data to a channel, and expects to get
  the same number of elements (each of them multiplied by three) from a different channel.
  It is expected to be used in order to test the server along with `uut.go` or `uut.c`.
- `uut.*`: client that receives any number of elements from a channel and returns each element multiplied by three to a different channel.
  It is expected to be used in order to test the server along with `gen`.
  Two functionally equivalent versions are provided: `uut.go` and `uut.c` (which depends on `libgrpc-go.so` and `libgrpc-go.h` built with go).
  NOTE: `LD_LIBRARY_PATH=$(pwd) ./uut-c`.

In order to build all the pieces, `protoc`, `golang` and `gcc` are required.
See [Dockerfile](Dockerfile) and [Makefile](Makefile).

Either after installing the dependencies natively or after starting the container, execute `make all` to install dependencies and build all the pieces.
Artifacts (server, three clients, shared library and header file) will be output to `./dist/`.

## References

- [gh:protocolbuffers/protobuf: examples#go](https://github.com/protocolbuffers/protobuf/tree/main/examples#go)
- [gh:campoy/justforfunc/31-grpc: The Basics of gRPC](https://github.com/campoy/justforfunc/tree/master/31-grpc)
- [gh:vladimirvivien/go-cshared-examples: Calling Go Functions from Other Languages using C Shared Libraries](https://github.com/vladimirvivien/go-cshared-examples)
