package database

import "app/models"

type UserRepository struct {
  SqlHandler
}

func (repo *UserRepository) Store(user models.User) (models.User, error) {
  new_user, err := repo.Insert(user)
  return new_user, err
}

func (repo *UserRepository) FindById(id int) (models.User, error) {
  user, err := repo.Find(id)
  return user, err
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
  users, err := repo.All()
  return users, err
}
