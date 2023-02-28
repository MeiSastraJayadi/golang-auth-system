package db

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateConnection(database string, driver string) *sql.DB {
  address := fmt.Sprintf("root:@tcp(localhost:3306)/%s", database)
  conn, err := sql.Open(driver, address)
  if err != nil {
    return nil
  }
  err = conn.Ping()
  if err != nil {
    return nil
  }
  conn.SetMaxIdleConns(20)
  conn.SetMaxOpenConns(100)
  conn.SetConnMaxLifetime(time.Hour * 3)
  conn.SetConnMaxIdleTime(time.Minute * 10)
  return conn 
}
