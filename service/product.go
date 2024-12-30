package service

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/validate"
)

type ProductService interface {
	GetByID(ID uint) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Create(product *model.Product) error
	Update(ID uint, product *model.Product) (*model.Product, error)
	Delete(ID uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetByID(ID uint) (*model.Product, error) {
	product, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching product with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to fetch product with ID %d: %w", ID, err)
	}

	return product, nil
}

func (s *productService) GetAll() ([]model.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return nil, fmt.Errorf("service: failed to fetch products: %w", err)
	}

	if validate.VerifyListEmpty(products) {
		log.Println("Products list is empty")
		return products, nil
	}

	return products, nil
}

func (s *productService) Create(product *model.Product) error {
	if err := s.repo.Create(product); err != nil {
		log.Printf("Error creating product: %+v, error: %v", product, err)
		return fmt.Errorf("service: failed to create product: %w", err)
	}

	return nil
}

func (s *productService) Update(ID uint, product *model.Product) (*model.Product, error) {
	productFound, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d for update: %v", ID, err)
		return nil, fmt.Errorf("service: failed to fetch product with ID %d for update: %w", ID, err)
	}

	productFound.Mark = product.Mark
	productFound.Name = product.Name
	productFound.Price = product.Price

	updatedProduct, err := s.repo.Update(productFound)
	if err != nil {
		log.Printf("Error updating product with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to update product with ID %d: %w", ID, err)
	}

	return updatedProduct, nil
}

func (s *productService) Delete(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		log.Printf("Error deleting product with ID %d: %v", ID, err)
		return fmt.Errorf("service: failed to delete product with ID %d: %w", ID, err)
	}

	return nil
}

//Obtener todos los productos cuya cantidad_disponible sea menor a 5 -> ruta: /falta_stock

// Obtener la lista de productos de una determinada venta -> ruta:/{codigo de venta}

//requerimiento: actualizar el stock de un producto (descontar) realizar una venta, y tampoco controlar si cuenta con la cantidad disponible para vender;
