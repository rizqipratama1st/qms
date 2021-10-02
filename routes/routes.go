package routes

import (
	"qms/handler"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	bookHandler handler.BookHandlerInterface
	authHandler handler.AuthHandlerInterface
}

func NewRoutes(bookHandler handler.BookHandlerInterface, authHandler handler.AuthHandlerInterface) *Routes {
	return &Routes{
		bookHandler: bookHandler,
		authHandler: authHandler,
	}
}

func (r *Routes) Setup(app *fiber.App) {

	// auth
	app.Post("/api/register", r.authHandler.Register)
	app.Post("/api/login", r.authHandler.Login)
	app.Get("/api/user", r.authHandler.User)
	app.Post("/api/logout", r.authHandler.Logout)

	// booking
	app.Post("/book/create", r.bookHandler.CreateBook)
	app.Get("/bank", r.bookHandler.GetBank)

}
