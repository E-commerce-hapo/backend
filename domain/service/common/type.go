package common

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/kiem-toan/pkg/common/cm_array"
	"github.com/kiem-toan/pkg/errorx"
)

type Paging struct {
	Limit  int
	Page   int
	Offset int
	Sorts  []string
}

func (p *Paging) Validate(model interface{}) error {
	if model == nil {
		return nil
	}
	r := reflect.ValueOf(model).Elem()
	typeOfT := r.Type()
	var normalizedColumns []string
	for i := 0; i < r.NumField(); i++ {
		fieldName := typeOfT.Field(i).Name
		normalizedColumns = append(normalizedColumns, strings.ToLower(fieldName))
	}

	if len(p.Sorts) > 0 {
		// Example Sort: "Created_at desc"
		for _, sort := range p.Sorts {
			lowerSort := strings.ToLower(sort)             // created_at desc
			lowerSortStrs := strings.Split(lowerSort, " ") // ["created_at", "desc"]
			if len(lowerSortStrs) != 2 {
				return errorx.Errorf(http.StatusBadRequest, nil, "Sort does not valid")
			}
			sortField := lowerSortStrs[0]                             // "created_at"
			normalizedField := strings.ReplaceAll(sortField, "_", "") // createdat
			isContained := cm_array.ListStringsContain(normalizedColumns, normalizedField)
			if !isContained {
				return errorx.Errorf(http.StatusBadRequest, nil, fmt.Sprintf("Sorted field %v does not exist in table", normalizedField))
			}
		}
	}
	return nil
}
