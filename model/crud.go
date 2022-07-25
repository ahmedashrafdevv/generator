package model

var percent = "'%'"

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
}

type Column struct {
	Name     string
	Type     string
	Required Required
	Filter   Filter
}
