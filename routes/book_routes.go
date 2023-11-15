package routes

import (
	"github.com/HanawuZ/gin-be-basics/controllers"
	"github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/usecases"

	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func BookRoutes(router *gin.Engine, db *gorm.DB) {
// 	bookRepository := repositories.NewBookRepository(db)

// 	bookController := &controllers.BookController{
// 		BookUsecase: usecases.NewBookUsecase(bookRepository),
// 	}

//		router.GET("/books", bookController.ListBook)
//		router.POST("/books", bookController.AddBook)
//		router.GET("/books/:isbn", bookController.GetBookByISBN)
//		router.PUT("/books/:isbn", bookController.UpdateBookByISBN)
//		router.DELETE("/books/:isbn", bookController.DeleteBookByISBN)
//	}
func BookRoutes(router *fiber.App, db *gorm.DB) {
	bookRepository := repositories.NewBookRepository(db)

	bookController := &controllers.BookController{
		BookUsecase: usecases.NewBookUsecase(bookRepository),
	}

	router.Get("/books", bookController.ListBook)
	router.Post("/books", bookController.AddBook)
	router.Get("/books/:isbn", bookController.GetBookByISBN)
	router.Put("/books/:isbn", bookController.UpdateBookByISBN)
	router.Delete("/books/:isbn", bookController.DeleteBookByISBN)
}
