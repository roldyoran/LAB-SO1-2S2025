# Usa una imagen base ligera con Python
FROM python:3.11-alpine

# Instala numpy
RUN pip install numpy

# Copia el script al contenedor
COPY liviana.py /liviana.py

# Ejecuta el script
CMD ["python", "/liviana.py"]