# Generate sources

From the root of the project:
```terminal
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/plan.proto
```