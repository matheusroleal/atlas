setup:
	go mod download

run:
	go run ozy/main.go

help:
	@echo "---------------- HELP ---------------------" 
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\://'| sed -e 's/##//'