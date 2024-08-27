protoc:
	cd internal/grpc/proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	*.proto

swag-init:
	swag init -g /cmd/main.go

start:
	go run cmd/main.go

test:
	go test ./...