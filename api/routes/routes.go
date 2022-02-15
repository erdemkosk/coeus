package routes

import (
	"time"

	handlers "github.com/erdemkosk/go-config-service/api/handlers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/erdemkosk/go-config-service/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GenerateFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	return app
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)

	api := app.Group("/api")
	api.Get("/config/:key", handlers.GetConfig)
	api.Put("/config/:key", middleware.Protected(), handlers.UpdateConfig)
	api.Post("/config/", middleware.Protected(), handlers.CreateConfig)
	api.Delete("/config/:key", middleware.Protected(), handlers.DeleteConfig)

}
