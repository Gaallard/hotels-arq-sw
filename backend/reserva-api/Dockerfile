# Dockerfile for hotels-api
FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o reservas .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/reservas .
EXPOSE 8082
ENTRYPOINT [ "/app/reservas" ]