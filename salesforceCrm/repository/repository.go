package repository

import "gorm.io/gorm"

type CrmRepositoryDB struct {
	CrmCollection *gorm.DB
}

func NewCampaignsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}

func NewCustomFieldsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewCustomRedirectsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewContactsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewFormHandlersRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewProspectsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewListsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewListMembershipsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewPingRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewLayoutTemplatesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewTasksRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewUsersRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
