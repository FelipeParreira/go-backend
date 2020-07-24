DOCKER_CMP_COMMAND=docker-compose
DOCKER_CMP_COMMAND_EXEC=$(DOCKER_CMP_COMMAND) exec

run:
	$(DOCKER_CMP_COMMAND) up --build

run-native:
	go run main.go

test:
	DATA_FILE_PATH="../data.json" go test ./...

audit:
	go mod tidy

lint:
	gofmt -w .

build:
	go build

doc:
	swag init -g app/url_mappings.go

pipeline/qa:
	make lint && \
	make audit && \
	make test

shutdown:
	$(DOCKER_CMP_COMMAND) down