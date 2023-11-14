package repositories

import (
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
	"gorm.io/gorm"
)

// UserRepository is a struct to store db connection
type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) interfaces.AuthorRepository {
	return &AuthorRepository{
		DB: db,
	}
}

func (authorRepo *AuthorRepository) ListAuthor() (authors []models.Author, err error) {
	if err := authorRepo.DB.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
