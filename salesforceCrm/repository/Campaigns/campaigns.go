package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Campaigns"

	"gorm.io/gorm"
)

type CampaignsDB struct {
	CrmCollection *gorm.DB
}
type CampaignsRepositoryInterface interface {
	GetCampaigns(model crm_structures.GetCampaigns) (bool, error)
	CreateCampaigns(model crm_structures.CreateCampaigns) (bool, error)
	UpdateCampaigns(model crm_structures.UpdateCampaigns) (bool, error)
}

func (t CampaignsDB) GetCampaigns(model crm_structures.GetCampaigns) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Campaigns |")
	return true, nil
}
func (t CampaignsDB) CreateCampaigns(model crm_structures.CreateCampaigns) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Campaigns |")
	return true, nil
}
func (t CampaignsDB) UpdateCampaigns(model crm_structures.UpdateCampaigns) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Campaigns |")
	return true, nil
}
