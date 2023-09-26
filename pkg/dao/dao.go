package dao

import (
	"api-example/model"
	"fmt"
	"time"
	"unicode"
)

type CustomerStorage interface {
	GetCustomers() (model.CustomerList, error)
	AddCustomer(customer *model.Customer) error
}

type CustomerDao struct {
	customerList model.CustomerList
	lastID       int
}

func NewCustomerDao() (CustomerStorage, error) {
	dao := &CustomerDao{
		lastID:       1,
		customerList: make(model.CustomerList, 0),
	}
	dao.initDao()
	return dao, nil
}

func (dao *CustomerDao) initDao() {
	dao.AddCustomer(&model.Customer{
		ID:           1,
		Name:         "John Doe",
		CreationDate: time.Now(),
	})
}

func (dao *CustomerDao) GetCustomers() (model.CustomerList, error) {
	return dao.customerList, nil
}

func (dao *CustomerDao) AddCustomer(customer *model.Customer) error {
	err := dao.ValidateCustomerInput(customer)
	if err != nil {
		return err
	}

	customer.ID = dao.lastID
	customer.CreationDate = time.Now()
	dao.customerList = append(dao.customerList, *customer)
	dao.lastID++
	return nil
}

func (dao *CustomerDao) ValidateCustomerInput(customer *model.Customer) error {
	if customer == nil {
		return fmt.Errorf("customer is nil")
	}

	if customer.Name == "" {
		return fmt.Errorf("customer name is empty")
	}

	for _, char := range customer.Name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return fmt.Errorf("customer name contains invalid characters")
		}
	}

	for _, existingCustomer := range dao.customerList {
		if existingCustomer.Name == customer.Name {
			return fmt.Errorf("customer name is not unique")
		}
	}
	return nil
}
