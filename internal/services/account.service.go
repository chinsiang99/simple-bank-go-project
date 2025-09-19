package services

import (
	"context"

	// "github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"
	"github.com/chinsiang99/simple-bank-go-project/internal/dto"
)

type IAccountService interface {
	// AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error)
	// GetUserByID(id uint) (*models.User, error)
	// GetAllUsers() ([]models.User, error)
	// UpdateUser(user *models.User) error
	// DeleteUser(id uint) error
	CreateAccount(ctx context.Context, req dto.CreateAccountRequest) (db.Account, error)
}

type accountService struct {
	store db.IStore
}

func NewAccountService(store db.IStore) IAccountService {
	return &accountService{store: store}
}

func (s *accountService) CreateAccount(ctx context.Context, req dto.CreateAccountRequest) (db.Account, error) {
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := s.store.CreateAccount(ctx, arg)
	return account, err
}

// func (s *userService) GetUserByID(id uint) (*models.User, error) {
// 	return s.store.FindByID(id)
// }

// func (s *userService) GetAllUsers() ([]models.User, error) {
// 	return s.store.FindAll()
// }

// func (s *userService) UpdateUser(user *models.User) error {
// 	return s.store.Update(user)
// }

// func (s *userService) DeleteUser(id uint) error {
// 	return s.store.Delete(id)
// }
