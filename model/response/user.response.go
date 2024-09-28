package response

import (
	"golang-fiber-gorm/model/entity"
	"time"

	"gorm.io/gorm"
)

type UserResponseGetAllUsers struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"column:password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Posts     []entity.Post  `json:"-" gorm:"column:posts"`
}

type UserResponseById struct {
	ID        uint              `json:"id" gorm:"primarykey"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Address   string            `json:"address"`
	Phone     string            `json:"phone"`
	Posts     []entity.Post     `json:"posts"`
	Language  []entity.Language `json:"languages"`
}

type UserResponseGetPosts struct {
	ID    uint          `json:"id" gorm:"primarykey"`
	Name  string        `json:"name"`
	Posts []entity.Post `json:"posts" gorm:"foreignKey:UserID"`
}
