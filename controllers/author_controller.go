package controllers

// import gorm
import (
	"net/http"

	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/gofiber/fiber/v2"
	// "github.com/gin-gonic/gin"
)

type AuthorController struct {
	AuthorUsecase interfaces.AuthorUsecase
}

// func (a *AuthorController) ListAuthor(c *gin.Context) {
// 	authors, err := a.AuthorUsecase.ListAuthor()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, authors)
// }

func (a *AuthorController) ListAuthor(c *fiber.Ctx) error {
	authors, err := a.AuthorUsecase.ListAuthor()
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	c.Status(http.StatusOK).JSON(authors)
	return nil
	// c.JSON(http.StatusOK, authors)
}
