package fiber

import (
	"github.com/S4mkiel/validation/infra/http/controller"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("fiber",
	FiberModule,
	fx.Invoke(RegisterControllers),
	fx.Provide(controller.NewUserController),
)

func RegisterControllers(app *fiber.App, userController *controller.UserController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// controllers
	userController.RegisterRoutes(v1)
}
