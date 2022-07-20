package database

import (
	"bookstore/config"
	"bookstore/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASS"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error While Connecting Database.")
	}
	fmt.Println("Connection Opened to Database.")

	if err = DB.AutoMigrate(&models.Author{}, &models.Publisher{}, &models.Book{}); err != nil {
		log.Fatal("Error While Migrating.")
	}
	fmt.Println("Database Migrated.")
}
