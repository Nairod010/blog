package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Init() error {
	dsn := "host=localhost database=blog user=postgres password=pass port=5432 sslmode=disable"
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil{
		fmt.Print("failed to connect to database: %v", err)
	
	}
	DB = db
	return nil
}
