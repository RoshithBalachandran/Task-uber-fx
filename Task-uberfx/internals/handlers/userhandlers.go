package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roshith/uber-fx/internals/service"
)

type UserHandlers struct {
	srvic *service.UserService
}

func NewUserHandlers(srv *service.UserService) *UserHandlers {
	return &UserHandlers{srvic: srv}
}

func (h *UserHandlers) HandlerRoutes(app *fiber.App) {
	app.Post("/users", h.Create)
	app.Get("/all", h.FindAll)
}
func (h *UserHandlers) Create(c *fiber.Ctx) error {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	// pass pointer
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data format",
		})
	}

	user, err := h.srvic.Create(body.Name, body.Email) // <-- use :=
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": user,
	})
}

func (h *UserHandlers) FindAll(c *fiber.Ctx) error {
	user, err := h.srvic.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}
