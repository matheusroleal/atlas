db-up: ## Start DB Strucutre
	docker-compose -f docker-compose.yml up

db-migration: ## Create DB Strucutre
	mysql --user="root" --password="password" --port 6603 --host 0.0.0.0 < "infra/db/structure.sql"

setup:  ## Install Go Modules
	go mod download

run:  ## Start project locally
	go run src/main.go

help:
	@echo "---------------- HELP ---------------------" 
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\://'| sed -e 's/##//'