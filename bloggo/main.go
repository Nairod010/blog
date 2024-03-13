package main

import (
	"fmt"
	"time"
	

	"github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Username string `gorm:"size:250;unique"`
	Email string  `gorm:"size:250;unique"`
	Password string	`gorm:"size:250"`
	BlogPosts []BlogPost `gorm:"foreignKey:AuthorID"`
	Comments []Comment `gorm:"foreignKey:AuthorID"`
	//WikiPages []WikiPage `gorm:"foreignKey:AuthorID"`
}


type BlogPost struct {
	gorm.Model
	PostID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Title string `gorm:"size:250"`
	Content	string `gorm:"type:text"`
	AuthorID uuid.UUID `gorm:"type:uuid"`
	Author User `gorm:"foreignKey:AuthorID"`
	//Categories []Category `gorm:"many2many:post_categories"`
	//Tags []Tag `gorm:"many2many:post_tags"`
}

type Comment struct {
	gorm.Model
	PostID string `gorm:"type:uuid"`
	Post BlogPost	`gorm:"foreignKey:PostID"`
	AuthorID string `gorm:"type:uuid"`
	Author User	`gorm:"foreignKey:AuthorID"`
	CommentText string `gorm:"type:text"`
	CommentDate time.Time `gorm::"autoCreateTime;type:timestamp with time zone"`
}


func main() {
	dsn := "host=localhost database=blog user=postgres password=pass port=5432 sslmode=disable"
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil{
		fmt.Print("failed to connect")
	}

	db.AutoMigrate(
					&User{},
				    &BlogPost{},
					&Comment{},
			   )
	
	
}
