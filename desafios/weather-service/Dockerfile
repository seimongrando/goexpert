# Build stage
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o weather-service

# Run stage
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/weather-service .
EXPOSE 8080
CMD ["./weather-service"]
