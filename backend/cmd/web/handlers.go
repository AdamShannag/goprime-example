package main

import (
	"backend/resource"
	"encoding/json"
	"net/http"
)

func (app *application) ListMysqlCustomers(w http.ResponseWriter, r *http.Request) {
	customersListRequest, err := app.parseParams(r)
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	customers, err := app.mysqlPrimeRepo.GetAllCustomers(customersListRequest)
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.WriteJSON(w, http.StatusOK, customers)
}

func (app *application) ListMysqlRepresentatives(w http.ResponseWriter, _ *http.Request) {
	reps, err := app.mysqlPrimeRepo.GetAllRepresentatives()
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.WriteJSON(w, http.StatusOK, reps)
}

func (app *application) ListPostgresCustomers(w http.ResponseWriter, r *http.Request) {
	customersListRequest, err := app.parseParams(r)
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	customers, err := app.postgresPrimeRepo.GetAllCustomers(customersListRequest)
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.WriteJSON(w, http.StatusOK, customers)
}

func (app *application) ListPostgresRepresentatives(w http.ResponseWriter, _ *http.Request) {
	reps, err := app.postgresPrimeRepo.GetAllRepresentatives()
	if err != nil {
		_ = app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.WriteJSON(w, http.StatusOK, reps)
}

func (app *application) parseParams(r *http.Request) (resource.CustomersListRequest, error) {
	req := resource.CustomersListRequest{
		First:     r.URL.Query().Get("first"),
		Rows:      r.URL.Query().Get("rows"),
		SortOrder: r.URL.Query().Get("sortOrder"),
		SortField: r.URL.Query().Get("sortField"),
		Filters:   nil,
	}

	if r.URL.Query().Get("filters") != "" {
		err := json.Unmarshal([]byte(r.URL.Query().Get("filters")), &req.Filters)

		if err != nil {
			return resource.CustomersListRequest{}, err
		}
	}

	return req, nil
}
