postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it  postgres17 createdb --username=root --owner=root nimbus_db

dropdb:
	docker exec -it postgres17 dropdb nimbus_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nimbus_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nimbus_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nimbus_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nimbus_db?sslmode=disable" -verbose down 1

migrateneon:
	migrate -path db/migration -database "$$DB_SOURCE" -verbose up

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

mock:
	mockgen -package mockdb -destination db/mockdb/store.go github.com/OmSingh2003/nimbus/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/OmSingh2003/nimbus/worker TaskDistributor

server:
	go run main.go

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative  --grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative   --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=vaultguard-api,disable_default_errors=true,simple_operation_ids=true  proto/*.proto
	statik -src=./doc/swagger -dest=./doc

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 db_docs db_schema sqlc test mock server proto clean_proto rebuild_proto evan evans redis

clean_proto:
	rm -f pb/*.pb.go

rebuild_proto: clean_proto proto

evan:
	 evans --host localhost --port 9090 -r repl

evans :
	evans --path proto --proto service_vaultguard_api.proto --host localhost --port 9090

redis:
	docker run --name redis -p 6379:6379 -d redis:8.0.2-alpine
