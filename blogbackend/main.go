package main

import (
	"fmt"
	
	"blogdb/models"
	"blogdb/db"
	
)	


func main() {
	if err:= db.Init(); err != nil{
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}	
	
	if err := models.AutoMigrate(); err != nil {
		fmt.Printf("Error automigrateing models: %v/n", err)
		return
	}
	
	models.AutoMigrate()
	
	dbInstance := db.GetDB()

	models.PopulateTestData(dbInstance)


	fmt.Println("Database and models initialized successfully")
}


