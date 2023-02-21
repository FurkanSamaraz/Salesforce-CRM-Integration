package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/CustomRedirects"
	crm_structures "main/structure/CustomRedirects"
)

type Service_CustomRedirects interface {
	CustomRedirectsCreate(model crm_structures.CreateCustomRedirects) (*dto.Crm_DTO, error)
	CustomRedirectsGet(model crm_structures.GetCustomRedirects) (*dto.Crm_DTO, error)
}
type CustomRedirectsCRMService struct {
	Repo crm_repository.CustomRedirectsRepositoryInterface
}

func (t CustomRedirectsCRMService) CustomRedirectsGet(model crm_structures.GetCustomRedirects) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - CustomRedirects |")
	result, err := t.Repo.GetCustomRedirects(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CustomRedirectsCRMService) CustomRedirectsCreate(model crm_structures.CreateCustomRedirects) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - CustomRedirects |")
	result, err := t.Repo.CreateCustomRedirects(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
