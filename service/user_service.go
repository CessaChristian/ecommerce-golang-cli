package service

import (
	"dealer_golang/models"
	"dealer_golang/repositories"
	"dealer_golang/utils"
	"errors"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// CONSTRUCTOR
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// REGISTER
func (s *UserService) Register(name, email, password string) error {
	// Validasi email
	if !utils.IsValidEmail(email) {
		return errors.New("invalid email format")
	}

	// Validasi password minimal 6 karakter
	if len(password) < 3 {
		return errors.New("password must be at least 6 characters")
	}

	// Cek apakah email sudah terdafta
	existing, _ := s.Repo.GetByEmail(email)
	if existing.ID != 0 {
		return errors.New("email already exists")
	}

	// Hash password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// Buat user baru
	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     "customer",
	}

	return s.Repo.Create(user)
}

// LOGIN
func (s *UserService) Login(email, password string) (models.User, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil || user.ID == 0 {
		return models.User{}, errors.New("user no found")
	}

	// cek password
	if !utils.CheckPassword(password, user.Password) {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAllUsers()
}

