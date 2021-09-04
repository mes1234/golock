package auth

import (
	stdjwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("test")
var Keys stdjwt.Keyfunc = func(token *stdjwt.Token) (interface{}, error) {
	return key, nil
}
