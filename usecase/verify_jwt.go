package usecase

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header["Authorization"] == nil {
      fmt.Fprintln(w, "Unauthorized request")
      log.Println("Unauthorized request")
      return
    }


    tkn := r.Header["Authorization"][0]
    log.Println(tkn)
    check := strings.Split(tkn, " ")
    if check[0] == "Bearer" {
      tkn = check[1]
    } else {
      fmt.Fprintln(w, "Unauthorized request : the token that provided is not a bearer token")
      log.Println("Unauthorized request : the token that provided is not a bearer token")
      return
    }

    token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
      _, ok := token.Method.(*jwt.SigningMethodECDSA)
      if !ok {
        return nil, errors.New("Token signing method is invalid")
      } 
      result, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
      if err != nil {
        return nil, err
      }
      return result, nil
    })

    if err != nil {
      fmt.Fprintln(w, "Unauthorized request : error when parsing token")
      log.Println("Unauthorized request : error when parsing token")
      return
    }

    if token.Valid {
      next(w, r)
    } else {
      fmt.Fprintln(w, "Unauthorized request : the token is invalid")
      log.Println("Unauthorized request : the token is invalid")
      return
    }

  })

}
