package repositories

// import "github.com/chinsiang99/simple-bank-go-project/internal/database"

// type AuthRepository struct {
// 	db *database.DB
// }

// func NewAuthRepository(db *database.DB) *AuthRepository {
// 	return &AuthRepository{db: db}
// }

// func (r *AuthRepository) FindUserByUsername(username string) (string, error) {
// 	// Example: query from DB
// 	var hashedPassword string
// 	err := r.db.Raw("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword).Error
// 	return hashedPassword, err
// }
