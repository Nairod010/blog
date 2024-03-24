package models

import (
		
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
//	Categories []Category `gorm:"many2many:post_categories"`
	Tags []*Tag `gorm:"many2many:post_tags;foreignKey:ID;joinForeignKey:TagID"`
}

type Comment struct {
	gorm.Model
	PostID string `gorm:"type:uuid"`
	Post BlogPost	`gorm:"foreignKey:PostID"`
	AuthorID string `gorm:"type:uuid"`
	Author User	`gorm:"foreignKey:AuthorID"`
	CommentText string `gorm:"type:text"`	
}

type WikiPage struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);"`
	Content string `gorm:"type:text;"`
	AuthorID uuid.UUID `gorm:"type:uuid"`
	Author User	`gorm:"foreignKey:AuthorID"`
	Tag []Tag `gorm:"many2many:page_tags"`
	Categories []*Category `gorm:"many2many:wiki_page_categories"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);"`
	Posts []*WikiPage `gorm:"many2many:wiki_page_categories"`
}

type Tag  struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
	BlogPost []*BlogPost `gorm:"many2many:post_tags;foreignKey:ID;joinForeignKey:PostID"`
	WikiPage []*WikiPage `gorm:"many2many:page_tags"`
}

type PostTag struct {
	PostID uuid.UUID `gorm:"type:uuid"`
	TagID uint `gorm:"type:bigint"`
}

func (PostTag) TableName() string {
	return "post_tags"
}


func AutoMigrate() error {
	return db.DB.AutoMigrate(&User{},&BlogPost{}, &Comment{},&WikiPage{},&Tag{},&PostTag{})
}

