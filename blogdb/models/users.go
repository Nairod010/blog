package models

import (
	"time"
	
	"github.com/satori/go.uuid"	
	"blogdb/db"
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
	WikiPages []WikiPage `gorm:"foreignKey:AuthorID"`
}


type BlogPost struct {
	gorm.Model
	PostID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Title string `gorm:"size:250"`
	Content	string `gorm:"type:text"`
	AuthorID uuid.UUID `gorm:"type:uuid"`
	Author User `gorm:"foreignKey:AuthorID"`
	Categories []Category `gorm:"many2many:post_categories"`
	Tags []Tag `gorm:"many2many:post_tags"`
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

type WikiPage struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);"`
	Content string `gorm:"type:text;"`
	AuthorID string `gorm:"type:uuid"`
	Author User	`gorm:"foreignKey:AuthorID"`
	Tag []Tag `gorm:"many2many:page_tags"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);"`
	Posts []BlogPost `gorm:"many2many:post_categories"`
}

type Tag  struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
	BlogPost []BlogPost `gorm:"many2many:post_tags"`
	WikiPage []WikiPage `gorm:"many2many:page_tags"`
}

func AutoMigrate() error {
	return db.DB.AutoMigrate(&User{},&BlogPost{}, &Comment{},&WikiPage{},&Category{},&Tag{})
}

