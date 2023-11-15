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
	Publisher   Publisher `json:"publisher" gorm:"references:Id"`

	Authors []Author `json:"authors" gorm:"many2many:book_authors;"`
}

// // Add a GORM callback to handle cascading delete
// func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
// 	err = tx.Model(b).Association("Authors").Clear()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
