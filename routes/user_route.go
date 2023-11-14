package routes

import (
	"github.com/HanawuZ/gin-be-basics/controllers"
	"github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUsersRepository(db)

	userController := &controllers.UserController{
		UserUsecase: usecases.NewUserUsecase(userRepository),
	}
	router.POST("/", userController.CreateUser)
}
