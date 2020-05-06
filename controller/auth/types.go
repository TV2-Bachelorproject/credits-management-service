package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Email, Password string
}

type Claims struct {
	ID    uint
	Admin bool
	jwt.StandardClaims
}
