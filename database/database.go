package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/golang_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database")

	DB = db
}
