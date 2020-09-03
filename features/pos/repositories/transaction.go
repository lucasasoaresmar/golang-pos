package repositories

import (
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

// TransactionRepository repository
type TransactionRepository struct{}

// GetAll categories
func (ur *TransactionRepository) GetAll() (categories []models.Transaction, err error) {
	categories = []models.Transaction{
		models.Transaction{
			Sales: []models.Sale{
				models.Sale{
					Item: models.Item{
						Name:  "Pizza",
						Price: 23.23,
						Category: models.Category{
							Type: "Comida",
						},
					},
					Price:    20.00,
					Quantity: 1,
				},
				models.Sale{
					Item: models.Item{
						Name:  "Coca",
						Price: 4.50,
						Category: models.Category{
							Type: "Bebida",
						},
					},
					Price:    50.00,
					Quantity: 1,
				},
			},
		},
	}

	return
}

// Create transaction
func (ur *TransactionRepository) Create(transaction *models.Transaction) (err error) {
	return nil
}

// Update transaction
func (ur *TransactionRepository) Update(transaction *models.Transaction) (err error) {
	return nil
}

// Delete transaction
func (ur *TransactionRepository) Delete(id string) (err error) {
	return nil
}
