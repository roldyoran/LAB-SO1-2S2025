FROM rust:alpine AS builder

# Crear directorio de trabajo
WORKDIR /app

RUN apk add --no-cache protobuf gcc musl-dev

# Copiar archivos de configuración
COPY . .

# Compilar la aplicación en modo release
RUN cargo install --path .



# Imagen final más pequeña
FROM alpine

# Crear directorio de trabajo
WORKDIR /app

# Copiar el binario compilado desde el builder
COPY --from=builder --chown=app_user:app_user /app/target/release/api_rust ./api_rust

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Agregar variable de entorno para la URL del servidor gRPC
ENV GRPC_SERVER_URL=server-go-service:50051

RUN adduser -D app_user
USER app_user

# Comando para ejecutar la aplicación
CMD ["./api_rust"]