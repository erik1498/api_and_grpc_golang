package http

import (
	"api_and_grpc_golang/golang_server_api/app/user/domain"
	"api_and_grpc_golang/golang_server_api/app/user/service"
	"api_and_grpc_golang/golang_server_api/utils"
	"api_and_grpc_golang/golang_server_api/utils/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
	cfg         utils.Config
	log         utils.Zlog
}

func NewUserHandler(cfg utils.Config, userService *service.UserService, log utils.Zlog) *UserHandler {
	return &UserHandler{
		cfg:         cfg,
		userService: userService,
		log:         log,
	}
}

func (h *UserHandler) GetMenu(c *fiber.Ctx) error {
	h.log.LogInfo(c, utils.StartHandler+utils.GetFunctionName())
	result, err := h.userService.ServiceGetMenu(c)
	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(
			response.NewError(fiber.StatusInternalServerError, err.Error()),
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		response.Response{
			Code:     result.Code,
			Status:   result.Status,
			Messages: result.Messages,
			Data:     result.Data,
		},
	)
}

func (h *UserHandler) CreateMenu(c *fiber.Ctx) error {
	h.log.LogInfo(c, utils.StartHandler+utils.GetFunctionName())
	req := new(domain.MenuRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(
			response.NewError(fiber.StatusInternalServerError, err.Error()),
		)
	}
	result, err := h.userService.ServiceCreateMenu(c, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewError(fiber.StatusInternalServerError, err.Error()),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.Response{
			Code:     result.Code,
			Status:   result.Status,
			Messages: result.Messages,
			Data:     result.Data,
		},
	)
}
