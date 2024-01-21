package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name,omitempty"`
	Email    string `gorm:"size:255;not null;unique" json:"email,omitempty"`
	Password string `gorm:"size:255;not null" json:"password,omitempty"`
}
