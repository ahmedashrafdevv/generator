package model

import (
	"fmt"
	"generator/utils"
)

// defining an interface
type Required interface {
	GenerateFilter(name string) string
	GenerateColumn(name string) string
	GenerateParameter(name string) string
}
type IntRequired struct {
	Required bool
}
type StringRequired struct {
	Required bool
}
type DateRequired struct {
	Required bool
}

func (filter IntRequired) GenerateFilter(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("dbo.ISZERO(@%s , %s)", name, name)
	}
	return isNullGen
}

func (filter StringRequired) GenerateFilter(name string) string {
	var isNullGen string = "@" + name
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s)", name, name)
	}
	return isNullGen
}

func (filter DateRequired) GenerateFilter(name string) string {
	var isNullGen string = "@" + name + " "
	if !filter.Required {
		newNameInstance := name
		utils.RemoveLastNChars(&newNameInstance, 5)
		isNullGen = fmt.Sprintf("ISNULL(@%s , %s)", name, newNameInstance)
	}
	return isNullGen
}

func (filter StringRequired) GenerateParameter(name string) string {
	if filter.Required {
		return ""
	}
	return "= NULL"
}

func (filter IntRequired) GenerateParameter(name string) string {
	if filter.Required {
		return ""
	}
	return "= 0"
}

func (filter DateRequired) GenerateParameter(name string) string {
	if filter.Required {
		return ""
	}
	return "= NULL"
}

func (filter IntRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , 0),\n", name)
	}
	return isNullGen
}

func (filter StringRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , ''),\n", name)
	}
	return isNullGen
}

func (filter DateRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(@%s , ''),\n", name)
	}
	return isNullGen
}
