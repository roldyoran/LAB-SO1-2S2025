#!/bin/bash

# Remove the containers.db file if it exists
if [ -f "containers.db" ]; then
    echo "[INFO] Borrando containers.db..."
    if ! rm -rf containers.db; then
        echo "[ERROR] No se pudo borrar containers.db. Abortando script."
        exit 1
    fi
    echo "[INFO] containers.db borrado."
else
    echo "[INFO] containers.db no existe."
fi


# Remove the daemon binary if it exists
if [ -f "daemon" ]; then
    echo "[INFO] Borrando binario de daemon..."
    if ! rm -rf daemon; then
        echo "[ERROR] No se pudo borrar el binario de daemon. Abortando script."
        exit 1
    fi
    echo "[INFO] Binario de daemon borrado."
else
    echo "[INFO] Binario de daemon no existe."
fi



# ejecutar docker-compose down en la carpeta grafana
echo "--------------------------------"
echo "🛑 Deteniendo contenedores de docker-compose..."
if ! (cd ./grafana && docker-compose down); then
    echo "[ERROR] No se pudo detener los contenedores de docker-compose. Abortando script."
    exit 1
fi
echo "✅ Contenedores detenidos."

if ! sudo systemctl stop containers-daemon; then
    echo "[ERROR] Falló  al parar el daemon."
    exit 1
fi

if ! sudo systemctl disable containers-daemon; then
    echo "[ERROR] Falló al deshabilitar el daemon."
    exit 1
fi

# Build the Go daemon
echo "[INFO] Construyendo Go daemon..."
if ! go build -o daemon main.go; then
    echo "[ERROR] Falló la construcción del Go daemon. Abortando script."
    exit 1
fi

echo "[INFO] Proceso completado exitosamente."

echo "--------------------------------"

echo "⚙️ Cargando servicio en systemd..."

if ! sudo systemctl daemon-reload; then
    echo "[ERROR] Falló la recarga de systemd. Abortando script."
    exit 1
fi

if ! sudo systemctl enable containers-daemon; then
    echo "[ERROR] Falló al habilitar el servicio containers-daemon. Abortando script."
    exit 1
fi

if ! sudo systemctl start containers-daemon; then
    echo "[ERROR] Falló al iniciar el servicio containers-daemon. Abortando script."
    exit 1
fi

echo "✅ Servicio systemd creado e iniciado."



echo "--------------------------------"


echo "🚀 Iniciando contenedores de docker-compose..."
if ! (cd ./grafana && docker-compose up -d); then
    echo "[ERROR] No se pudo iniciar los contenedores de docker-compose. Abortando script."
    exit 1
fi
echo "✅ Contenedores iniciados."

echo "--------------------------------"
echo "✅ Script completado exitosamente."