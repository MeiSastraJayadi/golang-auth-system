package main

import (
	"fmt"
	"os"

	database "github.com/MeiSastraJayadi/golang-auth-system.git/db"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db := database.CreateConnection("go-auth", "mysql")
  defer db.Close()
  if db == nil {
    os.Exit(1)
  }
  err := database.AddUser("dekmei_13", "meisastra", db)
  if err != nil {
    fmt.Printf("%s", err.Error())
  }
}
