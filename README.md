# Order management

Order management service to serve customers more efficiently.

## Installation:

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
