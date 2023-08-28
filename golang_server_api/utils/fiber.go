package utils

import (
	"api_and_grpc_golang/golang_server_api/utils/response"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func NewFiber(cfg Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.GetString(AppName),
		AppName:               cfg.GetString(AppName),
		Prefork:               cfg.GetBool(Prefork),
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           5 * time.Second,
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
		BodyLimit:             100 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./storage/public"),
	}))

	return app
}
