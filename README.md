# [DBHI] gRPC

This subdir contains sources to provide inter-process communication through [gRPC Remote Procedure Calls](https://grpc.io/). gRPC generates cross-platform client and server bindings for many languages. We used [go(lang)](https://golang.org/) to write both the server and two example clients. The server provides a go channel for each identifier added to the list through the API. Then, read and write methods are available for clients to interchange information with FIFO-alike interfaces. Moreover, helper functions for clients can be built to a shared library and it can be used in C applications.

- `lib/`: definition of the API. Language-specific sources are generated with `protoc` (see [run](./run)).
- `server/`: server that implements the API defined in `lib/lib.proto`, and provides channels as FIFO interfaces.
- `client/`:
    - `*.go`: helper functions that depend on `lib`, and functions exported to C.
    - `gen/`: client that waits until two channels exist in the server, then pushes data to a channel, and expects to get the same number of elements (each of them multiplied by three) from a different channel. It is expected to be used in order to test the server along with `eg`. This is written in golang, and depends on `lib`.
    - `eg-*/`: client that receives any number of elements from a channel and returns each element multiplied by three to a different channel. It is expected to be used in order to test the server along with `gen`. Two functionally equivalent versions are provided:
        - `eg-go/` in golang. In order to run/build it, `client/client.go` must be copied to `client/eg-go/`.
        - `eg-c` in C. In order to build it, `libgrpc-go.so` and `libgrpc-go.h` must be generated first (see [run](./run)).

In order to build all the pieces, `protoc`, `golang` and `gcc` are required. If docker is available, image `aptman/dbhi:stretch-gRPC` can be used. Start a container as follows:

``` bash
# WORK_DIR="/go/src/github.com/umarcor/dbhi/gRPC"
# $(command -v winpty) docker run --rm -itv "/$(pwd):/$WORK_DIR" -w "/$WORK_DIR" aptman/dbhi:stretch-gRPC bash
```

Either after installing the dependencies natively or after starting the container, execute [`./run`](./run) to build all the pieces. The artifacts (server, three clients, shared library and header file) will be output to `./dist/`.

## References

- [campoy/justforfunc/31-grpc: The Basics of gRPC](https://github.com/campoy/justforfunc/tree/master/31-grpc)
- [vladimirvivien/go-cshared-examples: Calling Go Functions from Other Languages using C Shared Libraries](https://github.com/vladimirvivien/go-cshared-examples)
