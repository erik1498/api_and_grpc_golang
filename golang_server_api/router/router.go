package router

import (
	"api_and_grpc_golang/golang_server_api/app/user"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App        fiber.Router
	UserRouter *user.Router
}

func NewRouter(
	fiber *fiber.App,
	UserRouter *user.Router,
) *Router {
	return &Router{
		App:        fiber,
		UserRouter: UserRouter,
	}
}

func (r *Router) RegisterS() {
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello Calon Orang Sukses")
	})
	r.UserRouter.RegisterUserRouter()
}
