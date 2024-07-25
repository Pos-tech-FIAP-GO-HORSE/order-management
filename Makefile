POSTGRESQL_URL ?= postgres://docker:123@localhost:5432/tech_challenge_fiap?sslmode=disable

run:
	go run cmd/main.go

migrate_create:
	migrate create -ext sql -dir internal/db/migrations -seq ${NAME}

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path internal/db/migrations down
