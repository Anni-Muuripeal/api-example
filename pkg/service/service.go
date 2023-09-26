package service

import (
	"api-example/model"
	"api-example/pkg/dao"
)

type CustomerManager interface {
	GetCustomers() (model.CustomerList, error)
	AddCustomer(customer *model.Customer) error
}

type CustomerService struct {
	customerDao dao.CustomerStorage
}

func NewCustomerService(customerDao dao.CustomerStorage) (CustomerManager, error) {
	srv := &CustomerService{
		customerDao: customerDao,
	}
	return srv, nil
}

func (c *CustomerService) GetCustomers() (model.CustomerList, error) {
	return c.customerDao.GetCustomers()
}

func (c *CustomerService) AddCustomer(customer *model.Customer) error {
	return c.customerDao.AddCustomer(customer)
}
