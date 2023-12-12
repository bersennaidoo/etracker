# database name
DB_NAME ?= postgres

# database type
DB_TYPE ?= postgres

# database username
DB_USER ?= postgres

# database password
DB_PWD ?= bersen

# psql URL
IP=127.0.0.1

PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

# sqlc yaml file
SQLC_YAML ?= ./sqlc.yaml

.PHONY : postgresup postgresdown psql createdb teardown_recreate generate

postgresup:
	docker run --name testpgdb -v $(PWD):/usr/share/infra -e POSTGRES_PASSWORD=$(DB_PWD) -p 5432:5432 -d $(DB_NAME)

postgresdown:
	docker stop testpgdb  || true && 	docker rm testpgdb || true

migrate-up:
	migrate -path server/infrastructure/storage/pgstore/migrations -database "postgresql://postgres:bersen@localhost/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path server/infrastructure/storage/pgstore/migrations -database "postgresql://postgres:bersen@localhost/postgres?sslmode=disable" -verbose down


psql:
	docker exec -it testpgdb psql $(PSQLURL)

# task to create database without typing it manually
createdb:
	docker exec -it testpgdb psql $(PSQLURL) -c "\i /usr/share/infra/server/infrastructure/storage/pgstore/schema/schema.sql"

teardown_recreate: postgresdown postgresup
	sleep 5
	$(MAKE) createdb

generate:
	@echo "Generating Go models with sqlc "
	sqlc generate -f $(SQLC_YAML)
