FROM golang:1.22-alpine as builder

RUN adduser --disabled-password fiap

USER fiap

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/api cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

ENTRYPOINT ["/app/api"]
