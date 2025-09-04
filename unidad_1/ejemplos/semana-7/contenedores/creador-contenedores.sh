#!/bin/bash

# Verificar si Docker está instalado
if ! [ -x "$(command -v docker)" ]; then
  echo 'Error: Docker no está instalado.' >&2
  exit 1
fi

# Crear 10 contenedores con nombres aleatorios 
# entre los pesados y los livianos

# Se necesita:
# 2 LIVIANOS (para diferenciar entre ellos)
# 1 CPU HIGH
# 1 RAM HIGH
list_contenedores=("img_pesado:v1" "img_liviana:v1")


for i in {1..4}; do

    # Obtener un índice aleatorio basado en la longitud de la lista de contenedores
    random_index=$((RANDOM % ${#list_contenedores[@]}))
    
    # Seleccionar el tipo de contenedor basado en el índice aleatorio
    tipo_contenedor=${list_contenedores[$random_index]}

    # Generar un nombre aleatorio para el contenedor
    random_name=$(cat /dev/urandom | tr -dc 'a-z0-9' | fold -w 8 | head -n 1)

    # Obtener la hora y minuto actuales
    current_time=$(date +"%H%M%S")

    # Combinar el tipo de contenedor y el nombre aleatorio
    container_name="${tipo_contenedor//:/-}-${i}-${random_name}-${current_time}"

    # Ejecutar los demás contenedores como antes
    docker run -d --name "$container_name" "$tipo_contenedor"
  

    echo "========================================"
    echo "Contenedor NO.$i"
    echo "Nombre: $container_name"
    echo "Tipo: $tipo_contenedor"
    echo "Estado: Creado exitosamente"
    echo "========================================"
    echo " "
    echo " "
    echo " "


done