package main

import (
	"fliper/postgresql"
	"fliper/routers"
	"fliper/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	repository := postgresql.RepositoryPostgreSQL{}
	service := service.NewService(repository)

	app := fiber.New()
	routers.InitRouter(service, app)

	app.Listen(":3000")
}
