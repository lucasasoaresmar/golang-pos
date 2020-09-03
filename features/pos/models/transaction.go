package models

import "time"

// Transaction model
type Transaction struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Sales     []Sale     `json:"Sales"`
}
