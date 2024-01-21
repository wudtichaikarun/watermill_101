# proto:
# 	protoc --go_out=paths=source_relative:. -I inputs inputs/events.proto 

tidy:
	go mod tidy

dev:
	go run .
