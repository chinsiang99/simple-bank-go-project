package handlers

import "github.com/chinsiang99/simple-bank-go-project/internal/services"

type HandlerManager struct {
	// AuthHandler    *AuthHandler
	// UserHandler    IUserHandler
	AccountHandler IAccountHandler
	// TicketHandler *TicketHandler
}

func NewHandlerManager(services *services.ServiceManager) *HandlerManager {
	return &HandlerManager{
		// AuthHandler:    NewAuthHandler(services),
		// UserHandler:    NewUserHandler(services.UserService),
		AccountHandler: NewAccountHandler(services.AccountService),
		// TicketHandler: NewTicketHandler(service),
	}
}
