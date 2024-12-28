package service

import (
	"log"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/validate"
)

type UserService interface {
	GetBydID(ID uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	RegisterUser(user *model.User) error
	Update(ID uint, user *model.User) (*model.User, error)
	Delete(ID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetBydID(ID uint) (*model.User, error) {
	user, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", ID, err)
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAll() ([]model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}

	if validate.VerifyListEmpty(users) {
		log.Println("Customers list is empty")
		return users, nil
	}

	return users, nil
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		log.Printf("Error fetching user with Email %s: %v", email, err)
		return nil, err
	}

	return user, nil
}

func (s *userService) RegisterUser(user *model.User) error {
	if err := s.repo.Create(user); err != nil {
		log.Printf("Error creating user: %+v, error: %v", user, err)
		return err
	}

	return nil
}

func (s *userService) Update(ID uint, user *model.User) (*model.User, error) {
	userFound, err := s.repo.GetByID(user.ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d for update: %v", ID, err)

		return nil, err
	}

	userFound.Email = user.Email
	userFound.Password = user.Password

	err = s.repo.Update(userFound)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", ID, err)
		return nil, err
	}

	return userFound, nil
}

func (s *userService) Delete(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		log.Printf("Error deleting customer with ID %d: %v", ID, err)
		return err
	}

	return nil
}
