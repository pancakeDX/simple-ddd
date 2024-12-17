package main

import (
	"os"
	"os/signal"
	domainUser "simple-ddd/domain/user"
	infraUser "simple-ddd/infra/user"
	interfaceUser "simple-ddd/interface/user"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			Simple DDD
//	@version		0.0.1
//	@description	Simple DDD
//	@BasePath		/
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// v1
	v1 := app.Group("/v1")

	// user
	repo := infraUser.NewMemoryRepository()
	service := domainUser.NewService(repo)
	handler := interfaceUser.NewHandler(service)
	handler.RegisterRoutes(v1)

	// swagger
	app.Get("/public/apidoc/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
