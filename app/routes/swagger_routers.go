package routes

import (
	"github.com/gofiber/fiber/v2"

	_ "dont/hexagonal/docs"

	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	//route.Get("*", swagger.HandlerDefault) // get one user by ID
	route.Get("*", swagger.New(
		swagger.Config{ // custom
			URL:       "doc.json",
			ConfigURL: "",
		}))
}