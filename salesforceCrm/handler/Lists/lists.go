package crm_handlers

import (
	crm_services "main/service/Lists"

	"github.com/gofiber/fiber/v2"
)

type ListsHandler struct {
	Service crm_services.Service_Lists
}

func (h ListsHandler) Lists(c *fiber.Ctx) error {
	return nil
}

func (h ListsHandler) ListsCreate(c *fiber.Ctx) error {
	return nil
}
func (h ListsHandler) ListsById(c *fiber.Ctx) error {
	return nil
}
func (h ListsHandler) ListsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h ListsHandler) ListsDelete(c *fiber.Ctx) error {
	return nil
}
