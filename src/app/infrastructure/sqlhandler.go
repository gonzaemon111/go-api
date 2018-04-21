package infrastructure

import (
	"app/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SQLHandler manage DB processing
type SQLHandler struct {
	Db *gorm.DB
}

// NewSQLHandler generates the instance of SQLHandler
func NewSQLHandler() *SQLHandler {
	authConnect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("AUTH_DB_HOST"),
		os.Getenv("AUTH_DB_PORT"),
		os.Getenv("AUTH_DB_USER"),
		os.Getenv("AUTH_DB_NAME"),
		os.Getenv("AUTH_DB_PASSWORD"))

	db, err := gorm.Open("postgres", authConnect)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("successly connect database")
	}

	// Migratie Users
	db.AutoMigrate(&models.User{})
	sqlHandler := new(SQLHandler)
	sqlHandler.Db = db
	return sqlHandler
}

// Insert is to insert models
func (handler *SQLHandler) Insert(user models.User) (models.User, error) {
	var newUser *models.User

	createResult := handler.Db.Create(&user)
	if createResult.Error != nil {
		return models.User{}, createResult.Error
	}
	newUser = createResult.Value.(*models.User)
	return *newUser, nil
}

// Find is to get user by id
func (handler *SQLHandler) Find(id int) (models.User, error) {
	user := models.User{}
	findResult := handler.Db.First(&user, id)

	if findResult.Error != nil {
		return models.User{}, findResult.Error
	}
	return user, nil
}

// All is to get all user
func (handler *SQLHandler) All() ([]models.User, error) {
	users := []models.User{}
	findResult := handler.Db.Find(&users)

	if findResult.Error != nil {
		return []models.User{}, findResult.Error
	}
	return users, nil
}

// Update is to update a user
func (handler *SQLHandler) Update(id int, user models.User) (models.User, error) {
	var userBefore, userAfter models.User
	// find a user by id
	findResult := handler.Db.First(&userBefore, id)
	if findResult.Error != nil {
		return models.User{}, findResult.Error
	}
	// update the found user
	updateResult := handler.Db.Model(&userBefore).Update(user)
	if updateResult.Error != nil {
		return models.User{}, updateResult.Error
	}
	// get the updated user
	handler.Db.First(&userAfter, id)
	return userAfter, nil
}

// Delete is to delete user by id
func (handler *SQLHandler) Delete(id int) error {
	user := models.User{}
	// find a user by id
	findResult := handler.Db.First(&user, id)
	if findResult.Error != nil {
		return findResult.Error
	}
	// delete the found user
	deleteResult := handler.Db.Delete(&user)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}
