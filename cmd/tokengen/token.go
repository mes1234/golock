package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type customClaims struct {
	ClientID string `json:"clientId"`
	jwt.StandardClaims
}

func generateToken(signingKey []byte, clientID string) (string, error) {
	claims := customClaims{
		clientID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1000).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func main() {

	key := []byte("test")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id: "witek",
		})
	tokenString, _ := token.SignedString(key)

	fmt.Println("TOKEN:", tokenString)

}
