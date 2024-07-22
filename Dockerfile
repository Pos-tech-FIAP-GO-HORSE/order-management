FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api cmd/api/main.go

# COPY /build/api /usr/local/bin

# ENTRYPOINT ["/usr/local/bin/api"]

FROM scratch

COPY --from=builder /build/api /usr/local/bin

ENTRYPOINT ["/usr/local/bin/api"]
