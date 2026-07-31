# OMS (Order Management System)

A Go-based microservices Order Management System (OMS) built with a `go.work` multi-module workspace.

## Architecture

| Module | Purpose |
|---|---|
| `gateway` | HTTP API gateway (default port `:8080`) that proxies order requests to the orders service via gRPC |
| `orders` | gRPC order domain service (default port `:2000`) exposing `CreateOrder` |
| `kitchen` | Kitchen service module (scaffold) |
| `payments` | Payments service module (scaffold) |
| `stock` | Stock service module (scaffold) |
| `common` | Shared library code, including protobuf-generated types and gRPC stubs |

## Inter-service Communication

- `gateway` → `orders`: gRPC (insecure)
- Protocol: protobuf (`common/api/oms.proto`)
- Endpoint exposed by gateway: `POST /api/customers/{customerID}/orders`

## Modules

- **gateway** — HTTP server (`main.go`), request routing, and handlers (`http_handler.go`). Deserializes the incoming JSON request, builds a `CreateOrderRequest`, and forwards it to the orders service over gRPC.
- **orders** — gRPC server (`main.go`) with domain interfaces and handlers (`grpc_handler.go`, `service.go`, `store.go`). Create flow is currently a stub.
- **kitchen** — Kitchen-specific business logic (module scaffold).
- **payments** — Payment processing logic (module scaffold).
- **stock** — Stock/inventory management logic (module scaffold).
- **common** — Shared utilities (`env.go`, `json.go`), protobuf schema (`api/oms.proto`), and generated Go/gRPC code.

## Tech Stack

- Go 1.26.5
- `go.work` workspace: `common`, `gateway`, `kitchen`, `orders`, `payments`, `stock`
- Protocol Buffers (`common/api/oms.proto`) with `make gen`
- gRPC (`google.golang.org/grpc`)
- godotenv for environment configuration (`.env`)
- [air](https://github.com/cosmtrek/air) for hot reloading (`.air.toml`)

## Getting Started

Prerequisites: Go 1.26+ and `protoc` installed.

Start the orders gRPC service:

```bash
go run ./orders
```

Start the gateway HTTP service (in another terminal):

```bash
go run ./gateway
```

Generate protobuf code after updating schema:

```bash
cd common && make gen
```

Hot reload development with `air`:

```bash
air
```

## Configuration

| Variable | Module | Default | Description |
|---|---|---|---|
| `HTTP_ADDR` | `gateway` | `:8080` | HTTP listener address |
| `GRPC_ADDR` | `orders` | `localhost:2000` | gRPC listener address |

## API

### Create Order

`POST /api/customers/{customerID}/orders`

Request body:

```json
[
  { "ID": "item-1", "Quantity": 2 },
  { "ID": "item-2", "Quantity": 1 }
]
```
