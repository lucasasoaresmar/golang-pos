package models

import "time"

// Item model
type Item struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"Name"`
	Price     float32    `json:"Price"`
	Category  Category   `json:"Category"`
}
