package domain

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CustomerRepositoryDB struct {
	client *gorm.DB
}

func NewCustomerRepositoryDB() *CustomerRepositoryDB {
	dsn := "root:123@tcp(localhost:3306)/banking?parseTime=false"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	return &CustomerRepositoryDB{client: db}
}

func (r CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	var customers []Customer
	var result *gorm.DB

	if status == "" {
		result = r.client.Find(&customers)
	} else {
		result = r.client.Where("status = ?", status).Find(&customers)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r CustomerRepositoryDB) ByID(id string) (*Customer, error) {
	var customer Customer
	result := r.client.First(&customer, "customer_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}
