package models

import "time"

// Category model
type Category struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Type      string     `json:"Type"`
}
