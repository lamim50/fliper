package routers

import (
	"fliper/routers/types"
	"fliper/service"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	service *service.Service
}

func InitRouter(service *service.Service, router fiber.Router) *Router {
	appRouter := &Router{service: service}
	router.Get("/", appRouter.handleFind)
	router.Get("/:id", appRouter.handleFindeByID)
	router.Post("/", appRouter.handleCreateFeature)
	router.Put("/:id", appRouter.handleUpdateFeature)
	router.Delete("/:id", appRouter.handleDeleteFeature)
	return appRouter
}

func (s Router) handleFindeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	feat, err := s.service.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	res := types.FromFeature(feat)
	return c.JSON(res)
}

func (s Router) handleFind(c *fiber.Ctx) error {
	feats, err := s.service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	res := types.FromFeatures(feats)
	return c.JSON(res)
}

func (s Router) handleCreateFeature(c *fiber.Ctx) error {
	cmd := new(types.CreateFeatureCMD)
	if err := c.BodyParser(cmd); err != nil {
		return err
	}

	feats, err := s.service.Create(cmd.Title, cmd.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	res := types.FromFeature(feats)
	return c.JSON(res)
}

func (s Router) handleUpdateFeature(c *fiber.Ctx) error {
	cmd := new(types.UpdateFeatureCMD)
	id := c.Params("id")
	if err := c.BodyParser(cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	feats, err := s.service.Update(id, cmd.Title, cmd.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	res := types.FromFeature(feats)
	return c.JSON(res)
}

func (s Router) handleDeleteFeature(c *fiber.Ctx) error {
	id := c.Params("id")

	err := s.service.DeleteByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
