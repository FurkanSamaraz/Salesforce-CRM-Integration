package crm_structures

// ! /list/{listid}/{objectname} kurulmadı
type CreateList struct {
	Attributes struct {
		Cost                       int    `json:"cost"`
		CostBase                   int    `json:"cost_base"`
		Createdby                  string `json:"createdby"`
		Createdbyname              string `json:"createdbyname"`
		Createdbyyominame          string `json:"createdbyyominame"`
		Createdfromcode            string `json:"createdfromcode"`
		Createdon                  int    `json:"createdon"`
		Createdonbehalfby          string `json:"createdonbehalfby"`
		Createdonbehalfbyname      string `json:"createdonbehalfbyname"`
		Createdonbehalfbyyominame  string `json:"createdonbehalfbyyominame"`
		Description                string `json:"description"`
		Donotsendonoptout          bool   `json:"donotsendonoptout"`
		Exchangerate               int    `json:"exchangerate"`
		Ignoreinactivelistmembers  bool   `json:"ignoreinactivelistmembers"`
		Importsequencenumber       string `json:"importsequencenumber"`
		Lastusedon                 int    `json:"lastusedon"`
		Listid                     string `json:"listid"`
		Listname                   string `json:"listname"`
		Lockstatus                 bool   `json:"lockstatus"`
		Membercount                int    `json:"membercount"`
		Membertype                 int    `json:"membertype"`
		Modifiedby                 string `json:"modifiedby"`
		Modifiedbyname             string `json:"modifiedbyname"`
		Modifiedbyyominame         string `json:"modifiedbyyominame"`
		Modifiedon                 int    `json:"modifiedon"`
		Modifiedonbehalfby         string `json:"modifiedonbehalfby"`
		Modifiedonbehalfbyname     string `json:"modifiedonbehalfbyname"`
		Modifiedonbehalfbyyominame string `json:"modifiedonbehalfbyyominame"`
		Overriddencreatedon        int    `json:"overriddencreatedon"`
		Ownerid                    string `json:"ownerid"`
		Owneridname                string `json:"owneridname"`
		Owneridtype                string `json:"owneridtype"`
		Owneridyominame            string `json:"owneridyominame"`
		Owningbusinessunit         string `json:"owningbusinessunit"`
		Owningteam                 string `json:"owningteam"`
		Owninguser                 string `json:"owninguser"`
		Processid                  string `json:"processid"`
		Purpose                    string `json:"purpose"`
		Query                      string `json:"query"`
		Source                     string `json:"source"`
		Stageid                    string `json:"stageid"`
		Statecode                  int    `json:"statecode"`
		Statuscode                 int    `json:"statuscode"`
		Timezoneruleversionnumber  int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid      string `json:"transactioncurrencyid"`
		Transactioncurrencyidname  string `json:"transactioncurrencyidname"`
		Traversedpath              string `json:"traversedpath"`
		Type                       bool   `json:"type"`
		Utcconversiontimezonecode  int    `json:"utcconversiontimezonecode"`
		Versionnumber              int    `json:"versionnumber"`
	} `json:"attributes"`
	ID string `json:"id"`
}
type GetList struct {
	Attributes struct {
		Cost                       int    `json:"cost"`
		CostBase                   int    `json:"cost_base"`
		Createdby                  string `json:"createdby"`
		Createdbyname              string `json:"createdbyname"`
		Createdbyyominame          string `json:"createdbyyominame"`
		Createdfromcode            string `json:"createdfromcode"`
		Createdon                  int    `json:"createdon"`
		Createdonbehalfby          string `json:"createdonbehalfby"`
		Createdonbehalfbyname      string `json:"createdonbehalfbyname"`
		Createdonbehalfbyyominame  string `json:"createdonbehalfbyyominame"`
		Description                string `json:"description"`
		Donotsendonoptout          bool   `json:"donotsendonoptout"`
		Exchangerate               int    `json:"exchangerate"`
		Ignoreinactivelistmembers  bool   `json:"ignoreinactivelistmembers"`
		Importsequencenumber       string `json:"importsequencenumber"`
		Lastusedon                 int    `json:"lastusedon"`
		Listid                     string `json:"listid"`
		Listname                   string `json:"listname"`
		Lockstatus                 bool   `json:"lockstatus"`
		Membercount                int    `json:"membercount"`
		Membertype                 int    `json:"membertype"`
		Modifiedby                 string `json:"modifiedby"`
		Modifiedbyname             string `json:"modifiedbyname"`
		Modifiedbyyominame         string `json:"modifiedbyyominame"`
		Modifiedon                 int    `json:"modifiedon"`
		Modifiedonbehalfby         string `json:"modifiedonbehalfby"`
		Modifiedonbehalfbyname     string `json:"modifiedonbehalfbyname"`
		Modifiedonbehalfbyyominame string `json:"modifiedonbehalfbyyominame"`
		Overriddencreatedon        int    `json:"overriddencreatedon"`
		Ownerid                    string `json:"ownerid"`
		Owneridname                string `json:"owneridname"`
		Owneridtype                string `json:"owneridtype"`
		Owneridyominame            string `json:"owneridyominame"`
		Owningbusinessunit         string `json:"owningbusinessunit"`
		Owningteam                 string `json:"owningteam"`
		Owninguser                 string `json:"owninguser"`
		Processid                  string `json:"processid"`
		Purpose                    string `json:"purpose"`
		Query                      string `json:"query"`
		Source                     string `json:"source"`
		Stageid                    string `json:"stageid"`
		Statecode                  int    `json:"statecode"`
		Statuscode                 int    `json:"statuscode"`
		Timezoneruleversionnumber  int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid      string `json:"transactioncurrencyid"`
		Transactioncurrencyidname  string `json:"transactioncurrencyidname"`
		Traversedpath              string `json:"traversedpath"`
		Type                       bool   `json:"type"`
		Utcconversiontimezonecode  int    `json:"utcconversiontimezonecode"`
		Versionnumber              int    `json:"versionnumber"`
	} `json:"attributes"`
	ID string `json:"id"`
}
type UpdateList struct {
	Attributes struct {
		Cost                       int    `json:"cost"`
		CostBase                   int    `json:"cost_base"`
		Createdby                  string `json:"createdby"`
		Createdbyname              string `json:"createdbyname"`
		Createdbyyominame          string `json:"createdbyyominame"`
		Createdfromcode            string `json:"createdfromcode"`
		Createdon                  int    `json:"createdon"`
		Createdonbehalfby          string `json:"createdonbehalfby"`
		Createdonbehalfbyname      string `json:"createdonbehalfbyname"`
		Createdonbehalfbyyominame  string `json:"createdonbehalfbyyominame"`
		Description                string `json:"description"`
		Donotsendonoptout          bool   `json:"donotsendonoptout"`
		Exchangerate               int    `json:"exchangerate"`
		Ignoreinactivelistmembers  bool   `json:"ignoreinactivelistmembers"`
		Importsequencenumber       string `json:"importsequencenumber"`
		Lastusedon                 int    `json:"lastusedon"`
		Listid                     string `json:"listid"`
		Listname                   string `json:"listname"`
		Lockstatus                 bool   `json:"lockstatus"`
		Membercount                int    `json:"membercount"`
		Membertype                 int    `json:"membertype"`
		Modifiedby                 string `json:"modifiedby"`
		Modifiedbyname             string `json:"modifiedbyname"`
		Modifiedbyyominame         string `json:"modifiedbyyominame"`
		Modifiedon                 int    `json:"modifiedon"`
		Modifiedonbehalfby         string `json:"modifiedonbehalfby"`
		Modifiedonbehalfbyname     string `json:"modifiedonbehalfbyname"`
		Modifiedonbehalfbyyominame string `json:"modifiedonbehalfbyyominame"`
		Overriddencreatedon        int    `json:"overriddencreatedon"`
		Ownerid                    string `json:"ownerid"`
		Owneridname                string `json:"owneridname"`
		Owneridtype                string `json:"owneridtype"`
		Owneridyominame            string `json:"owneridyominame"`
		Owningbusinessunit         string `json:"owningbusinessunit"`
		Owningteam                 string `json:"owningteam"`
		Owninguser                 string `json:"owninguser"`
		Processid                  string `json:"processid"`
		Purpose                    string `json:"purpose"`
		Query                      string `json:"query"`
		Source                     string `json:"source"`
		Stageid                    string `json:"stageid"`
		Statecode                  int    `json:"statecode"`
		Statuscode                 int    `json:"statuscode"`
		Timezoneruleversionnumber  int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid      string `json:"transactioncurrencyid"`
		Transactioncurrencyidname  string `json:"transactioncurrencyidname"`
		Traversedpath              string `json:"traversedpath"`
		Type                       bool   `json:"type"`
		Utcconversiontimezonecode  int    `json:"utcconversiontimezonecode"`
		Versionnumber              int    `json:"versionnumber"`
	} `json:"attributes"`
	ID string `json:"id"`
}
