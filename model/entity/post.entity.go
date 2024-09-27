package entity

type Post struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	UserID uint
}
