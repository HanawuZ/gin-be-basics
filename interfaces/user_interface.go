package interfaces

import (
	"github.com/HanawuZ/gin-be-basics/models"
)

type UserUseCase interface {
	// GetAllUser() (t []models.User, err error)
	CreateUser(t *models.User) (err error)
	// GetUser(t *models.User, username string) (err error)
	// UpdateUser(t *models.User, username string) (err error)
	// DeleteUser(t *models.User, username string) (err error)
}

type UserRepository interface {
	// GetAllUser() (t []models.User, err error)
	CreateUser(t *models.User) (err error)
	// GetUser(t *models.User, username string) (err error)
	// UpdateUser(t *models.User, username string) (err error)
	// DeleteUser(t *models.User, username string) (err error)
}
