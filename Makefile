include .env
# include ./scripts/protoc-generate.sh

empty:
	echo "empty"

dcb:
	docker compose build
dcu:
	docker compose up -d bff
dcd:
	docker compose down

# enter the container
## db container
db-ssh:
	docker exec -it $(MYSQL_HOST) sh
## bff container
bff-ssh:
	docker exec -it $(BFF_CONTAINER_NAME) sh
## user service container
user-ssh:
	docker exec -it $(USER_SERVICE_CONTAINER_NAME) sh
## batch service container
batch-ssh:
	docker exec -it $(BATCH_SERVICE_CONTAINER_NAME) sh


# db migration
db-migrate:
	docker exec -it ${BATCH_SERVICE_CONTAINER_NAME} sh -c "migrate -source file://migrations -database 'mysql://user:password@tcp(auto_stock_trading_system_db:3306)/AUTO_STOCK_TRADING_SYSTEM?multiStatements=true' up"

db-seed:
	docker exec -it ${BATCH_SERVICE_CONTAINER_NAME} sh -c "go run ./cmd/seed/seed.go"

db-rollback:
	docker exec -it ${BATCH_SERVICE_CONTAINER_NAME} sh -c "migrate -source file://migrations -database 'mysql://user:password@tcp(auto_stock_trading_system_db:3306)/AUTO_STOCK_TRADING_SYSTEM?multiStatements=true' down"

db-reset:
	make db-rollback && make db-migrate && make db-seed

# generate
## generate proto
gen-proto:
	$(shell ./scripts/protoc-generate.sh)


# create entity
create-entity:
	$(shell ./scripts/create-entity.sh)
# create-entity-batch:
# 	docker exec -it ${BATCH_SERVICE_CONTAINER_NAME} sh -c "sqlboiler mysql -c database/sqlboiler.toml -o database/entity -p entity --no-tests --wipe"

# go library install
go-add-library-batch:
	docker exec -it ${BATCH_SERVICE_CONTAINER_NAME} sh -c "go get ${name}"