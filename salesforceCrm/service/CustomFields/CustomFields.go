package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/CustomFields"
	crm_structures "main/structure/CustomFields"
)

type Service_CustomFields interface {
	CustomFieldsCreate(model crm_structures.CreateCustomFields) (*dto.Crm_DTO, error)
	CustomFieldsGet(model crm_structures.GetCustomFields) (*dto.Crm_DTO, error)
	CustomFieldsUpdate(model crm_structures.UpdateCustomFields) (*dto.Crm_DTO, error)
}
type CustomFieldsCRMService struct {
	Repo crm_repository.CustomFieldsRepositoryInterface
}

func (t CustomFieldsCRMService) CustomFieldsGet(model crm_structures.GetCustomFields) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - CustomFields |")
	result, err := t.Repo.GetCustomFields(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CustomFieldsCRMService) CustomFieldsCreate(model crm_structures.CreateCustomFields) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - CustomFields |")
	result, err := t.Repo.CreateCustomFields(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CustomFieldsCRMService) CustomFieldsUpdate(model crm_structures.UpdateCustomFields) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - CustomFields |")
	result, err := t.Repo.UpdateCustomFields(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
