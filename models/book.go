package models

import "time"

type Book struct {
	/*
	   JSON format
	   {
	          "isbn": "",
	          "title": "",
	          "genre": "",
	          "publicationYear": "",
	          "copiesAvailable": "",
	          "price": "",
	          "authorId": "",
	          "publisherId": ""
	   }
	*/
	Isbn  string `json:"isbn" gorm:"primaryKey"`
	Title string `json:"title"`
	Genre string `json:"genre"`

	PublicationYear time.Time `json:"publicationYear"`
	CopiesAvailable uint      `json:"copiesAvailable"`
	Price           float32   `json:"price"`

	PublisherId *uint     `json:"publisherId"`
	Publisher   Publisher `gorm:"references:Id"`

	Authors []Author `gorm:"many2many:book_authors;"`
}
