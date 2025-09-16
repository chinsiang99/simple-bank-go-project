package services

import "github.com/chinsiang99/simple-bank-go-project/internal/repositories"

type ServiceManager struct {
	AuthService *AuthService
	UserService IUserService
	// TicketService *TicketService
}

func NewServiceManager(repositories *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		AuthService: NewAuthService(repositories),
		UserService: NewUserService(repositories.UserRepository),
		// TicketService: NewTicketService(repositories),
	}
}
