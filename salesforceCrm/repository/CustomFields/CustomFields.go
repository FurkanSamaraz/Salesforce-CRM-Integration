package crm_repository

import (
	"fmt"
	crm_structures "main/structure/CustomFields"

	"gorm.io/gorm"
)

type CustomFieldsDB struct {
	CrmCollection *gorm.DB
}
type CustomFieldsRepositoryInterface interface {
	GetCustomFields(model crm_structures.GetCustomFields) (bool, error)
	CreateCustomFields(model crm_structures.CreateCustomFields) (bool, error)

	UpdateCustomFields(model crm_structures.UpdateCustomFields) (bool, error)
}

func (t CustomFieldsDB) GetCustomFields(model crm_structures.GetCustomFields) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - CustomFields |")
	return true, nil
}
func (t CustomFieldsDB) CreateCustomFields(model crm_structures.CreateCustomFields) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - CustomFields |")
	return true, nil
}
func (t CustomFieldsDB) UpdateCustomFields(model crm_structures.UpdateCustomFields) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - CustomFields |")
	return true, nil
}
