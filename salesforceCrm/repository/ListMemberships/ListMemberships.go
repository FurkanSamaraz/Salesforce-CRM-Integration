package crm_repository

import (
	"fmt"
	crm_structures "main/structure/ListMemberships"

	"gorm.io/gorm"
)

type ListMembershipsDB struct {
	CrmCollection *gorm.DB
}
type ListMembershipsRepositoryInterface interface {
	GetListMemberships(model crm_structures.GetListMemberships) (bool, error)
	CreateListMemberships(model crm_structures.CreateListMemberships) (bool, error)
	UpdateListMemberships(model crm_structures.UpdateListMemberships) (bool, error)
}

func (t ListMembershipsDB) GetListMemberships(model crm_structures.GetListMemberships) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - ListMemberships |")
	return true, nil
}
func (t ListMembershipsDB) CreateListMemberships(model crm_structures.CreateListMemberships) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - ListMemberships |")
	return true, nil
}
func (t ListMembershipsDB) UpdateListMemberships(model crm_structures.UpdateListMemberships) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - ListMemberships |")
	return true, nil
}
