package paging

import (
	"net/http"
	"strings"

	"github.com/kiem-toan/domain/enums/order_direction"
	"github.com/kiem-toan/pkg/errorx"

	service_common "github.com/kiem-toan/domain/service/common"
)

type Paging struct {
	Limit int      `json:"limit"`
	Page  int      `json:"page"`
	Sort  []string `json:"sort"`
}

type PagingInfo struct {
	Total int      `json:"total"`
	Limit int      `json:"limit"`
	Page  int      `json:"page"`
	Sort  []string `json:"sort"`
}

func (p *Paging) Convert_api_Paging_to_service_Paging() (*service_common.Paging, error) {
	if p == nil {
		return &service_common.Paging{
			Limit:  100,
			Page:   1,
			Offset: 0,
		}, nil
	}
	// Maximum limit = 100
	if p.Limit == 0 || p.Limit > 100 {
		p.Limit = 100
	}

	if p.Page == 0 {
		p.Page = 1
	}

	// Example sort :
	// ["created_at desc", "updated_at ascz"]
	for _, s := range p.Sort {
		splitedSortStrs := strings.Split(s, " ")
		if len(splitedSortStrs) != 2 {
			return nil, errorx.Errorf(http.StatusBadRequest, nil, "Sort does not valid")
		}
		orderDirection := order_direction.OrderDirection(strings.Split(s, " ")[1])
		if orderDirection != order_direction.Asc && orderDirection != order_direction.Desc {
			return nil, errorx.Errorf(http.StatusBadRequest, nil, "Sort direction does not valid (enum: asc, desc)")
		}
	}
	return &service_common.Paging{
		Limit:  p.Limit,
		Page:   p.Page,
		Offset: (p.Page - 1) * p.Limit,
		Sorts:  p.Sort,
	}, nil
}

func (p *Paging) Convert_api_Paging_to_api_PagingInfo(total int) *PagingInfo {
	return &PagingInfo{
		Total: total,
		Limit: p.Limit,
		Page:  p.Page,
		Sort:  p.Sort,
	}
}
