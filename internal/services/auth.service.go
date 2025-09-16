package services

import (
	"context"
	"errors"

	"github.com/chinsiang99/simple-bank-go-project/internal/repositories"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repositories *repositories.RepositoryManager) *AuthService {
	return &AuthService{repo: repositories.AuthRepository}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	storedHash, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Normally compare password hash here
	if storedHash != password {
		return "", errors.New("invalid password")
	}

	return "fake-jwt-token", nil
}
