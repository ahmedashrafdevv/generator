package required

import (
	"fmt"
)

// // defining an interface
// type Required interface {
// 	GenerateFilter() string
// 	GenerateColumn() string
// 	// GeneratColumn(name string) string
// }

// declaring a struct
type IntFilter struct {
	Required bool
}
type StringFilter struct {
	Required bool
}

func (filter IntFilter) GenerateFilter(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISZERO(@%s , %s)", name, name)
	}
	return isNullGen
}

func (filter StringFilter) GenerateFilter(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s)", name, name)
	}
	return isNullGen
}

func (filter IntFilter) GenerateColumn(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , 0),\n", name)
	}
	return isNullGen
}

func (filter StringFilter) GenerateColumn(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s),\n", name, name)
	}
	return isNullGen
}
