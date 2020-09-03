package models

import "time"

// User model
type User struct {
	ID        uint       `json:"ID" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Roles     []string   `json:"roles"`
}

// IsValidRole validate a user role
func (u *User) IsValidRole(roleToValidate string) (ok bool) {
	for _, role := range u.Roles {
		if role == roleToValidate {
			return true
		}
	}
	return false
}
