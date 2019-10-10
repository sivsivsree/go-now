package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sivsivsree/go-now/practices"
	"log"
	"os"
)

func CreateJWT(tk *practices.Token) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN")))
	return tokenString
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {

	var toeknError error = nil

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN")), nil
	})

	if token.Valid {
		log.Println("Token valid")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			toeknError = errors.New("Malformed token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			toeknError = errors.New("Token is either expired or not active yet")
		} else {
			toeknError = errors.New("Couldn't handle this token")
		}
	} else {
		log.Println("Couldn't handle this token:-", err)
		toeknError = errors.New("Couldn't handle this token")
	}

	return token, toeknError

}
