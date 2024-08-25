package main

import (
	"backend/repository"
	"flag"
	"fmt"
	"github.com/AdamShannag/goprime/filter"
	"github.com/AdamShannag/goprime/filters"
	"github.com/AdamShannag/goprime/placeholder"
	"github.com/AdamShannag/goprime/prime"
	"github.com/AdamShannag/toolkit/v2"
	"log"
	"net/http"
	"time"
)

type config struct {
	port        string
	mysqlDsn    string
	postgresDsn string
}

type application struct {
	mysqlPrimeRepo    repository.PrimeRepository
	postgresPrimeRepo repository.PrimeRepository
	toolkit.Tools
}

func main() {
	conf := config{}

	flag.StringVar(&conf.port, "port", "8080", "server port")
	flag.StringVar(&conf.mysqlDsn, "mysql", "mysql:myverysecretpassword@tcp(localhost:3306)/prime?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "MYSQL DSN")
	flag.StringVar(&conf.postgresDsn, "postgres", "host=localhost port=5432 user=postgres password=myverysecretpassword dbname=prime sslmode=disable timezone=UTC connect_timeout=5", "POSTGRES DSN")
	flag.Parse()

	mysqlDB, err := newDB("mysql", conf.mysqlDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()

	postgresDB, err := newDB("postgres", conf.postgresDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	f := map[filter.MatchMode]filter.Filter{

		filter.STARTS_WITH:         filters.NewPatternMatchFilter("LIKE", filters.POST),
		filter.ENDS_WITH:           filters.NewPatternMatchFilter("LIKE", filters.PRE),
		filter.CONTAINS:            filters.NewPatternMatchFilter("LIKE", filters.AROUND),
		filter.NOT_CONTAINS:        filters.NewPatternMatchFilter("NOT LIKE", filters.AROUND),
		filter.EQUALS:              filters.ValueFilter("="),
		filter.NOT_EQUALS:          filters.ValueFilter("<>"),
		filter.GREATER_THAN:        filters.ValueFilter(">"),
		filter.GREATER_THAN_EQUALS: filters.ValueFilter(">="),
		filter.LESS_THAN:           filters.ValueFilter("<"),
		filter.LESS_THAN_EQUALS:    filters.ValueFilter("<="),
		filter.DATE_AFTER:          filters.DateAfterFilter(0),
		filter.DATE_BEFORE:         filters.DateBeforeFilter(0),
		filter.DATE_IS:             filters.DateIsFilter(0),
		filter.DATE_IS_NOT:         filters.DateIsNotFilter(0),
		filter.IN:                  filters.InFilter(0),
		filter.BETWEEN:             filters.BetweenFilter(0),
	}

	app := application{
		mysqlPrimeRepo: repository.NewMysqlPrimeRepository(mysqlDB,
			prime.NewWithFilters(placeholder.UnNumbered("?"), f)),
		postgresPrimeRepo: repository.NewPostgresPrimeRepository(postgresDB,
			prime.NewWithFilters(placeholder.Numbered("$"), f)),
		Tools: toolkit.Tools{},
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", conf.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("starting web application on port", conf.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
