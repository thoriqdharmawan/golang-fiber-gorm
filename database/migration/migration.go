package migration

import (
	"fmt"
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Database Migrated")
}
