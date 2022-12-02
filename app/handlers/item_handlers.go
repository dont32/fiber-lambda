package handlers

import (
	"dont/hexagonal/dto"
	"dont/hexagonal/errs"
	"dont/hexagonal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ItemHandlers struct {
	Service service.ItemService
}

// GetItem func gets item by given ID.
// @Description Get item by given ID.
// @Summary get item by given ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} dto.ItemResponse
// @Router /api/v1/item/{id} [get]
func (h *ItemHandlers) GetByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errs.NewValidationError(err.Error()).AsMessage())
	}
	var req dto.ItemRequest
	req.ID = id
	res, apperr := h.Service.GetByID(req)
	if apperr != nil {
		return ctx.Status(apperr.Code).JSON(apperr.AsMessage())
	}
	return ctx.JSON(res)
}

// CreateItem func for creates a new item.
// @Description Create a new item.
// @Summary create a new item
// @Tags Item
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Success 200 {object} dto.ItemResponse
// @Router /api/v1/item [post]
func (h *ItemHandlers) AddNew(ctx *fiber.Ctx) error {
	item := dto.ItemRequest{}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(&item); err != nil {
		// Return status 400 and error message.
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errs.NewUnexpectedError(err.Error()))
	}
	item.ID = uuid.New()
	res, apperr := h.Service.AddNew(item)
	if apperr != nil {
		return ctx.Status(apperr.Code).JSON(apperr.AsMessage())
	}
	return ctx.JSON(res)
}
