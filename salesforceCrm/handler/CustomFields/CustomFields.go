package crm_handlers

import (
	crm_integration "main/package"
	crm_services "main/service/CustomFields"

	"github.com/gofiber/fiber/v2"
)

type CustomFieldsHandler struct {
	Service crm_services.Service_CustomFields
}
type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func (h CustomFieldsHandler) CustomFields(c *fiber.Ctx) error {
	return nil
}

func (h CustomFieldsHandler) CustomFieldsCreate(c *fiber.Ctx) error {
	return nil
}
func (h CustomFieldsHandler) CustomFieldsById(c *fiber.Ctx) error {
	return nil
}
func (h CustomFieldsHandler) CustomFieldsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CustomFieldsHandler) CustomFieldsDelete(c *fiber.Ctx) error {
	return nil
}
