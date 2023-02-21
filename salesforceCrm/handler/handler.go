package crm_handlers

import (
	"fmt"
	configs "main/config"
	handler_Campaigns "main/handler/Campaigns"
	handler_CustomFields "main/handler/CustomFields"
	crm_integration "main/package"
	repository "main/repository"
	crm_repository_Campaigns "main/repository/Campaigns"
	crm_repository_CustomFields "main/repository/CustomFields"

	handler_CustomRedirects "main/handler/CustomRedirects"
	crm_repository_CustomRedirects "main/repository/CustomRedirects"

	handler_FormHandlers "main/handler/FormHandlers"
	crm_repository_FormHandlers "main/repository/FormHandlers"

	handler_Prospects "main/handler/Prospects"
	crm_repository_Prospects "main/repository/Prospects"

	handler_lists "main/handler/Lists"
	crm_repository_lists "main/repository/Lists"

	handler_LayoutTemplates "main/handler/LayoutTemplates"
	crm_repository_LayoutTemplates "main/repository/LayoutTemplates"

	handler_Users "main/handler/Users"
	crm_repository_Users "main/repository/Users"

	handler_ListMemberships "main/handler/ListMemberships"
	crm_repository_ListMemberships "main/repository/ListMemberships"

	service "main/service"
)

type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func App() {
	dbClient := configs.GetCollection()

	CampaignsRepositoryDb := repository.NewCampaignsRepositoryDB(dbClient)
	Campaigns := handler_Campaigns.CampaignsHandler{Service: service.NewCampaignsService(crm_repository_Campaigns.CampaignsDB{CrmCollection: CampaignsRepositoryDb.CrmCollection})}

	CustomFieldsRepositoryDb := repository.NewCustomFieldsRepositoryDB(dbClient)
	CustomFields := handler_CustomFields.CustomFieldsHandler{Service: service.NewCustomFieldsService(crm_repository_CustomFields.CustomFieldsDB{CrmCollection: CustomFieldsRepositoryDb.CrmCollection})}

	CustomRedirectsRepositoryDb := repository.NewCustomRedirectsRepositoryDB(dbClient)
	CustomRedirects := handler_CustomRedirects.CustomRedirectsHandler{Service: service.NewCustomRedirectsService(crm_repository_CustomRedirects.CustomRedirectsDB{CrmCollection: CustomRedirectsRepositoryDb.CrmCollection})}

	FormHandlersRepositoryDb := repository.NewFormHandlersRepositoryDB(dbClient)
	FormHandlers := handler_FormHandlers.FormHandlersHandler{Service: service.NewFormHandlersService(crm_repository_FormHandlers.FormHandlersDB{CrmCollection: FormHandlersRepositoryDb.CrmCollection})}

	ProspectsRepositoryDb := repository.NewProspectsRepositoryDB(dbClient)
	Prospects := handler_Prospects.ProspectsHandler{Service: service.NewProspectsService(crm_repository_Prospects.ProspectsDB{CrmCollection: ProspectsRepositoryDb.CrmCollection})}

	listsRepositorysDb := repository.NewListsRepositoryDB(dbClient)
	Lists := handler_lists.ListsHandler{Service: service.NewListsService(crm_repository_lists.ListsDB{CrmCollection: listsRepositorysDb.CrmCollection})}

	LayoutTemplatesRepositoryDb := repository.NewLayoutTemplatesRepositoryDB(dbClient)
	LayoutTemplates := handler_LayoutTemplates.LayoutTemplatesHandler{Service: service.NewLayoutTemplatesService(crm_repository_LayoutTemplates.LayoutTemplatesDB{CrmCollection: LayoutTemplatesRepositoryDb.CrmCollection})}

	UsersRepositoryDb := repository.NewUsersRepositoryDB(dbClient)
	Users := handler_Users.UsersHandler{Service: service.NewUsersService(crm_repository_Users.UsersDB{CrmCollection: UsersRepositoryDb.CrmCollection})}

	ListMembershipsRepositoryDb := repository.NewListMembershipsRepositoryDB(dbClient)
	ListMemberships := handler_ListMemberships.ListMembershipsHandler{Service: service.NewListMembershipsService(crm_repository_ListMemberships.ListMembershipsDB{CrmCollection: ListMembershipsRepositoryDb.CrmCollection})}

	fmt.Println(
		Campaigns,
		CustomFields,
		CustomRedirects,

		FormHandlers,
		Prospects,
		Lists,

		LayoutTemplates,
		Users,

		ListMemberships,
	)

}
