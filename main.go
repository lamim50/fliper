package main

import (
	"fliper/postgresql"
	"fliper/routers"
	"fliper/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&postgresql.FeatureModel{})
	repository := postgresql.NewRepositoryPostgreSQL(*db)
	service := service.NewService(repository)

	app := fiber.New()
	routers.InitRouter(service, app)

	app.Listen(":3000")
}
