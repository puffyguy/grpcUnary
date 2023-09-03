# gRPC Unary Service in Go (Golang)

## Overview

This repository contains a simple gRPC Unary Service implemented in Go. The service demonstrates a basic client-server interaction using gRPC, where the client sends a single request, and the server responds with a single response.

### Features

- Unary RPC (Request-Response) communication.
- Sample implementation of a client and server.
- Protocol Buffers (Proto) for defining the service and messages.

## Prerequisites

Before you can run the gRPC Unary Service, ensure that you have the following prerequisites installed on your system:

- Go (Golang): [Install Go](https://golang.org/doc/install)
- Protocol Buffers (Proto): [Install Protobuf](https://developers.google.com/protocol-buffers/docs/gotutorial)

## Getting Started

Follow these steps to set up and run the gRPC Unary Service:

1. Clone the repository:

   ```bash
   git clone https://github.com/puffyguy/grpcUnary.git
   cd grpcUnary
   ```

2. Start the server:
    ```
    ./server
    ```

3. In another terminal, run the client:
    ```
    ./client
    ```