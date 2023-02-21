package crm_handlers

import (
	crm_services "main/service/ListMemberships"

	"github.com/gofiber/fiber/v2"
)

type ListMembershipsHandler struct {
	Service crm_services.Service_ListMemberships
}

func (h ListMembershipsHandler) ListMemberships(c *fiber.Ctx) error {
	return nil
}

func (h ListMembershipsHandler) ListMembershipsCreate(c *fiber.Ctx) error {
	return nil
}
func (h ListMembershipsHandler) ListMembershipsById(c *fiber.Ctx) error {
	return nil
}
func (h ListMembershipsHandler) ListMembershipsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h ListMembershipsHandler) ListMembershipsDelete(c *fiber.Ctx) error {
	return nil
}
