package service

import (
	"log"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/validate"
)

type CustomerService interface {
	GetByID(ID uint) (*model.Customer, error)
	GetAll() ([]model.Customer, error)
	Create(customer *model.Customer) error
	Update(ID uint, customer *model.Customer) (*model.Customer, error)
	Delete(ID uint) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetByID(ID uint) (*model.Customer, error) {
	customer, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching customer with ID %d: %v", ID, err)
		return nil, err
	}

	return customer, nil
}

func (s *customerService) GetAll() ([]model.Customer, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching customers: %v", err)
		return nil, err
	}

	if validate.VerifyListEmpty(customers) {
		log.Println("Customers list is empty")
		return customers, nil
	}

	return customers, nil
}

func (s *customerService) Create(customer *model.Customer) error {
	if err := s.repo.Create(customer); err != nil {
		log.Printf("Error creating customer: %+v, error: %v", customer, err)
		return err
	}

	return nil
}

func (s *customerService) Update(ID uint, customer *model.Customer) (*model.Customer, error) {
	customerFound, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching customer with ID %d for update: %v", ID, err)
		return nil, err
	}

	customerFound.DNI = customer.DNI
	customerFound.Name = customer.Name
	customerFound.LastName = customer.LastName

	updatedCustomer, err := s.repo.Update(customerFound)
	if err != nil {
		log.Printf("Error updating customer with ID %d: %v", ID, err)
		return nil, err
	}

	return updatedCustomer, nil
}

func (s *customerService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Printf("Error deleting customer with ID %d: %v", id, err)
		return err
	}

	return nil
}
