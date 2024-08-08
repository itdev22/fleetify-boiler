package config

import (
	"boilerplate/app/database"
	"boilerplate/app/routes"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func GetConfig() {
	//Read ENV
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)

	if err != nil {
		log.Fatal(err)
	}
	//Read ENV

	// database.ConnectDbMysql()
	database.ConnectDbMysql()
	app := fiber.New(fiber.Config{
		BodyLimit: 500 * 1024 * 1024,
	})

	// Add recover middleware
	app.Use(recover.New())

	// Configuring CORS middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.SetUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}
