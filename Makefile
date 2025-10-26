run: contract
	go run cmd/debez/main.go

contract:
	protoc --go_out=. --go-grpc_out=. ./api/user.proto