package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Prospects"
	crm_structures "main/structure/Prospects"
)

type Service_Prospects interface {
	ProspectsGet(model crm_structures.GetProspects) (*dto.Crm_DTO, error)
}
type ProspectsCRMService struct {
	Repo crm_repository.ProspectsRepositoryInterface
}

func (t ProspectsCRMService) ProspectsGet(model crm_structures.GetProspects) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Prospects |")
	result, err := t.Repo.GetProspects(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
