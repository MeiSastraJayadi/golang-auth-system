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

type DeliverAuth struct {
  db *sql.DB 
}

func NewDeliverAuth(db *sql.DB) *DeliverAuth {
  return &DeliverAuth{
    db : db, 
  }
}

func LoginRouter(db *sql.DB) *multiplexer.Router {
  newDeliver := NewDeliverAuth(db)
  router := multiplexer.NewRouter("/").SetPrefix("auth")  
  router.Methods(http.MethodPost).HandleFunc("/login", newDeliver.LoginHandler)
  router.Methods(http.MethodPost).HandleFunc("/register", newDeliver.Registration)
  return router
}

func(lg *DeliverAuth) LoginHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("/auth/login")
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

func(lg *DeliverAuth) Registration(w http.ResponseWriter, r *http.Request) {
  log.Println("/auth/register")
  user := &mdl.UserRegis{}
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(user)
  if err != nil {
    log.Println(err.Error())
    fmt.Fprintln(w, err.Error())
    return
  }
  err = usecase.Regis(user.Username, user.Password, user.Password2, lg.db)
  if err != nil {
    log.Println(err.Error())
    fmt.Fprintln(w, err.Error())
    return
  }
}









