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
	res := handler.Db.Create(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	newUser = res.Value.(*models.User)
	return *newUser, nil
}

// Find is to get user by id
func (handler *SQLHandler) Find(id int) (models.User, error) {
	user := models.User{}
	res := handler.Db.First(&user, id)

	if res.Error != nil {
		return models.User{}, res.Error
	}
	return user, nil
}

// All is to get all user
func (handler *SQLHandler) All() ([]models.User, error) {
	users := []models.User{}
	res := handler.Db.Find(&users)

	if res.Error != nil {
		return []models.User{}, res.Error
	}
	return users, nil
}
