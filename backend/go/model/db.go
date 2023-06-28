package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")
	
	db *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print(dsn + "failed to connect database")
	}

	db.AutoMigrate(&Todo{})
}