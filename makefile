.PHONY: migrate clean-migrations run cleanpg build
include .env

run:
	go run cmd/app/main.go

build:
	 go build -o ./build/main ./cmd/app/main.go && ./build/main

migrate:
	migrate -source file://migrations \
	-database ${DATABASE_URL} up

clean-migrations: 
	migrate -source file://migrations \
	-database ${DATABASE_URL} down

cleanpg:
	docker stop postgres && docker remove postgres && docker volume remove pgdata