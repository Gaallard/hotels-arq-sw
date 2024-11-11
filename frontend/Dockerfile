# Etapa de construcción
FROM node:16-alpine AS builder

WORKDIR /app

COPY package*.json ./

# Instala las dependencias
RUN npm install

COPY . .

# Construye la aplicación
RUN npm run build

# Etapa de producción
FROM nginx:alpine

# Copia los archivos compilados de la etapa de construcción
COPY --from=builder /app/build /usr/share/nginx/html

# Copia el archivo de configuración personalizado para nginx
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
