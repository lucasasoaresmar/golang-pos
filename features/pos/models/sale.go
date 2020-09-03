package models

import "time"

// Sale model
type Sale struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Item      Item       `json:"Item"`
	Price     float64    `json:"Price"`
	Quantity  int        `json:"Quantity"`
}
