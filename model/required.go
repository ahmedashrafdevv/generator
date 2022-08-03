package model

import (
	"fmt"
	"generator/utils"
)

var RequiredFactory map[string]Required = map[string]Required{
	"int_YES":      IntRequired{Required: true},
	"real_YES":     IntRequired{Required: true},
	"varchar_YES":  StringRequired{Required: true},
	"binary_YES":   BinaryRequired{Required: true},
	"datetime_YES": DateRequired{Required: true},
	"date_YES":     DateRequired{Required: true},
	"int_NO":       IntRequired{Required: false},
	"real_NO":      IntRequired{Required: false},
	"varchar_NO":   StringRequired{Required: false},
	"binary_NO":    BinaryRequired{Required: false},
	"datetime_NO":  DateRequired{Required: false},
	"date_NO":      DateRequired{Required: false},
}

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
type BinaryRequired struct {
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

func (filter BinaryRequired) GenerateFilter(name string) string {
	var isNullGen string = "@" + name + " "
	if !filter.Required {
		newNameInstance := name
		// utils.RemoveLastNChars(&newNameInstance, 5)
		isNullGen = fmt.Sprintf("ISNULL(CAST(@%s AS VARCHAR) , %s)", name, newNameInstance)
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

func (filter BinaryRequired) GenerateParameter(name string) string {
	if filter.Required {
		return ""
	}
	return "= NULL"
}

func (filter IntRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(%s , 0),\n", name)
	}
	return isNullGen
}

func (filter StringRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(%s , ''),\n", name)
	}
	return isNullGen
}

func (filter DateRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(%s , ''),\n", name)
	}
	return isNullGen
}

func (filter BinaryRequired) GenerateColumn(name string) string {
	var isNullGen string = name + ",\n"
	if !filter.Required {
		isNullGen = fmt.Sprintf("ISNULL(CAST(%s AS VARCHAR) , ''),\n", name)
	}
	return isNullGen
}
