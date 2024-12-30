package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/IsraelTeo/api-store-go/dto"
	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/IsraelTeo/api-store-go/validate"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetBydID(ID uint) (*dto.UserResponse, error)
	GetByEmail(email string) (*dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	RegisterUser(user *model.User) error
	Update(ID uint, user *model.User) (*dto.UserResponse, error)
	Delete(ID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetBydID(ID uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to fetch user with ID %d: %w", ID, err)
	}

	userDto := util.ToUserDTO(user)

	return userDto, nil
}

func (s *userService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, fmt.Errorf("service: failed to fetch users: %w", err)
	}

	if validate.VerifyListEmpty(users) {
		log.Println("Customers list is empty")
		return []dto.UserResponse{}, nil
	}

	usersDto := util.ToListUserDTO(users)
	return usersDto, nil
}

func (s *userService) GetByEmail(email string) (*dto.UserResponse, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		log.Printf("Error fetching user with Email %s: %v", email, err)
		return nil, fmt.Errorf("service: failed to fetch user with email %s: %w", email, err)
	}

	userDto := util.ToUserDTO(user)

	return userDto, nil
}

func (s *userService) RegisterUser(user *model.User) error {
	userData, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		log.Printf("Error fetching user by email: %s, error: %v", user.Email, err)
		return fmt.Errorf("service: failed to check user by email: %w", err)
	}

	if userData != nil {
		log.Printf("User with email %s already exists", user.Email)
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", user.Email, err)
		return fmt.Errorf("service: failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	if err := s.repo.Create(user); err != nil {
		log.Printf("Error creating user: %+v, error: %v", user, err)
		return fmt.Errorf("service: failed to create user: %w", err)
	}

	return nil
}

func (s *userService) Update(ID uint, user *model.User) (*dto.UserResponse, error) {
	userFound, err := s.repo.GetByID(ID)
	if err != nil {
		log.Printf("Error fetching user with ID %d for update: %v", ID, err)
		return nil, err
	}

	if user.Password != "" {
		log.Printf("Password update detected for user with ID %d. Hashing new password.", ID)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password for user with ID %d: %v", ID, err)
			return nil, fmt.Errorf("service: Error hashing password for user with ID %d %w", ID, err)
		}

		userFound.Password = string(hashedPassword)
	}

	userFound.Email = user.Email

	userUpdated, err := s.repo.Update(userFound)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", ID, err)
		return nil, fmt.Errorf("service: failed to updating user with ID %d %w", ID, err)
	}

	userDto := util.ToUserDTO(userUpdated)

	return userDto, nil
}

func (s *userService) Delete(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		log.Printf("Error deleting customer with ID %d: %v", ID, err)
		return err
	}

	return nil
}
