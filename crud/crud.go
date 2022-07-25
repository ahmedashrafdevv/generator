package crud

import (
	"fmt"
	"generator/model"
	"generator/repo"
)

type CrudHandler struct {
	tablesRepo repo.TablesRepo
}

func NewCrudHandler(tablesRepo repo.TablesRepo) *CrudHandler {
	return &CrudHandler{
		tablesRepo: tablesRepo,
	}
}

func (h *CrudHandler) GeneratCrud(params *model.Crud) error {
	err := h.GeneratListProc(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *CrudHandler) GeneratListProc(params *model.Crud) error {
	// final result = SELECT name , phone FROM customers WHERE name

	columns := generateColumnsForSelect(params.Columns)
	filters := generateFiltersForSelect(params.Filters)

	open := createProc(&params.Table, params.Filters)
	close := closeProc()
	proc := fmt.Sprintf("%s \n SELECT \n %s \n FROM %s WHERE \n %s \n %s", *open, *columns, params.Table, *filters, *close)
	fmt.Print(proc)
	return nil
}
