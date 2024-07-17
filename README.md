# backend-golang

## Overview

This project is built using Go and follows the principles of Clean Architecture. The structure of the project ensures separation of concerns, making the codebase maintainable, testable, and scalable.

## Directory Structure

### 1. app
Contains application configurations such as middleware and routing.

- **middlewares**
  - **auth.go**: Middleware for authentication, verifying tokens or user credentials for accessing specific endpoints.
  - **logger.go**: Middleware for logging, recording requests and responses for debugging and audit purposes.

- **routes**
  - **routes.go**: Sets up the application's routes, connecting HTTP endpoints with the appropriate handlers.

### 2. businesses
Contains the business logic and domain rules.

- **admin**
  - **domain.go**: Defines the domain entities and repository interfaces for the admin, specifying the contracts that must be implemented by the repository layer.
  - **usecase.go**: Contains use cases or application services, implementing the business logic related to admin.

### 3. controller
Handles HTTP requests and directs them to the appropriate use case.

- **admin**
  - **request**
    - **json.go**: Defines the JSON request structures for admin, used for parsing data from client requests.
  - **response**
    - **json.go**: Defines the JSON response structures for admin, used for sending data back to clients.
  - **http.go**: Contains the HTTP handlers for requests related to admin, connecting routes with use cases.
- **base_response.go**: Defines the base response structure used throughout the application.

### 4. drivers
Handles interactions with external resources such as databases.

- **mysql**
  - **admin**
    - **mysql.go**: Implementation of the repository for admin using MySQL, connecting to the database and performing CRUD operations.
    - **record.go**: Defines the database model for admin, mapping domain entities to database tables.
  - **domain_factory.go**: Contains code for initializing and connecting all domain parts to the corresponding repositories.

### 5. mariadb
Contains schema files for the MariaDB database.

- **schema.sql**: SQL file that defines the MariaDB database schema, including tables, columns, and indexes.

### 6. utils
Contains utility functions used throughout the application.

- **utils.go**: General utility functions that can be used by various parts of the application.

### 7. .env
Configuration file that stores environment variables, such as database credentials and other settings.

### 8. docker-compose.yml
Docker Compose configuration file that defines the services, networks, and volumes needed to run the application in a Docker environment.

### 9. Dockerfile
File that defines how to build the Docker image for the application, including installation and configuration steps.

### 10. go.mod
Go module file that defines the module name and dependencies used by the project.

### 11. go.sum
Checksum file to ensure the integrity of dependencies defined in `go.mod`.

### 12. main.go
The main application file containing the `main()` function, the entry point of the application.

### 13. Readme.md
Documentation file that contains information about setting up, running, and using the application.

## Summary
This project structure follows clean architecture principles, separating business logic, application logic, and infrastructure details. Each directory and file has a specific role:

- **app**: Application setup and middleware.
- **businesses**: Business logic and domain rules.
- **controller**: HTTP request handlers and request/response handling.
- **drivers**: Interaction with external resources like databases.
- **mariadb**: Database schema.
- **utils**: General utility functions.
- **.env**: Environment variables.
- **docker-compose.yml**: Docker Compose configuration.
- **Dockerfile**: Docker image definition.
- **go.mod**: Go module dependencies.
- **go.sum**: Dependency checksums.
- **main.go**: Application entry point.
- **Readme.md**: Project documentation.
