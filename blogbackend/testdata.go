package main

import (
//	"blogdb/db"
	"blogdb/models"	
	"gorm.io/gorm"
)


func PopulateTestdata(db *gorm.DB) {	
	
	//create users
	user1 := models.User{Username: "user1"}
	user2 := models.User{Username: "user2"}	
	db.Create(&user1)
	db.Create(&user2)
	
	//create Categories
	category1 := models.Category{Name :"Category1"}
	category2 := models.Category{Name :"Category2"}
	db.Create(&category1)
	db.Create(&category2)

	//create Tags
	tag1 := models.Tag{Name: "Tag1"}
	tag2 := models.Tag{Name: "Tag2"}
	db.Create(&tag1)
	db.Create(&tag2)

	//create posts
	post1 := models.BlogPost{Title: "Post 1", Content: "Content for post 1", AuthorID: user1.ID, Tags: []*models.Tag{&tag1}}
	post2 := models.BlogPost{Title: "Post 2", Content: "Content for post 2", AuthorID: user2.ID, Tags: []*models.Tag{&tag1, &tag2}}
	db.Create(&post1)	
	db.Create(&post2)

	//create wiki page and associate with categories
	page1 := models.WikiPage{Title: "Page 1", Content: "Content for page 1", AuthorID: user1.ID, Categories: []*models.Category{&category1}}
	page2 := models.WikiPage{Title: "Page 2", Content: "Content for page 2", AuthorID: user2.ID, Categories: []*models.Category{&category1, &category2}}
	db.Create(&page1)
	db.Create(&page2)
}
	





