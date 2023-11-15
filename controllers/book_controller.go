package controllers

// import gorm
import (
	"fmt"
	"net/http"

	"github.com/HanawuZ/gin-be-basics/dto"
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookUsecase interfaces.BookUsecase
}

func (b *BookController) ListBook(c *gin.Context) {
	books, err := b.BookUsecase.ListBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (b *BookController) AddBook(c *gin.Context) {
	// var book models.Book
	var bookReuqest dto.BookRequest
	if err := c.ShouldBindJSON(&bookReuqest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bookReuqest)
	err := b.BookUsecase.AddBook(&bookReuqest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "test",
	})
}

func (b *BookController) GetBookByISBN(c *gin.Context) {
	var isbn = c.Param("isbn")
	book, err := b.BookUsecase.GetBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *BookController) UpdateBookByISBN(c *gin.Context) {
	var isbn = c.Param("isbn")
	book, err := b.BookUsecase.GetBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
	// err = BookUsecase.BookRepository.UpdateBookByISBN(isbn, book)
}

func (b *BookController) DeleteBookByISBN(c *gin.Context) {
	var isbn = c.Param("isbn")
	err := b.BookUsecase.DeleteBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "deleted",
	})
}
