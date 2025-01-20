# Build stage
FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-service

# Run stage
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/weather-service .
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
EXPOSE 8080
CMD ["./weather-service"]
