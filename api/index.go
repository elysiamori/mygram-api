package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.App) {
	fmt.Printf("<h1>Golang Fiber</h1>")
}
