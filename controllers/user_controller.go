package controllers

// import gorm
import (
	"net/http"

	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserUsecase interfaces.UserUsecase
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User

	// * Using gin
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// * Using fiber
	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPassword)

	if err := u.UserUsecase.CreateUser(&user); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	// c.JSON(http.StatusCreated, gin.H{
	// 	"data":    user,
	// 	"message": "User created successfully",
	// })
	c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": user,
	})
	return nil

}
