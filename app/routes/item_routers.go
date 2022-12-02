package routes

import (
	"dont/hexagonal/app/handlers"
	"dont/hexagonal/platform/database"
	"dont/hexagonal/repository"
	"dont/hexagonal/service"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	db := database.ConnectDb()
	iservice := service.NewItemService(repository.NewItemRepositoryDb(db))
	ihandlers := handlers.ItemHandlers{Service: iservice}

	route.Get("/item/:id", ihandlers.GetByID)
	route.Post("/item", ihandlers.AddNew)
}
