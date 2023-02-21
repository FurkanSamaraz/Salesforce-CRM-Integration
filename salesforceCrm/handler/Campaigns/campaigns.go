package crm_handlers

import (
	crm_services "main/service/Campaigns"

	"github.com/gofiber/fiber/v2"
)

type CampaignsHandler struct {
	Service crm_services.Service_Campaigns
}

func (h CampaignsHandler) Campaigns(c *fiber.Ctx) error {
	return nil
}

func (h CampaignsHandler) CampaignsCreate(c *fiber.Ctx) error {
	return nil
}
func (h CampaignsHandler) CampaignsById(c *fiber.Ctx) error {
	return nil
}
func (h CampaignsHandler) CampaignsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CampaignsHandler) CampaignsDelete(c *fiber.Ctx) error {
	return nil
}
