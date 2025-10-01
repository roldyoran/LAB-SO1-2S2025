# Multi-stage build para optimizar el tamaño de la imagen
FROM rust:1.90 AS builder

# Crear directorio de trabajo
WORKDIR /app

RUN apt-get update && apt-get install -y protobuf-compiler ca-certificates

# Copiar archivos de configuración
COPY Cargo.toml Cargo.lock ./
COPY build.rs ./
COPY proto/ ./proto/

# Copiar código fuente
COPY src/ ./src/

# Compilar la aplicación en modo release
RUN cargo build --release

# Imagen final más pequeña
FROM ubuntu:jammy

# Instalar dependencias de runtime necesarias
RUN apt-get update && apt-get install -y locales ca-certificates && rm -rf /var/lib/apt/lists/* \
	&& localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

ENV LANG=en_US.utf8

# Crear directorio de trabajo
WORKDIR /app

# Copiar el binario compilado desde el builder
COPY --from=builder /app/target/release/api_rust ./api_rust

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Agregar variable de entorno para la URL del servidor gRPC
ENV GRPC_SERVER_URL=server_go_service:50051

# Comando para ejecutar la aplicación
CMD ["./api_rust"]