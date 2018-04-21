package models

import (
  "time"
)

type User struct {
  ID    uint `json:"id" gorm:"primary_key"`
  Name  string `json:"name" gorm:"not null;size:255"`
  Email string `json:"email" gorm:"not null;size:255"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
