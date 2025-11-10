## 1. Código de un daemon **muy simple** en Go

Este daemon solo escribe un log cada 5 segundos, así podrás ver si corre en segundo plano.

```go
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// Abrir archivo de log
	f, err := os.OpenFile("/var/log/mydaemon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("Daemon iniciado")

	for {
		log.Println("Daemon sigue vivo...")
		time.Sleep(5 * time.Second)
	}
}
```

Compílalo:

```bash
go build -o /usr/local/bin/mydaemon main.go
```

---

## 2. Crear el archivo systemd

Crea `/etc/systemd/system/mydaemon.service` con este contenido:

```ini
[Unit]
Description=Mi daemon en Go
After=network.target

[Service]
ExecStart=/usr/local/bin/mydaemon
Restart=always

[Install]
WantedBy=multi-user.target
```

---

## 3. Activar y arrancar el daemon

```bash
sudo systemctl daemon-reload        # recargar systemd
sudo systemctl enable --now mydaemon
```

---

## 4. Verificar que está cargado y corriendo

* **Estado del servicio:**

  ```bash
  systemctl status mydaemon
  ```

* **Logs en journald:**

  ```bash
  journalctl -u mydaemon -f
  ```

* **Logs en el archivo** que configuramos (`/var/log/mydaemon.log`):

  ```bash
  tail -f /var/log/mydaemon.log
  ```

Deberías ver algo como:

```
Aug 21 16:10:12 myhost mydaemon[1234]: Daemon iniciado
Aug 21 16:10:17 myhost mydaemon[1234]: Daemon sigue vivo...
Aug 21 16:10:22 myhost mydaemon[1234]: Daemon sigue vivo...
```



# APAGAR
  ```bash
sudo systemctl stop mydaemon
sudo systemctl disable mydaemon
  ```