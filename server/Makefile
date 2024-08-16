postgres_init:
	docker run --name postgres15 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/go-chat?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/go-chat?sslmode=disable" -verbose down

.PHONY: postgres_init postgres migrateup migratedown