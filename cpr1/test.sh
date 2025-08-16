#!/bin/bash

# Verificar si se proporcionó un argumento
if [ $# -eq 0 ]; then
    echo "Uso: $0 <1|2>"
    exit 1
fi

# Evaluar el argumento
case "$1" in
    1)
        echo "¡Hola!"
        ;;
    2)
        echo "¡Adiós!"
        ;;
    *)
        echo "Opción no válida. Usa 1 o 2."
        exit 1
        ;;
esac