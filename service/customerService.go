package service

import "github.com/parthvinchhi/bank-app/domain"

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return d.repo.ByID(id)
}

func NewCustomerService(respository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: respository,
	}
}
