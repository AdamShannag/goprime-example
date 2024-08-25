package resource

import (
	"backend/model"
	"github.com/AdamShannag/goprime/prime"
)

type CustomersListRequest struct {
	First     string      `json:"first"`
	Rows      string      `json:"rows"`
	SortOrder string      `json:"sort_order"`
	SortField string      `json:"sort_field"`
	Filters   prime.Specs `json:"filters,omitempty"`
}

type CustomersListResponse struct {
	Customers    []*model.Customer `json:"customers"`
	TotalRecords int               `json:"totalRecords"`
}
