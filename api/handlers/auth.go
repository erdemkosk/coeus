package handlers

import (
	"github.com/dgrijalva/jwt-go"
	models "github.com/erdemkosk/go-config-service/api/models"
	plugin "github.com/erdemkosk/go-config-service/api/plugin"
	"github.com/gofiber/fiber/v2"
)

// Login is a function to get a book by ID
// @Summary Post config with using key
// @Description Get config with using key
// @Tags auth
// @Accept json
// @Produce json
// @Param cridentinials body models.LoginInput true "Cridentinials"
// @Success 200 {object} models.AuthToken
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {

	var input models.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	identity := input.Identity
	pass := input.Password
	if identity != plugin.GetEnvConfig("ADMIN_USER_ID") || pass != plugin.GetEnvConfig("ADMIN_USER_PASS") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["admin"] = true

	t, err := token.SignedString([]byte(plugin.GetEnvConfig("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)

	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
