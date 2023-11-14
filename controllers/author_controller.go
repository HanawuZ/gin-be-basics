package controllers

// import gorm
import (
	"net/http"

	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	AuthorUsecase interfaces.AuthorUsecase
}

func (a *AuthorController) ListAuthor(c *gin.Context) {
	authors, err := a.AuthorUsecase.ListAuthor()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": authors,
	})
}
