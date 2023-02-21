package crm_handlers

import (
	crm_services "main/service/Prospects"

	"github.com/gofiber/fiber/v2"
)

type ProspectsHandler struct {
	Service crm_services.Service_Prospects
}

func (h ProspectsHandler) Prospects(c *fiber.Ctx) error {
	return nil
}

func (h ProspectsHandler) ProspectsCreate(c *fiber.Ctx) error {
	return nil
}
func (h ProspectsHandler) ProspectsById(c *fiber.Ctx) error {
	return nil
}
func (h ProspectsHandler) ProspectsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h ProspectsHandler) ProspectsDelete(c *fiber.Ctx) error {
	return nil
}
