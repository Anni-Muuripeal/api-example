package dao

import (
	"api-example/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerList_GetCustomers(t *testing.T) {
	customer1 := model.Customer{ID: 1, Name: "John Doe", CreationDate: time.Now()}
	customer2 := model.Customer{ID: 2, Name: "Jane Doe", CreationDate: time.Now()}

	customerDao, err := NewCustomerDao()
	if err != nil {
		t.Fatal(err)
	}

	err = customerDao.AddCustomer(&customer1)
	if err != nil {
		t.Fatal(err)
	}

	err = customerDao.AddCustomer(&customer2)
	if err != nil {
		t.Fatal(err)
	}

	expectedCustomers := model.CustomerList{customer1, customer2}
	actualCustomers, err := customerDao.GetCustomers()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedCustomers, actualCustomers, "customer lists are not equal")
}
