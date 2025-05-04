package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0067h/gofiber-auth-server/ent"
	"github.com/n0067h/gofiber-auth-server/internal/apperror"
	"github.com/n0067h/gofiber-auth-server/internal/dto"
	"github.com/n0067h/gofiber-auth-server/internal/models"
	"github.com/n0067h/gofiber-auth-server/internal/utils"
)

func Register(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body dto.RegisterRequest
		if err := c.BodyParser(&body); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": apperror.ErrInvalidFields.Error(),
				})
		}

		if err := body.Validate(); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": apperror.ErrMissingFields.Error(),
				})
		}

		email, name, pass := body.Email, body.Name, body.Pass

		hashed, err := utils.HashPassword(pass)
		if err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"error": apperror.ErrInternalServerError.Error(),
				})
		}

		u, err := models.CreateUser(c.Context(), client, email, name, hashed)
		if err != nil {
			if err == apperror.ErrEmailExists {
				return c.
					Status(fiber.StatusConflict).
					JSON(fiber.Map{
						"error": apperror.ErrEmailExists.Error(),
					})
			}

			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"error": apperror.ErrInternalServerError.Error(),
				})
		}

		return c.
			Status(fiber.StatusCreated).
			JSON(fiber.Map{
				"id":    u.ID,
				"email": u.Email,
				"name":  u.Name,
			})
	}

}

func Login(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body dto.LoginRequest

		if err := c.BodyParser(&body); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": apperror.ErrInvalidFields.Error(),
				})
		}

		if err := body.Validate(); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": apperror.ErrMissingFields.Error(),
				})
		}

		email, pass := body.Email, body.Pass

		u, err := models.ReadUser(c.Context(), client, email)
		if err != nil {
			return c.
				Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"error": apperror.ErrInvalidCredentials.Error(),
				})
		}

		if err := utils.CheckPasswordHash(u.Password, pass); err != nil {
			return c.
				Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"error": apperror.ErrInvalidCredentials.Error(),
				})
		}

		return c.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"id":    u.ID,
				"email": u.Email,
				"name":  u.Name,
			})
	}
}
