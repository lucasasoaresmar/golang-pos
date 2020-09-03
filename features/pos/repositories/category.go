package repositories

import (
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

// CategoryRepository repository
type CategoryRepository struct{}

// GetAll categories
func (ur *CategoryRepository) GetAll() (categories []models.Category, err error) {
	categories = []models.Category{
		models.Category{
			Type: "Comida",
		},
		models.Category{
			Type: "Bebida",
		},
	}

	return
}

// Create category
func (ur *CategoryRepository) Create(category *models.Category) (err error) {
	return nil
}

// Update category
func (ur *CategoryRepository) Update(category *models.Category) (err error) {
	return nil
}

// Delete category
func (ur *CategoryRepository) Delete(id string) (err error) {
	return nil
}
