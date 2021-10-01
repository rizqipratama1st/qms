package handler

import (
	"qms/models"
	"qms/service"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookService service.BookServiceInterface
}

func NewBookHandler(bookService service.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

type BookHandlerInterface interface {
	CreateBook(c *fiber.Ctx) error
	GetBank(c *fiber.Ctx) error
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := models.SlotBooking{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	response, err := h.bookService.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": response,
	})

}

func (h *BookHandler) GetBank(c *fiber.Ctx) error {
	banks, err := h.bookService.GetBank()

	if err != nil {
		c.Status(fiber.StatusForbidden).JSON(err)
	}

	return c.JSON(banks)
}
