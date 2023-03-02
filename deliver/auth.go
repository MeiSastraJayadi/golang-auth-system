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
  privateKey string
}

func NewDeliverAuth(db *sql.DB, privateKey string) *DeliverAuth {
  return &DeliverAuth{
    db : db, 
    privateKey: privateKey,
  }
}

func LoginRouter(db *sql.DB, privateKey string) *multiplexer.Router {
  newDeliver := NewDeliverAuth(db, privateKey)
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
  err = usecase.Login(user.Username, user.Password, lg.db, w, lg.privateKey)
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









