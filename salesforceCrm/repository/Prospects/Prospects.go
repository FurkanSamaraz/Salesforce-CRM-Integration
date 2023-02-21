package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Prospects"

	"gorm.io/gorm"
)

type ProspectsDB struct {
	CrmCollection *gorm.DB
}
type ProspectsRepositoryInterface interface {
	GetProspects(model crm_structures.GetProspects) (bool, error)
	CreateProspects(model crm_structures.CreateProspects) (bool, error)
	UpdateProspects(model crm_structures.UpdateProspects) (bool, error)
}

func (t ProspectsDB) GetProspects(model crm_structures.GetProspects) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Prospects |")
	return true, nil
}
func (t ProspectsDB) CreateProspects(model crm_structures.CreateProspects) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Prospects |")
	return true, nil
}
func (t ProspectsDB) UpdateProspects(model crm_structures.UpdateProspects) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Prospects |")
	return true, nil
}
