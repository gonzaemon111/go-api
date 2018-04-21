package usecase

import "app/models"

type UserRepository interface {
  Store(models.User) (models.User, error)
  FindById(int) (models.User, error)
  FindAll() ([]models.User, error)
}
