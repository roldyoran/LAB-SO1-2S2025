# Creacion de la imagen con Docker y su uso

## Construcción de la imagen Docker
Para construir la imagen Docker, ejecuta el siguiente comando en el directorio donde se encuentra el archivo `Dockerfile`:

```bash
docker build -t api-go:pesada .
```

## Ejecución del contenedor
Para ejecutar el contenedor basado en la imagen creada, utiliza el siguiente comando:

```bash
docker run -d -p 8080:80 --name mi-api-go api-go:pesada
```
***--name mi-api-go*** es opcional

Esto expondrá el servidor en el puerto 8080 de su máquina local.

## Prueba del servidor
puede probar el servidor utilizando `curl` con el siguiente comando:

```bash
curl http://localhost:8080
```

Esto debería devolver el mensaje de respuesta del servidor.
Si eso pasa significa que el contenedor esta funcionando como debe.

