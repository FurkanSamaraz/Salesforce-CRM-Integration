package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/LayoutTemplates"
	crm_structures "main/structure/LayoutTemplates"
)

type Service_LayoutTemplates interface {
	LayoutTemplatesCreate(model crm_structures.CreateLayoutTemplates) (*dto.Crm_DTO, error)
	LayoutTemplatesGet(model crm_structures.GetLayoutTemplates) (*dto.Crm_DTO, error)
	LayoutTemplatesUpdate(model crm_structures.UpdateLayoutTemplates) (*dto.Crm_DTO, error)
}
type LayoutTemplatesCRMService struct {
	Repo crm_repository.LayoutTemplatesRepositoryInterface
}

func (t LayoutTemplatesCRMService) LayoutTemplatesGet(model crm_structures.GetLayoutTemplates) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LayoutTemplates |")
	result, err := t.Repo.GetLayoutTemplates(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LayoutTemplatesCRMService) LayoutTemplatesCreate(model crm_structures.CreateLayoutTemplates) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LayoutTemplates |")
	result, err := t.Repo.CreateLayoutTemplates(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LayoutTemplatesCRMService) LayoutTemplatesUpdate(model crm_structures.UpdateLayoutTemplates) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LayoutTemplates |")
	result, err := t.Repo.UpdateLayoutTemplates(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
