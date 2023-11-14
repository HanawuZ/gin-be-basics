package repositories

import (
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
	"gorm.io/gorm"
)

// UserRepository is a struct to store db connection
type UserRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) CreateUser(user *models.User) (err error) {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
