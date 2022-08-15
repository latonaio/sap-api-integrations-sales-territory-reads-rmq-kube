package responses

type SalesTerritoryCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID               string `json:"ObjectID"`
			ID                     string `json:"Id"`
			Name                   string `json:"Name"`
			ParentID               string `json:"ParentID"`
			EmployeeResponsible    string `json:"EmployeeResponsible"`
			HierarchyLevelCode     string `json:"HierarchyLevelCode"`
			HierarchyLevelCodeText string `json:"HierarchyLevelCodeText"`
			EntityLastChangedOn    string `json:"EntityLastChangedOn"`
			ETag                   string `json:"ETag"`
			SalesTerritoryAccount  struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesTerritoryAccount"`
			SalesTerritoryTeam struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesTerritoryTeam"`
		} `json:"results"`
	} `json:"d"`
}
