DOCKER_CMP_COMMAND=docker-compose
DOCKER_CMP_COMMAND_EXEC=$(DOCKER_CMP_COMMAND) exec

run:
	$(DOCKER_CMP_COMMAND) up --build

test:
	$(DOCKER_CMP_COMMAND_EXEC) api go test -v ./...

audit:
	$(DOCKER_CMP_COMMAND_EXEC) api go mod tidy

lint:
	$(DOCKER_CMP_COMMAND_EXEC) api gofmt -w .

doc:
	swag init -g app/url_mappings.go

pipeline/qa:
	make lint && \
	make audit && \
	make test

shutdown:
	$(DOCKER_CMP_COMMAND) down