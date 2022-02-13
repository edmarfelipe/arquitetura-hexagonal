package handler

import (
	"github.com/edmarfelipe/go-hexagonal/application"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrParseBodyError = "Could not parse body"
	ErrFailExecute    = "Could not complete the request"
	ErrMissingParms   = "The parameter ID is missing"
	ErrNotFound       = "Could not found the product"
)

type ProductHandler struct {
	Service application.ProductServiceInterface
}

func (handler *ProductHandler) MakeProductHandlers(app *fiber.App) {
	app.Patch("/products/:id/enable", handler.EnableHandler)
	app.Patch("/products/:id/disable", handler.DisableHandler)
	app.Get("/products/:id", handler.FindHandler)
	app.Post("/products", handler.CreateHandler)
}

func (handler *ProductHandler) EnableHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			ErrNotFound,
		)
	}

	product, err := handler.Service.Get(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusNotFound,
			err,
		)
	}

	result, err := handler.Service.Enable(product)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err,
		)
	}

	return c.Status(200).JSON(result)
}

func (handler *ProductHandler) DisableHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			ErrNotFound,
		)
	}

	product, err := handler.Service.Get(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusNotFound,
			err,
		)
	}

	result, err := handler.Service.Disable(product)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err,
		)
	}

	return c.Status(200).JSON(result)
}

func (handler *ProductHandler) CreateHandler(c *fiber.Ctx) error {
	body := new(ProductDTO)

	err := c.BodyParser(body)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			ErrParseBodyError,
		)
	}

	result, err := handler.Service.Create(body.Name, body.Price)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err,
		)
	}

	return c.Status(201).JSON(result)
}

func (handler *ProductHandler) FindHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return fiber.NewError(
			fiber.StatusNotFound,
			ErrNotFound,
		)
	}

	result, err := handler.Service.Get(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err,
		)
	}

	return c.Status(200).JSON(result)
}
