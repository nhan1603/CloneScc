# ----------------------------
# Env Variables
# ----------------------------
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
DATABASE_CONTAINER ?= database
API_CONTAINER ?= server

# ----------------------------
# api Methods
# ----------------------------
api-setup: pg-create pg-migrate kafka sleep kafka-topic
api-run:
	docker compose up -d ${API_CONTAINER}

api-test: 
	cd api; \
		env $$(cat ./local.env | egrep -v '^#' | xargs) \
	 	sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./... | grep -v pkg" \
	cd..

## Execute CLI command to generator new database models
api-gen-models:
	 cd api && sh -c "sqlboiler psql"

api-gen-mocks:
	docker compose -f ${DOCKER_COMPOSE_FILE} run --name mockery --rm -w /api --entrypoint '' mockery /bin/sh -c "\
		mockery --dir internal/controller --all --recursive --inpackage && \
		mockery --dir internal/repository --all --recursive --inpackage"

# ----------------------------
# simulator
# ----------------------------
api-simulator-run:
	docker compose up -d simulator

# ----------------------------
# database Methods
# ----------------------------
## Start postgres database container only
pg-create:
	docker compose up -d ${DATABASE_CONTAINER}

## Create a DB migration files e.g `make migrate-create name=test`
migrate-create:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

## Run migrations UP
pg-migrate:
	docker compose --profile tools run --rm migrate up

## Run migrations DROP
pg-drop:
	docker compose --profile tools run --rm migrate drop

## Rollback migrations against non test DB
pg-redo:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down

# ----------------------------
# Base Methods
# ----------------------------
zookeeper:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d zookeeper

kafka:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d kafka

kafka-topic:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d kafka-topic

sleep:
	sleep 5

teardown:
	docker compose -f ${DOCKER_COMPOSE_FILE} down -v
	docker compose -f ${DOCKER_COMPOSE_FILE} rm --force --stop -v
