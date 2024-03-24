package models

import (
		
	"github.com/google/uuid"	
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

func PopulateTestData(db *gorm.DB) {	
	
	//create users
	user1 := User{Username: "user1", ID: uuid.New()}
	user2 := User{Username: "user2", ID: uuid.New()}	
	db.Create(&user1)
	db.Create(&user2)
	
	//create Categories
	category1 := Category{Name :"Category1"}
	category2 := Category{Name :"Category2"}
	db.Create(&category1)
	db.Create(&category2)

	//create Tags
	tag1 := Tag{Name: "Tag1"}
	tag2 := Tag{Name: "Tag2"}
	db.Create(&tag1)
	db.Create(&tag2)

	//create posts
	post1 := BlogPost{Title: "Post 1", Content: "Content for post 1", AuthorID: user1.ID, Tags: []*Tag{&tag1}}
	post2 := BlogPost{Title: "Post 2", Content: "Content for post 2", AuthorID: user2.ID, Tags: []*Tag{&tag1, &tag2}}
	db.Create(&post1)	
	db.Create(&post2)

	//create wiki page and associate with categories
	page1 := WikiPage{Title: "Page 1", Content: "Content for page 1", AuthorID: user1.ID, Categories: []*Category{&category1}}
	page2 := WikiPage{Title: "Page 2", Content: "Content for page 2", AuthorID: user2.ID, Categories: []*Category{&category1, &category2}}
	db.Create(&page1)
	db.Create(&page2)
}

