package services

import "fmt"

type InfoService interface {
	GetInfo() string
}

type Availability interface {
	CheckAvailability(requiredAmount int, availableAmount int) bool
}

type SupplierService interface {
	InfoService
	Availability
}

type Supplier struct {
	Number  string
	Contact string
	City    string
}

func (s Supplier) GetInfo() string {
	return fmt.Sprintf("Number: %s | Contact: %s | City: %s", s.Number, s.Contact, s.City)
}

func (s Supplier) CheckAvailability(requiredAmount int, availableAmount int) bool {
	return availableAmount >= requiredAmount
}
