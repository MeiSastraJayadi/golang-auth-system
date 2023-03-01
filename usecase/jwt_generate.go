package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/MeiSastraJayadi/golang-auth-system.git/mdl"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string, exp time.Duration) (*mdl.TokenInfo, error) {
  secretKey := []byte(os.Getenv("SECRET_KEY"))
  token := jwt.New(jwt.SigningMethodEdDSA)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(exp).Unix()
  claims["user"] = username
  claims["authorized"] = true
  tokenKey, err := token.SignedString(secretKey)
  if err != nil {
    return nil, errors.New("Failed to generate token")
  }
  tkn := &mdl.TokenInfo{Token: tokenKey}
  return tkn, nil 
}
