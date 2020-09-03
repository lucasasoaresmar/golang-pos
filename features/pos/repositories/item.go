package repositories

import (
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

// ItemRepository repository
type ItemRepository struct{}

// GetAll items
func (irep *ItemRepository) GetAll() (items []models.Item, err error) {
	items = []models.Item{
		models.Item{
			Name:  "Pizza",
			Price: 23.23,
			Category: models.Category{
				Type: "Comida",
			},
		},
		models.Item{
			Name:  "Coca",
			Price: 4.50,
			Category: models.Category{
				Type: "Bebida",
			},
		},
	}

	return
}

// Create item
func (irep *ItemRepository) Create(item *models.Item) (err error) {
	return nil
}

// Update item
func (irep *ItemRepository) Update(item *models.Item) (err error) {
	return nil
}

// Delete item
func (irep *ItemRepository) Delete(id string) (err error) {
	return nil
}
