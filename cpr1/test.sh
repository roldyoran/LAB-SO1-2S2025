#!/bin/bash

# Función para verificar si un comando está instalado
check_command() {
    if command -v "$1" &> /dev/null; then
        return 0  # verdadero (está instalado)
    else
        return 1  # falso (no está instalado)
    fi
}

# Procesar argumentos
if [ "$1" == "1" ]; then
    echo -e "\n\nVerificando runtime de contenedores VM1 y VM2...\n"
    if check_command "docker"; then
        echo -e "❌ Docker está instalado en la VM cuando no deberia. ❌\n"
    elif check_command "containerd"; then
        echo -e "✅ Containerd está instalado (pero no Docker). ✅\n"
    else
        echo -e "❌ No se encontró ni Docker ni Containerd. ❌\n"
    fi
elif [ "$1" == "2" ]; then
    echo -e "\n\nVerificando runtime de contenedores... VM3\n"
    if check_command "docker"; then
        echo -e "✅ Docker está instalado. ✅\n"
    else
        echo -e "❌ No se encontró ni Docker ni Containerd. ❌\n"
    fi
else
    echo "Uso: $0 <1|2>"
    echo "  1 = Verificar Docker/Containerd VM1 y VM2"
    echo "  2 = Verificar Docker VM3"
    exit 1
fi

