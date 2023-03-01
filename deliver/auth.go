package deliver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MeiSastraJayadi/acacia/multiplexer"
	"github.com/MeiSastraJayadi/golang-auth-system.git/mdl"
	"github.com/MeiSastraJayadi/golang-auth-system.git/usecase"
)

type DeliverLogin struct {
  db *sql.DB 
}

func NewDeliverLogin(db *sql.DB) *DeliverLogin {
  return &DeliverLogin{
    db : db, 
  }
}

func LoginRouter(db *sql.DB) *multiplexer.Router {
  newDeliver := NewDeliverLogin(db)
  router := multiplexer.NewRouter("/").SetPrefix("auth")  
  router.Methods(http.MethodPost).HandleFunc("/login", newDeliver.LoginHandler)
  return router
}

func(lg *DeliverLogin) LoginHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("/login")
  user := &mdl.User{}
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(user)
  if err != nil {
    fmt.Fprint(w, err.Error())
    log.Println("Login error")
  }
  err = usecase.Login(user.Username, user.Password, lg.db)
  if err != nil {
    fmt.Fprint(w, err.Error())
    log.Println("Login error")
  }
}









