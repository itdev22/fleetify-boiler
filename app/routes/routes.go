package routes

import (
	"boilerplate/app/src/tests"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	//Your routes here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Gateway!")
	})
	//Example1
	api := app.Group("/api")

	TestsGroup := api.Group("/tests")
	tests.TestsRoute(TestsGroup)

	// JWT Middleware
	// To use middleware uncomment this
	//middleware.JWTConfig(app)
	// Put your endpoint below here to Make auth with JWT
}
