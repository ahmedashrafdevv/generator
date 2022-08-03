package model

var percent = "'%'"

type CrudBasic struct {
	Table   string
	Columns []string
	Filters []FilterBasic
}

type FilterBasic struct {
	Name     string
	Required bool
	Type     string
}
type Crud struct {
	Table   string
	Joins   []Join
	Columns []Column
	Filters []Column
}
type Join struct {
	Table   string
	Primary string
	Foreign string
	Type    string
}

type Column struct {
	Name         string
	Type         string
	RequiredType Required
	FilterType   Filter
}
