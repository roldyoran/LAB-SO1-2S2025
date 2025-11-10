# Multi-stage build para optimizar el tamaño de la imagen
FROM golang:1.25.1-alpine AS builder

# Instalar git (necesario para algunas dependencias de Go)
RUN apk add --no-cache git

# Crear directorio de trabajo
WORKDIR /app

# Copiar archivos de módulo Go
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server_go .

# Imagen final más pequeña
FROM alpine:latest

# Instalar certificados CA para conexiones HTTPS/TLS
RUN apk --no-cache add ca-certificates

# Crear directorio de trabajo
WORKDIR /root/

# Copiar el binario compilado desde el builder
COPY --from=builder /app/server_go .

# Exponer el puerto gRPC
EXPOSE 50051

# Comando para ejecutar la aplicación
CMD ["./server_go"]