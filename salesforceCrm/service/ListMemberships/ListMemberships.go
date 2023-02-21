package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/ListMemberships"
	crm_structures "main/structure/ListMemberships"
)

type Service_ListMemberships interface {
	ListMembershipsCreate(model crm_structures.CreateListMemberships) (*dto.Crm_DTO, error)
	ListMembershipsGet(model crm_structures.GetListMemberships) (*dto.Crm_DTO, error)
	ListMembershipsUpdate(model crm_structures.UpdateListMemberships) (*dto.Crm_DTO, error)
}
type ListMembershipsCRMService struct {
	Repo crm_repository.ListMembershipsRepositoryInterface
}

func (t ListMembershipsCRMService) ListMembershipsGet(model crm_structures.GetListMemberships) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - ListMemberships |")
	result, err := t.Repo.GetListMemberships(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ListMembershipsCRMService) ListMembershipsCreate(model crm_structures.CreateListMemberships) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - ListMemberships |")
	result, err := t.Repo.CreateListMemberships(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ListMembershipsCRMService) ListMembershipsUpdate(model crm_structures.UpdateListMemberships) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - ListMemberships |")
	result, err := t.Repo.UpdateListMemberships(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
