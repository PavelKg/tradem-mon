package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
)

// SwaggerRoute func to add swagger group route to src routes
func SwaggerRoute(router fiber.Router) {
	// Create routes group.
	route := router.Group("/swagger")
	// Routes for GET method:
	route.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/index.html",
		},
		StatusCode: 301,
	}))
	route.Get("/*", swagger.HandlerDefault)
}
