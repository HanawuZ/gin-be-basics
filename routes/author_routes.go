package routes

import (
	"github.com/HanawuZ/gin-be-basics/controllers"
	"github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorRoutes(router *gin.Engine, db *gorm.DB) {
	authorRepository := repositories.NewAuthorRepository(db)

	authorController := &controllers.AuthorController{
		AuthorUsecase: usecases.NewAuthorUsecase(authorRepository),
	}
	router.GET("/authors", authorController.ListAuthor)
}
