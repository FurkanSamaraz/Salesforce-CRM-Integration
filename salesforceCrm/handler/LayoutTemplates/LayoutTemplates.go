package crm_handlers

import (
	crm_services "main/service/LayoutTemplates"

	"github.com/gofiber/fiber/v2"
)

type LayoutTemplatesHandler struct {
	Service crm_services.Service_LayoutTemplates
}

func (h LayoutTemplatesHandler) LayoutTemplates(c *fiber.Ctx) error {
	return nil
}

func (h LayoutTemplatesHandler) LayoutTemplatesCreate(c *fiber.Ctx) error {
	return nil
}
func (h LayoutTemplatesHandler) LayoutTemplatesById(c *fiber.Ctx) error {
	return nil
}
func (h LayoutTemplatesHandler) LayoutTemplatesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h LayoutTemplatesHandler) LayoutTemplatesDelete(c *fiber.Ctx) error {
	return nil
}
