package repository

const countSql = `SELECT COUNT(*) AS total_count
            FROM customer
            LEFT JOIN country ON customer.country_id = country.country_id
            LEFT JOIN representative ON customer.representative_id = representative.representative_id`

const selectCustomersSql = `
SELECT
    customer.id,
    customer.name,
    customer.company,
    customer.date,
    customer.status,
    customer.verified,
    customer.activity,
    customer.balance,
    country.name AS country_name,
    country.code AS country_code,
    representative.representative_id,
    representative.name AS representative_name,
    representative.image AS representative_image
FROM
    customer
        LEFT JOIN
    country ON customer.country_id = country.country_id
        LEFT JOIN
    representative ON customer.representative_id = representative.representative_id`

const limitSql = " LIMIT %s OFFSET %s"

const selectRepresentativesSql = `SELECT * FROM representative`
