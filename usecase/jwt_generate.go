package usecase

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"

	// "os"
	"time"

	"github.com/MeiSastraJayadi/golang-auth-system.git/mdl"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey string = os.Getenv("PRIVATE_KEY") 

func GenerateJWT(username string, exp time.Duration, privatekey string) (*mdl.TokenInfo, error) {

  privK := Decode(privatekey)

  token := jwt.New(jwt.SigningMethodES256)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(exp).Unix()
  claims["user"] = username
  claims["authorized"] = true
  tokenKey, err := token.SignedString(privK)
  if err != nil {
    log.Println(err.Error())
    return nil, errors.New("Failed to generate token")
  }
  tkn := &mdl.TokenInfo{Token: tokenKey}
  return tkn, nil 
}

func Decode(privateKey string) *ecdsa.PrivateKey {
  block, _ := pem.Decode([]byte(privateKey))
  x509Block := block.Bytes
  privKey, _ := x509.ParseECPrivateKey(x509Block)
  return privKey
}



