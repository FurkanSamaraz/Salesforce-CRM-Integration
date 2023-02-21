package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Campaigns"
	crm_structures "main/structure/Campaigns"
)

type CampaignsCRMService struct {
	Repo crm_repository.CampaignsRepositoryInterface
}

type Service_Campaigns interface {
	CampaignsCreate(model crm_structures.CreateCampaigns) (*dto.Crm_DTO, error)
	CampaignsGet(model crm_structures.GetCampaigns) (*dto.Crm_DTO, error)
	CampaignsUpdate(model crm_structures.UpdateCampaigns) (*dto.Crm_DTO, error)
}

func (t CampaignsCRMService) CampaignsGet(model crm_structures.GetCampaigns) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Campaigns |")
	result, err := t.Repo.GetCampaigns(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CampaignsCRMService) CampaignsCreate(model crm_structures.CreateCampaigns) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Campaigns |")
	result, err := t.Repo.CreateCampaigns(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CampaignsCRMService) CampaignsUpdate(model crm_structures.UpdateCampaigns) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Campaigns |")
	result, err := t.Repo.UpdateCampaigns(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
