package main

import "fmt"
import "gorm.io/driver/postgres"
import "gorm.io/gorm"


type Users struct {
	ID uint
	Name string
	Age uint8
	Email string
	password string
}

func main() {
	dsn := "host=localhost database=blog user=postgres password=pass port=5432 sslmode=disable"
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil{
		fmt.Print("failed to connect")
	}

	db.AutoMigrate(&Users{})

	db.Create(&Users{ID: 1 ,Name: "Dorian", Age: 27, Email: "dorian@mail.com", password: "nam"})
		
}
