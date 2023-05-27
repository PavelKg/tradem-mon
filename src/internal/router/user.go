package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
	"github.com/pavelkg/tradem-mon-api/internal/middleware"
	// _ "github.com/pavelkg/tradem-mon-api/internal/middleware"
)

func NewUserRoutes(router fiber.Router, userHandlers model.UserPresenter, key string) {

	//var featureName = "user"
	// Create "auth" route group.
	routeAuth := router.Group("/auth")
	routeAuth.Post("/login", userHandlers.LoginUser)
	// Only JWT protected allow
	routeAuth.Use(middleware.JWTProtected(key))

	routeAuth.Get("/me",
		userHandlers.GetUserPersonalProps)

	// Create "users" route group
	routeUsers := router.Group("/users")

	// Only JWT protected allow
	routeUsers.Use(middleware.JWTProtected(key))

	routeUsers.Get("/", userHandlers.Get)        // get list of the users
	routeUsers.Get("/:id", userHandlers.GetById) // get user by id (login)

	// Only for a user with permissions

	// Creates a new user
	routeUsers.Post("/",
		userHandlers.Create)

	// Updates user data
	routeUsers.Put("/:id",
		userHandlers.Update)

	// Deletes a user
	routeUsers.Delete("/:id",
		userHandlers.Delete)
}
