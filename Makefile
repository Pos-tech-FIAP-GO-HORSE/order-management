MONGO_URL ?= mongodb://localhost:27017/tech_challenge_fiap

run:
	go run cmd/main.go

swag-init:
	swag init -g ./cmd/main.go -o cmd/docs

migrate_create:
	migrate create -ext json -dir internal/db/migrations/mongo -seq ${NAME}

migrate_up:
	migrate -database ${MONGO_URL} -path internal/db/migrations/mongo up

migrate_down:
	migrate -database ${MONGO_URL} -path internal/db/migrations/mongo down
