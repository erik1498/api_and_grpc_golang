package service

import (
	"api_and_grpc_golang/golang_server_api/app/user/domain"
	"api_and_grpc_golang/golang_server_api/app/user/repository"
	"api_and_grpc_golang/golang_server_api/utils"
	"api_and_grpc_golang/golang_server_api/utils/response"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	UserRepo *repository.UserRepo
	log      utils.Zlog
}

func NewUserService(repo *repository.UserRepo, log utils.Zlog) *UserService {
	return &UserService{
		UserRepo: repo,
		log:      log,
	}
}

func (s *UserService) ServiceGetMenu(c *fiber.Ctx) (response.Response, error) {
	data, err := s.UserRepo.RepoGetMenu()
	if err != nil {
		return response.Response{
			Code:     fiber.ErrInternalServerError,
			Messages: err.Error(),
		}, err
	}
	return response.Response{
		Code:     fiber.StatusOK,
		Messages: response.ResponseSuccess,
		Data:     data,
	}, nil
}

func (s *UserService) ServiceCreateMenu(c *fiber.Ctx, req *domain.MenuRequest) (response.Response, error) {
	err := s.UserRepo.RepoCreateMenu(req)
	if err != nil {
		return response.Response{
			Code:     fiber.ErrInternalServerError,
			Messages: err.Error(),
		}, err
	}
	return response.Response{
		Code:     fiber.StatusCreated,
		Messages: response.ResponseSuccess,
	}, nil
}
