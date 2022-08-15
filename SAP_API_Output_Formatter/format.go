package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-territory-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSalesTerritoryCollection(raw []byte, l *logger.Logger) ([]SalesTerritoryCollection, error) {
	pm := &responses.SalesTerritoryCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesTerritoryCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesTerritoryCollection := make([]SalesTerritoryCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesTerritoryCollection = append(salesTerritoryCollection, SalesTerritoryCollection{
			ObjectID:               data.ObjectID,
			ID:                     data.ID,
			Name:                   data.Name,
			ParentID:               data.ParentID,
			EmployeeResponsible:    data.EmployeeResponsible,
			HierarchyLevelCode:     data.HierarchyLevelCode,
			HierarchyLevelCodeText: data.HierarchyLevelCodeText,
			EntityLastChangedOn:    data.EntityLastChangedOn,
			ETag:                   data.ETag,
		})
	}

	return salesTerritoryCollection, nil
}
