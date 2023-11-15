package controllers

// import gorm
import (
	"fmt"
	"net/http"

	"github.com/HanawuZ/gin-be-basics/dto"
	"github.com/HanawuZ/gin-be-basics/interfaces"

	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	BookUsecase interfaces.BookUsecase
}

func (b *BookController) ListBook(c *fiber.Ctx) error {
	books, err := b.BookUsecase.ListBooks()
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}

	// c.JSON(http.StatusOK, books)
	c.Status(http.StatusOK).JSON(books)
	return nil
}

func (b *BookController) AddBook(c *fiber.Ctx) error {
	// var book models.Book
	var bookReuqest dto.BookRequest

	// if err := c.ShouldBindJSON(&bookReuqest); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := c.BodyParser(&bookReuqest); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	fmt.Println(bookReuqest)
	err := b.BookUsecase.AddBook(&bookReuqest)
	if err != nil {

		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": "test",
	// })
	c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": bookReuqest,
	})
	return nil
}

func (b *BookController) GetBookByISBN(c *fiber.Ctx) error {
	isbn := c.Params("isbn")
	book, err := b.BookUsecase.GetBookByISBN(isbn)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// c.JSON(http.StatusOK, book)
	c.Status(http.StatusOK).JSON(book)
	return nil
}

func (b *BookController) UpdateBookByISBN(c *fiber.Ctx) error {
	var isbn = c.Params("isbn")
	book, err := b.BookUsecase.GetBookByISBN(isbn)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}

	c.Status(http.StatusOK).JSON(book)
	return nil
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": book,
	// })
}

func (b *BookController) DeleteBookByISBN(c *fiber.Ctx) error {
	var isbn = c.Params("isbn")
	err := b.BookUsecase.DeleteBookByISBN(isbn)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": "deleted",
	// }
	c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "deleted",
	})
	return nil
}
