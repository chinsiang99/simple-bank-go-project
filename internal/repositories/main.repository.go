package repositories

import "github.com/chinsiang99/simple-bank-go-project/internal/database"

// import "myapp/internal/database"

type RepositoryManager struct {
	AuthRepository *AuthRepository
	UserRepository IUserRepository
	// TicketRepository *TicketRepository
}

func NewRepositoryManager(db *database.DB) *RepositoryManager {
	return &RepositoryManager{
		AuthRepository: NewAuthRepository(db),
		UserRepository: NewUserRepository(db.DB),
		// TicketRepository: NewTicketRepository(db),
	}
}
