package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/pavelkg/tradem-mon-api/internal/presenter"
)

// SetupRoutes func to set up src routes
func SetupRoutes(app *fiber.App, presenters *presenter.Presenters, hostPrefix string, jwtKey string) {

	// FixMe it should be changed after fixing docker routing
	api := app.Group("/api")
	//api := app.Group(hostPrefix)

	api.Get("/dashboard", monitor.New())

	SwaggerRoute(api)
	NewUserRoutes(api, presenters.UserPresenter, jwtKey)
	NotFoundRoute(app)
}
