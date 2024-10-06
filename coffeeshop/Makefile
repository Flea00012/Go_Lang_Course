build-proto:
	protoc --go_out=./coffeeshop_proto --go_opt=paths=source_relative \
	--go-grpc_out=./coffeeshop_proto --go-grpc_opt=paths=source_relative \
	coffee_shop.proto

server:
	go run server.go

client:
	go run ./client/client.go

