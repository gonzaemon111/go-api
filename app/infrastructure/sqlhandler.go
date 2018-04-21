package infrastructure

import (
  "os"
  "fmt"
  "app/models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type SqlHandler struct {
  Db *gorm.DB
}

func NewSqlHandler() *SqlHandler {
  authConnect := fmt.Sprintf( "host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
    os.Getenv( "AUTH_DB_HOST" ),
    os.Getenv( "AUTH_DB_PORT" ),
    os.Getenv( "AUTH_DB_USER" ),
    os.Getenv( "AUTH_DB_NAME" ),
    os.Getenv( "AUTH_DB_PASSWORD" ))

  db, err := gorm.Open( "postgres", authConnect )
  
  if err != nil {
    fmt.Println(err)
    panic("failed to connect database")
  } else {
    fmt.Println("successly connect database")
  }

  // Migratie Users
  db.AutoMigrate(&models.User{})
  sqlHandler := new(SqlHandler)
  sqlHandler.Db = db
  return sqlHandler
}

func (handler *SqlHandler) Insert(user models.User) (models.User, error) {
  var new_user *models.User
  res := handler.Db.Create(&user)
  if res.Error != nil {
    return models.User{}, res.Error
  }
  new_user = res.Value.(*models.User)
  return *new_user, nil
}

func (handler *SqlHandler) Find(id int) (models.User, error) {
  user := models.User{}
  res := handler.Db.First(&user, id)

  if res.Error != nil {
    return models.User{}, res.Error
  }
  return user, nil
}

func (handler *SqlHandler) All() ([]models.User, error) {
  users := []models.User{}
  res := handler.Db.Find(&users)

  if res.Error != nil {
    return []models.User{}, res.Error
  }
  return users, nil
}
