package routes

import (
	"github.com/HanawuZ/gin-be-basics/controllers"
	"github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/usecases"

	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func AuthorRoutes(router *gin.Engine, db *gorm.DB) {
// 	authorRepository := repositories.NewAuthorRepository(db)

// 	authorController := &controllers.AuthorController{
// 		AuthorUsecase: usecases.NewAuthorUsecase(authorRepository),
// 	}
// 	router.GET("/authors", authorController.ListAuthor)
// }

func AuthorRoutes(router *fiber.App, db *gorm.DB) {
	authorRepository := repositories.NewAuthorRepository(db)

	authorController := &controllers.AuthorController{
		AuthorUsecase: usecases.NewAuthorUsecase(authorRepository),
	}
	router.Get("/authors", authorController.ListAuthor)
}
