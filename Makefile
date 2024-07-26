POSTGRESQL_URL ?= postgres://docker:123@localhost:5432/tech_challenge_fiap?sslmode=disable
MONGO_URL ?= mongodb://localhost:27017/tech_challenge_fiap

run:
	go run cmd/main.go

migrate_create_pg:
	migrate create -ext sql -dir internal/db/migrations/postgres -seq ${NAME}

migrate_up_pg:
	migrate -database ${POSTGRESQL_URL} -path internal/db/migrations/postgres up

migrate_down_pg:
	migrate -database ${POSTGRESQL_URL} -path internal/db/migrations/postgres down

migrate_create_mongo:
	migrate create -ext json -dir internal/db/migrations/mongo -seq ${NAME}

migrate_up_mongo:
	migrate -database ${MONGO_URL} -path internal/db/migrations/mongo up

migrate_down_mongo:
	migrate -database ${MONGO_URL} -path internal/db/migrations/mongo down
