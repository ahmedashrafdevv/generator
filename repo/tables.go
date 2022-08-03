package repo

import (
	"fmt"
	"generator/model"
	"generator/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

type TablesRepo struct {
	db *gorm.DB
}

func NewTablesRepo(db *gorm.DB) TablesRepo {
	return TablesRepo{
		db: db,
	}
}

func (ur *TablesRepo) GetneratColumnFromString(params *model.CrudBasic) (*[]model.Column, *[]model.Column, error) {
	columns := (params.Columns)
	table := (params.Table)
	filters := (params.Filters)
	var columnsResp []model.Column
	var filtersResp []model.Column
	cols := strings.Join(columns, ",")
	rows, err := ur.db.Raw("EXEC columns_find_by_names @columns = ? , @table = ?", cols, table).Rows()
	defer rows.Close()
	if err != nil {
		return nil, nil, err
	}
	for rows.Next() {
		var rec model.Column
		var isNull string
		var length string
		err := rows.Scan(&rec.Name, &rec.Type, &isNull, &length)
		if err != nil {
			return nil, nil, err
		}
		rec.RequiredType = model.RequiredFactory[fmt.Sprintf("%s_%s", rec.Type, isNull)]
		if rec.Type == "varchar" {
			rec.Type += fmt.Sprintf("(%s)", length)
		}
		columnsResp = append(columnsResp, rec)
	}

	if filters != nil {
		var colNames []string
		for _, fil := range filters {
			colNames = append(colNames, fil.Name)
		}
		colNamesArg := strings.Join(colNames, ",")
		filterRows, err := ur.db.Raw("EXEC columns_find_by_names @columns = ? , @table = ?", colNamesArg, table).Rows()
		defer filterRows.Close()
		if err != nil {
			return nil, nil, err
		}
		for filterRows.Next() {
			var filterRec model.Column
			var filterLength string
			var isFilterNull string
			err := filterRows.Scan(&filterRec.Name, &filterRec.Type, &isFilterNull, &filterLength)
			if err != nil {
				return nil, nil, err
			}
			filterBasic := findFilterBasicFromName(&filters, &filterRec.Name)
			requiredTypeKey := fmt.Sprintf("%s_%s", filterRec.Type, utils.ConvertBoolToYesOrNo(&filterBasic.Required))
			fmt.Println(requiredTypeKey)
			filterRec.RequiredType = model.RequiredFactory[requiredTypeKey]
			filterRec.FilterType = model.FilterFactory(&filterBasic.Type, &filterRec.RequiredType)
			filtersResp = append(filtersResp, filterRec)
		}
	}
	return &columnsResp, &filtersResp, nil
}

func findFilterBasicFromName(list *[]model.FilterBasic, name *string) model.FilterBasic {
	for _, filter := range *(list) {
		if filter.Name == *name {
			return filter
		}
	}
	return model.FilterBasic{}
}

func generateCSV(arr []string) string {
	var res string
	for i := 0; i < len(arr); i++ {
		res += fmt.Sprintf("'%s',", arr[i])
	}
	utils.RemoveLastNChars(&res, 2)
	utils.RemoveFirstNChars(&res, 1)
	return res
}
