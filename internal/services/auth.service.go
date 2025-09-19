package services

import (
	db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"
)

type IAuthService interface {
	// AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error)
	// GetUserByID(id uint) (*models.User, error)
	// GetAllUsers() ([]models.User, error)
	// UpdateUser(user *models.User) error
	// DeleteUser(id uint) error
}

type AuthService struct {
	store db.IStore
}

func NewAuthService(store db.IStore) IAuthService {
	return &AuthService{store: store}
}

// func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
// 	// storedHash, err := s.store.FindUserByUsername(username)
// 	// if err != nil {
// 	// 	return "", errors.New("user not found")
// 	// }

// 	// // Normally compare password hash here
// 	// if storedHash != password {
// 	// 	return "", errors.New("invalid password")
// 	// }

// 	// return "fake-jwt-token", nil
// }
