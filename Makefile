.PHONY: up, down, lint

up:
	docker-compose up -d --build

down:
	docker-compose down -v

lint:
	go vet ./...
	golangci-lint run ./...