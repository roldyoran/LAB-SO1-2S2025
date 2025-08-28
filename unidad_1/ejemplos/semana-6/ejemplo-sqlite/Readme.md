Este ejemplo hace lo siguiente:

1. Crea una **base de datos SQLite** en la misma carpeta del programa.
2. Define una tabla `containers` para guardar la información.
3. Cada **20 segundos**, genera hasta **4 contenedores con datos aleatorios** (nombre, CPU, memoria, estado) y los guarda (Esto es solo para simular lo que ustedes deben hacer con /proc/continfo).
4. Se ejecuta como **daemon** en segundo plano.

---

## 📌 Compilación

En la carpeta del código:

```bash
go mod init daemon-test
go get github.com/mattn/go-sqlite3
go build -o daemon main.go
```

---

## 📌 Ejecución manual

```bash
./daemon &
```

(`&` lo manda a segundo plano)

Ver logs en tiempo real:

```bash
tail -f nohup.out
```

Detenerlo:

```bash
pkill daemon
```

---

## 📌 Instalar como Daemon en Linux (systemd)

Crea el archivo de servicio:

```bash
sudo nano /etc/systemd/system/grafana-db-daemon.service
```

Contenido:

```txt
[Unit]
Description=Daemon en Go - (descripcion corta)
After=network.target

[Service]
ExecStart=/ruta/a/tu/daemon
WorkingDirectory=/ruta/a/tu/carpeta
Restart=always

[Install]
WantedBy=multi-user.target
```

Guardar y habilitar:

```bash
sudo systemctl daemon-reload
sudo systemctl enable grafana-db-daemon
sudo systemctl start grafana-db-daemon
```

Ver estado:

```bash
sudo systemctl status grafana-db-daemon
```

Parar y quitar el daemon
```bash
sudo systemctl stop grafana-db-daemon
sudo systemctl disable grafana-db-daemon
sudo rm /etc/systemd/system/grafana-db-daemon.service
sudo systemctl daemon-reload
```

---

## GRAFANA
### 📂 Crear carpeta para Grafana

Crea una carpeta llamada `grafana` en la misma ubicación donde está tu código de Golang. Si decides crearla en otra ruta, asegúrate de actualizar la ruta correspondiente en el archivo `docker-compose.yml`.

```bash
mkdir grafana
cd grafana
```

---

## 📌 Paso 2: Dockerfile de Grafana con SQLite plugin

Grafana por defecto **no soporta SQLite** como datasource, pero hay un **plugin** de la comunidad llamado `frser-sqlite-datasource`.

Crea el archivo `Dockerfile`:

```dockerfile
FROM grafana/grafana:11.0.0

# Instalar plugin de SQLite
RUN grafana-cli plugins install frser-sqlite-datasource

# Hacer persistente la DB (montada desde host)
VOLUME ["/var/lib/grafana", "/db"]

# Puerto por defecto
EXPOSE 3000
```

---

## 📌 Paso 3: docker-compose.yml

Crea un `docker-compose.yml`:

```yaml
version: "3.9"

services:
  grafana:
    build: .
    container_name: grafana-sqlite
    ports:
      - "3000:3000"
	user: "0" 
    volumes:
      - ./grafana-data:/var/lib/grafana   # Datos de grafana
      - ../containers.db:/db/containers.db  # Montamos la DB creada por GO
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS=frser-sqlite-datasource
```

---

## 📌 Paso 4: Levantar Grafana

En la carpeta de grafana:

```bash
docker-compose up -d
```

Grafana estará en:
*http://localhost:3000*
Usuario: `admin`
Password: `admin`

---

## 📌 Paso 5: Configurar Data Source SQLite

1. En Grafana, ve a **Connections → Data Sources → Add new data source**.
2. Busca **SQLite** (plugin `frser-sqlite-datasource`).
3. Configura el path al archivo dentro del contenedor:

   ```
   /db/containers.db
   ```
4. Guarda y prueba.

---

## 📌 Paso 6: Crear Dashboard

Ya puedes hacer queries como:

```sql
SELECT 
  name,
  cpu,
  memory,
  status,
  created_at / 1000 as time  -- ← DIVIDIR entre 1000 para que Grafana lo tome en segundos
FROM containers
ORDER BY memory DESC
LIMIT 10
```


