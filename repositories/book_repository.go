package repositories

import (
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
	"gorm.io/gorm"
)

// UserRepository is a struct to store db connection
type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (bookRepo *BookRepository) ListBooks() (books []models.Book, err error) {
	if err := bookRepo.DB.Preload("Authors").Preload("Publisher").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bookRepo *BookRepository) AddBook(book *models.Book) (err error) {
	if err := bookRepo.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (bookRepo *BookRepository) GetBookByISBN(isbn string) (book models.Book, err error) {
	if err := bookRepo.DB.Where("isbn = ?", isbn).Preload("Authors").Preload("Publisher").First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (bookRepo *BookRepository) UpdateBookByISBN(isbn string, book *models.Book) (err error) {
	if err := bookRepo.DB.Model(&models.Book{}).Where("isbn = ?", isbn).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (bookRepo *BookRepository) DeleteBookByISBN(isbn string) (err error) {
	if err := bookRepo.DB.Where("isbn = ?", isbn).Delete(&models.Book{}).Error; err != nil {
		return err
	}
	return nil
}
