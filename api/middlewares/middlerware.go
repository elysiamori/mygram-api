package middlewares

import (
	"net/http"

	"github.com/elysiamori/mygram-api/helpers"
	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := helpers.ExtractTokenID(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		c.Locals("user_id", userID)
		return c.Next()
	}
}
