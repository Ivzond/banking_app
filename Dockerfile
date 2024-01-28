# Stage 1: Build the application
FROM golang:1.21.6 AS builder

WORKDIR /app

COPY cmd/app cmd/app
COPY internal internal
COPY go.mod go.sum ./

WORKDIR /app/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go build -o fintech_app .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/app/fintech_app .

EXPOSE 8080

CMD ["./fintech_app"]
