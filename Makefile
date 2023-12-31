DOCKER_COMPOSE_FILE = build/docker-compose.yaml

.SILENT:

down-compose:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d


unit-tests:
	go test -v ./tests/unit
