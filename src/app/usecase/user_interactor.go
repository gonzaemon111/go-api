package usecase

import "app/models"

// UserInteractor is the gateway to controllers
type UserInteractor struct {
	UserRepository UserRepository
}

// Add is to add a new user
func (interactor *UserInteractor) Add(u models.User) (models.User, error) {
	newUser, err := interactor.UserRepository.Store(u)
	return newUser, err
}

// UserByID is to get a user by ID
func (interactor *UserInteractor) UserByID(id int) (models.User, error) {
	user, err := interactor.UserRepository.FindByID(id)
	return user, err
}

// Users is to get all users
func (interactor *UserInteractor) Users() ([]models.User, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}
