package sap_api_output_formatter

type SalesTerritory struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	SalesTerritoryCode string `json:"sales_territory_code"`
	Deleted            bool   `json:"deleted"`
}

type SalesTerritoryCollection struct {
	ObjectID               string `json:"ObjectID"`
	ID                     string `json:"Id"`
	Name                   string `json:"Name"`
	ParentID               string `json:"ParentID"`
	EmployeeResponsible    string `json:"EmployeeResponsible"`
	HierarchyLevelCode     string `json:"HierarchyLevelCode"`
	HierarchyLevelCodeText string `json:"HierarchyLevelCodeText"`
	EntityLastChangedOn    string `json:"EntityLastChangedOn"`
	ETag                   string `json:"ETag"`
}