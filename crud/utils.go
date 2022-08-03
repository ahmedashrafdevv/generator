package crud

import (
	"fmt"
	"generator/model"
	"generator/utils"

	"github.com/iancoleman/strcase"
)

func nullValueFromType(typeKey *string) string {
	if *typeKey == "INT" {
		return "0"
	}
	if *typeKey == "BOOL" {
		return "FALSE"
	}
	return "''"
}
func createProc(name *string, filters []model.Column) *string {
	var filtersGen string
	action := "list"
	for i := 0; i < len(filters); i++ {
		currentFilter := filters[i]
		filter := currentFilter.FilterType.GenerateParameterQuery(currentFilter.Name, currentFilter.Type)
		filtersGen += filter
	}
	filtersGen += "@rows INT OUTPUT"
	procName := generateProcName(name, &action)
	str := fmt.Sprintf(
		"DROP PROC IF EXISTS %s \nGO \n CREATE PROC %s (\n %s \n) \n AS \n BEGIN \n", procName, procName, filtersGen)
	return &str
}

func generateColumnsForSelect(columns []model.Column) *string {
	var columnsGen string
	for i := 0; i < len(columns); i++ {
		currentColumn := columns[i]
		columnName := currentColumn.RequiredType.GenerateColumn(currentColumn.Name)
		columnsGen += columnName
	}
	utils.RemoveLastNChars(&columnsGen, 2)

	return &columnsGen
}

func generateJoinsForSelect(joins []model.Join, table string) *string {
	var joinsGen string
	for i := 0; i < len(joins); i++ {
		currentJoin := joins[i]
		join := fmt.Sprintf("\n %s JOIN %s \n ON %s.%s = %s.%s \n", currentJoin.Type, currentJoin.Table, table, currentJoin.Primary, currentJoin.Table, currentJoin.Foreign)
		// Join := currentJoin.Join.GenerateJoinQuery(currentJoin.Name)
		joinsGen += join
	}
	return &joinsGen
}
func generateFiltersForSelect(filters []model.Column) *string {
	if filters == nil {
		return nil
	}
	var filtersGen string
	for i := 0; i < len(filters); i++ {
		currentFilter := filters[i]
		filter := currentFilter.FilterType.GenerateFilterQuery(currentFilter.Name)
		filtersGen += filter
	}
	utils.RemoveFirstNChars(&filtersGen, 3)
	utils.RemoveLastNChars(&filtersGen, 1)
	return &filtersGen
}

func convertBoolToYesOrNo(isTrue *bool) string {
	if *isTrue {
		return "YES"
	}
	return "NO"
}
func closeProc() *string {
	str := "END"
	return &str
}
func generateProcName(table *string, action *string) string {
	procName := strcase.ToCamel(fmt.Sprintf("%s%s", *table, *action))
	return procName
}
