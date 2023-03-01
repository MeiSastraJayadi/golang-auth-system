package usecase

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	// "os"
	"time"

	"github.com/MeiSastraJayadi/golang-auth-system.git/mdl"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string, exp time.Duration) (*mdl.TokenInfo, error) {
  // key, keyErr := jwt.ParseEdPrivateKeyFromPEM([]byte(os.Getenv("SECRET_KEY")))
  // if keyErr != nil {
  //   return nil, keyErr
  // }
  key, keyErr := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if keyErr != nil {
    return nil, keyErr
  }

  token := jwt.New(jwt.SigningMethodES256)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(exp).Unix()
  claims["user"] = username
  claims["authorized"] = true
  tokenKey, err := token.SignedString(key)
  if err != nil {
    return nil, errors.New("Failed to generate token")
  }
  tkn := &mdl.TokenInfo{Token: tokenKey}
  return tkn, nil 
}
