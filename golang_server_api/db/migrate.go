package db

import (
	user "api_and_grpc_golang/golang_server_api/app/user/domain"
	"api_and_grpc_golang/golang_server_api/utils"

	"gorm.io/gorm"
)

func AutoMigrateDatabaseDomain(DB *gorm.DB, log utils.Zlog) {
	err := DB.AutoMigrate(
		user.MenuModel{},
	)
	if err != nil {
		log.LogError(err, nil, "Error to migrate domain")
	} else {
		log.LogInfo(nil, "Migrate database success")
	}
}
