package service

import (
	"log"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/validate"
)

type SaleService interface {
	GetByID(ID uint) (*model.Sale, error)
	GetAll() ([]model.Sale, error)
	Create(sale *model.Sale) error
	Update(ID uint, sale *model.Sale) (*model.Sale, error)
	Delete(ID uint) error
}

type saleService struct {
	repo repository.SaleRepository
}

func NewSaleRepository(repo repository.SaleRepository) SaleService {
	return &saleService{repo: repo}
}

func (s *saleService) GetByID(ID uint) (*model.Sale, error) {
	sale, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching sale with ID %d: %v", ID, err)
		return nil, err
	}

	return sale, nil
}

func (s *saleService) GetAll() ([]model.Sale, error) {
	sales, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching sales: %v", err)
		return nil, err
	}

	if validate.VerifyListEmpty(sales) {
		log.Println("Sales list is empty")
		return sales, nil
	}

	return sales, nil
}

func (s *saleService) Create(sale *model.Sale) error {
	amount, err := calculateAmount(sale.Products)
	if err != nil {
		log.Printf("Error when adding total amount, error: %v", err)
		return err
	}

	sale.TotalAmount = amount

	if err := s.repo.Create(sale); err != nil {
		log.Printf("Error creating sale: %+v, error: %v", sale, err)
		return err
	}

	return nil
}

func (s *saleService) Update(ID uint, sale *model.Sale) (*model.Sale, error) {
	saleFound, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d for update: %v", ID, err)
		return nil, err
	}

	amount, err := calculateAmount(sale.Products)
	if err != nil {
		log.Printf("Error to updating total amount, error: %v", err)
		return nil, err
	}

	saleFound.Customer = sale.Customer
	saleFound.Products = sale.Products
	saleFound.TotalAmount = amount

	saleUpdated, err := s.repo.Update(saleFound)
	if err != nil {
		log.Printf("Error updating sale with ID %d: %v", ID, err)
		return nil, err
	}

	return saleUpdated, nil
}

func (s *saleService) Delete(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		log.Printf("Error deleting sale with ID %d: %v", ID, err)
		return err
	}

	return nil
}

func calculateAmount(products []model.Product) (float64, error) {
	if validate.VerifyListEmpty(products) {
		log.Println("The product list is empty, there is no total amount.")
		return 0, nil
	}

	var totalAmount float64

	for _, p := range products {
		totalAmount += p.Price
	}

	return totalAmount, nil
}
