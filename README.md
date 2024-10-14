# Order management

Order management app to serve customers more efficiently.

## Table of Contents

- [Background](#background)
- [Diagrams](#diagrams)
- [Tech Stack](#tech-stack)
- [Requirements](#requirements)
- [Install](#install)
  - [Dev Environment](#dev-environment)
  - [Migrations](#migrations)
  - [Dev Environment](#dev-environment)
  - [Application startup (using Docker)](#application-startup-using-docker)
  - [Swagger](#swagger)
- [Maintainers](#maintainers)

## Background

The goal of this app is to provide an order management service for fast-food stores.

## Diagrams

- Brain Storming
- Event Storming
- Vertical Flowchart

https://miro.com/app/board/uXjVK6bllSY=/

## Tech Stack

- Golang
- MongoDB
- Swagger
- Docker

## Requirements

- Docker

## Install

In order to install the app locally, ensure you have Docker installed and then run the command bellow:

```
docker compose up
```

The Swagger will be available on http://localhost:8080/swagger/index.html.

### Dev Environment

The instructions below are only required for development environments.

- From the root folder, run `docker-compose up` to install the required images.

- Create a `.env` file based off the `.env.example` on the root folder and fill up the variables accordingly:

```
# API
APP_ENV=""

# Database
DB_STORAGE=""
DB_HOST=""
DB_PORT=""
DB_USER=""
DB_PASSWORD=""
DB_NAME=""

```

- Create the database with `make migrate_up` (see more details in the [Migrations](#migrations) section).

- Run `make run` to start the server.

### Migrations

In order to run migrations please install [migrate](https://github.com/golang-migrate/migrate).
The migration commands can be found on the `Makefile`.

### Application startup (using Docker):

- From the root folder, execute the command `docker compose up -d`.
- The API will be running on port `8080`.

### Swagger

The Swagger is available on the path `/swagger/index.html`.

Whenever new annotations are added in the codebase please run the command below and commit the changes to this repository.
Ensure you have the Swag CLI installed.

```
swag init -g ./cmd/main.go -o cmd/docs
```

## Maintainers

| Member                        | Info     |
| ----------------------------- | -------- |
| Caio Martins Pereira          | RM357712 |
| Maria Eduarda da Luz Meregali | RM356466 |
| Rafael de Souza Ribeiro       | RM357622 |
| Thaís Oliveira de Moura       | RM357737 |
| Victor Toschi                 | RM356847 |
