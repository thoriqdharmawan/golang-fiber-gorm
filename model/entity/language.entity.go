package entity

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Users     []*User        `json:"-" gorm:"many2many:user_languages;"`
}