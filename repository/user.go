package repository

import (
	"github.com/IsraelTeo/api-store-go/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(ID uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(ID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ID uint) (*model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ID uint) error {
	if err := r.db.Delete(&model.User{}, ID).Error; err != nil {
		return err
	}
	return nil
}
