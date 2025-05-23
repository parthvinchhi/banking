package domain

type Customer struct {
	ID      string `gorm:"column:customer_id;primaryKey"`
	Name    string
	City    string
	Zipcode string
	DOB     string `gorm:"column:date_of_birth"`
	Status  string
}

type CustomerRepository interface {
	// string == 1 or string == 0 or string == ""
	FindAll(string) ([]Customer, error)
	ByID(string) (*Customer, error)
}
