package filter

import "fmt"

var percent = "'%'"

// defining an interface
type Filter interface {
	GenerateFilter(name string) string
}

// declaring a struct
type SearchFilter struct {
	Required bool
}
type EqualFilter struct {
	Required bool
}

// function to calculate
// area of a rectangle
func (filter SearchFilter) GenerateFilter(name string) string {
	var generatedFilter string
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s)", name, name)
	}
	generatedFilter = fmt.Sprintf("AND %s LIKE CONCAT(%s , %s , %s) \n", name, percent, isNullGen, percent)
	return generatedFilter
}

// function to calculate
// area of a square
func (filter EqualFilter) GenerateFilter(name string) string {
	var generatedFilter string
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s)", name, name)
	}
	generatedFilter = fmt.Sprintf("AND %s =  %s \n", name, isNullGen)
	return generatedFilter
}
