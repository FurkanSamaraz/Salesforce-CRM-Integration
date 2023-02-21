package crm_services

import (
	crm_repository_Campaigns "main/repository/Campaigns"
	crm_service_Campaigns "main/service/Campaigns"

	crm_repository_CustomFields "main/repository/CustomFields"
	crm_service_CustomFields "main/service/CustomFields"

	crm_repository_Prospects "main/repository/Prospects"
	crm_service_Prospects "main/service/Prospects"

	crm_repository_CustomRedirects "main/repository/CustomRedirects"
	crm_service_CustomRedirects "main/service/CustomRedirects"

	crm_repository_FormHandlers "main/repository/FormHandlers"
	crm_service_FormHandlers "main/service/FormHandlers"

	crm_repository_Lists "main/repository/Lists"
	crm_service_Lists "main/service/Lists"

	crm_repository_LayoutTemplates "main/repository/LayoutTemplates"
	crm_service_LayoutTemplates "main/service/LayoutTemplates"

	crm_repository_ListMemberships "main/repository/ListMemberships"
	crm_service_ListMemberships "main/service/ListMemberships"

	crm_repository_Users "main/repository/Users"
	crm_service_Users "main/service/Users"
)

func NewCampaignsService(Repo crm_repository_Campaigns.CampaignsRepositoryInterface) crm_service_Campaigns.CampaignsCRMService {
	return crm_service_Campaigns.CampaignsCRMService{Repo: Repo}
}

func NewProspectsService(Repo crm_repository_Prospects.ProspectsRepositoryInterface) crm_service_Prospects.ProspectsCRMService {
	return crm_service_Prospects.ProspectsCRMService{Repo: Repo}
}

func NewCustomFieldsService(Repo crm_repository_CustomFields.CustomFieldsRepositoryInterface) crm_service_CustomFields.CustomFieldsCRMService {
	return crm_service_CustomFields.CustomFieldsCRMService{Repo: Repo}
}

func NewCustomRedirectsService(Repo crm_repository_CustomRedirects.CustomRedirectsRepositoryInterface) crm_service_CustomRedirects.CustomRedirectsCRMService {
	return crm_service_CustomRedirects.CustomRedirectsCRMService{Repo: Repo}
}
func NewFormHandlersService(Repo crm_repository_FormHandlers.FormHandlersRepositoryInterface) crm_service_FormHandlers.FormHandlersCRMService {
	return crm_service_FormHandlers.FormHandlersCRMService{Repo: Repo}
}
func NewListsService(Repo crm_repository_Lists.ListsRepositoryInterface) crm_service_Lists.ListsCRMService {
	return crm_service_Lists.ListsCRMService{Repo: Repo}
}
func NewLayoutTemplatesService(Repo crm_repository_LayoutTemplates.LayoutTemplatesRepositoryInterface) crm_service_LayoutTemplates.LayoutTemplatesCRMService {
	return crm_service_LayoutTemplates.LayoutTemplatesCRMService{Repo: Repo}
}
func NewListMembershipsService(Repo crm_repository_ListMemberships.ListMembershipsRepositoryInterface) crm_service_ListMemberships.ListMembershipsCRMService {
	return crm_service_ListMemberships.ListMembershipsCRMService{Repo: Repo}
}

func NewUsersService(Repo crm_repository_Users.UsersRepositoryInterface) crm_service_Users.UsersCRMService {
	return crm_service_Users.UsersCRMService{Repo: Repo}
}
