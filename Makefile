postgres:
	docker run --name mpi-draw -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=d4ERBLCFmfK10U80Z0o4lGD2s9l9Tlxr -d postgres
createdb:
	docker exec -it mpi-draw createdb --username=root --owner=root mpi_draw
dropdb:
	docker exec -it mpi-draw dropdb mpi_draw
migrateup:
	migrate -path db/migration -database "postgresql://root:d4ERBLCFmfK10U80Z0o4lGD2s9l9Tlxr@localhost:5432/mpi_draw?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:d4ERBLCFmfK10U80Z0o4lGD2s9l9Tlxr@localhost:5432/mpi_draw?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test