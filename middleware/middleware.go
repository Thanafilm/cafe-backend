package middleware

import (
	"cafe-backend/util"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthoricated",
		})
	}
	return c.Next()
}
