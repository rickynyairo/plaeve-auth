build:
	protoc --proto_path=proto/auth --proto_path=third_party --go_out=plugins=micro:proto/auth auth.proto
	@echo 'Build Docker image'
	docker build -t plaeve-auth .

run:
	docker run --net="host" \
			-p 50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=etcd \
			-e DB_HOST=localhost \
			-e DB_NAME=postgres \
			-e DB_USER=postgres \
			-e DB_PASSWORD="" \
			-e DB_PORT=5432 \
			plaeve-auth

run-go:
	go run main.go database.go repository.go handler.go token_service.go