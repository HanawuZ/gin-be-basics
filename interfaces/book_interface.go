package interfaces

import (
	"github.com/HanawuZ/gin-be-basics/models"
)

type BookRepository interface {
	ListBooks() (books []models.Book, err error)
	AddBook(book *models.Book) (err error)
	GetBookByISBN(isbn string) (book models.Book, err error)
	UpdateBookByISBN(isbn string, book *models.Book) (err error)
	DeleteBookByISBN(isbn string) (err error)
}
type BookUsecase interface {
	ListBooks() (books []models.Book, err error)
	AddBook(book *models.Book) (err error)
	GetBookByISBN(isbn string) (book models.Book, err error)
	UpdateBookByISBN(isbn string, book *models.Book) (err error)
	DeleteBookByISBN(isbn string) (err error)
}
