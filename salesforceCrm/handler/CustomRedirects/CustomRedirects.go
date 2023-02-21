package crm_handlers

import (
	crm_services "main/service/CustomRedirects"

	"github.com/gofiber/fiber/v2"
)

type CustomRedirectsHandler struct {
	Service crm_services.Service_CustomRedirects
}

func (h CustomRedirectsHandler) CustomRedirects(c *fiber.Ctx) error {
	return nil
}

func (h CustomRedirectsHandler) CustomRedirectsCreate(c *fiber.Ctx) error {
	return nil
}
func (h CustomRedirectsHandler) CustomRedirectsById(c *fiber.Ctx) error {
	return nil
}
func (h CustomRedirectsHandler) CustomRedirectsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CustomRedirectsHandler) CustomRedirectsDelete(c *fiber.Ctx) error {
	return nil
}
