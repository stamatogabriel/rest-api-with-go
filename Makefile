include .env

up:
	@echo "Starting MongoDB and application..."
	docker-compose up -d --remove-orphans

down:
	@echo "Stopping and removing containers..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/main.go

start:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} \
		MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} \
		./${BINARY}

restart: build start