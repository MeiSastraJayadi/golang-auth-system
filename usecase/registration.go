package usecase

import (
	"database/sql"
	"errors"
  database "github.com/MeiSastraJayadi/golang-auth-system.git/db"
)

func Regis(username string, password1 string, password2 string, db *sql.DB) error {
  if password1 != password2 {
    return errors.New("Confirmation password is wrong")
  }
  err := database.AddUser(username, password1, db) 
  if err != nil {
    return errors.New("Failed to add user")
  }
  return nil
}
