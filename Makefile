# proto:
# 	protoc --go_out=paths=source_relative:. -I inputs inputs/events.proto 

tidy:
	go mod tidy

dev:
	go run .

compose-dev:
	@docker-compose \
		-f docker-compose.yml \
		up --build