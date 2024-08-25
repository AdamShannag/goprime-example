package repository

import (
	"backend/model"
	"backend/resource"
)

type PrimeRepository interface {
	GetAllCustomers(request resource.CustomersListRequest) (*resource.CustomersListResponse, error)
	GetAllRepresentatives() ([]*model.Representative, error)
}
