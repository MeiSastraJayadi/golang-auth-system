package usecase

import (
	"database/sql"
	"errors"

	database "github.com/MeiSastraJayadi/golang-auth-system.git/db"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string, db *sql.DB) error {
  user := database.FetchUser(username, db)
  if user == nil {
    return errors.New("The user doesnt exist")
  }
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if err != nil {
    return err
  }
  return nil
} 
