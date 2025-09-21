package services

import (
	"context"

	// "github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"
	"github.com/chinsiang99/simple-bank-go-project/internal/dto"
	"golang.org/x/sync/errgroup"
)

type IAccountService interface {
	// AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error)
	// GetUserByID(id uint) (*models.User, error)
	// GetAllUsers() ([]models.User, error)
	// UpdateUser(user *models.User) error
	// DeleteUser(id uint) error
	CreateAccount(ctx context.Context, req dto.CreateAccountRequest) (db.Account, error)
	GetAccount(ctx context.Context, id int64) (db.Account, error)
	GetAccounts(ctx context.Context, req dto.GetAccountsRequest) (dto.GetAccountsResponse, error)
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

func (s *accountService) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	account, err := s.store.GetAccount(ctx, id)
	return account, err
}

func (s *accountService) GetAccounts(ctx context.Context, req dto.GetAccountsRequest) (dto.GetAccountsResponse, error) {
	// without count
	// account, err := s.store.ListAccounts(ctx, db.ListAccountsParams{
	// 	Limit:  req.Limit,
	// 	Offset: req.Offset,
	// })
	// return account, err
	var (
		accounts []db.Account
		total    int64
	)

	g, ctx := errgroup.WithContext(ctx)

	// Fetch accounts
	g.Go(func() error {
		var err error
		accounts, err = s.store.ListAccounts(ctx, db.ListAccountsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
		return err
	})

	// Fetch count
	g.Go(func() error {
		var err error
		total, err = s.store.GetAccountsCount(ctx)
		return err
	})

	// Wait for both
	if err := g.Wait(); err != nil {
		return dto.GetAccountsResponse{}, err
	}

	return dto.GetAccountsResponse{
		Accounts:   accounts,
		TotalCount: total,
	}, nil
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
