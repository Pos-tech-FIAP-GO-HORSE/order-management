FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/api cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

ENTRYPOINT ["/app/api"]
