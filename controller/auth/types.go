package auth

import jwt "github.com/dgrijalva/jwt-go"

var secret = []byte("my secrect key...")

type Credentials struct {
	Email, Password string
}

type Claims struct {
	ID uint
	jwt.StandardClaims
}
