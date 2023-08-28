package user

import (
	"api_and_grpc_golang/golang_server_api/app/user/delivery/http"
	"api_and_grpc_golang/golang_server_api/app/user/repository"
	"api_and_grpc_golang/golang_server_api/app/user/service"
	"api_and_grpc_golang/golang_server_api/middleware"
	"api_and_grpc_golang/golang_server_api/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	App     fiber.Router
	Handler *http.UserHandler
	Config  utils.Config
}

func NewUserRouter(fiber *fiber.App, handler *http.UserHandler, config utils.Config) *Router {
	return &Router{
		App:     fiber,
		Handler: handler,
		Config:  config,
	}
}

var NewUserModule = fx.Options(
	fx.Provide(repository.NewUserRepository),
	fx.Provide(service.NewUserService),
	fx.Provide(http.NewUserHandler),
	fx.Provide(NewUserRouter),
)

func (r *Router) RegisterUserRouter() {
	r.App.Route("/app/menu", func(router fiber.Router) {
		router.Get("/", middleware.PermissionCheck("admin"), r.Handler.GetMenu)
		router.Post("/", middleware.PermissionCheck("admin"), r.Handler.CreateMenu)
	})
}
