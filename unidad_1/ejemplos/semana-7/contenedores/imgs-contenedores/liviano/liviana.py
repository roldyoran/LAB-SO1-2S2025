import time
import numpy as np

# Alocar una cantidad específica de memoria
size = 800 * 1024 * 1024  # 100 MB
array = np.zeros(size)

# memory_hog = []
# for i in array:
#     memory_hog.append('x')  # Agrega cadenas de 1 KB

# Mantener el contenedor en ejecución durante un tiempo
time.sleep(1200)