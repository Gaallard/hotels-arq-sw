FROM golang:1.23.2-alpine AS builder
WORKDIR /app

# Copia go.mod y go.sum para instalar dependencias
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copia el resto de los archivos y construye la aplicación
COPY . .
RUN go build -o users .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/users .
EXPOSE 8081
ENTRYPOINT [ "/app/users" ]
