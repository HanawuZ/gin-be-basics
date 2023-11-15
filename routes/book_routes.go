package routes

import (
	"github.com/HanawuZ/gin-be-basics/controllers"
	"github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRoutes(router *gin.Engine, db *gorm.DB) {
	bookRepository := repositories.NewBookRepository(db)

	bookController := &controllers.BookController{
		BookUsecase: usecases.NewBookUsecase(bookRepository),
	}

	router.GET("/books", bookController.ListBook)
	router.POST("/books", bookController.AddBook)
	router.GET("/books/:isbn", bookController.GetBookByISBN)
	router.PUT("/books/:isbn", bookController.UpdateBookByISBN)
	router.DELETE("/books/:isbn", bookController.DeleteBookByISBN)
	// router.GET("/authors", authorController.ListAuthor)
}
