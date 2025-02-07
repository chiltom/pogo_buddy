# TODO List

## Frontend

## Backend

- Seed containerized database and test CRUD capabilities
- Build out handlers and write tests as each model gets CRUD handlers
- Build out routes to connect to mux router
- Setup pulling data from 3rd party APIs with requests and sending it to front end

## DB

- Get DB seeded with test data

> To migrate: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up`
> To undo last migration: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1`
> To drop all migrations: `migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down`

## Docker

- Write production environment Dockerfiles and compose file to ensure smooth transition to production site
    - Build in .env.production file to ensure that secrets are well maintained
- Ensure hot reloading is working for development

## Misc

- Create CI/CD pipeline with GitHub Actions to lint and test code on push
- Finish ERD and other Design Documents
