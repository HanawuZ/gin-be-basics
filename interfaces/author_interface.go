package interfaces

import (
	"github.com/HanawuZ/gin-be-basics/models"
)

type AuthorRepository interface {
	ListAuthor() (authors []models.Author, err error)
}
type AuthorUsecase interface {
	ListAuthor() (authors []models.Author, err error)
}
