package repositories

import "github.com/lucasasoaresmar/golang-pos/features/auth/models"

// UserRepository repository
type UserRepository struct{}

// Create User
func (ur *UserRepository) Create(user *models.User) error {
	return nil
}

// GetByEmailAndPassword an User
func (ur *UserRepository) GetByEmailAndPassword(email string, password string) (user models.User, err error) {
	user = models.User{
		Email:    email,
		Password: password,
		Roles:    []string{"user", "admin"},
	}

	return
}
