package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Mensajes de prueba
var mensajes = []string{
	"Daemon en ejecución...",
	"Proceso funcionando correctamente",
	"Escribiendo logs de prueba",
	"Todo sigue bien en el sistema",
	"Mensaje aleatorio generado",
}

// Función para crear un cronjob simple
func crearCronJob() {
	// Script que el cronjob ejecutará
	script := "/home/rol/Documentos/sopes1/semana5DAEMON/ejemplo-docker-cronjob/test_scripty.sh"

	// Verificar que el script existe y tiene permisos de ejecución
	if _, err := os.Stat(script); os.IsNotExist(err) {
		log.Fatalf("El script %s no existe", script)
	}

	// Hacer el script ejecutable
	if err := os.Chmod(script, 0755); err != nil {
		log.Printf("Advertencia: no se pudieron cambiar permisos del script: %v", err)
	}

	// Agregar cronjob (cada minuto)
	cronCommand := fmt.Sprintf("* * * * * %s", script)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("(crontab -l 2>/dev/null; echo \"%s\") | crontab -", cronCommand))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error agregando cronjob: %v, output: %s", err, string(output))
	}

	log.Printf("Cronjob creado correctamente: %s", cronCommand)

	// Verificar que se agregó correctamente
	verifyCmd := exec.Command("crontab", "-l")
	verifyOutput, err := verifyCmd.CombinedOutput()
	if err != nil {
		log.Printf("Error verificando cronjob: %v", err)
	} else {
		log.Printf("Cronjobs actuales:\n%s", string(verifyOutput))
	}
}

func iniciarContenedor(nombre string) {
	cmd := exec.Command("docker", "start", nombre)

	// Establecer el directorio de trabajo del comando
	// Esto le dice al comando cmd que se ejecute en la ruta especificada
	// ruta_carpeta := "su/ruta/a/su/carpeta"
	// cmd.Dir = ruta_carpeta

	if err := cmd.Run(); err != nil {
		log.Printf("Error iniciando contenedor %s: %v", nombre, err)
	} else {
		log.Printf("Contenedor %s iniciado.", nombre)
	}
}

func main() {
	// 1. Iniciar el contenedor api-go
	iniciarContenedor("api-go")

	// 2. Crear cronjob
	crearCronJob()

	// 3. Loop infinito: escribir logs en un archivo
	logFile := "/home/rol/Documentos/sopes1/semana5DAEMON/ejemplo-docker-cronjob/daemon_log.txt"
	for {
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error abriendo archivo: %v", err)
		}

		mensaje := mensajes[rand.Intn(len(mensajes))]
		_, _ = f.WriteString(fmt.Sprintf("%s - %s\n", time.Now().Format(time.RFC3339), mensaje))
		f.Close()

		log.Println("Escribiendo en log:", mensaje)

		time.Sleep(10 * time.Second) // espera 10s entre mensajes
	}
}
