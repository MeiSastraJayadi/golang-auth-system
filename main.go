package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/MeiSastraJayadi/acacia/multiplexer"
	database "github.com/MeiSastraJayadi/golang-auth-system.git/db"
	"github.com/MeiSastraJayadi/golang-auth-system.git/deliver"
	"github.com/MeiSastraJayadi/golang-auth-system.git/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  db := database.CreateConnection("go-auth", "mysql")
  err := godotenv.Load()
  if err != nil {
    log.Println(".env file doesn't exist")
  }
  defer db.Close()
  if db == nil {
    os.Exit(1)
  }

  privKey := os.Getenv("PRIVATE_KEY")
  privKey = strings.ReplaceAll(privKey, "\\n", "\n")

  // err := database.AddUser("dekmei_13", "meisastra", db)
  // if err != nil {
  //   fmt.Printf("%s", err.Error())
  // }
  // err := usecase.Login("dekmei_13", "meisastra", Database)
  // if err != nil {
  //   fmt.Println(err.Error())
  // }



  authjwt := usecase.NewAuth(privKey)
  mainRouter := multiplexer.NewRouter("/")
  mainRouter.Methods(http.MethodGet).HandleFunc("/slow", SlowHandler)
  mainRouter.Methods(http.MethodPost, http.MethodGet).HandleFunc("/", authjwt.VerifyJWT(MainHandler))
  loginRouter := deliver.LoginRouter(db, privKey)
  err = mainRouter.SubRouter(loginRouter)
  if err != nil {
    os.Exit(1)
  }

  server := http.Server{
    Addr: "localhost:9090",
    Handler: mainRouter,
    IdleTimeout: time.Hour,
    WriteTimeout: time.Minute*3,
    ReadTimeout: time.Minute*3,
  }

  go func(){
    log.Println("Server runing...")
    runError := server.ListenAndServe()
    if runError != nil {
      os.Exit(1)
    }
  }()
  ch := make(chan os.Signal)
  signal.Notify(ch, os.Interrupt)
  signal.Notify(ch, os.Kill)
  interruptString := <-ch
  log.Printf("Interrupt happen : %s", interruptString.String())

  tc, _ := context.WithTimeout(context.Background(), time.Hour*3)
  server.Shutdown(tc)
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
  time.Sleep(time.Second * 15)
  log.Println("/slow")
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  token := ctx.Value(usecase.ContextValue{}).(*jwt.Token)
  claims := token.Claims.(jwt.MapClaims)
  if r.Method == http.MethodPost {
    fmt.Fprintf(w, "/home POST")
  }
  log.Println("/home")
  log.Printf("Hello %s", claims["user"].(string))
}

