package usecases

import (
	"github.com/HanawuZ/gin-be-basics/dto"
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
)

type BookUsecase struct {
	BookRepository interfaces.BookRepository
}

func NewBookUsecase(bookRepo interfaces.BookRepository) interfaces.BookUsecase {
	return &BookUsecase{
		BookRepository: bookRepo,
	}
}

func (bookUsecase *BookUsecase) ListBooks() (books []models.Book, err error) {
	books, err = bookUsecase.BookRepository.ListBooks()
	return books, err
}

func (BookUsecase *BookUsecase) AddBook(bookReuqest *dto.BookRequest) (err error) {
	err = BookUsecase.BookRepository.AddBook(bookReuqest)
	return err
}

func (BookUsecase *BookUsecase) GetBookByISBN(isbn string) (book models.Book, err error) {
	book, err = BookUsecase.BookRepository.GetBookByISBN(isbn)
	return book, err
}

func (BookUsecase *BookUsecase) UpdateBookByISBN(isbn string, book *models.Book) (err error) {
	err = BookUsecase.BookRepository.UpdateBookByISBN(isbn, book)
	return err
}

func (BookUsecase *BookUsecase) DeleteBookByISBN(isbn string) (err error) {
	err = BookUsecase.BookRepository.DeleteBookByISBN(isbn)
	return err
}
