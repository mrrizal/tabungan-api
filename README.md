# Tabungan API

## Run database
`docker-compose up`

## Migrate up
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/tabungan?sslmode=disable" up`

## Migrate down
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/tabungan?sslmode=disable" down`

## Run
`go run main.go`


todo:
- write unittest
- use interfaces/abstract for database
- dockerize golang app
- pagination for mutasi endpoint
