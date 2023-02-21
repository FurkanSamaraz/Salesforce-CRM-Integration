package crm_repository

import (
	"fmt"
	crm_structures "main/structure/CustomRedirects"

	"gorm.io/gorm"
)

type CustomRedirectsDB struct {
	CrmCollection *gorm.DB
}
type CustomRedirectsRepositoryInterface interface {
	GetCustomRedirects(model crm_structures.GetCustomRedirects) (bool, error)
	CreateCustomRedirects(model crm_structures.CreateCustomRedirects) (bool, error)
}

func (t CustomRedirectsDB) GetCustomRedirects(model crm_structures.GetCustomRedirects) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - CustomRedirects |")
	return true, nil
}
func (t CustomRedirectsDB) CreateCustomRedirects(model crm_structures.CreateCustomRedirects) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - CustomRedirects |")
	return true, nil
}
