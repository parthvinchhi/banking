package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

// func Connect() (*CustomerRepositoryDB, error) {
// 	dsn := "root:123@tcp(localhost:3306)/banking?parseTime=true"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CustomerRepositoryDB{client: db}, nil
// }

func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	client, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(3 * time.Minute)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	var rows *sql.Rows
	// var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = client.Query(findAllSql, status)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var cust Customer
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.City, &cust.Zipcode, &cust.DOB, &cust.Status); err != nil {
			log.Println(err)
			return nil, err
		}

		customers = append(customers, cust)
	}

	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(3 * time.Minute)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{
		client: client,
	}
}

func (db CustomerRepositoryDB) ByID(id string) (*Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := db.client.QueryRow(customerSql, id)
	var cust Customer

	if err := row.Scan(&cust.ID, &cust.Name, &cust.City, &cust.Zipcode, &cust.DOB, &cust.Status); err != nil {
		return nil, err
	}

	return &cust, nil
}
