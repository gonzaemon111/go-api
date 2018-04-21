package database

import "app/models"

// UserRepository manages DB processing
type UserRepository struct {
	SQLHandler
}

// Store is to insert a new user to DB
func (repo *UserRepository) Store(user models.User) (models.User, error) {
	newUser, err := repo.Insert(user)
	return newUser, err
}

// FindByID is to get a user by ID
func (repo *UserRepository) FindByID(id int) (models.User, error) {
	user, err := repo.Find(id)
	return user, err
}

// FindAll is to get all users
func (repo *UserRepository) FindAll() ([]models.User, error) {
	users, err := repo.All()
	return users, err
}
