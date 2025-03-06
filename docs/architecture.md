# Architecture Documentation

## Overview

This application is a web server built using **Go**, with **HTMX** for dynamic frontend interactions and **PostgreSQL** as the database. It follows a **layered architecture** to separate concerns, enhance testability, and support scalability. The key layers are:

- **Models**: Simple data structures representing database entities.
- **Database (DB)**: Manages raw database interactions.
- **Services**: Encapsulates business logic and validation.
- **Handlers**: Handles HTTP requests and responses, including template rendering.

The core logic resides in the `internal` directory to ensure privacy and prevent external packages from accessing it directly.

## Directory Structure

```txt
├── cmd                 # Application entry point
├── docs                # Application documentation
├── internal
│   ├── db              # Database connection setup and low-level database operations
│   ├── handlers        # Handler initialization and HTTP method handlers
│   ├── models          # Data models
│   ├── services        # Business logic
│   └── utils           # Miscellaneous utility methods
└── static
    ├── css             # Stylesheets
    └── html            # HTMX templates
```

#### Reasoning

- `cmd`: Houses `main.go` as the entry point, keeping it separate from the core logic.
- `internal`: Enforces Go's convention of keeping core logic private, preventing external misuse.
- `db`: Splits database setup (`db.go`) from entity-specific operations (`users.go`) for clarity and reusability.
- `handlers`: Centralizes shared resources (e.g. templates) in `handlers.go` and organizes endpoint logic (e.g. `users.go`) for modularity.
- `models`: Defines plain structs like `User` for data consistency across layers.
- `services`: Isolates business logic (e.g., password hashing) to keep handlers lightweight.
- `static`: Standard locations for frontend assets, simplifying management and serving.

## Database Layer (`internal/db`)

- `db.go`: Contains a `Config` struct for connection settings and a `DB` struct wrapping `*sql.DB` for database interactions.
- `users.go`: Implements `UserStore` for user-specific SQL operations (e.g., CRUD).

#### Reasoning

- **Config Struct**: Enables flexible, configurable database connections without hardcoding values.
- **DB Wrapper**: Allows future enhancements (e.g., connection pooling) while keeping the global namespace clean.
- **Separation**: Isolates database logic from business rules, easing testing and potential database swaps.

## Models Layer (`internal/models`)

- `users.go`: Defines the `User` struct with fields matching the database schema, acting as a data container without logic.

#### Reasoning

- **DTO Role**: Ensures type-safe data transfer between layers.
- **Simplicity**: Keeps structs free of behavior, making them versatile for serialization or mapping.

## Services Layer (`internal/services`)

- `users.go`: Implements `UserService` for business logic, such as password hashing, input validation, and coordinating database calls.

#### Reasoning

- **Encapsulation**: Centralizes rules (e.g., "validate email") to keep handlers focused on HTTP tasks.
- **Reusability**: Services can support multiple contexts (e.g., web and CLI)
- **Testability**: Easy to test independently of HTTP or database layers.

## Handlers Layer (`internal/handlers`)

- `handlers.go`: Initializes all handler groups (e.g., `UserHandlers`) and shared resources like templates.
- `users.go`: Contains `UserHandlers` with HTTP endpoint functions for user operations.

#### Reasoning

- **Centralized Setup**: Avoids redundant initialization of templates or services across files.
- **Modularity**: Handler groups (e.g., users) can evolve independently.
- **Thin Logic**: Handlers delegate to services, focusing on request/response handling.

## Dependency Flow

- `main.go`: Loads configuration, sets up the database (`db.DB`), and passes it to `handlers.New`.
- `handlers.New`: Instantiates services and stores with the database connection, then configures handler groups.
- **Handler Functions**: Call services for logic and use templates for rendering.

#### Reasoning

- **Dependency Injection**: Explicitly passing dependencies improves testability and flexibility.
- **Clarity**: Each layer only accesses what it needs, reducing coupling.

## Template Rendering

- Templates are parsed once in `handlers.New` and stored in a shared `tmpl` field.
- Handlers render specific templates using `ExecuteTemplate` with relevant data.

#### Reasoning

- **Efficiency**: Parsing at startup avoids repeated I/O during requests.
- **Consistency**: A single template set ensures uniform rendering across handlers.

## Error handling

- Lower layers (database, services) return errors to handlers.
- Handlers log errors and respond with appropriate HTTP status codes (e.g., 500 for server errors, 400 for client errors).

#### Reasoning

- **Centralized Control**: Handlers manage user-facing responses, keeping lower layers focused on their tasks.
- **Security**: Detailed errors are logged but not exposed to users.

## Scalability Considerations

- **New Entities**: Add files for the entity (e.g., `pokemon.go`) in `models`, `db`, `services`, and `handlers`, then update the `Handlers` struct.
- **Shared Resources**: Templates and database connections are initialized once and reused.

#### Reasoning

- **Modularity**: New features integrate without disrupting existing code.
- **Performance**: Reusing resources minimizes overhead.

## Testing

- **Unit Tests**: Test services and database layers by mocking dependencies (e.g., mock `UserStore` for `UserService`)
- **Integration Tests**: Test handlers with a real or in-memory database.

#### Reasoning

- **Isolation**: Layer-specific tests are simpler and faster.
- **Mocking**: Dependency injection enables reliable, quick tests.

