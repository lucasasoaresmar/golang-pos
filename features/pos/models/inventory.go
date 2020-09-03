package models

import "time"

// Inventory model
type Inventory struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Item      Item       `json:"item"`
	Quantity  int        `json:"quantity"`
}
