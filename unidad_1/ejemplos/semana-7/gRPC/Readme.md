# gRPC con Golang

## 1. Instalar dependencias necesarias

Asegúrate de tener Go y lo siguiente:

### Instalar Protocol Buffers (protoc)

1. Abre una terminal y actualiza los paquetes de tu sistema:

   ```bash
   sudo apt update && sudo apt upgrade -y
   ```

2. Instala las herramientas necesarias para compilar `protoc`:

   ```bash
   sudo apt install -y build-essential libtool pkg-config protobuf-compiler
   ```

4. Verifica la instalación ejecutando:

   ```bash
   protoc --version
   ```

   Deberías ver la versión instalada de `protoc`.

### Instalar plugins de Go para gRPC

1. Instala los plugins necesarios para generar código Go:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
---

## 2. Generar código Go desde el `.proto`

1. Coloca tu archivo `.proto` (por ejemplo, `tweet.proto`) en el directorio de tu proyecto.

2. Genera el código necesario ejecutando:

   ```bash
   protoc --go_out=. --go-grpc_out=. tweet.proto
   ```

   Esto generará dos archivos en /proto:

   * `tweet.pb.go`
   * `tweet_grpc.pb.go`

---

## 3. Configurar el codigo de Go
Copia el codigo de server.go y client.go

```bash
go mod init grpc
# luego
go mod tidy
```   

---

## 4. Ejecutar el proyecto

1. Inicia el servidor en una terminal:

   ```bash
   go run server.go
   ```

2. En otra terminal, ejecuta el cliente:

   ```bash
   go run client.go
   ```

Deberías ver los mensajes de interacción entre el cliente y el servidor en las respectivas terminales.

