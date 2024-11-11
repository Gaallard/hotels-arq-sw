FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . . 
RUN go build -o users .  

# Usa una imagen base de Alpine para el contenedor final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/users .  
EXPOSE 8081  
ENTRYPOINT ["/app/users"] 