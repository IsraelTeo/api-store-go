package repository

import (
	"github.com/IsraelTeo/api-store-go/model"
	"gorm.io/gorm"
)

type SaleRepository interface {
	GetByID(ID uint) (*model.Sale, error)
	GetAll() ([]model.Sale, error)
	Create(sale *model.Sale) error
	Update(sale *model.Sale) (*model.Sale, error)
	Delete(ID uint) error
}

// Struct de sale repository, simula ser una clase, tiene como atributo o campo un puntero a la estructura DB de gorm
type saleRepository struct {
	db *gorm.DB
}

// Método constructor que sirve para crear una instancia del struct repository
// recibe de parametro un puntero de gorm, ya que la instancia que se creará del struct tiene un puntero de gorm
// Retorna cualquier tipo que implemente la interfaz repository
func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &saleRepository{db: db}
}

// Creamos los métodos del struct, este struct implementa a la interfaz y por eso el constructor retorna un struct que implementa dicha interfaz
func (r *saleRepository) GetByID(ID uint) (*model.Sale, error) {
	sale := model.Sale{}
	if err := r.db.First(&sale, ID).Error; err != nil {
		return nil, err
	}

	return &sale, nil
}

func (r *saleRepository) GetAll() ([]model.Sale, error) {
	var sales []model.Sale
	if err := r.db.Find(&sales).Error; err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *saleRepository) Create(sale *model.Sale) error {
	if err := r.db.Create(sale).Error; err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) Update(sale *model.Sale) (*model.Sale, error) {
	if err := r.db.Save(sale).Error; err != nil {
		return nil, err
	}

	return sale, nil
}

func (r *saleRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.Sale{}, ID).Error; err != nil {
		return err
	}

	return nil
}
