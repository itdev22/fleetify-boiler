package tests

import (
	"github.com/gofiber/fiber/v2"
)

func TestsRoute(r fiber.Router) {
	r.Get("testing-query", GetTest)

}
