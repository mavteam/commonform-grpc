# CommonForm Engine (as a Server)

This repository provides a daemon engine that wraps some of the most oft used [Common Form](https://github.com/commonform) utilities.

The daemon is designed to be operated as a standalone [Docker](./Dockerfile) container that just runs within a server farm or cluster using a [GRPC](https://grpc.io) server which accepts inbound requests via the [protobuf standard file](./requests/commonform.proto).

## Install

There are two ways to install the engine within your system.

The first is to clone the repository, build the Docker image, push to a repository, and then to turn it on via the instructions below.

The second is to leverage the [auto-built Docker images](https://quay.io/repository/monax/commonform-engine) within your cluster.

## Operate

The engine can be operated via node.js on metal with `yarn start` although it has been primarily built to operate in a cluster via the Docker image.

The configuration of the engine is handled via environment variables:

| **Variable** | **Notes** | **Default** |
|------------|--------|------------------|
| `LOG_LEVEL` | Log level for the server(s) | `info` |
| `GRPC_SERVER_HOST` | Host the GRPC server should listen on (generally this should be blank unless you really know what you're doing) | `127.0.0.1` |
| `GRPC_SERVER_PORT` | Port the GRPC server should listen on | `8081` |
