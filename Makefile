include .env

empty:
	echo "empty"

dcb:
	docker compose build
dcu:
	docker compose up
dcd:
	docker compose down

# enter the container
## db container
db-ssh:
	docker exec -it $(MYSQL_HOST) sh
## bff container
bff-ssh:
	docker exec -it $(BFF_CONTAINER) sh
## user service container
user-ssh:
	docker exec -it $(USER_SERVICE_CONTAINER) sh