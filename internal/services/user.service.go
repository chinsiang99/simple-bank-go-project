package services

import (
	"github.com/chinsiang99/simple-bank-go-project/internal/models"
	"github.com/chinsiang99/simple-bank-go-project/internal/repositories"
)

type IUserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
