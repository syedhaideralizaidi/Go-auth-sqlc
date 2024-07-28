DB_URL=postgresql://root:secret@localhost:5432/simple_auth?sslmode=disable

postgres:
	docker run --name postgres_auth -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13

created_db:
	docker exec -it postgres_auth createdb --username=root --owner=root simple_auth

drop_db:
	docker exec -it postgres_auth dropdb simple_auth

migrate_up:
	migrate -path database/migration -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path database/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres created_db drop_db migrate_up migrate_down sqlc