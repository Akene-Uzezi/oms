# OMS (Order Management System)

A Go-based microservices Order Management System (OMS) using a `go.work` multi-module workspace.

## Architecture

| Module | Purpose |
|---|---|
| `gateway` | HTTP API gateway (default port 8080) exposing order endpoints |
| `orders` | Order domain service (store/service interfaces + create flow) |
| `kitchen` | Kitchen service module |
| `payments` | Payments service module |
| `stock` | Stock service module |
| `common` | Shared library code, including protobuf-generated types and gRPC stubs |

## Modules

- **gateway** — HTTP server (`main.go`), request routing, and handlers (`http_handler.go`). Exposes `POST /api/customers/{customerID}/orders`.
- **orders** — Domain layer with `OrdersStore`, `OrdersService`, and a `main` entrypoint. Create flow is currently a stub.
- **kitchen** — Kitchen-specific business logic (module scaffold).
- **payments** — Payment processing logic (module scaffold).
- **stock** — Stock/inventory management logic (module scaffold).
- **common** — Shared types, utilities, and protobuf/gRPC definitions under `common/api/`.

## Tech Stack

- Go 1.26.5
- `go.work` workspace linking: `common`, `gateway`, `kitchen`, `orders`, `payments`, `stock`
- Protocol Buffers for shared API types (`common/api/oms.proto`)
- gRPC generated stubs in `common/api/`
- godotenv for environment configuration (`gateway/.env`)

## Getting Started

Prerequisites: Go 1.26+ installed.

Run the gateway service:

```bash
go run ./gateway
```

Run the orders service:

```bash
go run ./orders
```

Generate protobuf code:

```bash
cd common && make gen
```

The HTTP server listens on the address configured by `HTTP_ADDR` (default `:8080`).
