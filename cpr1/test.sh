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
    echo "Verificando runtime de contenedores VM1 y VM2..."
    if check_command "docker"; then
        echo "❌ Docker está instalado."
    elif check_command "containerd"; then
        echo "✅ Containerd está instalado (pero no Docker)."
    else
        echo "❌ No se encontró ni Docker ni Containerd."
    fi
elif [ "$1" == "2" ]; then
    echo "Verificando runtime de contenedores... VM3"
    if check_command "docker"; then
        echo "✅ Docker está instalado."
    else
        echo "❌ No se encontró ni Docker ni Containerd."
    fi
else
    echo "Uso: $0 <1|2>"
    echo "  1 = Verificar Docker/Containerd VM1 y VM2"
    echo "  2 = Verificar Docker VM3"
    exit 1
fi