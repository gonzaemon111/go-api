package database

import (
  "app/models"
)

type SqlHandler interface {
  Insert(models.User) (models.User, error)
  Find(int) (models.User, error)
  All() ([]models.User, error)
}
