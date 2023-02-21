package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/FormHandlers"
	crm_structures "main/structure/FormHandlers"
)

type Service_FormHandlers interface {
	FormHandlersCreate(model crm_structures.Create_FormHandlers) (*dto.Crm_DTO, error)
	FormHandlersGet(model crm_structures.Get_FormHandlers) (*dto.Crm_DTO, error)
	FormHandlersUpdate(model crm_structures.Update_FormHandlers) (*dto.Crm_DTO, error)
}
type FormHandlersCRMService struct {
	Repo crm_repository.FormHandlersRepositoryInterface
}

func (t FormHandlersCRMService) FormHandlersGet(model crm_structures.Get_FormHandlers) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - FormHandlers |")
	result, err := t.Repo.GetFormHandlers(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t FormHandlersCRMService) FormHandlersCreate(model crm_structures.Create_FormHandlers) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - FormHandlers |")
	result, err := t.Repo.CreateFormHandlers(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t FormHandlersCRMService) FormHandlersUpdate(model crm_structures.Update_FormHandlers) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - FormHandlers |")
	result, err := t.Repo.UpdateFormHandlers(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
