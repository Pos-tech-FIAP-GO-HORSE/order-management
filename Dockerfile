FROM golang:1.22-alpine AS builder

RUN adduser --disabled-password fiap

USER fiap

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/api cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/api .
COPY --from=builder /app/internal/db/migrations/mongo ./internal/db/migrations/mongo

EXPOSE 8080

ENTRYPOINT ["/app/api"]
