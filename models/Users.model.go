package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name  *string `gorm:"size:255;unique;not null" json:"name"`
	Email *string `gorm:"size:255;unique;not null;" json:"email"`
	Tasks []Task  `json:"tasks"`
}
