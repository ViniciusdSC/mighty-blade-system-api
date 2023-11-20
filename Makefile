help:
	@echo "Available commands: \n"
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
	@echo "\n"

run-api: # run api
	go run cmd/api/main.go

docker-compose-develop-up: # run develop containers
	docker-compose -f develop-docker-compose.yml up -d

migrate-db-dev: # run migration for developement enviroment
	sql-migrate up -env="development" -config=dbconfig.yml

revert-migration-db-dev: # revert migration for development enviroment
	sql-migrate down -env="development" -config=dbconfig.yml

mockery: # generate all mocks
	docker run -v "$(PWD)":/src -w /src vektra/mockery