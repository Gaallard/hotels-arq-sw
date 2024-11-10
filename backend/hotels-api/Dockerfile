# Dockerfile for Go API

# Usa una imagen base de Golang
FROM golang:1.20-alpine

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod ./
COPY go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia todo el código fuente al directorio de trabajo
COPY . .

# Compila la aplicación
RUN go build -o hotels-api

# Expone el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./hotels-api"]
