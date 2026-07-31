# OMS (Order Management System)

A Go-based microservices Order Management System (OMS) laid out as a `go.work` multi-module workspace.

## Architecture

| Module | Purpose |
|---|---|
| `gateway` | HTTP API gateway (port 8080) exposing order endpoints |
| `orders` | Order domain service (store/service interfaces + create flow) |
| `kitchen` | Kitchen service module |
| `payments` | Payments service module |
| `stock` | Stock service module |
| `common` | Shared library code |

## Modules

- **gateway** — HTTP server and request routing (`gateway/http_handler.go`)
- **orders** — Domain layer with `OrdersStore`, `OrdersService`, and `main` entrypoint
- **kitchen** — Kitchen-specific business logic
- **payments** — Payment processing logic
- **stock** — Stock/inventory management logic
- **common** — Shared types and utilities

## Tech Stack

- Go 1.26.5
- `go.work` workspace linking: `common`, `gateway`, `kitchen`, `orders`, `payments`, `stock`

## Getting Started

Prerequisites: Go 1.26+ installed.

Run each service from its module directory, for example:

```
go run ./gateway
```
