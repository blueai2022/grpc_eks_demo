# DB_URL=postgresql://root:P55geiSp411IH2cVasYV@appsubmission.ckwftjwmrhkw.us-east-1.rds.amazonaws.com:5432/app_submission?sslmode=disable
DB_URL=postgresql://root:secretpwd@localhost:5432/app_submission?sslmode=disable

k8_version:
	curl http://storage.googleapis.com/kubernetes-release/release/stable.txt

network:
	docker network create lifeai-network

postgres:
	docker run --name postgres14 --network lifeai-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpwd -d postgres:14-alpine
	docker exec -it postgres14 createdb --username=root --owner=root app_submission

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

server_dkr:
	docker run --name appsubmission --network lifeai-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secretpwd@postgres14:5432/app_submission?sslmode=disable" appsubmission:latest

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root app_submission

dropdb:
	docker exec -it postgres14 dropdb app_submission

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: k8_version network postgres server_dkr createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans
