package crm_handlers

import (
	crm_services "main/service/FormHandlers"

	"github.com/gofiber/fiber/v2"
)

type FormHandlersHandler struct {
	Service crm_services.Service_FormHandlers
}

func (h FormHandlersHandler) FormHandlers(c *fiber.Ctx) error {
	return nil
}

func (h FormHandlersHandler) FormHandlersCreate(c *fiber.Ctx) error {
	return nil
}
func (h FormHandlersHandler) FormHandlersById(c *fiber.Ctx) error {
	return nil
}
func (h FormHandlersHandler) FormHandlersUpdate(c *fiber.Ctx) error {
	return nil
}
func (h FormHandlersHandler) FormHandlersDelete(c *fiber.Ctx) error {
	return nil
}
