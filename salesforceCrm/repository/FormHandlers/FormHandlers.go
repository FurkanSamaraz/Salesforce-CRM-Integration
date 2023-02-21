package crm_repository

import (
	"fmt"
	crm_structures "main/structure/FormHandlers"

	"gorm.io/gorm"
)

type FormHandlersDB struct {
	CrmCollection *gorm.DB
}
type FormHandlersRepositoryInterface interface {
	GetFormHandlers(model crm_structures.Get_FormHandlers) (bool, error)
	CreateFormHandlers(model crm_structures.Create_FormHandlers) (bool, error)
	UpdateFormHandlers(model crm_structures.Update_FormHandlers) (bool, error)
}

func (t FormHandlersDB) GetFormHandlers(model crm_structures.Get_FormHandlers) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - FormHandlers |")
	return true, nil
}
func (t FormHandlersDB) CreateFormHandlers(model crm_structures.Create_FormHandlers) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - FormHandlers |")
	return true, nil
}
func (t FormHandlersDB) UpdateFormHandlers(model crm_structures.Update_FormHandlers) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - FormHandlers |")
	return true, nil
}
