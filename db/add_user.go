package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(username string, password string, db *sql.DB) error {
  ctx := context.Background()
  hashedPassword, encErr := bcrypt.GenerateFromPassword([]byte(password), 10)
  if encErr != nil {
    return encErr
  }
  query := fmt.Sprintf("INSERT INTO user_table (username, user_password) VALUES ('%s', '%s')", username, hashedPassword)
  _, err := db.ExecContext(ctx, query)
  if err != nil {
    return errors.New("Error when try inserting data to table")
  }
  return nil
}




