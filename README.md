# auto-stock-trading-system

This system allows you to buy and sell stocks automatically.

# System Composition

- web
  - nextjs
- backend for frontend
  - nextjs
- micro service

  - go

#### gRPC command

```
protoc --go_out=micro-service/user-service/pkg/grpc --go_opt=paths=source_relative --go-grpc_out=micro-service/user-service/pkg/grpc --go-grpc_opt=paths=source_relative pkg/proto/hello.proto
```
