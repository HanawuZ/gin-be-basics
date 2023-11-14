package controllers

// import gorm
import (
	"net/http"

	"github.com/HanawuZ/gin-be-basics/models"
	"github.com/HanawuZ/gin-be-basics/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecases.UserUseCase
}

func NewUsersController(r *gin.Engine, UserUsecase usecases.UserUseCase) {
	controllers := &UserController{
		UserUsecase: UserUsecase,
	}
	r.POST("/", controllers.CreateUser)

}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.UserUsecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}
