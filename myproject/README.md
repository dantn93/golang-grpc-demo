protoc -I src/protos/ src/protos/student.proto --go_out=. --go-grpc_out=.

go build -o bin/server src/server/server.go

grpcui -plaintext -import-path ./src/protos -proto student.proto localhost:50051