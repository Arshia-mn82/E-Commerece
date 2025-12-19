package handler

import (
	"E-Commerce/internal/dto"
	"E-Commerce/internal/repository"
	"E-Commerce/internal/service"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	auth  *service.AuthService
	users *repository.UserRepository
}

func NewAuthHandler(auth *service.AuthService, users *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		auth:  auth,
		users: users,
	}
}

func (ah *AuthHandler) SignUp(c fiber.Ctx) error {
	var req dto.SignUpRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "inavlid request body"})
	}

	token, err := ah.auth.Signup(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.AuthResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.AuthResponse{
		Message:     "signup successful",
		AccessToken: token,
	})
}

func (ah *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	token, err := ah.auth.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.AuthResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(dto.AuthResponse{
		Message:     "login successful",
		AccessToken: token,
	})
}

func (ah *AuthHandler) Me(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	u, err := ah.users.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to load user"})
	}
	return c.JSON(fiber.Map{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"role":  u.Role,
	})

}
