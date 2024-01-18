proto:
	protoc --proto_path=inputs --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	inputs/*.proto

tidy:
	go mod tidy

dev:
	go run .