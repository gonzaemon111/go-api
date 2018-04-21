package database

import (
	"app/models"
)

// SQLHandler is the definition of infrastructur/sqlhandler's behavior
type SQLHandler interface {
	Insert(models.User) (models.User, error)
	Find(int) (models.User, error)
	All() ([]models.User, error)
	Update(int, models.User) (models.User, error)
	Delete(int) error
}
