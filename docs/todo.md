# TODO List

## Frontend

## Backend

- Learn more Gin and PGX to set up handlers and SQL seeders
- Create schema design with pgx files and seed database with migrations
- Create handlers for http requests to the backend
- Setup pulling data from 3rd party APIs with requests and sending it to front end

## DB

- Get DB seeded with test data

To migrate: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up`
To undo last migration: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1`
To drop all migrations: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down`

## Docker

- Build and ensure Compose works, that I can ping back end for message and receive front end.
- Ensure hot reloading is working for development

## Misc

- Finish ERD and other Design Documents
