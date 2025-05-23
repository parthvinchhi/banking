package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:      "1001",
			Name:    "Parth",
			City:    "Sayla",
			Zipcode: "363430",
			DOB:     "2000-01-01",
			Status:  "1",
		},
		{
			ID:      "1002",
			Name:    "Rob",
			City:    "Sayla",
			Zipcode: "363430",
			DOB:     "2000-01-01",
			Status:  "1",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
