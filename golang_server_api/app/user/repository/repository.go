package repository

import (
	"api_and_grpc_golang/golang_server_api/app/user/domain"
	"api_and_grpc_golang/golang_server_api/utils"

	"gorm.io/gorm"
)

type UserRepo struct {
	db  *gorm.DB
	log utils.Zlog
}

func NewUserRepository(database *gorm.DB, zlog utils.Zlog) *UserRepo {
	return &UserRepo{
		db:  database,
		log: zlog,
	}
}

const (
	selectMenu = "SELECT * FROM menu_models"
	createMenu = "INSERT INTO `menu_models` (`name`) VALUES (?);"
)

func (r *UserRepo) RepoGetMenu() ([]domain.MenuModel, error) {
	result := new([]domain.MenuModel)
	r.db.Raw(selectMenu).Scan(&result)
	return *result, nil
}

func (r *UserRepo) RepoCreateMenu(req *domain.MenuRequest) error {
	r.db.Exec(createMenu, req.Name)
	return nil
}
