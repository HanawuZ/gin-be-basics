package usecases

import (
	"github.com/HanawuZ/gin-be-basics/interfaces"
	"github.com/HanawuZ/gin-be-basics/models"
)

type UserUseCase struct {
	userRepo interfaces.UserRepository
}

// NewUserUseCase is a constructor function that creates and returns a new instance
// of a type that implements the interfaces.UserUseCase interface.
func NewUserUseCase(repo interfaces.UserRepository) interfaces.UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

// Implement the interfaces.UserUseCase interface methods
// GetAllUser retrieves all users from the database.
// It returns a slice of User models and an error if any.
// func (u *UserUseCase) GetAllUser() (t []models.User, err error) {
// 	users, err := u.userRepo.GetAllUser()
// 	return users, err
// }

func (u *UserUseCase) CreateUser(t *models.User) (err error) {
	err = u.userRepo.CreateUser(t)
	// Implement your logic here
	return err
}

// func (u *UserUseCase) GetUser(t *models.User, username string) (err error) {
// 	// Implement your logic here
// 	err = u.userRepo.GetUser(t, username)
// 	return err
// }

// func (u *UserUseCase) UpdateUser(t *models.User, username string) (err error) {
// 	// Implement your logic here
// 	err = u.userRepo.UpdateUser(t, username)
// 	return err
// }

// func (u *UserUseCase) DeleteUser(t *models.User, username string) (err error) {
// 	// Implement your logic here
// 	err = u.userRepo.DeleteUser(t, username)
// 	return err
// }
