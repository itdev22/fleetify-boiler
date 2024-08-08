package tests

import (
	"github.com/gofiber/fiber/v2"
)

func GetTest(c *fiber.Ctx) error {
	test := GetTestQuery()

	if len(test) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "No data found",
		})
	}

	return c.Status(200).JSON(test)
}
