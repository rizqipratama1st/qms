package main

import (
	"os"
	"qms/database"
	"qms/handler"
	"qms/repository"
	"qms/routes"
	"qms/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func Book(c *fiber.Ctx) error {
// 	book := models.SlotBooking{
// 		ID:               1,
// 		TanggalPelayanan: time.Now(),
// 		JamPelayanan:     time.Now(),
// 		KeperluanLayanan: "transaksi",
// 		Status:           "done",
// 		BankID:           1,
// 		UserID:           1,
// 	}
// 	return c.JSON(book)
// }

// func PostBook(c *fiber.Ctx) error {
// 	book := &models.SlotBooking{}

// 	if err := c.BodyParser(book); err != nil {
// 		c.Status(403).JSON(err)
// 	}
// 	return c.JSON(book)
// }

func main() {

	//set up connection
	DB := database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Get("/", Book)
	// app.Post("/", PostBook)

	// book services
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// auth services
	var authHandler = &handler.AuthHandler{}

	routes := routes.NewRoutes(bookHandler, authHandler)
	routes.Setup(app)

	//untuk Local Development
	// app.Listen(":8080")

	// Untuk Production
	app.Listen(":" + os.Getenv("PORT"))
}