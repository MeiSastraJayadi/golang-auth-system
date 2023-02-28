package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MeiSastraJayadi/golang-auth-system.git/mdl"
)

func FetchUser(username string, db *sql.DB) *mdl.User {
  ctx := context.Background()
  query := fmt.Sprintf("SELECT * FROM user_table WHERE username = '%s'", username)
  result, err := db.QueryContext(ctx, query)
  if err != nil {
    return nil
  }

  user := &mdl.User{}

  if result.Next() {
    var (
      id int
      username string
      password string
    )
    scanErr := result.Scan(&id, &username, &password) 
    if scanErr != nil {
      return nil
    }
    user.Id = id
    user.Username = username
    user.Password = password
  } else {
    return nil
  }
  return user
}



