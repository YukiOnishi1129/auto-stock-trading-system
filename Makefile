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


# db migration
db-migrate:
	migrate -source file://micro-service/batch-service/migrations -database 'mysql://user:password@tcp(localhost:3306)/AUTO_STOCK_TRADING_SYSTEM?multiStatements=true' up

db-rollback:
	migrate -source file://micro-service/batch-service/migrations -database 'mysql://user:password@tcp(localhost:3306)/AUTO_STOCK_TRADING_SYSTEM?multiStatements=true' down

# generate
## generate proto
gen-proto:
	$(shell ./scripts/protoc-generate.sh)