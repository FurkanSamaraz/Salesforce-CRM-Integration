package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Lists"
	crm_structures "main/structure/Lists"
)

type Service_Lists interface {
	ListCreate(model crm_structures.CreateList) (*dto.Crm_DTO, error)
	ListGet(model crm_structures.GetList) (*dto.Crm_DTO, error)
	ListUpdate(model crm_structures.UpdateList) (*dto.Crm_DTO, error)
}
type ListsCRMService struct {
	Repo crm_repository.ListsRepositoryInterface
}

func (t ListsCRMService) ListGet(model crm_structures.GetList) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Lists |")
	result, err := t.Repo.GetLists(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ListsCRMService) ListCreate(model crm_structures.CreateList) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Lists |")
	result, err := t.Repo.CreateLists(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ListsCRMService) ListUpdate(model crm_structures.UpdateList) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Lists |")
	result, err := t.Repo.UpdateLists(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
