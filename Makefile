mod install: export GO111MODULE=on
mod install: export GOPROXY=direct
mod install: export GOSUMDB=off

migration-run: export POSTGRESQL_URL=postgresql://diagnosis_svc:diagnosis_svc@localhost:5434/diagnosis_svc?sslmode=disable
migration-test-run: export POSTGRESQL_URL=postgresql://diagnosis_svc:diagnosis_svc@localhost:5435/diagnosis_svc?sslmode=disable

.PHONY: mod
## Install project dependencies using go mod. Usage 'make mod'
mod:
	@go mod tidy
	@go mod vendor

.PHONY: mock
## Generate mock files. Usage: 'make mock'
mock: ; $(info Generating mock files)
	@chmod +x generate-mocks.sh 
	@./generate-mocks.sh

.PHONY: run
## Run service. Usage: 'make run'
run: ; $(info Starting svc...)
	go run --tags dev ./cmd/server/.

.PHONY: migration-create
migration-create: ## Creates a new migration usage: `migration-create name=<migration name>`
	@migrate create -dir ./migrations -ext sql $(name)

.PHONY: migration-run
migration-run: ## Running migrations: `migration-run dir=[up,down] (optional count=[number of migrations])`
	$(info Running migrations...)
	@migrate -database ${POSTGRESQL_URL} -path ./migrations $(dir) $(count)

.PHONY: test
## Run tests. Usage: 'make test' Options: path=./some-path/... [and/or] func=TestFunctionName
test: ; $(info running tests…) @
	@docker compose -f ./docker-compose_test.yml up --quiet-pull --force-recreate -d --wait;
	make migration-test-run dir=up;
	@if [ -z $(path) ]; then \
		path='./...'; \
	else \
		path=$(path); \
	fi; \
	if [ -z $(func) ]; then \
		go test -v -failfast $$path; \
	else \
		go test -v -failfast -run $$func $$path; \
	fi;
	docker compose -f ./docker-compose_test.yml down;

.PHONY: migration-test-run
migration-test-run: ## test purposes
	$(info Running migrations...)
	@migrate -database ${POSTGRESQL_URL} -path ./migrations $(dir) $(count)

.PHONY: seed
seed:
	go run cmd/seed/main.go

.PHONY: lint
lint:
	@golangci-lint run ./...
	