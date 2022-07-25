package model

import (
	"fmt"
)

// defining an interface
type Filter interface {
	GenerateFilterQuery(name string) string
	GenerateParameterQuery(name string, dataType string) string
}

// declaring a struct
type SearchFilter struct {
	Required Required
}
type EqualFilter struct {
	Required Required
}

type DateFilter struct {
	Required Required
}

func _defaultParameter(name string, dataType string, filterStr string) string {
	f := fmt.Sprintf("@%s %s %s,\n", name, dataType, filterStr)
	return f
}
func (filter EqualFilter) GenerateParameterQuery(name string, dataType string) string {
	var filterStr string = filter.Required.GenerateParameter(name)
	return _defaultParameter(name, dataType, filterStr)
}
func (filter SearchFilter) GenerateParameterQuery(name string, dataType string) string {
	var filterStr string = filter.Required.GenerateParameter(name)
	return _defaultParameter(name, dataType, filterStr)
}

func (filter DateFilter) GenerateParameterQuery(name string, dataType string) string {
	var filterStr string = filter.Required.GenerateParameter(name)
	fromParam := _defaultParameter(name+"_from", dataType, filterStr)
	toParam := _defaultParameter(name+"_to", dataType, filterStr)
	generatedFilter := fmt.Sprintf("%s \n %s", fromParam, toParam)
	return generatedFilter
}

func (filter SearchFilter) GenerateFilterQuery(name string) string {
	var generatedFilter string
	var filterStr string = filter.Required.GenerateFilter(name)
	generatedFilter = fmt.Sprintf("AND %s LIKE CONCAT(%s , %s , %s) \n", name, percent, filterStr, percent)
	return generatedFilter
}

func (filter EqualFilter) GenerateFilterQuery(name string) string {
	var generatedFilter string
	var filterStr string = filter.Required.GenerateFilter(name)
	generatedFilter = fmt.Sprintf("AND %s =  %s \n", name, filterStr)
	return generatedFilter
}

func (filter DateFilter) GenerateFilterQuery(name string) string {
	var generatedFilter string
	var fromFilter string = filter.Required.GenerateFilter(name + "_from")

	var toFilter string = filter.Required.GenerateFilter(name + "_to  ")

	generatedFilter += fmt.Sprintf("AND %s >=  %s \n", name, fromFilter)
	generatedFilter += fmt.Sprintf("AND %s <=  %s \n", name, toFilter)
	return generatedFilter
}
