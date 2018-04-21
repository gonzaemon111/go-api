package usecase

import "app/models"

// UserRepository is input from interfaces/database
type UserRepository interface {
	Store(models.User) (models.User, error)
	FindByID(int) (models.User, error)
	FindAll() ([]models.User, error)
}
