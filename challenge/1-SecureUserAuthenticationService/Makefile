DBNAME = testapp_db
CONTAINER_NAME = postgres-test-app
IMAGE = postgres:12.0-alpine
PORTS = 5432:5432
ENV_VARS = POSTGRES_PASSWORD=secret -e POSTGRES_USER=root
POSTGRESQL_URL = "postgresql://root:secret@localhost:5432/testapp_db?sslmode=disable"

postgres:
	@if [ "$$(docker ps -aq -f name=$(CONTAINER_NAME))" ]; then \
		echo "Container $(CONTAINER_NAME) exists"; \
		if [ "$$(docker ps -q -f name=$(CONTAINER_NAME))" ]; then \
			echo "Docker $(CONTAINER_NAME) is running"; \
		else \
			echo "Docker $(CONTAINER_NAME) is stopped"; \
			docker start $(CONTAINER_NAME); \
		fi; \
	else \
		docker run --name $(CONTAINER_NAME) -p $(PORTS) -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d $(IMAGE); \
	fi

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DBNAME)

dropdb:
	docker exec -it $(CONTAINER_NAME)  dropdb $(DBNAME)

run:
	go run main.go	