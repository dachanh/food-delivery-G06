package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Activate() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return err
	}
	dsn := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	sqlDB.Ping()
	return nil
}

func main() {
	if err := Activate(); err != nil {
		log.Fatal(err)
	}
}
