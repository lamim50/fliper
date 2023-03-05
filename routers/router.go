package routers

import (
	"fliper/routers/types"
	"fliper/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Router struct {
	service *service.Service
}

func InitRouter(service *service.Service, router fiber.Router) *Router {
	appRouter := &Router{service: service}
	router.Get("/", appRouter.handleFind)
	router.Get("/:id", appRouter.handleFindeByID)
	return appRouter
}

func (s Router) handleFindeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	feat, err := s.service.FindByID(id)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	res := types.FromFeature(feat)
	return c.JSON(res)
}

func (s Router) handleFind(c *fiber.Ctx) error {
	feats, err := s.service.FindAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	res := types.FromFeatures(feats)
	return c.JSON(res)
}
