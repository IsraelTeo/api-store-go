package repository

import (
	"github.com/IsraelTeo/api-store-go/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetByID(ID uint) (*model.Customer, error)
	GetAll() ([]model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer) (*model.Customer, error)
	Delete(ID uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetByID(ID uint) (*model.Customer, error) {
	customer := model.Customer{}
	if err := r.db.First(&customer, ID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) Create(customer *model.Customer) error {
	if err := r.db.Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Update(customer *model.Customer) (*model.Customer, error) {
	if err := r.db.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Customer{}, ID).Error; err != nil {
		return err
	}

	return nil
}
