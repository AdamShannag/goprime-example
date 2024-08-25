package repository

import (
	"backend/model"
	"backend/resource"
	"database/sql"
	"fmt"
	"github.com/AdamShannag/goprime/prime"
)

type mysqlPrimeRepository struct {
	db    *sql.DB
	prime *prime.Filter
}

func NewMysqlPrimeRepository(db *sql.DB, prime *prime.Filter) PrimeRepository {
	return &mysqlPrimeRepository{db: db, prime: prime}
}

func (m *mysqlPrimeRepository) GetAllCustomers(request resource.CustomersListRequest) (*resource.CustomersListResponse, error) {
	vals, condition, err := m.prime.Sql(request.Filters)
	if err != nil {
		return nil, err
	}

	var totalCount int
	cQuery := countSql
	if condition != "" {
		cQuery = cQuery + fmt.Sprintf(" WHERE %s", condition)
	}

	err = m.db.QueryRow(cQuery, vals...).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	dQuery := selectCustomersSql
	if condition != "" {
		dQuery = dQuery + fmt.Sprintf(" WHERE %s", condition)
	}

	if request.First != "" && request.Rows != "" {
		dQuery = dQuery + fmt.Sprintf(limitSql, "?", "?")
		vals = append(vals, request.Rows, request.First)
	}

	rows, err := m.db.Query(dQuery, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var c model.Customer
		var countryName, countryCode, repName, repImage sql.NullString
		var repID int
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Company,
			&c.Date,
			&c.Status,
			&c.Verified,
			&c.Activity,
			&c.Balance,
			&countryName,
			&countryCode,
			&repID,
			&repName,
			&repImage,
		)
		if err != nil {
			return nil, err
		}

		c.Country = model.Country{
			Name: countryName.String,
			Code: countryCode.String,
		}
		c.Representative = model.Representative{
			ID:    repID,
			Name:  repName.String,
			Image: repImage.String,
		}

		customers = append(customers, &c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &resource.CustomersListResponse{
		Customers:    customers,
		TotalRecords: totalCount,
	}, nil

}

func (m *mysqlPrimeRepository) GetAllRepresentatives() ([]*model.Representative, error) {
	rows, err := m.db.Query(selectRepresentativesSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var representatives []*model.Representative
	for rows.Next() {
		var representative model.Representative
		err := rows.Scan(
			&representative.ID,
			&representative.Name,
			&representative.Image)
		if err != nil {
			return nil, err
		}

		representatives = append(representatives, &representative)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return representatives, nil
}
