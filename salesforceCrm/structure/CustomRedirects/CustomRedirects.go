package crm_structures

type CreateCustomRedirects struct {
	Attributes struct {
		Contactscomplete           bool   `json:"Contactscomplete"`
		Actualserviceunits         int    `json:"actualserviceunits"`
		Billedserviceunits         int    `json:"billedserviceunits"`
		Blockedprofile             bool   `json:"blockedprofile"`
		Caseorigincode             int    `json:"caseorigincode"`
		Casetypecode               int    `json:"casetypecode"`
		Checkemail                 bool   `json:"checkemail"`
		Contractdetailid           string `json:"contractdetailid"`
		Contractid                 string `json:"contractid"`
		Contractservicelevelcode   int    `json:"contractservicelevelcode"`
		Customercontacted          bool   `json:"customercontacted"`
		Customerid                 string `json:"customerid"`
		Customeridtype             string `json:"customeridtype"`
		Customersatisfactioncode   int    `json:"customersatisfactioncode"`
		Decremententitlementterm   bool   `json:"decremententitlementterm"`
		Description                string `json:"description"`
		Entitlementid              string `json:"entitlementid"`
		Existingcase               string `json:"existingcase"`
		Firstresponsebykpiid       string `json:"firstresponsebykpiid"`
		Firstresponsesent          bool   `json:"firstresponsesent"`
		Firstresponseslastatus     int    `json:"firstresponseslastatus"`
		Followupby                 int    `json:"followupby"`
		Followuptaskcreated        bool   `json:"followuptaskcreated"`
		Importsequencenumber       int    `json:"importsequencenumber"`
		Incidentid                 string `json:"incidentid"`
		Incidentstagecode          int    `json:"incidentstagecode"`
		Influencescore             int    `json:"influencescore"`
		IntActualsatisfaction      string `json:"int_actualsatisfaction"`
		IntAssociatedproduct       string `json:"int_associatedproduct"`
		IntCasecategory            string `json:"int_casecategory"`
		IntCustomereffort          int    `json:"int_customereffort"`
		IntEffortnum               string `json:"int_effortnum"`
		IntParaturecasenumber      string `json:"int_paraturecasenumber"`
		IntSurveryparticipation    bool   `json:"int_surveryparticipation"`
		IntUpsellreferral          bool   `json:"int_upsellreferral"`
		Isdecrementing             bool   `json:"isdecrementing"`
		Isescalated                bool   `json:"isescalated"`
		Kbarticleid                string `json:"kbarticleid"`
		Lastonholdtime             int    `json:"lastonholdtime"`
		Masterid                   string `json:"masterid"`
		Messagetypecode            int    `json:"messagetypecode"`
		MsdynIncidenttype          string `json:"msdyn_incidenttype"`
		NewActualsatisfaction      string `json:"new_actualsatisfaction"`
		NewCasecategory            string `json:"new_casecategory"`
		NewConversationid          string `json:"new_conversationid"`
		NewUpsellreferral          bool   `json:"new_upsellreferral"`
		Overriddencreatedon        int    `json:"overriddencreatedon"`
		Ownerid                    string `json:"ownerid"`
		Owneridtype                string `json:"owneridtype"`
		Parentcaseid               string `json:"parentcaseid"`
		Primarycontactid           string `json:"primarycontactid"`
		Prioritycode               int    `json:"prioritycode"`
		Processid                  string `json:"processid"`
		Productid                  string `json:"productid"`
		LayoutTemplateserialnumber string `json:"LayoutTemplateserialnumber"`
		Resolveby                  int    `json:"resolveby"`
		Resolvebykpiid             string `json:"resolvebykpiid"`
		Resolvebyslastatus         int    `json:"resolvebyslastatus"`
		Responseby                 int    `json:"responseby"`
		Responsiblecontactid       string `json:"responsiblecontactid"`
		Routecase                  bool   `json:"routecase"`
		Sentimentvalue             int    `json:"sentimentvalue"`
		Servicestage               int    `json:"servicestage"`
		Severitycode               int    `json:"severitycode"`
		Slaid                      string `json:"slaid"`
		Socialprofileid            string `json:"socialprofileid"`
		Stageid                    string `json:"stageid"`
		Statuscode                 int    `json:"statuscode"`
		Subjectid                  string `json:"subjectid"`
		Ticketnumber               string `json:"ticketnumber"`
		Timezoneruleversionnumber  int    `json:"timezoneruleversionnumber"`
		Title                      string `json:"title"`
		Transactioncurrencyid      string `json:"transactioncurrencyid"`
		Traversedpath              string `json:"traversedpath"`
		Utcconversiontimezonecode  int    `json:"utcconversiontimezonecode"`
	} `json:"attributes"`
}
type GetCustomRedirects struct {
	Attributes struct {
		Accountid                       string `json:"accountid"`
		Accountidname                   string `json:"accountidname"`
		Accountidyominame               string `json:"accountidyominame"`
		Contactscomplete                bool   `json:"Contactscomplete"`
		Actualserviceunits              int    `json:"actualserviceunits"`
		Billedserviceunits              int    `json:"billedserviceunits"`
		Blockedprofile                  bool   `json:"blockedprofile"`
		Caseorigincode                  int    `json:"caseorigincode"`
		Casetypecode                    int    `json:"casetypecode"`
		Checkemail                      bool   `json:"checkemail"`
		Contactid                       string `json:"contactid"`
		Contactidname                   string `json:"contactidname"`
		Contactidyominame               string `json:"contactidyominame"`
		Contractdetailid                string `json:"contractdetailid"`
		Contractdetailidname            string `json:"contractdetailidname"`
		Contractid                      string `json:"contractid"`
		Contractidname                  string `json:"contractidname"`
		Contractservicelevelcode        int    `json:"contractservicelevelcode"`
		Createdby                       string `json:"createdby"`
		Createdbyexternalparty          string `json:"createdbyexternalparty"`
		Createdbyexternalpartyname      string `json:"createdbyexternalpartyname"`
		Createdbyexternalpartyyominame  string `json:"createdbyexternalpartyyominame"`
		Createdbyname                   string `json:"createdbyname"`
		Createdbyyominame               string `json:"createdbyyominame"`
		Createdon                       int    `json:"createdon"`
		Createdonbehalfby               string `json:"createdonbehalfby"`
		Createdonbehalfbyname           string `json:"createdonbehalfbyname"`
		Createdonbehalfbyyominame       string `json:"createdonbehalfbyyominame"`
		Customercontacted               bool   `json:"customercontacted"`
		Customerid                      string `json:"customerid"`
		Customeridname                  string `json:"customeridname"`
		Customeridtype                  string `json:"customeridtype"`
		Customeridyominame              string `json:"customeridyominame"`
		Customersatisfactioncode        int    `json:"customersatisfactioncode"`
		Decremententitlementterm        bool   `json:"decremententitlementterm"`
		Description                     string `json:"description"`
		Entitlementid                   string `json:"entitlementid"`
		Entitlementidname               string `json:"entitlementidname"`
		EntityimageTimestamp            int    `json:"entityimage_timestamp"`
		EntityimageURL                  string `json:"entityimage_url"`
		Entityimageid                   string `json:"entityimageid"`
		Escalatedon                     int    `json:"escalatedon"`
		Exchangerate                    int    `json:"exchangerate"`
		Existingcase                    string `json:"existingcase"`
		Firstresponsebykpiid            string `json:"firstresponsebykpiid"`
		Firstresponsebykpiidname        string `json:"firstresponsebykpiidname"`
		Firstresponsesent               bool   `json:"firstresponsesent"`
		Firstresponseslastatus          int    `json:"firstresponseslastatus"`
		Followupby                      int    `json:"followupby"`
		Followuptaskcreated             bool   `json:"followuptaskcreated"`
		Importsequencenumber            int    `json:"importsequencenumber"`
		Incidentid                      string `json:"incidentid"`
		Incidentstagecode               int    `json:"incidentstagecode"`
		Influencescore                  int    `json:"influencescore"`
		IntActualsatisfaction           string `json:"int_actualsatisfaction"`
		IntAssociatedproduct            string `json:"int_associatedproduct"`
		IntAssociatedproductname        string `json:"int_associatedproductname"`
		IntCasecategory                 string `json:"int_casecategory"`
		IntCustomereffort               int    `json:"int_customereffort"`
		IntEffortnum                    string `json:"int_effortnum"`
		IntParaturecasenumber           string `json:"int_paraturecasenumber"`
		IntSurveryparticipation         bool   `json:"int_surveryparticipation"`
		IntUpsellreferral               bool   `json:"int_upsellreferral"`
		Isdecrementing                  bool   `json:"isdecrementing"`
		Isescalated                     bool   `json:"isescalated"`
		Kbarticleid                     string `json:"kbarticleid"`
		Kbarticleidname                 string `json:"kbarticleidname"`
		Lastonholdtime                  int    `json:"lastonholdtime"`
		Masterid                        string `json:"masterid"`
		Masteridname                    string `json:"masteridname"`
		Merged                          bool   `json:"merged"`
		Messagetypecode                 int    `json:"messagetypecode"`
		Modifiedby                      string `json:"modifiedby"`
		Modifiedbyexternalparty         string `json:"modifiedbyexternalparty"`
		Modifiedbyexternalpartyname     string `json:"modifiedbyexternalpartyname"`
		Modifiedbyexternalpartyyominame string `json:"modifiedbyexternalpartyyominame"`
		Modifiedbyname                  string `json:"modifiedbyname"`
		Modifiedbyyominame              string `json:"modifiedbyyominame"`
		Modifiedon                      int    `json:"modifiedon"`
		Modifiedonbehalfby              string `json:"modifiedonbehalfby"`
		Modifiedonbehalfbyname          string `json:"modifiedonbehalfbyname"`
		Modifiedonbehalfbyyominame      string `json:"modifiedonbehalfbyyominame"`
		MsdynIncidenttype               string `json:"msdyn_incidenttype"`
		MsdynIncidenttypename           string `json:"msdyn_incidenttypename"`
		NewActualsatisfaction           string `json:"new_actualsatisfaction"`
		NewCasecategory                 string `json:"new_casecategory"`
		NewConversationid               string `json:"new_conversationid"`
		NewUpsellreferral               bool   `json:"new_upsellreferral"`
		Numberofchildincidents          int    `json:"numberofchildincidents"`
		Onholdtime                      int    `json:"onholdtime"`
		Overriddencreatedon             int    `json:"overriddencreatedon"`
		Ownerid                         string `json:"ownerid"`
		Owneridname                     string `json:"owneridname"`
		Owneridtype                     string `json:"owneridtype"`
		Owneridyominame                 string `json:"owneridyominame"`
		Owningbusinessunit              string `json:"owningbusinessunit"`
		Owningteam                      string `json:"owningteam"`
		Owninguser                      string `json:"owninguser"`
		Parentcaseid                    string `json:"parentcaseid"`
		Parentcaseidname                string `json:"parentcaseidname"`
		Primarycontactid                string `json:"primarycontactid"`
		Primarycontactidname            string `json:"primarycontactidname"`
		Primarycontactidyominame        string `json:"primarycontactidyominame"`
		Prioritycode                    int    `json:"prioritycode"`
		Processid                       string `json:"processid"`
		Productid                       string `json:"productid"`
		Productidname                   string `json:"productidname"`
		LayoutTemplateserialnumber      string `json:"LayoutTemplateserialnumber"`
		Resolveby                       int    `json:"resolveby"`
		Resolvebykpiid                  string `json:"resolvebykpiid"`
		Resolvebykpiidname              string `json:"resolvebykpiidname"`
		Resolvebyslastatus              int    `json:"resolvebyslastatus"`
		Responseby                      int    `json:"responseby"`
		Responsiblecontactid            string `json:"responsiblecontactid"`
		Responsiblecontactidname        string `json:"responsiblecontactidname"`
		Responsiblecontactidyominame    string `json:"responsiblecontactidyominame"`
		Routecase                       bool   `json:"routecase"`
		Sentimentvalue                  int    `json:"sentimentvalue"`
		Servicestage                    int    `json:"servicestage"`
		Severitycode                    int    `json:"severitycode"`
		Slaid                           string `json:"slaid"`
		Slainvokedid                    string `json:"slainvokedid"`
		Slainvokedidname                string `json:"slainvokedidname"`
		Slaname                         string `json:"slaname"`
		Socialprofileid                 string `json:"socialprofileid"`
		Socialprofileidname             string `json:"socialprofileidname"`
		Stageid                         string `json:"stageid"`
		Statecode                       int    `json:"statecode"`
		Statuscode                      int    `json:"statuscode"`
		Subjectid                       string `json:"subjectid"`
		Subjectidname                   string `json:"subjectidname"`
		Ticketnumber                    string `json:"ticketnumber"`
		Timezoneruleversionnumber       int    `json:"timezoneruleversionnumber"`
		Title                           string `json:"title"`
		Transactioncurrencyid           string `json:"transactioncurrencyid"`
		Transactioncurrencyidname       string `json:"transactioncurrencyidname"`
		Traversedpath                   string `json:"traversedpath"`
		Utcconversiontimezonecode       int    `json:"utcconversiontimezonecode"`
		Versionnumber                   int    `json:"versionnumber"`
	} `json:"attributes"`
	ID string `json:"id"`
}
