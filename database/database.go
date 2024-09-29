package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database")

	DB = db
}
