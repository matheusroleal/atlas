db-up: ## Start DB Strucutre
	docker-compose -f docker-compose.yml up

db-migration: ## Create DB Strucutre
	mysql --user="root" --password="password" --port 6603 --host 0.0.0.0 < "infra/db/structure.sql"

setup:  ## Install Go Modules
	go mod download

run-local: ## Start project locally
	go run src/main.go

unit-test: ## Start unit test
	printf "\n${COLOR_YELLOW}Executing Tests\n${COLOR_RESET}" && \
	go test -v -p 1 -count=1 -cover -coverprofile=/tmp/coverage.tmp ${TEST_FILE} && \
	go tool cover -html=/tmp/coverage.tmp -o coverage.html && \
	rm /tmp/coverage.tmp \

help:
	@echo "---------------- HELP ---------------------" 
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\://'| sed -e 's/##//'