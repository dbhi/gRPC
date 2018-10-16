# DBHI-gRPC

This subdir contains sources to provide inter-process communication through [gRPC Remote Procedure Calls](https://grpc.io/). gRPC generates cross-platform client and server bindings for many languages. We used [go(lang)](https://golang.org/) to write both the server and a few example clients. The server provides a go channel for each identifier added to the list through the API. Then, read and write methods are available for clients to interchange information with FIFO-alike interfaces. Moreover, helper functions for clients can be built to a shared library and it can be used in C applications.

- `lib/`: definition of the API. Language-specific sources are generated with `protoc` (see [run](./run)).
- `server/`: server that implements the API defined in `lib/lib.proto`, and provides channels as FIFO interfaces.
- `client/`:
    - `*.go`: helper functions that depend on `lib`, and functions exported to C.
    - `gen/`: client that waits until two channels exist in the server, then pushes data to a channel, and expects to get the same number of elements (each of them multiplied by three) from a different channel. It is expected to be used in order to test the server along with `eg`. This is written in golang, and depends on `lib`.
    - `eg/`: client that receives any number of elements from a channel and returns each element multiplied by three to a different channel. It is expected to be used in order to test the server along with `gen`. Two functionally equivalent versions are provided:
        - `main.go` in golang. In order to run/build it, `client/client.go` must be copied to `client/eg/`.
        - `main.c` in C. In order to build it, execute `client/run` and `client/eg/run`.

# References

- [campoy/justforfunc/31-grpc: The Basics of gRPC](https://github.com/campoy/justforfunc/tree/master/31-grpc)
- [vladimirvivien/go-cshared-examples: Calling Go Functions from Other Languages using C Shared Libraries](https://github.com/vladimirvivien/go-cshared-examples)


