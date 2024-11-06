generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./client --go-grpc_opt=paths=source_relative /pkg/api/order.proto