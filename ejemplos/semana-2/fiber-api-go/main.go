package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inicializar generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Listas de nombres y edades posibles
	names := []string{"Juan", "María", "Carlos", "Ana", "Luis", "Sofía"}
	ages := []int{25, 30, 22, 35, 28, 40}

	// Crear app Fiber
	app := fiber.New()

	// Middleware de logging (cada petición se registra y se muestra en consola con los logs)
	app.Use(func(c *fiber.Ctx) error {
		println("Petición recibida:", c.Method(), c.Path())
		return c.Next()
	})

	// Ruta raíz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola, estoy usando Fiber!")
	})

	// Ruta random
	app.Get("/random", func(c *fiber.Ctx) error {
		// Seleccionar nombre y edad aleatorios
		randomName := names[rand.Intn(len(names))]
		randomAge := ages[rand.Intn(len(ages))]

		// Devolver JSON
		return c.JSON(fiber.Map{
			"nombre": randomName,
			"edad":   randomAge,
		})
	})

	// Iniciar servidor en puerto 8081 en 0.0.0.0 que es accesible desde cualquier IP
	app.Listen("0.0.0.0:8081")
}