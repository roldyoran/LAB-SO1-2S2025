
# Cómo construir la imagen Docker

1. Abre una terminal y navega a la carpeta donde está este Dockerfile.

2. Ejecuta el siguiente comando para construir la imagen y nombrarla `fiber-api-go:liviana`:

```sh
docker build -t fiber-api-go:liviana .
```

Esto creará una imagen optimizada y ligera lista para ejecutarse.

# Cómo ejecutar el contenedor

Puedes iniciar el contenedor con:

```sh
docker run -d -p  8081:8081 fiber-api-go:liviana
```

Esto expondrá el servicio en el puerto 8081 de tu máquina.

# Cómo probar el contenedor con curl

Puedes probar que el contenedor está funcionando correctamente usando los siguientes comandos en otra terminal:

Para probar la ruta principal:

```sh
curl http://localhost:8081/
```

Para probar la ruta random (devuelve un nombre y edad aleatorios en formato JSON):

```sh
curl http://localhost:8081/random
```
