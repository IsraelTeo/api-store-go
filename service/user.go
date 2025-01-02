package service

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-store-go/auth"
	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/util"
)

type UserService interface {
	GetByID(ID uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	RegisterUser(user *model.RegisterUserPayload) error
	Update(ID uint, user model.RegisterUserPayload) (*model.User, error)
	Delete(ID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetByID(ID uint) (*model.User, error) {
	user, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to fetch user with ID %d: %w", ID, err)
	}

	return user, nil
}

func (s *userService) GetAll() ([]model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, fmt.Errorf("service: failed to fetch users: %w", err)
	}

	if util.VerifyListEmpty(users) {
		log.Println("Customers list is empty")
		return []model.User{}, nil
	}

	return users, nil
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		log.Printf("Error fetching user with Email %s: %v", email, err)
		return nil, fmt.Errorf("service: failed to fetch user with email %s: %w", email, err)
	}

	return user, nil
}

func (s *userService) RegisterUser(user *model.RegisterUserPayload) error {
	isExists, err := util.CheckEmailExists("email", user.Email, &model.User{})
	if err != nil {
		log.Printf("Error checking if email exists %s: %v", user.Email, err)
		return fmt.Errorf("service: error checking email existence: %w", err)
	}

	if isExists {
		log.Printf("Email already exists: %s", user.Email)
		return fmt.Errorf("service: email already exists")
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", user.Email, err)
		return fmt.Errorf("service: failed to hash password: %w", err)
	}

	user.Password = hashedPassword
	userSave := util.ToUser(user)

	if err := s.repo.Create(userSave); err != nil {
		log.Printf("Error creating user: %+v, error: %v", user, err)
		return fmt.Errorf("service: failed to create user: %w", err)
	}

	return nil
}

func (s *userService) Update(ID uint, user model.RegisterUserPayload) (*model.User, error) {
	userFound, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d for update: %v", ID, err)
		return nil, fmt.Errorf("user with ID %d not found", ID)
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", user.Email, err)
		return nil, fmt.Errorf("service: failed to hash password: %w", err)
	}

	// Actualizar campos
	userFound.FirstName = user.FirstName
	userFound.LastName = user.LastName
	userFound.Email = user.Email
	userFound.Password = hashedPassword
	userFound.IsAdmin = user.IsAdmin

	// Llamada al repositorio con el modelo correcto
	userUpdated, err := s.repo.Update(userFound)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to update user with ID %d: %w", ID, err)
	}

	return userUpdated, nil
}

func (s *userService) Delete(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		log.Printf("Error deleting customer with ID %d: %v", ID, err)
		return fmt.Errorf("service: failed to delete user with ID %d: %w", ID, err)
	}

	return nil
}
