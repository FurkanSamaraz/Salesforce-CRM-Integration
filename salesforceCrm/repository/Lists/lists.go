package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Lists"

	"gorm.io/gorm"
)

type ListsDB struct {
	CrmCollection *gorm.DB
}
type ListsRepositoryInterface interface {
	GetLists(model crm_structures.GetList) (bool, error)
	CreateLists(model crm_structures.CreateList) (bool, error)
	UpdateLists(model crm_structures.UpdateList) (bool, error)
}

func (t ListsDB) GetLists(model crm_structures.GetList) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Lists |")
	return true, nil
}
func (t ListsDB) CreateLists(model crm_structures.CreateList) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Lists |")
	return true, nil
}
func (t ListsDB) UpdateLists(model crm_structures.UpdateList) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Lists |")
	return true, nil
}
