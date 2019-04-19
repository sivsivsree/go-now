package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sivsivsree/go-now/web/modals"
	"os"
)

func CreateJWT(tk *modals.Token) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN")))
	return tokenString
}

func VerifyJWT(token string) bool {
	return false
}
