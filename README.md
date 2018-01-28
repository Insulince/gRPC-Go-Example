# gRPC Go Example

This is a short project showing the four different classes of [gRPC](https://grpc.io/) calls (in terms of data-streaming) that can be made from both client and server side using Golang with *minimal noise involved*, meaning this is a minimalistic implementation of the calls simply to show how each of them work.

This project was created using a Windows environment. When executing any commands in other environments, you may need to make the necessary changes to accomodate for things like the `exe` extension used.

All commands are from the perspective of the project root.

## Concepts

**RPC Types**
- **Unary** - Neither the **client** nor the **server** will be streaming any more than **one** message to each other.
- **Server Stream** - The **server** will **stream** responses to the client, the **client** will only send **one** request.
- **Client Stream** - The **client** will **stream** requests to the server, the **server** will only send **one** response.
- **Bidirectional Stream** - Both the **client** and the **server** will be **streaming** messages to each other.

**Resources**
- There is only one resource that is sent through messages, it is called `FooBarBaz`. As the name suggests, this is a completely arbitrary data structure. It is composed of three fields, `Foo` of type `string`, `Bar` of type `int64`, and `baz` of type `bool`. Each of these are included to show how different data types can be sent encapsulated in a struct.

**Requests**
- Requests from the client are composed of only one field, `FooBarBaz` of type `FooBarBaz`. The data sent is purely to show that data can be sent.

**Responses**
- Responses from the server are composed of only one field, `Success` of type `bool`. This field is used to indicate that all went well on the server.

**Processing Time**
- In an effort to show what is actually going on during the streaming events, I added in a function to the `./src/util.go` file called `SimulateProcessing`. The purpose of this function is solely to put the current thread to sleep for a random amount of time while the streaming is occurring. This allows one to observe the requests and responses as they are happening in real time, and also lends itself to a more realistic model of gRPC in which requests and responses would indeed be slowed down by some sort of processing time.

## Project Structure
- `bin`
    - Compiled project files will go here.
- `pkg`
    - Empty
- `src`
    - `client`
        - All code specific to the gRPC client is here.
    - `pb`
        - `foo-bar-baz.proto` This is the protocol buffer file.
        - `foo-bar-baz.pb.go` This is the Go code generated from `protoc` on the preceding file.
    - `server`
        - All code specific to the gRPC server is here.
    - `util.go` contains some code useful to both the client and server.

## The Server

#### Building
- `go.exe build -o "./bin/server.exe" "./src/server"`

#### Running
- `./bin/server.exe`

The server listens on port `9000` for incoming requests. The client is aware of this by default, no additional configuration is needed.


## The Client

#### Building
- `go.exe build -o "./bin/client.exe" "./src/client"`

#### Running
- `./bin/client.exe <request-type>`

    Where `<request-type>` is one of `unary`, `server-stream`, `client-stream`, or `bidirectional-stream`.

The request type you provide is what type of request the client will issue to the server. This is how you can test each type of request available.

## Protocol Buffer Compilation

The protocol buffer file in this project (`./src/pb/foo-bar-baz.proto`) has already been compiled into it's Go counterpart, but if you are making changes and would like to see them reflected in the project, continue reading.

1. Install [protoc](https://github.com/google/protobuf/releases) and ensure it is in your `Path`.

2. Install the `grpc` plugin for `protoc`:

    `go.exe get -u github.com/golang/protobug/protoc-gen-go`

3. Perform the `protoc` compilation

    `protoc.exe "./src/pb/foo-bar-baz.proto" --go_out=plugins=grpc:"./src"`

This tells `protoc` to compile the `proto` file located at `./src/pb/foo-bar-baz.proto` into Go using the `grpc` plugin, and placing the output at `./src` (and then into its given package directory, defined in the `proto` file).

Everything should be updated now in the generated code, all previous compilations are overwritten (unless you changed any names or directories).

## Notes
- I am not an authority on anything RPC or gRPC related. I made this for fun and as a way to flesh out these calls for myself. I am putting it online in order to keep it safe from hardware failures, and in case anyone else may be looking for a pure implementation of the gRPC calls without any unnecessary example business logic.
- A majority of these concepts I learned through Michael Van Sickle's PluralSight course [Enhancing Application Communication with gRPC](https://app.pluralsight.com/library/courses/grpc-enhancing-application-communication/table-of-contents) and I highly recommend it if you want to see a more real world example of these RPCs in action.
