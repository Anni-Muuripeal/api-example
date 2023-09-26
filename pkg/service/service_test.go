package service

import (
	"api-example/model"
	"testing"
	"time"
)

func Test_GetCustomers(t *testing.T) {

	cdao := &MocDao{}
	srv, err := NewCustomerService(cdao)
	if err != nil {
		t.Fatal("NewCustomerService error: %v", err)
	}

	ret, _ := srv.GetCustomers()
	for _, c := range ret {
		t.Logf("customer %d has name %s", c.ID, c.Name)
	}

}

func Test_AddCustomer(t *testing.T) {
	cdao := NewMocDao()
	srv, err := NewCustomerService(cdao)
	if err != nil {
		t.Fatal("NewCustomerService error: %v", err)
	}

	customer := &model.Customer{
		ID:           3,
		Name:         "Alice",
		CreationDate: time.Now(),
	}

	err = srv.AddCustomer(customer)
	if err != nil {
		t.Fatal("AddCustomer error: %v", err)
	}

	customers, err := srv.GetCustomers()
	if err != nil {
		t.Fatal("GetCustomers error: %v", err)
	}
	for _, c := range customers {
		if c.ID == customer.ID && c.Name == customer.Name {
			return
		}
	}

	t.Fatal("failed to add customer")
}

func NewMocDao() *MocDao {
	customer1 := model.Customer{ID: 1, Name: "John Doe", CreationDate: time.Now()}
	customer2 := model.Customer{ID: 2, Name: "Jane Doe", CreationDate: time.Now()}
	customers := model.CustomerList{customer1, customer2}

	return &MocDao{
		customers: customers,
	}
}

func (c *MocDao) GetCustomers() (model.CustomerList, error) {
	return c.customers, nil
}

func (c *MocDao) AddCustomer(customer *model.Customer) error {
	c.customers = append(c.customers, *customer)
	return nil
}
