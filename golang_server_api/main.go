package main

import (
	"api_and_grpc_golang/golang_server_api/app/user"
	"api_and_grpc_golang/golang_server_api/db"
	"api_and_grpc_golang/golang_server_api/router"
	"api_and_grpc_golang/golang_server_api/utils"
	"context"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		fx.Provide(utils.NewLogger),
		fx.Provide(utils.NewViper),
		fx.Provide(utils.NewMysqlDB),
		fx.Provide(utils.NewFiber),
		fx.Provide(router.NewRouter),
		user.NewUserModule,
		fx.Invoke(Start),
	).Run()
}

func Start(
	lc fx.Lifecycle,
	logger utils.Zlog,
	cfg utils.Config,
	fiber *fiber.App,
	router *router.Router,
	DB *gorm.DB,
) {
	AppsName := cfg.GetString(utils.AppName)
	AppsPort := cfg.GetString(utils.AppPort)
	AppsProd := cfg.GetBool(utils.AppProduction)
	AppPrefork := cfg.GetBool(utils.Prefork)
	AppHost := cfg.GetString(utils.AppHost)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				router.RegisterS()
				db.AutoMigrateDatabaseDomain(DB, logger)
				if AppHost == "" {
					AppHost = "0.0.0.0"
				}

				logger.LogInfo(nil, "%s %s", fiber.Config().AppName, " is running at the moment!")
				if !AppsProd {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					logger.LogInfo(nil, "Prefork %t", AppPrefork)
					if !AppPrefork {
						procs = 1
						prefork = "Disabled"
					}

					logger.LogInfo(nil, "Version: %s", "-")
					logger.LogInfo(nil, "Host: %s", AppsName)
					logger.LogInfo(nil, "Port: %s", AppsPort)
					logger.LogInfo(nil, "Prefork: %s", prefork)
					logger.LogInfo(nil, "Handlers: %d", fiber.HandlersCount())
					logger.LogInfo(nil, "Processes: %d", procs)
					logger.LogInfo(nil, "PID: %d", os.Getpid())
				}

				go func() {
					if err := fiber.Listen(AppsPort); err != nil {
						logger.LogError(err, nil, "%s", "An unknown error occurred when to run server !")
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.LogInfo(nil, "%s", "Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					logger.LogPanic(err, nil, "Err Panic")
				}

				logger.LogInfo(nil, "%s", "Running cleanup tasks...")
				logger.LogInfo(nil, "%s", "1- Shutdown the db")
				logger.LogInfo(nil, "%s was successfull shutdown", AppsName)

				return nil
			},
		},
	)
}
