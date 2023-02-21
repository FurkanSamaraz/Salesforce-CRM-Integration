package crm_repository

import (
	"fmt"
	crm_structures "main/structure/LayoutTemplates"

	"gorm.io/gorm"
)

type LayoutTemplatesDB struct {
	CrmCollection *gorm.DB
}
type LayoutTemplatesRepositoryInterface interface {
	GetLayoutTemplates(model crm_structures.GetLayoutTemplates) (bool, error)
	CreateLayoutTemplates(model crm_structures.CreateLayoutTemplates) (bool, error)
	UpdateLayoutTemplates(model crm_structures.UpdateLayoutTemplates) (bool, error)
}

func (t LayoutTemplatesDB) GetLayoutTemplates(model crm_structures.GetLayoutTemplates) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LayoutTemplates |")
	return true, nil
}
func (t LayoutTemplatesDB) CreateLayoutTemplates(model crm_structures.CreateLayoutTemplates) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LayoutTemplates |")
	return true, nil
}
func (t LayoutTemplatesDB) UpdateLayoutTemplates(model crm_structures.UpdateLayoutTemplates) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LayoutTemplates |")
	return true, nil
}
