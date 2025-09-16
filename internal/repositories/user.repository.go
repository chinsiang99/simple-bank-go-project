package repositories

import (
	"github.com/chinsiang99/simple-bank-go-project/internal/models"
	"gorm.io/gorm"
)

// Interface so we can mock this in unit tests
type IUserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// ✅ db.Create() will insert a new row into users.
// If success → user.ID gets auto-populated with the new primary key.
func (r *userRepository) Create(user *models.User) error {
	// INSERT INTO users (...) VALUES (...)
	return r.db.Create(user).Error
}

// ✅ db.First() finds the first record that matches the condition (id).
// If record not found → returns gorm.ErrRecordNotFound
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	// SELECT * FROM users WHERE id = ? LIMIT 1
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ✅ db.Find() loads all rows into a slice
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	// SELECT * FROM users
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ✅ db.Save() updates the record.
// If the record has a primary key, it updates.
// If not, it will insert a new one (careful: always ensure user.ID is set)
func (r *userRepository) Update(user *models.User) error {
	// UPDATE users SET ... WHERE id = ?
	return r.db.Save(user).Error
}

// ✅ db.Delete() deletes by primary key.
// Pass an empty struct (&models.User{}) to tell GORM which table.
func (r *userRepository) Delete(id uint) error {
	// DELETE FROM users WHERE id = ?
	return r.db.Delete(&models.User{}, id).Error
}

// ✅ db.Where("column = ?", value).First(&record)
// Safe from SQL injection (uses placeholders ?)
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	// SELECT * FROM users WHERE email = ? LIMIT 1
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
