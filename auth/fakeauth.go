package auth

import (
	stdjwt "github.com/dgrijalva/jwt-go"
	"os"
)

var key = []byte(os.Getenv("go_key"))
var Keys stdjwt.Keyfunc = func(token *stdjwt.Token) (interface{}, error) {
	return key, nil
}
