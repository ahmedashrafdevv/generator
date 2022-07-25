package crud

import (
	"fmt"
	"generator/model"

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
		filter := currentFilter.Filter.GenerateParameterQuery(currentFilter.Name, currentFilter.Type)
		filtersGen += filter
	}
	removeLastNChars(&filtersGen, 2)
	procName := generateProcName(name, &action)
	str := fmt.Sprintf(
		"DROP PROC IF EXISTS %s \nGO \n CREATE PROC %s (\n %s \n) \n AS \n BEGIN \n", procName, procName, filtersGen)
	return &str
}

func generateColumnsForSelect(columns []model.Column) *string {
	var columnsGen string
	for i := 0; i < len(columns); i++ {
		currentColumn := columns[i]
		columnName := currentColumn.Required.GenerateColumn(currentColumn.Name)
		columnsGen += columnName
	}
	removeLastNChars(&columnsGen, 2)

	return &columnsGen
}

func generateFiltersForSelect(filters []model.Column) *string {
	var filtersGen string
	for i := 0; i < len(filters); i++ {
		currentFilter := filters[i]
		filter := currentFilter.Filter.GenerateFilterQuery(currentFilter.Name)
		filtersGen += filter
	}
	filtersGen = removeFirstNChars(filtersGen, 3)
	removeLastNChars(&filtersGen, 1)
	return &filtersGen
}

func removeLastNChars(str *string, n int) {
	*str = string([]rune(*str)[:len(*str)-n])
}
func removeFirstNChars(str string, n int) string {
	return str[n:]
}
func closeProc() *string {
	str := "END"
	return &str
}
func generateProcName(table *string, action *string) string {
	procName := strcase.ToCamel(fmt.Sprintf("%s%s", *table, *action))
	return procName
}
