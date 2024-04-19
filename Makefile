include .env
## 
#=================#
# HELPERS         #
#=================#

## help: print this help message
.PHONY: help # if there was a file ./help , it would be called with make help. with a PHONY, it won't
help:
	@echo 'Needs Golang Migrate: go install -tags "postgres,mysql,sqlite,sqlite3" github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0'
	@echo
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# Create a confirmation prerequisite
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

#=================#
# DEVELOPMENT     #
#=================#

## run: run the web application
.PHONY: run
run:
	go run ./cmd/app

## test: run all tests in the application
.PHONY: test
test:
	go test ./...

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${DATABASE_DNS}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm # needs confirmation
	@echo 'Running up migrations...'
	migrate -path=./migrations --database=${DATABASE_DNS} up

## db/migrations/down: apply down database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	migrate -path=./migrations --database=${DATABASE_DNS} down

