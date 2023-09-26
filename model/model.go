package model

import "time"

type Customer struct {
	ID           int
	Name         string
	CreationDate time.Time
}

type CustomerList []Customer
