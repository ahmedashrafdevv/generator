package main

import (
	"generator/crud"
	"generator/db"
	"generator/model"
	"generator/repo"
)

func main() {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	tablesRepo := repo.NewTablesRepo(db)
	cr := crud.NewCrudHandler(tablesRepo)
	// var columns []model.Column
	// columns = append(columns, model.Column{Name: "customer_id", Type: "INT", Required: false})
	// columns = append(columns, model.Column{Name: "created_at", Type: "DATETIME", Required: false})
	// columns = append(columns, model.Column{Name: "name", Type: "VARCHAR(250)", Required: false})
	// columns = append(columns, model.Column{Name: "phone", Type: "VARCHAR(250)"})
	// params := &model.Crud{
	// 	Table:   "staging.customers",
	// 	Columns: columns,
	// 	Filters: columns,
	// }
	// cr.GeneratCrud(params)

	// var searchFilter filter.Filter = filter.SearchFilter{
	// 	Required: false,
	// }

	// // declaring a square instance
	// var equalFilter filter.Filter = filter.EqualFilter{
	// 	Required: false,
	// }
	// filter := "name"
	var columns []model.Column
	column := model.Column{Name: "name", Type: "VARCHAR(250)", Filter: model.SearchFilter{Required: model.StringRequired{Required: false}}, Required: model.StringRequired{Required: true}}
	column2 := model.Column{Name: "customer_id", Type: "INT", Filter: model.EqualFilter{Required: model.IntRequired{Required: false}}, Required: model.IntRequired{Required: true}}
	column3 := model.Column{Name: "created_at", Type: "DATETIME", Filter: model.DateFilter{Required: model.DateRequired{Required: true}}, Required: model.DateRequired{Required: true}}
	columns = append(columns, column)
	columns = append(columns, column2)
	columns = append(columns, column3)
	params := &model.Crud{
		Table:   "staging.customers",
		Columns: columns,
		Filters: columns,
	}
	cr.GeneratCrud(params)

}
