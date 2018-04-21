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

