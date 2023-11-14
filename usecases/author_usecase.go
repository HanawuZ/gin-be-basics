package usecases

import (
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
)

type AuthorUsecase struct {
	AuthorRepository interfaces.AuthorRepository
}

func NewAuthorUsecase(authorRepo interfaces.AuthorRepository) interfaces.AuthorUsecase {
	return &AuthorUsecase{
		AuthorRepository: authorRepo,
	}
}

func (authorUsecase *AuthorUsecase) ListAuthor() (authors []models.Author, err error) {
	authors, err = authorUsecase.AuthorRepository.ListAuthor()
	return authors, err
}
