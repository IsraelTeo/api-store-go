package repository

import (
	"github.com/IsraelTeo/api-store-go/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetByID(ID uint) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) (*model.Product, error)
	Delete(ID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetByID(ID uint) (*model.Product, error) {
	product := model.Product{}
	if err := r.db.First(&product, ID).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Update(product *model.Product) (*model.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Product{}, ID).Error; err != nil {
		return err
	}

	return nil
}
