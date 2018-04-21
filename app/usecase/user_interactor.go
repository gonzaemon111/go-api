package usecase

import "app/models"

type UserInteractor struct {
  UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u models.User) (models.User, error) {
  new_user, err := interactor.UserRepository.Store(u)
  return new_user, err
}

func (interactor *UserInteractor) UserById(id int) (models.User, error) {
  user, err := interactor.UserRepository.FindById(id)
  return user, err
}

func (interactor *UserInteractor) Users() ([]models.User, error) {
  users, err := interactor.UserRepository.FindAll()
  return users, err
}
