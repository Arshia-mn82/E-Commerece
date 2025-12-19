package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(jwtSecret string) fiber.Handler {
	secret := []byte(jwtSecret)

	return func(c fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid authorization header"})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return secret, nil
		})

		if err != nil || !tok.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "inavlid token claim"})
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token subject",
			})
		}
		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token subject",
			})
		}

		c.Locals("userID", uint(sub))
		c.Locals("role", role)

		return c.Next()
	}
}
