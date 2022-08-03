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

func (h *CrudHandler) GeneratCrud(params *model.CrudBasic) error {
	columns, filters, err := h.tablesRepo.GetneratColumnFromString(params)
	if err != nil {
		fmt.Println("erer")
		fmt.Println(err.Error())
		return err
	}
	crud := &model.Crud{
		Table:   params.Table,
		Columns: *columns,
		Filters: *filters,
	}
	// if params.Filters != nil {
	// 	var filters []model.Column
	// 	for _, filter := range params.Filters {
	// 		var required model.Required
	// 		required = model.RequiredFactory[fmt.Sprintf("%s_%s", filter.Type, convertBoolToYesOrNo(&filter.Required))]
	// 		fmt.Println("required")
	// 		fmt.Println(required)
	// 		fmt.Println(fmt.Sprintf("%s_%s", filter.Type, convertBoolToYesOrNo(&filter.Required)))
	// 		filterType := model.FilterFactory(&filter.Type, &required)
	// 		filterColumn := model.Column{Name: filter.Name, Type: filter.Type, RequiredType: required, FilterType: filterType}
	// 		filters = append(filters, filterColumn)
	// 	}
	// 	crud.Filters = filters
	// }
	err = h.GeneratListProc(crud)
	if err != nil {
		return err
	}
	return nil
}

func (h *CrudHandler) GeneratListProc(params *model.Crud) error {
	columns := generateColumnsForSelect(params.Columns)
	filters := generateFiltersForSelect(params.Filters)
	joins := generateJoinsForSelect(params.Joins, params.Table)
	open := createProc(&params.Table, params.Filters)
	close := closeProc()
	var genratedFilters string
	if filters != nil {
		genratedFilters = fmt.Sprintf("WHERE \n %s", *filters)
	}
	proc := fmt.Sprintf("%s \n SELECT \n %s \n FROM %s %s %s \n  SET @rows = @@ROWCOUNT \n %s", *open, *columns, params.Table, *joins, genratedFilters, *close)
	fmt.Print(proc)
	return nil
}

func (h *CrudHandler) GeneratListProcTest(params *model.Crud) error {

	return nil
}
