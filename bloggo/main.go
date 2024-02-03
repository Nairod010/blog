package main

	import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq"
)

func main() {
	connStr := "user=dorian dbname=test password=pass port = 5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err =db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

}
