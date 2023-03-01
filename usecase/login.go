package usecase

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"time"

	database "github.com/MeiSastraJayadi/golang-auth-system.git/db"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string, db *sql.DB, w io.Writer) error {
  user := database.FetchUser(username, db)
  if user == nil {
    return errors.New("The user doesnt exist")
  }
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if err != nil {
    return err
  }
  token, err := GenerateJWT(username, time.Minute*10) 
  if err != nil {
    return err
  }
  encoder := json.NewEncoder(w)
  encoder.Encode(token)
  return nil
} 
