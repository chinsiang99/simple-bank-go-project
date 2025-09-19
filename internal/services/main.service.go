package services

import (
	db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"
)

type ServiceManager struct {
	// AuthService    IAuthService
	// UserService    IUserService
	AccountService IAccountService
	// TicketService *TicketService
}

func NewServiceManager(store *db.Store) *ServiceManager {
	return &ServiceManager{
		// AuthService:    NewAuthService(store),
		// UserService:    NewUserService(store),
		AccountService: NewAccountService(store),
		// TicketService: NewTicketService(repositories),
	}
}
