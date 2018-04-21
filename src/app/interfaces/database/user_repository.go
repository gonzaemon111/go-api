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

// Edit is to edit a user
func (repo *UserRepository) Edit(id int, user models.User) (models.User, error) {
	updatedUser, err := repo.Update(id, user)
	return updatedUser, err
}

// Destroy is to destroy a user
func (repo *UserRepository) Destroy(id int) error {
	err := repo.Delete(id)
	return err
}
